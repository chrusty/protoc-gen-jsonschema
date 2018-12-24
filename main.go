// protoc plugin which converts .proto to JSON schema
// It is spawned by protoc and generates JSON-schema files.
// "Heavily influenced" by Google's "protog-gen-bq-schema"
//
// usage:
//  $ bin/protoc --jsonschema_out=path/to/outdir foo.proto
//
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	jsonschema "github.com/alecthomas/jsonschema"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	gojsonschema "github.com/xeipuuv/gojsonschema"
)

const (
	LOG_DEBUG = 0
	LOG_INFO  = 1
	LOG_WARN  = 2
	LOG_ERROR = 3
	LOG_FATAL = 4
	LOG_PANIC = 5
)

const (
	keyMapFieldIndex   = 0
	valueMapFieldIndex = 1
)

var (
	allowNullValues              bool = false
	disallowAdditionalProperties bool = false
	disallowBigIntsAsStrings     bool = false
	debugLogging                 bool = false
	globalPkg                         = &ProtoPackage{
		name:     "",
		parent:   nil,
		children: make(map[string]*ProtoPackage),
		types:    make(map[string]*descriptor.DescriptorProto),
	}
	logLevels = map[LogLevel]string{
		0: "DEBUG",
		1: "INFO",
		2: "WARN",
		3: "ERROR",
		4: "FATAL",
		5: "PANIC",
	}
)

// ProtoPackage describes a package of Protobuf, which is an container of message types.
type ProtoPackage struct {
	name     string
	parent   *ProtoPackage
	children map[string]*ProtoPackage
	types    map[string]*descriptor.DescriptorProto
}

type LogLevel int

func init() {
	flag.BoolVar(&allowNullValues, "allow_null_values", false, "Allow NULL values to be validated")
	flag.BoolVar(&disallowAdditionalProperties, "disallow_additional_properties", false, "Disallow additional properties")
	flag.BoolVar(&disallowBigIntsAsStrings, "disallow_bigints_as_strings", false, "Disallow bigints to be strings (eg scientific notation)")
	flag.BoolVar(&debugLogging, "debug", false, "Log debug messages")
}

func logWithLevel(logLevel LogLevel, logFormat string, logParams ...interface{}) {
	// If we're not doing debug logging then just return:
	if logLevel <= LOG_INFO && !debugLogging {
		return
	}

	// Otherwise log:
	logMessage := fmt.Sprintf(logFormat, logParams...)
	log.Printf(fmt.Sprintf("[%v] %v", logLevels[logLevel], logMessage))
}

func registerType(pkgName *string, msg *descriptor.DescriptorProto) {
	pkg := globalPkg
	if pkgName != nil {
		for _, node := range strings.Split(*pkgName, ".") {
			if pkg == globalPkg && node == "" {
				// Skips leading "."
				continue
			}
			child, ok := pkg.children[node]
			if !ok {
				child = &ProtoPackage{
					name:     pkg.name + "." + node,
					parent:   pkg,
					children: make(map[string]*ProtoPackage),
					types:    make(map[string]*descriptor.DescriptorProto),
				}
				pkg.children[node] = child
			}
			pkg = child
		}
	}
	pkg.types[msg.GetName()] = msg
}

func (pkg *ProtoPackage) lookupType(name string) (*descriptor.DescriptorProto, bool) {
	if strings.HasPrefix(name, ".") {
		return globalPkg.relativelyLookupType(name[1:len(name)])
	}

	for ; pkg != nil; pkg = pkg.parent {
		if desc, ok := pkg.relativelyLookupType(name); ok {
			return desc, ok
		}
	}
	return nil, false
}

func relativelyLookupNestedType(desc *descriptor.DescriptorProto, name string) (*descriptor.DescriptorProto, bool) {
	components := strings.Split(name, ".")
componentLoop:
	for _, component := range components {
		for _, nested := range desc.GetNestedType() {
			if nested.GetName() == component {
				desc = nested
				continue componentLoop
			}
		}
		logWithLevel(LOG_INFO, "no such nested message %s in %s", component, desc.GetName())
		return nil, false
	}
	return desc, true
}

