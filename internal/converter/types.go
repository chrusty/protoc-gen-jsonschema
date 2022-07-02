package converter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iancoleman/orderedmap"
	"github.com/invopop/jsonschema"
	"github.com/xeipuuv/gojsonschema"
	"google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"

	"github.com/chrusty/protoc-gen-jsonschema/internal/protos"
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
		"Struct":      true,
	}
)

func (c *Converter) registerEnum(pkgName string, enum *descriptor.EnumDescriptorProto) {
	pkg := globalPkg
	if pkgName != "" {
		for _, node := range strings.Split(pkgName, ".") {
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

func (c *Converter) registerType(pkgName string, msgDesc *descriptor.DescriptorProto) {
	pkg := globalPkg
	if pkgName != "" {
		for _, node := range strings.Split(pkgName, ".") {
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
	pkg.types[msgDesc.GetName()] = msgDesc
}

// Convert a proto "field" (essentially a type-switch with some recursion):
func (c *Converter) convertField(curPkg *ProtoPackage, desc *descriptor.FieldDescriptorProto, msgDesc *descriptor.DescriptorProto, duplicatedMessages map[*descriptor.DescriptorProto]string, messageFlags ConverterFlags) (*jsonschema.Schema, error) {

	// Prepare a new jsonschema.Type for our eventual return value:
	jsonSchema := &jsonschema.Schema{}

	// Generate a description from src comments (if available)
	if src := c.sourceInfo.GetField(desc); src != nil {
		jsonSchema.Title, jsonSchema.Description = c.formatTitleAndDescription(nil, src)
	}

	// Switch the types, and pick a JSONSchema equivalent:
	switch desc.GetType() {

	// Float32:
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_NUMBER},
			}
		} else {
			jsonSchema.Type = gojsonschema.TYPE_NUMBER
		}

	// Int32:
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_INTEGER},
			}
		} else {
			jsonSchema.Type = gojsonschema.TYPE_INTEGER
		}

	// Int64:
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:

		// As integer:
		if c.Flags.DisallowBigIntsAsStrings {
			if messageFlags.AllowNullValues {
				jsonSchema.OneOf = []*jsonschema.Schema{
					{Type: gojsonschema.TYPE_INTEGER},
					{Type: gojsonschema.TYPE_NULL},
				}
			} else {
				jsonSchema.Type = gojsonschema.TYPE_INTEGER
			}
		}

		// As string:
		if !c.Flags.DisallowBigIntsAsStrings {
			if messageFlags.AllowNullValues {
				jsonSchema.OneOf = []*jsonschema.Schema{
					{Type: gojsonschema.TYPE_STRING},
					{Type: gojsonschema.TYPE_NULL},
				}
			} else {
				jsonSchema.Type = gojsonschema.TYPE_STRING
			}
		}

	// String:
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		stringDef := &jsonschema.Schema{Type: gojsonschema.TYPE_STRING}
		// Check for custom options
		opts := desc.GetOptions()
		if opts != nil && proto.HasExtension(opts, protos.E_FieldOptions) {
			if opt := proto.GetExtension(opts, protos.E_FieldOptions); opt != nil {
				if fieldOptions, ok := opt.(*protos.FieldOptions); ok {
					if fieldOptions.GetMinLength() > 0 {
						stringDef.MinLength = int(fieldOptions.GetMinLength())
					}
					if fieldOptions.GetMaxLength() > 0 {
						stringDef.MaxLength = int(fieldOptions.GetMaxLength())
					}
					if fieldOptions.GetPattern() != "" {
						stringDef.Pattern = fieldOptions.GetPattern()
					}
				}
			}
		}

		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_NULL},
				stringDef,
			}
		} else {
			jsonSchema.Type = stringDef.Type
			if stringDef.MinLength != 0 {
				jsonSchema.MinLength = stringDef.MinLength
			}
			if stringDef.MaxLength != 0 {
				jsonSchema.MaxLength = stringDef.MaxLength
			}
			if stringDef.Pattern != "" {
				jsonSchema.Pattern = stringDef.Pattern
			}
		}

	// Bytes:
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_NULL},
				{
					Type:           gojsonschema.TYPE_STRING,
					Format:         "binary",
					BinaryEncoding: "base64",
				},
			}
		} else {
			jsonSchema.Type = gojsonschema.TYPE_STRING
			jsonSchema.Format = "binary"
			jsonSchema.BinaryEncoding = "base64"
		}

	// ENUM:
	case descriptor.FieldDescriptorProto_TYPE_ENUM:

		// Go through all the enums we have, see if we can match any to this field.
		fullEnumIdentifier := strings.TrimPrefix(desc.GetTypeName(), ".")
		matchedEnum, _, ok := c.lookupEnum(curPkg, fullEnumIdentifier)
		if !ok {
			return nil, fmt.Errorf("unable to resolve enum type: %s", desc.GetType().String())
		}

		// We already have a converter for standalone ENUMs, so just use that:
		enumSchema, err := c.convertEnumType(matchedEnum, messageFlags)
		if err != nil {
			return nil, err
		}

		jsonSchema = &enumSchema

	// Bool:
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_BOOLEAN},
			}
		} else {
			jsonSchema.Type = gojsonschema.TYPE_BOOLEAN
		}

	// Group (object):
	case descriptor.FieldDescriptorProto_TYPE_GROUP, descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		switch desc.GetTypeName() {
		// Make sure that durations match a particular string pattern (eg 3.4s):
		case ".google.protobuf.Duration":
			jsonSchema.Type = gojsonschema.TYPE_STRING
			jsonSchema.Format = "regex"
			jsonSchema.Pattern = `^([0-9]+\.?[0-9]*|\.[0-9]+)s$`
		case ".google.protobuf.Timestamp":
			jsonSchema.Type = gojsonschema.TYPE_STRING
			jsonSchema.Format = "date-time"
		default:
			jsonSchema.Type = gojsonschema.TYPE_OBJECT
			if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_OPTIONAL {
				jsonSchema.AdditionalProperties = jsonschema.TrueSchema
			}
			if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
				jsonSchema.AdditionalProperties = jsonschema.FalseSchema
			}
			if messageFlags.DisallowAdditionalProperties {
				jsonSchema.AdditionalProperties = jsonschema.FalseSchema
			}
		}

	default:
		return nil, fmt.Errorf("unrecognized field type: %s", desc.GetType().String())
	}

	// Recurse array of primitive types:
	if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED && jsonSchema.Type != gojsonschema.TYPE_OBJECT {
		jsonSchema.Items = &jsonschema.Type{}

		if len(jsonSchema.Enum) > 0 {
			jsonSchema.Items.Enum = jsonSchema.Enum
			jsonSchema.Enum = nil
			jsonSchema.Items.OneOf = nil
		} else {
			jsonSchema.Items.Type = jsonSchema.Type
			jsonSchema.Items.OneOf = jsonSchema.OneOf
		}

		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_NULL},
				{Type: gojsonschema.TYPE_ARRAY},
			}
		} else {
			jsonSchema.Type = gojsonschema.TYPE_ARRAY
			jsonSchema.OneOf = []*jsonschema.Schema{}
		}
		return jsonSchema, nil
	}

	// Recurse nested objects / arrays of objects (if necessary):
	if jsonSchema.Type == gojsonschema.TYPE_OBJECT {

		recordType, pkgName, ok := c.lookupType(curPkg, desc.GetTypeName())
		if !ok {
			return nil, fmt.Errorf("no such message type named %s", desc.GetTypeName())
		}

		// Recurse the recordType:
		recursedjsonSchema, err := c.recursiveConvertMessageType(curPkg, recordType, pkgName, duplicatedMessages, false)
		if err != nil {
			return nil, err
		}

		// Maps, arrays, and objects are structured in different ways:
		switch {

		// Maps:
		case recordType.Options.GetMapEntry():
			c.logger.
				WithField("field_name", recordType.GetName()).
				WithField("msgDesc_name", *msgDesc.Name).
				Tracef("Is a map")

			if recursedjsonSchema.Properties == nil {
				return nil, fmt.Errorf("Unable to find properties of MAP type")
			}

			// Make sure we have a "value":
			value, valuePresent := recursedjsonSchema.Properties.Get("value")
			if !valuePresent {
				return nil, fmt.Errorf("Unable to find 'value' property of MAP type")
			}

			// Marshal the "value" properties to JSON (because that's how we can pass on AdditionalProperties):
			additionalPropertiesJSON, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			jsonSchema.AdditionalProperties = additionalPropertiesJSON

		// Arrays:
		case desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED:
			jsonSchema.Items = recursedjsonSchema
			jsonSchema.Type = gojsonschema.TYPE_ARRAY

			// Build up the list of required fields:
			if messageFlags.AllFieldsRequired && len(recursedjsonSchema.OneOf) == 0 && recursedjsonSchema.Properties != nil {
				for _, property := range recursedjsonSchema.Properties.Keys() {
					jsonSchema.Items.Required = append(jsonSchema.Items.Required, property)
				}
			}
			jsonSchema.Items.Required = dedupe(jsonSchema.Items.Required)

		// Not maps, not arrays:
		default:

			// If we've got optional types then just take those:
			if recursedjsonSchema.OneOf != nil {
				return recursedjsonSchema, nil
			}

			// If we're not an object then set the type from whatever we recursed:
			if recursedjsonSchema.Type != gojsonschema.TYPE_OBJECT {
				jsonSchema.Type = recursedjsonSchema.Type
			}

			// Assume the attrbutes of the recursed value:
			jsonSchema.Properties = recursedjsonSchema.Properties
			jsonSchema.Ref = recursedjsonSchema.Ref
			jsonSchema.Required = recursedjsonSchema.Required

			// Build up the list of required fields:
			if messageFlags.AllFieldsRequired && len(recursedjsonSchema.OneOf) == 0 && recursedjsonSchema.Properties != nil {
				for _, property := range recursedjsonSchema.Properties.Keys() {
					jsonSchema.Required = append(jsonSchema.Required, property)
				}
			}
		}

		// Optionally allow NULL values:
		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_NULL},
				{Type: jsonSchema.Type},
			}
			jsonSchema.Type = ""
		}
	}

	jsonSchema.Required = dedupe(jsonSchema.Required)

	return jsonSchema, nil
}

