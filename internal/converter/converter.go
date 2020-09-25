package converter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	"github.com/alecthomas/jsonschema"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/sirupsen/logrus"
)

const (
	messageDelimiter = "+"
)

// Converter is everything you need to convert protos to JSONSchemas:
type Converter struct {
	AllFieldsRequired            bool
	AllowNullValues              bool
	DisallowAdditionalProperties bool
	DisallowBigIntsAsStrings     bool
	PrefixSchemaFilesWithPackage bool
	UseProtoAndJSONFieldnames    bool
	logger                       *logrus.Logger
	sourceInfo                   *sourceCodeInfo
	messageTargets               []string
}

// New returns a configured *Converter:
func New(logger *logrus.Logger) *Converter {
	return &Converter{
		logger: logger,
	}
}

// ConvertFrom tells the convert to work on the given input:
func (c *Converter) ConvertFrom(rd io.Reader) (*plugin.CodeGeneratorResponse, error) {
	c.logger.Debug("Reading code generation request")
	input, err := ioutil.ReadAll(rd)
	if err != nil {
		c.logger.WithError(err).Error("Failed to read request")
		return nil, err
	}

	req := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(input, req)
	if err != nil {
		c.logger.WithError(err).Error("Can't unmarshal input")
		return nil, err
	}

	c.logger.Debug("Converting input")
	return c.convert(req)
}

func (c *Converter) parseGeneratorParameters(parameters string) {
	for _, parameter := range strings.Split(parameters, ",") {
		switch parameter {
		case "all_fields_required":
			c.AllFieldsRequired = true
		case "allow_null_values":
			c.AllowNullValues = true
		case "debug":
			c.logger.SetLevel(logrus.DebugLevel)
		case "disallow_additional_properties":
			c.DisallowAdditionalProperties = true
		case "disallow_bigints_as_strings":
			c.DisallowBigIntsAsStrings = true
		case "prefix_schema_files_with_package":
			c.PrefixSchemaFilesWithPackage = true
		case "proto_and_json_fieldnames":
			c.UseProtoAndJSONFieldnames = true
		}

		// look for specific message targets
		// message types are separated by messageDelimiter "+"
		// examples:
		// 		messages=[foo+bar]
		// 		messages=[foo]
		rx := regexp.MustCompile(`messages=\[([^\]]+)\]`)
		if matches := rx.FindStringSubmatch(parameter); len(matches) == 2 {
			c.messageTargets = strings.Split(matches[1], messageDelimiter)
		}
	}
}

// Converts a proto "ENUM" into a JSON-Schema:
func (c *Converter) convertEnumType(enum *descriptor.EnumDescriptorProto) (jsonschema.Type, error) {

	// Prepare a new jsonschema.Type for our eventual return value:
	jsonSchemaType := jsonschema.Type{
		Version: jsonschema.Version,
	}

	// Generate a description from src comments (if available)
	if src := c.sourceInfo.GetEnum(enum); src != nil {
		jsonSchemaType.Description = formatDescription(src)
	}

	// Allow both strings and integers:
	jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: "string"})
	jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: "integer"})

	// Add the allowed values:
	for _, enumValue := range enum.Value {
		jsonSchemaType.Enum = append(jsonSchemaType.Enum, enumValue.Name)
		jsonSchemaType.Enum = append(jsonSchemaType.Enum, enumValue.Number)
	}

	return jsonSchemaType, nil
}

