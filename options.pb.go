// Custom options for protoc-gen-jsonschema
// Allocated range is 1125-1129
// See https://github.com/protocolbuffers/protobuf/blob/master/docs/options.md

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: options.proto

package protoc_gen_jsonschema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Custom VendorExtension definition, e.g., x-go-custom-tag: validate:"required"
type VendorExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of vendor extension key, e.g., x-go-custom-tag
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Value of vendor extension, e.g., validate:"required"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *VendorExtension) Reset() {
	*x = VendorExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VendorExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VendorExtension) ProtoMessage() {}

func (x *VendorExtension) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VendorExtension.ProtoReflect.Descriptor instead.
func (*VendorExtension) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

func (x *VendorExtension) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VendorExtension) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Custom FieldOptions
type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fields tagged with this will be omitted from generated schemas
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Fields tagged with this will be marked as "required" in generated schemas
	Required bool `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	// Fields tagged with this will constrain strings using the "minLength" keyword in generated schemas
	MinLength int32 `protobuf:"varint,3,opt,name=min_length,json=minLength,proto3" json:"min_length,omitempty"`
	// Fields tagged with this will constrain strings using the "maxLength" keyword in generated schemas
	MaxLength int32 `protobuf:"varint,4,opt,name=max_length,json=maxLength,proto3" json:"max_length,omitempty"`
	// Fields tagged with this will constrain strings using the "pattern" keyword in generated schemas
	Pattern string `protobuf:"bytes,5,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// Fields tagged with this will have accociated "vendor extension" object in generated schemas
	VendorExt *VendorExtension `protobuf:"bytes,6,opt,name=vendor_ext,json=vendorExt,proto3" json:"vendor_ext,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

func (x *FieldOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *FieldOptions) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *FieldOptions) GetMinLength() int32 {
	if x != nil {
		return x.MinLength
	}
	return 0
}

func (x *FieldOptions) GetMaxLength() int32 {
	if x != nil {
		return x.MaxLength
	}
	return 0
}

func (x *FieldOptions) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *FieldOptions) GetVendorExt() *VendorExtension {
	if x != nil {
		return x.VendorExt
	}
	return nil
}

// Custom FileOptions
type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Files tagged with this will not be processed
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Override the default file extension for schemas generated from this file
	Extension string `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{2}
}

func (x *FileOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *FileOptions) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