// Converts a proto "MESSAGE" into a JSON-Schema:
func (c *Converter) convertMessageType(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto) (*jsonschema.Schema, error) {

	// Get a list of any nested messages in our schema:
	duplicatedMessages, err := c.findNestedMessages(curPkg, msgDesc)
	if err != nil {
		return nil, err
	}

	// Build up a list of JSONSchema type definitions for every message:
	definitions := jsonschema.Definitions{}
	for refmsgDesc, name := range duplicatedMessages {
		refType, err := c.recursiveConvertMessageType(curPkg, refmsgDesc, "", duplicatedMessages, true)
		if err != nil {
			return nil, err
		}

		// Add the schema to our definitions:
		definitions[name] = refType
	}

	// Put together a JSON schema with our discovered definitions, and a $ref for the root type:
	newJSONSchema := &jsonschema.Schema{
		Type: &jsonschema.Schema{
			Ref:     fmt.Sprintf("%s%s", c.refPrefix, msgDesc.GetName()),
			Version: c.schemaVersion,
		},
		Definitions: definitions,
	}

	return newJSONSchema, nil
}

// findNestedMessages takes a message, and returns a map mapping pointers to messages nested within it:
// these messages become definitions which can be referenced (instead of repeating them every time they're used)
func (c *Converter) findNestedMessages(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto) (map[*descriptor.DescriptorProto]string, error) {

	// Get a list of all nested messages, and how often they occur:
	nestedMessages := make(map[*descriptor.DescriptorProto]string)
	if err := c.recursiveFindNestedMessages(curPkg, msgDesc, msgDesc.GetName(), nestedMessages); err != nil {
		return nil, err
	}

	// Now filter them:
	result := make(map[*descriptor.DescriptorProto]string)
	for message, messageName := range nestedMessages {
		if !message.GetOptions().GetMapEntry() && !strings.HasPrefix(messageName, ".google.protobuf.") {
			result[message] = strings.TrimLeft(messageName, ".")
		}
	}

	return result, nil
}

