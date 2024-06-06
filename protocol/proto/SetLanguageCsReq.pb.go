// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetLanguageCsReq.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetLanguageCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language LanguageType `protobuf:"varint,1,opt,name=language,proto3,enum=LanguageType" json:"language,omitempty"`
}

func (x *SetLanguageCsReq) Reset() {
	*x = SetLanguageCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetLanguageCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLanguageCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLanguageCsReq) ProtoMessage() {}

func (x *SetLanguageCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetLanguageCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLanguageCsReq.ProtoReflect.Descriptor instead.
func (*SetLanguageCsReq) Descriptor() ([]byte, []int) {
	return file_SetLanguageCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SetLanguageCsReq) GetLanguage() LanguageType {
	if x != nil {
		return x.Language
	}
	return LanguageType_LANGUAGE_NONE
}

var File_SetLanguageCsReq_proto protoreflect.FileDescriptor

var file_SetLanguageCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x10,
	0x53, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x29, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetLanguageCsReq_proto_rawDescOnce sync.Once
	file_SetLanguageCsReq_proto_rawDescData = file_SetLanguageCsReq_proto_rawDesc
)

func file_SetLanguageCsReq_proto_rawDescGZIP() []byte {
	file_SetLanguageCsReq_proto_rawDescOnce.Do(func() {
		file_SetLanguageCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetLanguageCsReq_proto_rawDescData)
	})
	return file_SetLanguageCsReq_proto_rawDescData
}

var file_SetLanguageCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetLanguageCsReq_proto_goTypes = []interface{}{
	(*SetLanguageCsReq)(nil), // 0: SetLanguageCsReq
	(LanguageType)(0),        // 1: LanguageType
}
var file_SetLanguageCsReq_proto_depIdxs = []int32{
	1, // 0: SetLanguageCsReq.language:type_name -> LanguageType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetLanguageCsReq_proto_init() }
func file_SetLanguageCsReq_proto_init() {
	if File_SetLanguageCsReq_proto != nil {
		return
	}
	file_LanguageType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetLanguageCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLanguageCsReq); i {
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
			RawDescriptor: file_SetLanguageCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetLanguageCsReq_proto_goTypes,
		DependencyIndexes: file_SetLanguageCsReq_proto_depIdxs,
		MessageInfos:      file_SetLanguageCsReq_proto_msgTypes,
	}.Build()
	File_SetLanguageCsReq_proto = out.File
	file_SetLanguageCsReq_proto_rawDesc = nil
	file_SetLanguageCsReq_proto_goTypes = nil
	file_SetLanguageCsReq_proto_depIdxs = nil
}