func (pkg *ProtoPackage) relativelyLookupType(name string) (*descriptor.DescriptorProto, bool) {
	components := strings.SplitN(name, ".", 2)
	switch len(components) {
	case 0:
		logWithLevel(LOG_DEBUG, "empty message name")
		return nil, false
	case 1:
		found, ok := pkg.types[components[0]]
		return found, ok
	case 2:
		logWithLevel(LOG_DEBUG, "looking for %s in %s at %s (%v)", components[1], components[0], pkg.name, pkg)
		if child, ok := pkg.children[components[0]]; ok {
			found, ok := child.relativelyLookupType(components[1])
			return found, ok
		}
		if msg, ok := pkg.types[components[0]]; ok {
			found, ok := relativelyLookupNestedType(msg, components[1])
			return found, ok
		}
		logWithLevel(LOG_INFO, "no such package nor message %s in %s", components[0], pkg.name)
		return nil, false
	default:
		logWithLevel(LOG_FATAL, "not reached")
		return nil, false
	}
}

func (pkg *ProtoPackage) relativelyLookupPackage(name string) (*ProtoPackage, bool) {
	components := strings.Split(name, ".")
	for _, c := range components {
		var ok bool
		pkg, ok = pkg.children[c]
		if !ok {
			return nil, false
		}
	}
	return pkg, true
}

