package converter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alecthomas/jsonschema"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/iancoleman/orderedmap"
	"github.com/xeipuuv/gojsonschema"
)

var (
	globalPkg = newProtoPackage(nil, "")

	wellKnownTypes = map[string]bool{
		"DoubleValue": true,
		"FloatValue":  true,
		"Int64Value":  true,
		"UInt64Value": true,
		"Int32Value":  true,
		"UInt32Value": true,
		"BoolValue":   true,
		"StringValue": true,
		"BytesValue":  true,
		"Value":       true,
		"Duration":    true,
	}
)

func (c *Converter) registerEnum(pkgName *string, enum *descriptor.EnumDescriptorProto) {
	pkg := globalPkg
	if pkgName != nil {
		for _, node := range strings.Split(*pkgName, ".") {
			if pkg == globalPkg && node == "" {
				// Skips leading "."
				continue
			}
			child, ok := pkg.children[node]
			if !ok {
				child = newProtoPackage(pkg, node)
				pkg.children[node] = child
			}
			pkg = child
		}
	}
	pkg.enums[enum.GetName()] = enum
}

func (c *Converter) registerType(pkgName *string, msg *descriptor.DescriptorProto) {
	pkg := globalPkg
	if pkgName != nil {
		for _, node := range strings.Split(*pkgName, ".") {
			if pkg == globalPkg && node == "" {
				// Skips leading "."
				continue
			}
			child, ok := pkg.children[node]
			if !ok {
				child = newProtoPackage(pkg, node)
				pkg.children[node] = child
			}
			pkg = child
		}
	}
	pkg.types[msg.GetName()] = msg
}

