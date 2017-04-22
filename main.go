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
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	jsonschema "github.com/alecthomas/jsonschema"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	AllowAdditionalProperties = true
)

var (
	globalPkg = &ProtoPackage{
		name:     "",
		parent:   nil,
		children: make(map[string]*ProtoPackage),
		types:    make(map[string]*descriptor.DescriptorProto),
	}
)

// ProtoPackage describes a package of Protobuf, which is an container of message types.
type ProtoPackage struct {
	name     string
	parent   *ProtoPackage
	children map[string]*ProtoPackage
	types    map[string]*descriptor.DescriptorProto
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
		glog.Infof("no such nested message %s in %s", component, desc.GetName())
		return nil, false
	}
	return desc, true
}

func (pkg *ProtoPackage) relativelyLookupType(name string) (*descriptor.DescriptorProto, bool) {
	components := strings.SplitN(name, ".", 2)
	switch len(components) {
	case 0:
		glog.V(1).Info("empty message name")
		return nil, false
	case 1:
		found, ok := pkg.types[components[0]]
		return found, ok
	case 2:
		glog.Infof("looking for %s in %s at %s (%v)", components[1], components[0], pkg.name, pkg)
		if child, ok := pkg.children[components[0]]; ok {
			found, ok := child.relativelyLookupType(components[1])
			return found, ok
		}
		if msg, ok := pkg.types[components[0]]; ok {
			found, ok := relativelyLookupNestedType(msg, components[1])
			return found, ok
		}
		glog.V(1).Infof("no such package nor message %s in %s", components[0], pkg.name)
		return nil, false
	default:
		glog.Fatal("not reached")
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

func convertField(curPkg *ProtoPackage, desc *descriptor.FieldDescriptorProto, msg *descriptor.DescriptorProto) (*jsonschema.Type, error) {
	// Prepare a new jsonschema.Type for our eventual return value:
	jsonSchemaType := &jsonschema.Type{
		Properties: make(map[string]*jsonschema.Type),
	}

	// Switch the types, and pick a JSONSchema equivalent:
	switch desc.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:
		jsonSchemaType.Type = "number"

	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		jsonSchemaType.Type = "integer"

	case descriptor.FieldDescriptorProto_TYPE_STRING,
		descriptor.FieldDescriptorProto_TYPE_BYTES:
		jsonSchemaType.Type = "string"

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: "string"})
		jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: "integer"})

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
		jsonSchemaType.Type = "boolean"

	case descriptor.FieldDescriptorProto_TYPE_GROUP,
		descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		jsonSchemaType.Type = "object"
		if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_OPTIONAL {
			jsonSchemaType.AdditionalProperties = []byte("true")
		}
		if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
			jsonSchemaType.AdditionalProperties = []byte("false")
		}

	default:
		return nil, fmt.Errorf("unrecognized field type: %s", desc.GetType().String())
	}

	// Recurse array of primitive types:
	if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED && jsonSchemaType.Type != "object" {
		jsonSchemaType.Items = &jsonschema.Type{
			Type: jsonSchemaType.Type,
		}
		jsonSchemaType.Type = "array"
		return jsonSchemaType, nil
	}

	// Recurse nested objects / arrays of objects (if necessary):
	if jsonSchemaType.Type == "object" {
		recordType, ok := curPkg.lookupType(desc.GetTypeName())
		if !ok {
			return nil, fmt.Errorf("no such message type named %s", desc.GetTypeName())
		}

		// Recurse:
		recursedJsonSchemaType, err := convertMessageType(curPkg, recordType)
		if err != nil {
			return nil, err
		}

		// The result is stored differently for arrays of objects (they become "items"):
		if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			jsonSchemaType.Items = &recursedJsonSchemaType
			jsonSchemaType.Type = "array"
		} else {
			// Nested objects are more straight-forward:
			jsonSchemaType.Properties = recursedJsonSchemaType.Properties
		}
	}

	return jsonSchemaType, nil
}

