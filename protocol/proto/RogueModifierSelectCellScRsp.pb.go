// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueModifierSelectCellScRsp.proto

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

type RogueModifierSelectCellScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellId      uint32    `protobuf:"varint,14,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	Retcode     uint32    `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GBHDPFJNKDO *ItemList `protobuf:"bytes,2,opt,name=GBHDPFJNKDO,proto3" json:"GBHDPFJNKDO,omitempty"`
}

func (x *RogueModifierSelectCellScRsp) Reset() {
	*x = RogueModifierSelectCellScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueModifierSelectCellScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueModifierSelectCellScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueModifierSelectCellScRsp) ProtoMessage() {}

func (x *RogueModifierSelectCellScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_RogueModifierSelectCellScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueModifierSelectCellScRsp.ProtoReflect.Descriptor instead.
func (*RogueModifierSelectCellScRsp) Descriptor() ([]byte, []int) {
	return file_RogueModifierSelectCellScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *RogueModifierSelectCellScRsp) GetCellId() uint32 {
	if x != nil {
		return x.CellId
	}
	return 0
}

func (x *RogueModifierSelectCellScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *RogueModifierSelectCellScRsp) GetGBHDPFJNKDO() *ItemList {
	if x != nil {
		return x.GBHDPFJNKDO
	}
	return nil
}

var File_RogueModifierSelectCellScRsp_proto protoreflect.FileDescriptor

var file_RogueModifierSelectCellScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x42, 0x48, 0x44, 0x50,
	0x46, 0x4a, 0x4e, 0x4b, 0x44, 0x4f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x47, 0x42, 0x48, 0x44, 0x50, 0x46, 0x4a,
	0x4e, 0x4b, 0x44, 0x4f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueModifierSelectCellScRsp_proto_rawDescOnce sync.Once
	file_RogueModifierSelectCellScRsp_proto_rawDescData = file_RogueModifierSelectCellScRsp_proto_rawDesc
)

func file_RogueModifierSelectCellScRsp_proto_rawDescGZIP() []byte {
	file_RogueModifierSelectCellScRsp_proto_rawDescOnce.Do(func() {
		file_RogueModifierSelectCellScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueModifierSelectCellScRsp_proto_rawDescData)
	})
	return file_RogueModifierSelectCellScRsp_proto_rawDescData
}

var file_RogueModifierSelectCellScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueModifierSelectCellScRsp_proto_goTypes = []interface{}{
	(*RogueModifierSelectCellScRsp)(nil), // 0: RogueModifierSelectCellScRsp
	(*ItemList)(nil),                     // 1: ItemList
}
var file_RogueModifierSelectCellScRsp_proto_depIdxs = []int32{
	1, // 0: RogueModifierSelectCellScRsp.GBHDPFJNKDO:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueModifierSelectCellScRsp_proto_init() }
func file_RogueModifierSelectCellScRsp_proto_init() {
	if File_RogueModifierSelectCellScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueModifierSelectCellScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueModifierSelectCellScRsp); i {
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
			RawDescriptor: file_RogueModifierSelectCellScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueModifierSelectCellScRsp_proto_goTypes,
		DependencyIndexes: file_RogueModifierSelectCellScRsp_proto_depIdxs,
		MessageInfos:      file_RogueModifierSelectCellScRsp_proto_msgTypes,
	}.Build()
	File_RogueModifierSelectCellScRsp_proto = out.File
	file_RogueModifierSelectCellScRsp_proto_rawDesc = nil
	file_RogueModifierSelectCellScRsp_proto_goTypes = nil
	file_RogueModifierSelectCellScRsp_proto_depIdxs = nil
}