// Convert a proto "field" (essentially a type-switch with some recursion):
func (c *Converter) convertField(curPkg *ProtoPackage, desc *descriptor.FieldDescriptorProto, msg *descriptor.DescriptorProto, duplicatedMessages map[*descriptor.DescriptorProto]string) (*jsonschema.Type, error) {
	// Prepare a new jsonschema.Type for our eventual return value:
	jsonSchemaType := &jsonschema.Type{}

	// Generate a description from src comments (if available)
	if src := c.sourceInfo.GetField(desc); src != nil {
		jsonSchemaType.Description = formatDescription(src)
	}

	// Switch the types, and pick a JSONSchema equivalent:
	switch desc.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if c.Flags.AllowNullValues {
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
		if c.Flags.AllowNullValues {
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
		if c.Flags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_STRING},
				{Type: gojsonschema.TYPE_NULL},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
		}

		if c.Flags.DisallowBigIntsAsStrings {
			jsonSchemaType.Type = gojsonschema.TYPE_INTEGER
		}

	case descriptor.FieldDescriptorProto_TYPE_STRING:
		if c.Flags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_STRING},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
		}

	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		if c.Flags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{
					Type:           gojsonschema.TYPE_STRING,
					Format:         "binary",
					BinaryEncoding: "base64",
				},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
			jsonSchemaType.Format = "binary"
			jsonSchemaType.BinaryEncoding = "base64"
		}

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_STRING})
		jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_INTEGER})
		if c.Flags.AllowNullValues {
			jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_NULL})
		}

		// Go through all the enums we have, see if we can match any to this field.
		fullEnumIdentifier := strings.TrimPrefix(desc.GetTypeName(), ".")
		matchedEnum, _, ok := c.lookupEnum(curPkg, fullEnumIdentifier)
		if !ok {
			return nil, fmt.Errorf("unable to resolve enum type: %s", desc.GetType().String())
		}

		// We have found an enum, append its values.
		for _, value := range matchedEnum.Value {
			jsonSchemaType.Enum = append(jsonSchemaType.Enum, value.Name)
			jsonSchemaType.Enum = append(jsonSchemaType.Enum, value.Number)
		}

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if c.Flags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_BOOLEAN},
			}
		} else {
			jsonSchemaType.Type = gojsonschema.TYPE_BOOLEAN
		}

	case descriptor.FieldDescriptorProto_TYPE_GROUP, descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		switch desc.GetTypeName() {
		// Make sure that durations match a particular string pattern (eg 3.4s):
		case ".google.protobuf.Duration":
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
			jsonSchemaType.Format = "regex"
			jsonSchemaType.Pattern = `^([0-9]+\.?[0-9]*|\.[0-9]+)s$`
		case ".google.protobuf.Timestamp":
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
			jsonSchemaType.Format = "date-time"
		default:
			jsonSchemaType.Type = gojsonschema.TYPE_OBJECT
			if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_OPTIONAL {
				jsonSchemaType.AdditionalProperties = []byte("true")
			}
			if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
				jsonSchemaType.AdditionalProperties = []byte("false")
			}
			if c.Flags.DisallowAdditionalProperties {
				jsonSchemaType.AdditionalProperties = []byte("false")
			}
		}

	default:
		return nil, fmt.Errorf("unrecognized field type: %s", desc.GetType().String())
	}

	// Recurse array of primitive types:
	if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED && jsonSchemaType.Type != gojsonschema.TYPE_OBJECT {
		jsonSchemaType.Items = &jsonschema.Type{}

		if len(jsonSchemaType.Enum) > 0 {
			jsonSchemaType.Items.Enum = jsonSchemaType.Enum
			jsonSchemaType.Enum = nil
			jsonSchemaType.Items.OneOf = nil
		} else {
			jsonSchemaType.Items.Type = jsonSchemaType.Type
			jsonSchemaType.Items.OneOf = jsonSchemaType.OneOf
		}

		if c.Flags.AllowNullValues {
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

		recordType, pkgName, ok := c.lookupType(curPkg, desc.GetTypeName())
		if !ok {
			return nil, fmt.Errorf("no such message type named %s", desc.GetTypeName())
		}

		// Recurse the recordType:
		recursedJSONSchemaType, err := c.recursiveConvertMessageType(curPkg, recordType, pkgName, duplicatedMessages, false)
		if err != nil {
			return nil, err
		}

		// Maps, arrays, and objects are structured in different ways:
		switch {

		// Maps:
		case recordType.Options.GetMapEntry():
			c.logger.
				WithField("field_name", recordType.GetName()).
				WithField("msg_name", *msg.Name).
				Tracef("Is a map")

			// Make sure we have a "value":
			value, valuePresent := recursedJSONSchemaType.Properties.Get("value")
			if !valuePresent {
				return nil, fmt.Errorf("Unable to find 'value' property of MAP type")
			}

			// Marshal the "value" properties to JSON (because that's how we can pass on AdditionalProperties):
			additionalPropertiesJSON, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			jsonSchemaType.AdditionalProperties = additionalPropertiesJSON

		// Arrays:
		case desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED:
			jsonSchemaType.Items = recursedJSONSchemaType
			jsonSchemaType.Type = gojsonschema.TYPE_ARRAY

			// Build up the list of required fields:
			if c.Flags.AllFieldsRequired && recursedJSONSchemaType.Properties != nil {
				for _, property := range recursedJSONSchemaType.Properties.Keys() {
					jsonSchemaType.Items.Required = append(jsonSchemaType.Items.Required, property)
				}
			}

		// Not maps, not arrays:
		default:

			// If we've got optional types then just take those:
			if recursedJSONSchemaType.OneOf != nil {
				return recursedJSONSchemaType, nil
			}

			// If we're not an object then set the type from whatever we recursed:
			if recursedJSONSchemaType.Type != gojsonschema.TYPE_OBJECT {
				jsonSchemaType.Type = recursedJSONSchemaType.Type
			}

			// Assume the attrbutes of the recursed value:
			jsonSchemaType.Properties = recursedJSONSchemaType.Properties
			jsonSchemaType.Ref = recursedJSONSchemaType.Ref
			jsonSchemaType.Required = recursedJSONSchemaType.Required

			// Build up the list of required fields:
			if c.Flags.AllFieldsRequired && recursedJSONSchemaType.Properties != nil {
				for _, property := range recursedJSONSchemaType.Properties.Keys() {
					jsonSchemaType.Required = append(jsonSchemaType.Required, property)
				}
			}
		}

		// Optionally allow NULL values:
		if c.Flags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_NULL},
				{Type: jsonSchemaType.Type},
			}
			jsonSchemaType.Type = ""
		}
	}

	jsonSchemaType.Required = dedupe(jsonSchemaType.Required)

	return jsonSchemaType, nil
}