func (c *Converter) recursiveFindNestedMessages(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto, typeName string, nestedMessages map[*descriptor.DescriptorProto]string) error {
	if _, present := nestedMessages[msgDesc]; present {
		return nil
	}
	nestedMessages[msgDesc] = typeName

	for _, desc := range msgDesc.GetField() {
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
		if err := c.recursiveFindNestedMessages(curPkg, recordType, typeName, nestedMessages); err != nil {
			return err
		}
	}

	return nil
}

func (c *Converter) recursiveConvertMessageType(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto, pkgName string, duplicatedMessages map[*descriptor.DescriptorProto]string, ignoreDuplicatedMessages bool) (*jsonschema.Schema, error) {

	// Prepare a new jsonschema:
	jsonSchema := new(jsonschema.Schema)

	// Set some per-message flags from config and options:
	messageFlags := c.Flags
	if opts := msgDesc.GetOptions(); opts != nil && proto.HasExtension(opts, protos.E_MessageOptions) {
		if opt := proto.GetExtension(opts, protos.E_MessageOptions); opt != nil {
			if messageOptions, ok := opt.(*protos.MessageOptions); ok {

				// AllFieldsRequired:
				if messageOptions.GetAllFieldsRequired() {
					messageFlags.AllFieldsRequired = true
				}

				// AllowNullValues:
				if messageOptions.GetAllowNullValues() {
					messageFlags.AllowNullValues = true
				}

				// DisallowAdditionalProperties:
				if messageOptions.GetDisallowAdditionalProperties() {
					messageFlags.DisallowAdditionalProperties = true
				}

				// ENUMs as constants:
				if messageOptions.GetEnumsAsConstants() {
					messageFlags.EnumsAsConstants = true
				}
			}
		}
	}

	// Generate a description from src comments (if available)
	if src := c.sourceInfo.GetMessage(msgDesc); src != nil {
		jsonSchema.Title, jsonSchema.Description = c.formatTitleAndDescription(strPtr(msgDesc.GetName()), src)
	}

	// Handle google's well-known types:
	if msgDesc.Name != nil && wellKnownTypes[*msgDesc.Name] && pkgName == ".google.protobuf" {
		switch *msgDesc.Name {
		case "DoubleValue", "FloatValue":
			jsonSchema.Type = gojsonschema.TYPE_NUMBER
		case "Int32Value", "UInt32Value", "Int64Value", "UInt64Value":
			jsonSchema.Type = gojsonschema.TYPE_INTEGER
		case "BoolValue":
			jsonSchema.Type = gojsonschema.TYPE_BOOLEAN
		case "BytesValue", "StringValue":
			jsonSchema.Type = gojsonschema.TYPE_STRING
		case "Value":
			jsonSchema.OneOf = []*jsonschema.Schema{
				{Type: gojsonschema.TYPE_ARRAY},
				{Type: gojsonschema.TYPE_BOOLEAN},
				{Type: gojsonschema.TYPE_NUMBER},
				{Type: gojsonschema.TYPE_OBJECT},
				{Type: gojsonschema.TYPE_STRING},
			}
		case "Duration":
			jsonSchema.Type = gojsonschema.TYPE_STRING
		case "Struct":
			jsonSchema.Type = gojsonschema.TYPE_OBJECT
		}

		// If we're allowing nulls then prepare a OneOf:
		if messageFlags.AllowNullValues {
			jsonSchema.OneOf = append(jsonSchema.OneOf, &jsonschema.Schema{Type: gojsonschema.TYPE_NULL}, &jsonschema.Schema{Type: jsonSchema.Type})
			return jsonSchema, nil
		}

		// Otherwise just return this simple type:
		return jsonSchema, nil
	}

	// Set defaults:
	jsonSchema.Properties = orderedmap.New()

	// Look up references:
	if refName, ok := duplicatedMessages[msgDesc]; ok && !ignoreDuplicatedMessages {
		return &jsonschema.Schema{
			Ref: fmt.Sprintf("%s%s", c.refPrefix, refName),
		}, nil
	}

	// Optionally allow NULL values:
	if messageFlags.AllowNullValues {
		jsonSchema.OneOf = []*jsonschema.Schema{
			{Type: gojsonschema.TYPE_NULL},
			{Type: gojsonschema.TYPE_OBJECT},
		}
	} else {
		jsonSchema.Type = gojsonschema.TYPE_OBJECT
	}

	// disallowAdditionalProperties will prevent validation where extra fields are found (outside of the schema):
	if messageFlags.DisallowAdditionalProperties {
		jsonSchema.AdditionalProperties = jsonschema.FalseSchema
	} else {
		jsonSchema.AdditionalProperties = jsonschema.TrueSchema
	}

	c.logger.WithField("message_str", msgDesc.String()).Trace("Converting message")
	for _, fieldDesc := range msgDesc.GetField() {

		// Check for our custom field options:
		opts := fieldDesc.GetOptions()
		if opts != nil && proto.HasExtension(opts, protos.E_FieldOptions) {
			if opt := proto.GetExtension(opts, protos.E_FieldOptions); opt != nil {
				if fieldOptions, ok := opt.(*protos.FieldOptions); ok {

					// "Ignored" fields are simply skipped:
					if fieldOptions.GetIgnore() {
						c.logger.WithField("field_name", fieldDesc.GetName()).WithField("message_name", msgDesc.GetName()).Debug("Skipping ignored field")
						continue
					}

					// "Required" fields are added to the list of required attributes in our schema:
					if fieldOptions.GetRequired() {
						c.logger.WithField("field_name", fieldDesc.GetName()).WithField("message_name", msgDesc.GetName()).Debug("Marking required field")
						if c.Flags.UseJSONFieldnamesOnly {
							jsonSchema.Required = append(jsonSchema.Required, fieldDesc.GetJsonName())
						} else {
							jsonSchema.Required = append(jsonSchema.Required, fieldDesc.GetName())
						}
					}
				}
			}
		}

		// Convert the field into a JSONSchema type:
		recursedjsonSchema, err := c.convertField(curPkg, fieldDesc, msgDesc, duplicatedMessages, messageFlags)
		if err != nil {
			c.logger.WithError(err).WithField("field_name", fieldDesc.GetName()).WithField("message_name", msgDesc.GetName()).Error("Failed to convert field")
			return nil, err
		}
		c.logger.WithField("field_name", fieldDesc.GetName()).WithField("type", recursedjsonSchema.Type).Trace("Converted field")

		// If this field is part of a OneOf declaration then build that here:
		if c.Flags.EnforceOneOf && fieldDesc.OneofIndex != nil {
			jsonSchema.OneOf = append(jsonSchema.OneOf, &jsonschema.Schema{Required: []string{fieldDesc.GetName()}})
		}

		// Figure out which field names we want to use:
		switch {
		case c.Flags.UseJSONFieldnamesOnly:
			jsonSchema.Properties.Set(fieldDesc.GetJsonName(), recursedjsonSchema)
		case c.Flags.UseProtoAndJSONFieldNames:
			jsonSchema.Properties.Set(fieldDesc.GetName(), recursedjsonSchema)
			jsonSchema.Properties.Set(fieldDesc.GetJsonName(), recursedjsonSchema)
		default:
			jsonSchema.Properties.Set(fieldDesc.GetName(), recursedjsonSchema)
		}

		// Enforce all_fields_required:
		if messageFlags.AllFieldsRequired && len(jsonSchema.OneOf) == 0 && jsonSchema.Properties != nil {
			for _, property := range jsonSchema.Properties.Keys() {
				jsonSchema.Required = append(jsonSchema.Required, property)
			}
		}

		// Look for required fields by the proto2 "required" flag:
		if fieldDesc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED && fieldDesc.OneofIndex == nil {
			if c.Flags.UseJSONFieldnamesOnly {
				jsonSchema.Required = append(jsonSchema.Required, fieldDesc.GetJsonName())
			} else {
				jsonSchema.Required = append(jsonSchema.Required, fieldDesc.GetName())
			}
		}
	}

	// Remove empty properties to keep the final output as clean as possible:
	if len(jsonSchema.Properties.Keys()) == 0 {
		jsonSchema.Properties = nil
	}

	// Dedupe required fields:
	jsonSchema.Required = dedupe(jsonSchema.Required)

	return jsonSchema, nil
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
