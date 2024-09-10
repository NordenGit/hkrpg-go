// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PEFKAOOBIHI.proto

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

type PEFKAOOBIHI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleInfo *GameRogueMiracle `protobuf:"bytes,4,opt,name=miracle_info,json=miracleInfo,proto3" json:"miracle_info,omitempty"`
	OEAPKCNGCKB uint32            `protobuf:"varint,9,opt,name=OEAPKCNGCKB,proto3" json:"OEAPKCNGCKB,omitempty"`
}

func (x *PEFKAOOBIHI) Reset() {
	*x = PEFKAOOBIHI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PEFKAOOBIHI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PEFKAOOBIHI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PEFKAOOBIHI) ProtoMessage() {}

func (x *PEFKAOOBIHI) ProtoReflect() protoreflect.Message {
	mi := &file_PEFKAOOBIHI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PEFKAOOBIHI.ProtoReflect.Descriptor instead.
func (*PEFKAOOBIHI) Descriptor() ([]byte, []int) {
	return file_PEFKAOOBIHI_proto_rawDescGZIP(), []int{0}
}

func (x *PEFKAOOBIHI) GetMiracleInfo() *GameRogueMiracle {
	if x != nil {
		return x.MiracleInfo
	}
	return nil
}

func (x *PEFKAOOBIHI) GetOEAPKCNGCKB() uint32 {
	if x != nil {
		return x.OEAPKCNGCKB
	}
	return 0
}

var File_PEFKAOOBIHI_proto protoreflect.FileDescriptor

var file_PEFKAOOBIHI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x45, 0x46, 0x4b, 0x41, 0x4f, 0x4f, 0x42, 0x49, 0x48, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0b, 0x50,
	0x45, 0x46, 0x4b, 0x41, 0x4f, 0x4f, 0x42, 0x49, 0x48, 0x49, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x45, 0x41, 0x50, 0x4b, 0x43, 0x4e, 0x47, 0x43, 0x4b, 0x42, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x45, 0x41, 0x50, 0x4b, 0x43, 0x4e, 0x47, 0x43,
	0x4b, 0x42, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PEFKAOOBIHI_proto_rawDescOnce sync.Once
	file_PEFKAOOBIHI_proto_rawDescData = file_PEFKAOOBIHI_proto_rawDesc
)

func file_PEFKAOOBIHI_proto_rawDescGZIP() []byte {
	file_PEFKAOOBIHI_proto_rawDescOnce.Do(func() {
		file_PEFKAOOBIHI_proto_rawDescData = protoimpl.X.CompressGZIP(file_PEFKAOOBIHI_proto_rawDescData)
	})
	return file_PEFKAOOBIHI_proto_rawDescData
}

var file_PEFKAOOBIHI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PEFKAOOBIHI_proto_goTypes = []interface{}{
	(*PEFKAOOBIHI)(nil),      // 0: PEFKAOOBIHI
	(*GameRogueMiracle)(nil), // 1: GameRogueMiracle
}
var file_PEFKAOOBIHI_proto_depIdxs = []int32{
	1, // 0: PEFKAOOBIHI.miracle_info:type_name -> GameRogueMiracle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PEFKAOOBIHI_proto_init() }
func file_PEFKAOOBIHI_proto_init() {
	if File_PEFKAOOBIHI_proto != nil {
		return
	}
	file_GameRogueMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PEFKAOOBIHI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PEFKAOOBIHI); i {
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
			RawDescriptor: file_PEFKAOOBIHI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PEFKAOOBIHI_proto_goTypes,
		DependencyIndexes: file_PEFKAOOBIHI_proto_depIdxs,
		MessageInfos:      file_PEFKAOOBIHI_proto_msgTypes,
	}.Build()
	File_PEFKAOOBIHI_proto = out.File
	file_PEFKAOOBIHI_proto_rawDesc = nil
	file_PEFKAOOBIHI_proto_goTypes = nil
	file_PEFKAOOBIHI_proto_depIdxs = nil
}