// Converts a proto "MESSAGE" into a JSON-Schema:
func (c *Converter) convertMessageType(curPkg *ProtoPackage, msg *descriptor.DescriptorProto) (*jsonschema.Schema, error) {

	// first, recursively find messages that appear more than once - in particular, that will break cycles
	duplicatedMessages, err := c.findDuplicatedNestedMessages(curPkg, msg)
	if err != nil {
		return nil, err
	}

	// main schema for the message
	rootType, err := c.recursiveConvertMessageType(curPkg, msg, "", duplicatedMessages, false)
	if err != nil {
		return nil, err
	}

	// and then generate the sub-schema for each duplicated message
	definitions := jsonschema.Definitions{}
	for refMsg, name := range duplicatedMessages {
		refType, err := c.recursiveConvertMessageType(curPkg, refMsg, "", duplicatedMessages, true)
		if err != nil {
			return nil, err
		}

		// need to give that schema an ID
		if refType.Extras == nil {
			refType.Extras = make(map[string]interface{})
		}
		refType.Extras["id"] = name
		definitions[name] = refType
	}

	newJSONSchema := &jsonschema.Schema{
		Type:        rootType,
		Definitions: definitions,
	}

	// Look for required fields (either by proto2 required flag, or the AllFieldsRequired option):
	for _, fieldDesc := range msg.GetField() {
		if c.Flags.AllFieldsRequired || fieldDesc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
			newJSONSchema.Required = append(newJSONSchema.Required, fieldDesc.GetName())
		}
	}

	newJSONSchema.Required = dedupe(newJSONSchema.Required)

	return newJSONSchema, nil
}

// findDuplicatedNestedMessages takes a message, and returns a map mapping pointers to messages that appear more than once
// (typically because they're part of a reference cycle) to the sub-schema name that we give them.
func (c *Converter) findDuplicatedNestedMessages(curPkg *ProtoPackage, msg *descriptor.DescriptorProto) (map[*descriptor.DescriptorProto]string, error) {
	all := make(map[*descriptor.DescriptorProto]*nameAndCounter)
	if err := c.recursiveFindDuplicatedNestedMessages(curPkg, msg, msg.GetName(), all); err != nil {
		return nil, err
	}

	result := make(map[*descriptor.DescriptorProto]string)
	for m, nameAndCounter := range all {
		if nameAndCounter.counter > 1 && !strings.HasPrefix(nameAndCounter.name, ".google.protobuf.") {
			result[m] = strings.TrimLeft(nameAndCounter.name, ".")
		}
	}

	return result, nil
}

type nameAndCounter struct {
	name    string
	counter int
}

func (c *Converter) recursiveFindDuplicatedNestedMessages(curPkg *ProtoPackage, msg *descriptor.DescriptorProto, typeName string, alreadySeen map[*descriptor.DescriptorProto]*nameAndCounter) error {
	if nameAndCounter, present := alreadySeen[msg]; present {
		nameAndCounter.counter++
		return nil
	}
	alreadySeen[msg] = &nameAndCounter{
		name:    typeName,
		counter: 1,
	}

	for _, desc := range msg.GetField() {
		descType := desc.GetType()
		if descType != descriptor.FieldDescriptorProto_TYPE_MESSAGE && descType != descriptor.FieldDescriptorProto_TYPE_GROUP {
			// no nested messages
			continue
		}

		typeName := desc.GetTypeName()
		recordType, _, ok := c.lookupType(curPkg, typeName)
		if !ok {
			return fmt.Errorf("no such message type named %s", typeName)
		}
		if err := c.recursiveFindDuplicatedNestedMessages(curPkg, recordType, typeName, alreadySeen); err != nil {
			return err
		}
	}

	return nil
}

