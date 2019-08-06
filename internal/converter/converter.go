package converter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/alecthomas/jsonschema"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/sirupsen/logrus"
)

// Converter is everything you need to convert protos to JSONSchemas:
type Converter struct {
	AllowNullValues              bool
	DisallowAdditionalProperties bool
	DisallowBigIntsAsStrings     bool
	UseProtoAndJSONFieldnames    bool
	logger                       *logrus.Logger
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

	c.parseGeneratorParameters(req.GetParameter())

	c.logger.Debug("Converting input")
	return c.convert(req)
	// return c.debugger(req)
}

func (c *Converter) parseGeneratorParameters(parameters string) {
	for _, parameter := range strings.Split(parameters, ",") {
		switch parameter {
		case "allow_null_values":
			c.AllowNullValues = true
		case "debug":
			c.logger.SetLevel(logrus.DebugLevel)
		case "disallow_additional_properties":
			c.DisallowAdditionalProperties = true
		case "disallow_bigints_as_strings":
			c.DisallowBigIntsAsStrings = true
		case "proto_and_json_fieldnames":
			c.UseProtoAndJSONFieldnames = true
		}
	}
}

// Converts a proto "ENUM" into a JSON-Schema:
func (c *Converter) convertEnumType(enum *descriptor.EnumDescriptorProto) (jsonschema.Type, error) {

	// Prepare a new jsonschema.Type for our eventual return value:
	jsonSchemaType := jsonschema.Type{
		Version: jsonschema.Version,
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
	response := []*plugin.CodeGeneratorResponse_File{}

	// Warn about multiple messages / enums in files:
	if len(file.GetMessageType()) > 1 {
		c.logger.WithField("schemas", len(file.GetMessageType())).WithField("proto_filename", protoFileName).Warn("protoc-gen-jsonschema will create multiple MESSAGE schemas from one proto file")
	}
	if len(file.GetEnumType()) > 1 {
		c.logger.WithField("schemas", len(file.GetMessageType())).WithField("proto_filename", protoFileName).Warn("protoc-gen-jsonschema will create multiple ENUM schemas from one proto file")
	}

	// Generate standalone ENUMs:
	if len(file.GetMessageType()) == 0 {
		for _, enum := range file.GetEnumType() {
			jsonSchemaFileName := fmt.Sprintf("%s.jsonschema", enum.GetName())
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
			jsonSchemaFileName := fmt.Sprintf("%s.jsonschema", msg.GetName())
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
	generateTargets := make(map[string]bool)
	for _, file := range req.GetFileToGenerate() {
		generateTargets[file] = true
	}

	res := &plugin.CodeGeneratorResponse{}
	for _, file := range req.GetProtoFile() {
		for _, msg := range file.GetMessageType() {
			c.logger.WithField("msg_name", msg.GetName()).WithField("package_name", file.GetPackage()).Debug("Loading a message")
			c.registerType(file.Package, msg)
		}
	}
	for _, file := range req.GetProtoFile() {
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