// Converts a proto file into a JSON-Schema:
func (c *Converter) convertFile(file *descriptor.FileDescriptorProto) ([]*plugin.CodeGeneratorResponse_File, error) {
	// Input filename:
	protoFileName := path.Base(file.GetName())

	// Prepare a list of responses:
	var response []*plugin.CodeGeneratorResponse_File

	// user wants specific messages
	genSpecificMessages := len(c.messageTargets) > 0

	// Warn about multiple messages / enums in files:
	if !genSpecificMessages && len(file.GetMessageType()) > 1 {
		c.logger.WithField("schemas", len(file.GetMessageType())).WithField("proto_filename", protoFileName).Warn("protoc-gen-jsonschema will create multiple MESSAGE schemas from one proto file")
	}
	if len(file.GetEnumType()) > 1 {
		c.logger.WithField("schemas", len(file.GetMessageType())).WithField("proto_filename", protoFileName).Warn("protoc-gen-jsonschema will create multiple ENUM schemas from one proto file")
	}

	// Generate standalone ENUMs:
	if len(file.GetMessageType()) == 0 {
		for _, enum := range file.GetEnumType() {
			jsonSchemaFileName := c.generateSchemaFilename(file, enum.GetName())
			c.logger.WithField("proto_filename", protoFileName).WithField("enum_name", enum.GetName()).WithField("jsonschema_filename", jsonSchemaFileName).Info("Generating JSON-schema for stand-alone ENUM")

			// Convert the ENUM:
			enumJSONSchema, err := c.convertEnumType(enum)
			if err != nil {
				c.logger.WithError(err).WithField("proto_filename", protoFileName).Error("Failed to convert")
				return nil, err
			}

			// Marshal the JSON-Schema into JSON:
			jsonSchemaJSON, err := json.MarshalIndent(enumJSONSchema, "", "    ")
			if err != nil {
				c.logger.WithError(err).Error("Failed to encode jsonSchema")
				return nil, err
			}

			// Add a response:
			resFile := &plugin.CodeGeneratorResponse_File{
				Name:    proto.String(jsonSchemaFileName),
				Content: proto.String(string(jsonSchemaJSON)),
			}
			response = append(response, resFile)
		}
	} else {
		// Otherwise process MESSAGES (packages):
		pkg, ok := c.relativelyLookupPackage(globalPkg, file.GetPackage())
		if !ok {
			return nil, fmt.Errorf("no such package found: %s", file.GetPackage())
		}
		for _, msg := range file.GetMessageType() {

			// skip if we are only generating schema for specific messages
			if genSpecificMessages && !contains(c.messageTargets, msg.GetName()) {
				continue
			}

			jsonSchemaFileName := c.generateSchemaFilename(file, msg.GetName())
			c.logger.WithField("proto_filename", protoFileName).WithField("msg_name", msg.GetName()).WithField("jsonschema_filename", jsonSchemaFileName).Info("Generating JSON-schema for MESSAGE")

			// Convert the message:
			messageJSONSchema, err := c.convertMessageType(pkg, msg)
			if err != nil {
				c.logger.WithError(err).WithField("proto_filename", protoFileName).Error("Failed to convert")
				return nil, err
			}

			// Marshal the JSON-Schema into JSON:
			jsonSchemaJSON, err := json.MarshalIndent(messageJSONSchema, "", "    ")
			if err != nil {
				c.logger.WithError(err).Error("Failed to encode jsonSchema")
				return nil, err
			}

			// Add a response:
			resFile := &plugin.CodeGeneratorResponse_File{
				Name:    proto.String(jsonSchemaFileName),
				Content: proto.String(string(jsonSchemaJSON)),
			}
			response = append(response, resFile)
		}
	}

	return response, nil
}

func (c *Converter) convert(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	c.parseGeneratorParameters(req.GetParameter())

	generateTargets := make(map[string]bool)
	for _, file := range req.GetFileToGenerate() {
		generateTargets[file] = true
	}

	c.sourceInfo = newSourceCodeInfo(req.GetProtoFile())
	res := &plugin.CodeGeneratorResponse{}
	for _, file := range req.GetProtoFile() {
		if file.GetPackage() == "" {
			c.logger.WithField("filename", file.GetName()).Warn("Proto file doesn't specify a package")
			continue
		}

		for _, msg := range file.GetMessageType() {
			c.logger.WithField("msg_name", msg.GetName()).WithField("package_name", file.GetPackage()).Debug("Loading a message")
			c.registerType(file.Package, msg)
		}

		if _, ok := generateTargets[file.GetName()]; ok {
			c.logger.WithField("filename", file.GetName()).Debug("Converting file")
			converted, err := c.convertFile(file)
			if err != nil {
				res.Error = proto.String(fmt.Sprintf("Failed to convert %s: %v", file.GetName(), err))
				return res, err
			}
			res.File = append(res.File, converted...)
		}
	}
	return res, nil
}

func (c *Converter) generateSchemaFilename(file *descriptor.FileDescriptorProto, protoName string) string {
	if c.PrefixSchemaFilesWithPackage {
		return fmt.Sprintf("%s/%s.jsonschema", file.GetPackage(), protoName)
	}
	return fmt.Sprintf("%s.jsonschema", protoName)
}

func contains(haystack []string, needle string) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}