func (c *Converter) recursiveConvertMessageType(curPkg *ProtoPackage, msg *descriptor.DescriptorProto, pkgName string, duplicatedMessages map[*descriptor.DescriptorProto]string, ignoreDuplicatedMessages bool) (*jsonschema.Type, error) {

	// Prepare a new jsonschema:
	jsonSchemaType := new(jsonschema.Type)

	// Generate a description from src comments (if available)
	if src := c.sourceInfo.GetMessage(msg); src != nil {
		jsonSchemaType.Description = formatDescription(src)
	}

	// Handle google's well-known types:
	if msg.Name != nil && wellKnownTypes[*msg.Name] && pkgName == ".google.protobuf" {
		switch *msg.Name {
		case "DoubleValue", "FloatValue":
			jsonSchemaType.Type = gojsonschema.TYPE_NUMBER
		case "Int32Value", "UInt32Value", "Int64Value", "UInt64Value":
			jsonSchemaType.Type = gojsonschema.TYPE_INTEGER
		case "BoolValue":
			jsonSchemaType.Type = gojsonschema.TYPE_BOOLEAN
		case "BytesValue", "StringValue":
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
		case "Value":
			jsonSchemaType.OneOf = []*jsonschema.Type{
				{Type: gojsonschema.TYPE_ARRAY},
				{Type: gojsonschema.TYPE_BOOLEAN},
				{Type: gojsonschema.TYPE_NUMBER},
				{Type: gojsonschema.TYPE_OBJECT},
				{Type: gojsonschema.TYPE_STRING},
			}
		case "Duration":
			jsonSchemaType.Type = gojsonschema.TYPE_STRING
		}

		// If we're allowing nulls then prepare a OneOf:
		if c.Flags.AllowNullValues {
			jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Type: gojsonschema.TYPE_NULL}, &jsonschema.Type{Type: jsonSchemaType.Type})
			return jsonSchemaType, nil
		}

		// Otherwise just return this simple type:
		return jsonSchemaType, nil
	}

	// Set defaults:
	jsonSchemaType.Properties = orderedmap.New()
	jsonSchemaType.Version = jsonschema.Version

	if refName, ok := duplicatedMessages[msg]; ok && !ignoreDuplicatedMessages {
		return &jsonschema.Type{
			Version: jsonschema.Version,
			Ref:     refName,
		}, nil
	}

	// Optionally allow NULL values:
	if c.Flags.AllowNullValues {
		jsonSchemaType.OneOf = []*jsonschema.Type{
			{Type: gojsonschema.TYPE_NULL},
			{Type: gojsonschema.TYPE_OBJECT},
		}
	} else {
		jsonSchemaType.Type = gojsonschema.TYPE_OBJECT
	}

	// disallowAdditionalProperties will prevent validation where extra fields are found (outside of the schema):
	if c.Flags.DisallowAdditionalProperties {
		jsonSchemaType.AdditionalProperties = []byte("false")
	} else {
		jsonSchemaType.AdditionalProperties = []byte("true")
	}

	c.logger.WithField("message_str", proto.MarshalTextString(msg)).Trace("Converting message")
	for _, fieldDesc := range msg.GetField() {
		recursedJSONSchemaType, err := c.convertField(curPkg, fieldDesc, msg, duplicatedMessages)
		if err != nil {
			c.logger.WithError(err).WithField("field_name", fieldDesc.GetName()).WithField("message_name", msg.GetName()).Error("Failed to convert field")
			return nil, err
		}
		c.logger.WithField("field_name", fieldDesc.GetName()).WithField("type", recursedJSONSchemaType.Type).Trace("Converted field")

		// If this field is part of a OneOf declaration then build that here:
		if c.Flags.EnforceOneOf && fieldDesc.OneofIndex != nil {
			jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Type{Required: []string{fieldDesc.GetName()}})
		}

		// Figure out which field names we want to use:
		switch {
		case c.Flags.UseJSONFieldnamesOnly:
			jsonSchemaType.Properties.Set(fieldDesc.GetJsonName(), recursedJSONSchemaType)
		case c.Flags.UseProtoAndJSONFieldNames:
			jsonSchemaType.Properties.Set(fieldDesc.GetName(), recursedJSONSchemaType)
			jsonSchemaType.Properties.Set(fieldDesc.GetJsonName(), recursedJSONSchemaType)
		default:
			jsonSchemaType.Properties.Set(fieldDesc.GetName(), recursedJSONSchemaType)
		}

		// Look for required fields (either by proto2 required flag, or the AllFieldsRequired option):
		if fieldDesc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
			jsonSchemaType.Required = append(jsonSchemaType.Required, fieldDesc.GetName())
		}
	}

	// Remove empty properties to keep the final output as clean as possible:
	if len(jsonSchemaType.Properties.Keys()) == 0 {
		jsonSchemaType.Properties = nil
	}

	return jsonSchemaType, nil
}

func formatDescription(sl *descriptor.SourceCodeInfo_Location) string {
	var lines []string
	for _, str := range sl.GetLeadingDetachedComments() {
		if s := strings.TrimSpace(str); s != "" {
			lines = append(lines, s)
		}
	}
	if s := strings.TrimSpace(sl.GetLeadingComments()); s != "" {
		lines = append(lines, s)
	}
	if s := strings.TrimSpace(sl.GetTrailingComments()); s != "" {
		lines = append(lines, s)
	}

	return strings.Join(lines, "\n\n")
}

func dedupe(inputStrings []string) []string {
	appended := make(map[string]bool)
	outputStrings := []string{}

	for _, inputString := range inputStrings {
		if !appended[inputString] {
			outputStrings = append(outputStrings, inputString)
			appended[inputString] = true
		}
	}
	return outputStrings
}