func convertMessageType(curPkg *ProtoPackage, msg *descriptor.DescriptorProto) (jsonschema.Type, error) {
	// Prepare a new jsonschema:
	jsonSchemaType := jsonschema.Type{
		Properties: make(map[string]*jsonschema.Type),
		Type:       "object",
		Version:    jsonschema.Version,
	}

	// AllowAdditionalProperties will prevent validation where extra fields are found (outside of the schema):
	if AllowAdditionalProperties {
		jsonSchemaType.AdditionalProperties = []byte("true")
	} else {
		jsonSchemaType.AdditionalProperties = []byte("false")
	}

	if glog.V(4) {
		glog.Info("Converting message: ", proto.MarshalTextString(msg))
	}
	for _, fieldDesc := range msg.GetField() {
		recursedJsonSchemaType, err := convertField(curPkg, fieldDesc, msg)
		if err != nil {
			glog.Errorf("Failed to convert field %s in %s: %v", fieldDesc.GetName(), msg.GetName(), err)
			return jsonSchemaType, err
		}
		jsonSchemaType.Properties[fieldDesc.GetName()] = recursedJsonSchemaType
	}
	return jsonSchemaType, nil
}

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

func convertFile(file *descriptor.FileDescriptorProto) ([]*plugin.CodeGeneratorResponse_File, error) {
	protoFileName := path.Base(file.GetName())
	var jsonSchemaName string

	// Prepare a new jsonschema:
	jsonSchema := jsonschema.Type{
		Properties: make(map[string]*jsonschema.Type),
		Version:    jsonschema.Version,
	}

	response := []*plugin.CodeGeneratorResponse_File{}

	// Generate standalone ENUMs:
	if len(file.GetMessageType()) == 0 {
		for _, enum := range file.GetEnumType() {
			glog.V(2).Info("Generating schema for ENUM type ", enum.GetName())
			enumJsonSchema, err := convertEnumType(enum)
			if err != nil {
				glog.Errorf("Failed to convert %s: %v", protoFileName, err)
				return nil, err
			} else {
				jsonSchema = enumJsonSchema
				jsonSchemaName = enum.GetName()
				break
			}
		}
	} else {
		pkg, ok := globalPkg.relativelyLookupPackage(file.GetPackage())
		if !ok {
			return nil, fmt.Errorf("no such package found: %s", file.GetPackage())
		}
		for _, msg := range file.GetMessageType() {
			glog.V(2).Info("Generating schema for MESSAGE type ", msg.GetName())
			messageJsonSchema, err := convertMessageType(pkg, msg)
			if err != nil {
				glog.Errorf("Failed to convert %s: %v", protoFileName, err)
				return nil, err
			} else {
				jsonSchema = messageJsonSchema
				jsonSchemaName = msg.GetName()
				break
			}
		}
	}

	jsonSchemaJson, err := json.MarshalIndent(jsonSchema, "", "    ")
	if err != nil {
		glog.Error("Failed to encode jsonSchema", err)
		return nil, err
	}

	resFile := &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(fmt.Sprintf("%s.jsonschema", jsonSchemaName)),
		Content: proto.String(string(jsonSchemaJson)),
	}
	response = append(response, resFile)

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
			glog.V(1).Infof("Loading a message type %s from package %s", msg.GetName(), file.GetPackage())
			registerType(file.Package, msg)
		}
	}
	for _, file := range req.GetProtoFile() {
		if _, ok := generateTargets[file.GetName()]; ok {
			glog.V(1).Info("Converting ", file.GetName())
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
	glog.V(1).Info("Reading code generation request")
	input, err := ioutil.ReadAll(rd)
	if err != nil {
		glog.Error("Failed to read request:", err)
		return nil, err
	}
	req := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(input, req)
	if err != nil {
		glog.Error("Can't unmarshal input:", err)
		return nil, err
	}

	glog.V(1).Info("Converting input")
	return convert(req)
}

func main() {
	flag.Parse()
	ok := true
	glog.Info("Processing code generator request")
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

	glog.Info("Serializing code generator response")
	data, err := proto.Marshal(res)
	if err != nil {
		glog.Fatal("Cannot marshal response", err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		glog.Fatal("Failed to write response", err)
	}

	if ok {
		glog.Info("Succeeded to process code generator request")
	} else {
		glog.Info("Failed to process code generator but successfully sent the error to protoc")
		os.Exit(1)
	}
}
