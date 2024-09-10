// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetStuffScNotify.proto

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

type GetStuffScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GBKGFOJIEAB uint32       `protobuf:"varint,13,opt,name=GBKGFOJIEAB,proto3" json:"GBKGFOJIEAB,omitempty"`
	LPKNCAKNNKN GetStuffType `protobuf:"varint,11,opt,name=LPKNCAKNNKN,proto3,enum=GetStuffType" json:"LPKNCAKNNKN,omitempty"`
}

func (x *GetStuffScNotify) Reset() {
	*x = GetStuffScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetStuffScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStuffScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStuffScNotify) ProtoMessage() {}

func (x *GetStuffScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GetStuffScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStuffScNotify.ProtoReflect.Descriptor instead.
func (*GetStuffScNotify) Descriptor() ([]byte, []int) {
	return file_GetStuffScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GetStuffScNotify) GetGBKGFOJIEAB() uint32 {
	if x != nil {
		return x.GBKGFOJIEAB
	}
	return 0
}

func (x *GetStuffScNotify) GetLPKNCAKNNKN() GetStuffType {
	if x != nil {
		return x.LPKNCAKNNKN
	}
	return GetStuffType_UNKNOW
}

var File_GetStuffScNotify_proto protoreflect.FileDescriptor

var file_GetStuffScNotify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x47, 0x42, 0x4b, 0x47, 0x46, 0x4f, 0x4a, 0x49, 0x45, 0x41, 0x42, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x42, 0x4b, 0x47, 0x46, 0x4f, 0x4a, 0x49, 0x45,
	0x41, 0x42, 0x12, 0x2f, 0x0a, 0x0b, 0x4c, 0x50, 0x4b, 0x4e, 0x43, 0x41, 0x4b, 0x4e, 0x4e, 0x4b,
	0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4c, 0x50, 0x4b, 0x4e, 0x43, 0x41, 0x4b, 0x4e,
	0x4e, 0x4b, 0x4e, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetStuffScNotify_proto_rawDescOnce sync.Once
	file_GetStuffScNotify_proto_rawDescData = file_GetStuffScNotify_proto_rawDesc
)

func file_GetStuffScNotify_proto_rawDescGZIP() []byte {
	file_GetStuffScNotify_proto_rawDescOnce.Do(func() {
		file_GetStuffScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetStuffScNotify_proto_rawDescData)
	})
	return file_GetStuffScNotify_proto_rawDescData
}

var file_GetStuffScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetStuffScNotify_proto_goTypes = []interface{}{
	(*GetStuffScNotify)(nil), // 0: GetStuffScNotify
	(GetStuffType)(0),        // 1: GetStuffType
}
var file_GetStuffScNotify_proto_depIdxs = []int32{
	1, // 0: GetStuffScNotify.LPKNCAKNNKN:type_name -> GetStuffType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetStuffScNotify_proto_init() }
func file_GetStuffScNotify_proto_init() {
	if File_GetStuffScNotify_proto != nil {
		return
	}
	file_GetStuffType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetStuffScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStuffScNotify); i {
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
			RawDescriptor: file_GetStuffScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetStuffScNotify_proto_goTypes,
		DependencyIndexes: file_GetStuffScNotify_proto_depIdxs,
		MessageInfos:      file_GetStuffScNotify_proto_msgTypes,
	}.Build()
	File_GetStuffScNotify_proto = out.File
	file_GetStuffScNotify_proto_rawDesc = nil
	file_GetStuffScNotify_proto_goTypes = nil
	file_GetStuffScNotify_proto_depIdxs = nil
}