// Convert a proto "field" (essentially a type-switch with some recursion):
func convertField(curPkg *ProtoPackage, desc *descriptor.FieldDescriptorProto, msg *descriptor.DescriptorProto) (*jsonschema.Type, error) {

	// Prepare a new jsonschema.Type for our eventual return value:
	jsonSchemaType := &jsonschema.Type{
		Properties: make(map[string]*jsonschema.Type),
	}

	// Switch the types, and pick a JSONSchema equivalent:
	switch desc.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if allowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_NUMBER},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_NUMBER
		}

	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		if allowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_INTEGER},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_INTEGER
		}

	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_INTEGER})
		if !disallowBigIntsAsStrings {
			jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_STRING})
		}
		if allowNullValues {
			jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_NULL})
		}

	case descriptor.FieldDescriptorProto_TYPE_STRING,
		descriptor.FieldDescriptorProto_TYPE_BYTES:
		if allowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_STRING},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
		}

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_STRING})
		jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_INTEGER})
		if allowNullValues {
			jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_NULL})
		}

		// Go through all the enums we have, see if we can match any to this field by name:
		for _, enumDescriptor := range msg.GetEnumType() {

			// Each one has several values:
			for _, enumValue := range enumDescriptor.Value {

				// Figure out the entire name of this field:
				fullFieldName := fmt.Sprintf(".%v.%v", *msg.Name, *enumDescriptor.Name)

				// If we find ENUM values for this field then put them into the JSONSchema list of allowed ENUM values:
				if strings.HasSuffix(desc.GetTypeName(), fullFieldName) {
					jsonSchemaType.Enum = append(jsonSchemaType.Enum, enumValue.Name)
					jsonSchemaType.Enum = append(jsonSchemaType.Enum, enumValue.Number)
				}
			}
		}

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if allowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_BOOLEAN},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_BOOLEAN
		}

	case descriptor.FieldDescriptorProto_TYPE_GROUP,
		descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		jsonSchemaType.Type = gojsonschema.TYPE_OBJECT
		if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_OPTIONAL {
			jsonSchemaType.AdditionalProperties = []byte("true")
		}
		if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
			jsonSchemaType.AdditionalProperties = []byte("false")
		}

	default:
		return nil, fmt.Errorf("unrecognized field type: %s", desc.GetType().String())
	}

	// Parse maps
	if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED && len(msg.NestedType) == 1 && msg.NestedType[0].Options.GetMapEntry() {
		/*
			// For maps fields:
			map<KeyType, ValueType> map_field = 1;
			The parsed descriptor looks like:
			message MapFieldEntry {
				option map_entry = true;
				optional KeyType key = 1;
				optional ValueType value = 2;
			}
			repeated MapFieldEntry map_field = 1;
		*/

		if msg.NestedType[0].Field[keyMapFieldIndex].GetType() != descriptor.FieldDescriptorProto_TYPE_STRING {
			return nil, errors.New("only strings are supported as keys in maps")
		}

		t := convertFieldDescriptorFieldType(msg.NestedType[0].Field[valueMapFieldIndex].GetType())
		jsonSchemaType = &jsonschema.Type{
			AdditionalProperties: json.RawMessage(fmt.Sprintf("{\"type\": \"%v\"}", t)),
			Type:                 jsonSchemaType.Type,
			OneOf:                jsonSchemaType.OneOf,
		}
		if allowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_OBJECT},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_OBJECT
			jsonSchemaType.OneOf = []*jsonschema.Type{}
		}

		return jsonSchemaType, nil
	}

	// Recurse array of primitive types:
	if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED && jsonSchemaType.Type != gojsonschema.TYPE_OBJECT {
		jsonSchemaType.Items = &jsonschema.Type{
			Type:  jsonSchemaType.Type,
			OneOf: jsonSchemaType.OneOf,
		}
		if allowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_ARRAY},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_ARRAY
			jsonSchemaType.OneOf = []*jsonschema.Type{}
		}

		return jsonSchemaType, nil
	}

	// Recurse nested objects / arrays of objects (if necessary):
	if jsonSchemaType.Type == gojsonschema.TYPE_OBJECT {

		recordType, ok := curPkg.lookupType(desc.GetTypeName())
		if !ok {
			return nil, fmt.Errorf("no such message type named %s", desc.GetTypeName())
		}

		// Recurse:
		recursedJSONSchemaType, err := convertMessageType(curPkg, recordType)
		if err != nil {
			return nil, err
		}

		// The result is stored differently for arrays of objects (they become "items"):
		if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			jsonSchemaType.Items = &recursedJSONSchemaType
			jsonSchemaType.Type = gojsonschema.TYPE_ARRAY
		} else {
			// Nested objects are more straight-forward:
			jsonSchemaType.Properties = recursedJSONSchemaType.Properties
		}

		// Optionally allow NULL values:
		if allowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: jsonSchemaType.Type},
			}
			jsonSchemaType.Type = ""
		}
	}

	return jsonSchemaType, nil
}

// Converts a proto "MESSAGE" into a JSON-Schema:
func convertMessageType(curPkg *ProtoPackage, msg *descriptor.DescriptorProto) (jsonschema.Type, error) {

	// Prepare a new jsonschema:
	jsonSchemaType := jsonschema.Type{
		Properties: make(map[string]*jsonschema.Type),
		Version:    jsonschema.Version,
	}

	// Optionally allow NULL values:
	if allowNullValues {
		jsonSchemaType.OneOf = []*jsonschema.Type{
			{Type: gojsonschema.TYPE_NULL},
			{Type: gojsonschema.TYPE_OBJECT},
		}
	} else {
		jsonSchemaType.Type = gojsonschema.TYPE_OBJECT
	}

	// disallowAdditionalProperties will prevent validation where extra fields are found (outside of the schema):
	if disallowAdditionalProperties {
		jsonSchemaType.AdditionalProperties = []byte("false")
	} else {
		jsonSchemaType.AdditionalProperties = []byte("true")
	}

	logWithLevel(LOG_DEBUG, "Converting message: %s", proto.MarshalTextString(msg))
	for _, fieldDesc := range msg.GetField() {
		recursedJSONSchemaType, err := convertField(curPkg, fieldDesc, msg)
		if err != nil {
			logWithLevel(LOG_ERROR, "Failed to convert field %s in %s: %v", fieldDesc.GetName(), msg.GetName(), err)
			return jsonSchemaType, err
		}
		jsonSchemaType.Properties[fieldDesc.GetName()] = recursedJSONSchemaType
	}
	return jsonSchemaType, nil
}