// Custom MessageOptions
type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Messages tagged with this will not be processed
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Messages tagged with this will have all fields marked as "required":
	AllFieldsRequired bool `protobuf:"varint,2,opt,name=all_fields_required,json=allFieldsRequired,proto3" json:"all_fields_required,omitempty"`
	// Messages tagged with this will additionally accept null values for all properties:
	AllowNullValues bool `protobuf:"varint,3,opt,name=allow_null_values,json=allowNullValues,proto3" json:"allow_null_values,omitempty"`
	// Messages tagged with this will have all fields marked as not allowing additional properties:
	DisallowAdditionalProperties bool `protobuf:"varint,4,opt,name=disallow_additional_properties,json=disallowAdditionalProperties,proto3" json:"disallow_additional_properties,omitempty"`
	// Messages tagged with this will have all nested enums encoded to use constants instead of simple types (supports value annotations):
	EnumsAsConstants bool `protobuf:"varint,5,opt,name=enums_as_constants,json=enumsAsConstants,proto3" json:"enums_as_constants,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{3}
}

func (x *MessageOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *MessageOptions) GetAllFieldsRequired() bool {
	if x != nil {
		return x.AllFieldsRequired
	}
	return false
}

func (x *MessageOptions) GetAllowNullValues() bool {
	if x != nil {
		return x.AllowNullValues
	}
	return false
}

func (x *MessageOptions) GetDisallowAdditionalProperties() bool {
	if x != nil {
		return x.DisallowAdditionalProperties
	}
	return false
}

func (x *MessageOptions) GetEnumsAsConstants() bool {
	if x != nil {
		return x.EnumsAsConstants
	}
	return false
}

// Custom EnumOptions
type EnumOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enums tagged with this will have be encoded to use constants instead of simple types (supports value annotations):
	EnumsAsConstants bool `protobuf:"varint,1,opt,name=enums_as_constants,json=enumsAsConstants,proto3" json:"enums_as_constants,omitempty"`
	// Enums tagged with this will only provide string values as options (not their numerical equivalents):
	EnumsAsStringsOnly bool `protobuf:"varint,2,opt,name=enums_as_strings_only,json=enumsAsStringsOnly,proto3" json:"enums_as_strings_only,omitempty"`
	// Enums tagged with this will have enum name prefix removed from values:
	EnumsTrimPrefix bool `protobuf:"varint,3,opt,name=enums_trim_prefix,json=enumsTrimPrefix,proto3" json:"enums_trim_prefix,omitempty"`
	// Enums tagged with this will not be processed
	Ignore bool `protobuf:"varint,4,opt,name=ignore,proto3" json:"ignore,omitempty"`
}

func (x *EnumOptions) Reset() {
	*x = EnumOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumOptions) ProtoMessage() {}

func (x *EnumOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumOptions.ProtoReflect.Descriptor instead.
func (*EnumOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{4}
}

func (x *EnumOptions) GetEnumsAsConstants() bool {
	if x != nil {
		return x.EnumsAsConstants
	}
	return false
}

func (x *EnumOptions) GetEnumsAsStringsOnly() bool {
	if x != nil {
		return x.EnumsAsStringsOnly
	}
	return false
}

func (x *EnumOptions) GetEnumsTrimPrefix() bool {
	if x != nil {
		return x.EnumsTrimPrefix
	}
	return false
}

func (x *EnumOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         1125,
		Name:          "protoc.gen.jsonschema.field_options",
		Tag:           "bytes,1125,opt,name=field_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         1126,
		Name:          "protoc.gen.jsonschema.file_options",
		Tag:           "bytes,1126,opt,name=file_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         1127,
		Name:          "protoc.gen.jsonschema.message_options",
		Tag:           "bytes,1127,opt,name=message_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*EnumOptions)(nil),
		Field:         1128,
		Name:          "protoc.gen.jsonschema.enum_options",
		Tag:           "bytes,1128,opt,name=enum_options",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional protoc.gen.jsonschema.FieldOptions field_options = 1125;
	E_FieldOptions = &file_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional protoc.gen.jsonschema.FileOptions file_options = 1126;
	E_FileOptions = &file_options_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional protoc.gen.jsonschema.MessageOptions message_options = 1127;
	E_MessageOptions = &file_options_proto_extTypes[2]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional protoc.gen.jsonschema.EnumOptions enum_options = 1128;
	E_EnumOptions = &file_options_proto_extTypes[3]
)

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0f, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x5f,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x69,
	0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x12, 0x45, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x45, 0x78, 0x74, 0x22, 0x43, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xf8, 0x01, 0x0a,
	0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x1e, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c, 0x64, 0x69, 0x73,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x41, 0x73, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x5f, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x41, 0x73, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x5f, 0x61,
	0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x41, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x5f, 0x74, 0x72, 0x69, 0x6d, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x54, 0x72, 0x69, 0x6d, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x3a, 0x68, 0x0a, 0x0d,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x64, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe6, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x70, 0x0a, 0x0f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xe7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x64,
	0x0a, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe8, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x72, 0x75, 0x73, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_proto_rawDescOnce sync.Once
	file_options_proto_rawDescData = file_options_proto_rawDesc
)

func file_options_proto_rawDescGZIP() []byte {
	file_options_proto_rawDescOnce.Do(func() {
		file_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_proto_rawDescData)
	})
	return file_options_proto_rawDescData
}

var file_options_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_options_proto_goTypes = []interface{}{
	(*VendorExtension)(nil),             // 0: protoc.gen.jsonschema.VendorExtension
	(*FieldOptions)(nil),                // 1: protoc.gen.jsonschema.FieldOptions
	(*FileOptions)(nil),                 // 2: protoc.gen.jsonschema.FileOptions
	(*MessageOptions)(nil),              // 3: protoc.gen.jsonschema.MessageOptions
	(*EnumOptions)(nil),                 // 4: protoc.gen.jsonschema.EnumOptions
	(*descriptorpb.FieldOptions)(nil),   // 5: google.protobuf.FieldOptions
	(*descriptorpb.FileOptions)(nil),    // 6: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 7: google.protobuf.MessageOptions
	(*descriptorpb.EnumOptions)(nil),    // 8: google.protobuf.EnumOptions
}
var file_options_proto_depIdxs = []int32{
	0, // 0: protoc.gen.jsonschema.FieldOptions.vendor_ext:type_name -> protoc.gen.jsonschema.VendorExtension
	5, // 1: protoc.gen.jsonschema.field_options:extendee -> google.protobuf.FieldOptions
	6, // 2: protoc.gen.jsonschema.file_options:extendee -> google.protobuf.FileOptions
	7, // 3: protoc.gen.jsonschema.message_options:extendee -> google.protobuf.MessageOptions
	8, // 4: protoc.gen.jsonschema.enum_options:extendee -> google.protobuf.EnumOptions
	1, // 5: protoc.gen.jsonschema.field_options:type_name -> protoc.gen.jsonschema.FieldOptions
	2, // 6: protoc.gen.jsonschema.file_options:type_name -> protoc.gen.jsonschema.FileOptions
	3, // 7: protoc.gen.jsonschema.message_options:type_name -> protoc.gen.jsonschema.MessageOptions
	4, // 8: protoc.gen.jsonschema.enum_options:type_name -> protoc.gen.jsonschema.EnumOptions
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	5, // [5:9] is the sub-list for extension type_name
	1, // [1:5] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VendorExtension); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_options_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		MessageInfos:      file_options_proto_msgTypes,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_rawDesc = nil
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}