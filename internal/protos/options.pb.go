// Custom options for protoc-gen-jsonschema
// Allocated range is 1125-1129
// See https://github.com/protocolbuffers/protobuf/blob/master/docs/options.md

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: options.proto

package protos

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

// Custom FieldOptions
type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fields tagged with this will be omitted from generated schemas
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Fields tagged with this will be marked as "required" in generated schemas
	Required bool `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
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

// Custom FileOptions
type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Files tagged with this will not be processed
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Override the default file extention for schemas generated from this file
	Extention string `protobuf:"bytes,2,opt,name=extention,proto3" json:"extention,omitempty"`
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

func (x *FileOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *FileOptions) GetExtention() string {
	if x != nil {
		return x.Extention
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
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{2}
}

func (x *MessageOptions) GetIgnore() bool {
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

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x43, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x28, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x3a, 0x68, 0x0a, 0x0d, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x64, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xe6, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x70, 0x0a, 0x0f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xe7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_options_proto_goTypes = []interface{}{
	(*FieldOptions)(nil),                // 0: protoc.gen.jsonschema.FieldOptions
	(*FileOptions)(nil),                 // 1: protoc.gen.jsonschema.FileOptions
	(*MessageOptions)(nil),              // 2: protoc.gen.jsonschema.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
	(*descriptorpb.FileOptions)(nil),    // 4: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 5: google.protobuf.MessageOptions
}
var file_options_proto_depIdxs = []int32{
	3, // 0: protoc.gen.jsonschema.field_options:extendee -> google.protobuf.FieldOptions
	4, // 1: protoc.gen.jsonschema.file_options:extendee -> google.protobuf.FileOptions
	5, // 2: protoc.gen.jsonschema.message_options:extendee -> google.protobuf.MessageOptions
	0, // 3: protoc.gen.jsonschema.field_options:type_name -> protoc.gen.jsonschema.FieldOptions
	1, // 4: protoc.gen.jsonschema.file_options:type_name -> protoc.gen.jsonschema.FileOptions
	2, // 5: protoc.gen.jsonschema.message_options:type_name -> protoc.gen.jsonschema.MessageOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 3,
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