// Converts a proto "ENUM" into a JSON-Schema:
func convertEnumType(enum *descriptor.EnumDescriptorProto) (jsonschema.Type, error) {

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
func convertFile(file *descriptor.FileDescriptorProto) ([]*plugin.CodeGeneratorResponse_File, error) {

	// Input filename:
	protoFileName := path.Base(file.GetName())

	// Prepare a list of responses:
	response := []*plugin.CodeGeneratorResponse_File{}

	// Warn about multiple messages / enums in files:
	if len(file.GetMessageType()) > 1 {
		logWithLevel(LOG_WARN, "protoc-gen-jsonschema will create multiple MESSAGE schemas (%d) from one proto file (%v)", len(file.GetMessageType()), protoFileName)
	}
	if len(file.GetEnumType()) > 1 {
		logWithLevel(LOG_WARN, "protoc-gen-jsonschema will create multiple ENUM schemas (%d) from one proto file (%v)", len(file.GetEnumType()), protoFileName)
	}

	// Generate standalone ENUMs:
	if len(file.GetMessageType()) == 0 {
		for _, enum := range file.GetEnumType() {
			jsonSchemaFileName := fmt.Sprintf("%s.jsonschema", enum.GetName())
			logWithLevel(LOG_INFO, "Generating JSON-schema for stand-alone ENUM (%v) in file [%v] => %v", enum.GetName(), protoFileName, jsonSchemaFileName)
			enumJsonSchema, err := convertEnumType(enum)
			if err != nil {
				logWithLevel(LOG_ERROR, "Failed to convert %s: %v", protoFileName, err)
				return nil, err
			} else {
				// Marshal the JSON-Schema into JSON:
				jsonSchemaJSON, err := json.MarshalIndent(enumJsonSchema, "", "    ")
				if err != nil {
					logWithLevel(LOG_ERROR, "Failed to encode jsonSchema: %v", err)
					return nil, err
				} else {
					// Add a response:
					resFile := &plugin.CodeGeneratorResponse_File{
						Name:    proto.String(jsonSchemaFileName),
						Content: proto.String(string(jsonSchemaJSON)),
					}
					response = append(response, resFile)
				}
			}
		}
	} else {
		// Otherwise process MESSAGES (packages):
		pkg, ok := globalPkg.relativelyLookupPackage(file.GetPackage())
		if !ok {
			return nil, fmt.Errorf("no such package found: %s", file.GetPackage())
		}
		for _, msg := range file.GetMessageType() {
			jsonSchemaFileName := fmt.Sprintf("%s.jsonschema", msg.GetName())
			logWithLevel(LOG_INFO, "Generating JSON-schema for MESSAGE (%v) in file [%v] => %v", msg.GetName(), protoFileName, jsonSchemaFileName)
			messageJSONSchema, err := convertMessageType(pkg, msg)
			if err != nil {
				logWithLevel(LOG_ERROR, "Failed to convert %s: %v", protoFileName, err)
				return nil, err
			} else {
				// Marshal the JSON-Schema into JSON:
				jsonSchemaJSON, err := json.MarshalIndent(messageJSONSchema, "", "    ")
				if err != nil {
					logWithLevel(LOG_ERROR, "Failed to encode jsonSchema: %v", err)
					return nil, err
				} else {
					// Add a response:
					resFile := &plugin.CodeGeneratorResponse_File{
						Name:    proto.String(jsonSchemaFileName),
						Content: proto.String(string(jsonSchemaJSON)),
					}
					response = append(response, resFile)
				}
			}
		}
	}

	return response, nil
}

func convert(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	generateTargets := make(map[string]bool)
	for _, file := range req.GetFileToGenerate() {
		generateTargets[file] = true
	}

	res := &plugin.CodeGeneratorResponse{}
	for _, file := range req.GetProtoFile() {
		for _, msg := range file.GetMessageType() {
			logWithLevel(LOG_DEBUG, "Loading a message type %s from package %s", msg.GetName(), file.GetPackage())
			registerType(file.Package, msg)
		}
	}
	for _, file := range req.GetProtoFile() {
		if _, ok := generateTargets[file.GetName()]; ok {
			logWithLevel(LOG_DEBUG, "Converting file (%v)", file.GetName())
			converted, err := convertFile(file)
			if err != nil {
				res.Error = proto.String(fmt.Sprintf("Failed to convert %s: %v", file.GetName(), err))
				return res, err
			}
			res.File = append(res.File, converted...)
		}
	}
	return res, nil
}

func convertFrom(rd io.Reader) (*plugin.CodeGeneratorResponse, error) {
	logWithLevel(LOG_DEBUG, "Reading code generation request")
	input, err := ioutil.ReadAll(rd)
	if err != nil {
		logWithLevel(LOG_ERROR, "Failed to read request: %v", err)
		return nil, err
	}

	req := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(input, req)
	if err != nil {
		logWithLevel(LOG_ERROR, "Can't unmarshal input: %v", err)
		return nil, err
	}

	commandLineParameter(req.GetParameter())

	logWithLevel(LOG_DEBUG, "Converting input")
	return convert(req)
}

func commandLineParameter(parameters string) {
	for _, parameter := range strings.Split(parameters, ",") {
		switch parameter {
		case "allow_null_values":
			allowNullValues = true
		case "debug":
			debugLogging = true
		case "disallow_additional_properties":
			disallowAdditionalProperties = true
		case "disallow_bigints_as_strings":
			disallowBigIntsAsStrings = true
		}
	}
}

func main() {
	flag.Parse()
	ok := true
	logWithLevel(LOG_DEBUG, "Processing code generator request")
	res, err := convertFrom(os.Stdin)
	if err != nil {
		ok = false
		if res == nil {
			message := fmt.Sprintf("Failed to read input: %v", err)
			res = &plugin.CodeGeneratorResponse{
				Error: &message,
			}
		}
	}

	logWithLevel(LOG_DEBUG, "Serializing code generator response")
	data, err := proto.Marshal(res)
	if err != nil {
		logWithLevel(LOG_FATAL, "Cannot marshal response: %v", err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		logWithLevel(LOG_FATAL, "Failed to write response: %v", err)
	}

	if ok {
		logWithLevel(LOG_DEBUG, "Succeeded to process code generator request")
	} else {
		logWithLevel(LOG_WARN, "Failed to process code generator but successfully sent the error to protoc")
		os.Exit(1)
	}
}

func convertFieldDescriptorFieldType(f descriptor.FieldDescriptorProto_Type) string {
	mapping := map[descriptor.FieldDescriptorProto_Type]string{
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:  "number",
		descriptor.FieldDescriptorProto_TYPE_FLOAT:   "number",
		descriptor.FieldDescriptorProto_TYPE_INT64:   "integer",
		descriptor.FieldDescriptorProto_TYPE_UINT64:  "integer",
		descriptor.FieldDescriptorProto_TYPE_INT32:   "integer",
		descriptor.FieldDescriptorProto_TYPE_FIXED64: "integer",
		descriptor.FieldDescriptorProto_TYPE_FIXED32: "integer",
		descriptor.FieldDescriptorProto_TYPE_BOOL:    "boolean",
		descriptor.FieldDescriptorProto_TYPE_STRING:  "string",
		descriptor.FieldDescriptorProto_TYPE_BYTES:   "string",
		descriptor.FieldDescriptorProto_TYPE_UINT32:  "integer",
		descriptor.FieldDescriptorProto_TYPE_SINT32:  "integer",
		descriptor.FieldDescriptorProto_TYPE_SINT64:  "integer",
	}

	v, ok := mapping[f]
	if !ok {
		return ""
	}

	return v
}
