// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetChessRogueNousStoryInfoScRsp.proto

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

type GetChessRogueNousStoryInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                 uint32                         `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ChessRogueSubStoryInfo  []*ChessRogueNousSubStoryInfo  `protobuf:"bytes,3,rep,name=chess_rogue_sub_story_info,json=chessRogueSubStoryInfo,proto3" json:"chess_rogue_sub_story_info,omitempty"`
	ChessRogueMainStoryInfo []*ChessRogueNousMainStoryInfo `protobuf:"bytes,11,rep,name=chess_rogue_main_story_info,json=chessRogueMainStoryInfo,proto3" json:"chess_rogue_main_story_info,omitempty"`
}

func (x *GetChessRogueNousStoryInfoScRsp) Reset() {
	*x = GetChessRogueNousStoryInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetChessRogueNousStoryInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChessRogueNousStoryInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChessRogueNousStoryInfoScRsp) ProtoMessage() {}

func (x *GetChessRogueNousStoryInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetChessRogueNousStoryInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChessRogueNousStoryInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetChessRogueNousStoryInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetChessRogueNousStoryInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetChessRogueNousStoryInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetChessRogueNousStoryInfoScRsp) GetChessRogueSubStoryInfo() []*ChessRogueNousSubStoryInfo {
	if x != nil {
		return x.ChessRogueSubStoryInfo
	}
	return nil
}

func (x *GetChessRogueNousStoryInfoScRsp) GetChessRogueMainStoryInfo() []*ChessRogueNousMainStoryInfo {
	if x != nil {
		return x.ChessRogueMainStoryInfo
	}
	return nil
}

var File_GetChessRogueNousStoryInfoScRsp_proto protoreflect.FileDescriptor

var file_GetChessRogueNousStoryInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x25, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e,
	0x6f, 0x75, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73, 0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f,
	0x75, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x1a, 0x63, 0x68,
	0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73, 0x53,
	0x75, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x16, 0x63, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x5a, 0x0a, 0x1b, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x17, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetChessRogueNousStoryInfoScRsp_proto_rawDescOnce sync.Once
	file_GetChessRogueNousStoryInfoScRsp_proto_rawDescData = file_GetChessRogueNousStoryInfoScRsp_proto_rawDesc
)

func file_GetChessRogueNousStoryInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetChessRogueNousStoryInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetChessRogueNousStoryInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetChessRogueNousStoryInfoScRsp_proto_rawDescData)
	})
	return file_GetChessRogueNousStoryInfoScRsp_proto_rawDescData
}

var file_GetChessRogueNousStoryInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetChessRogueNousStoryInfoScRsp_proto_goTypes = []interface{}{
	(*GetChessRogueNousStoryInfoScRsp)(nil), // 0: GetChessRogueNousStoryInfoScRsp
	(*ChessRogueNousSubStoryInfo)(nil),      // 1: ChessRogueNousSubStoryInfo
	(*ChessRogueNousMainStoryInfo)(nil),     // 2: ChessRogueNousMainStoryInfo
}
var file_GetChessRogueNousStoryInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetChessRogueNousStoryInfoScRsp.chess_rogue_sub_story_info:type_name -> ChessRogueNousSubStoryInfo
	2, // 1: GetChessRogueNousStoryInfoScRsp.chess_rogue_main_story_info:type_name -> ChessRogueNousMainStoryInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetChessRogueNousStoryInfoScRsp_proto_init() }
func file_GetChessRogueNousStoryInfoScRsp_proto_init() {
	if File_GetChessRogueNousStoryInfoScRsp_proto != nil {
		return
	}
	file_ChessRogueNousSubStoryInfo_proto_init()
	file_ChessRogueNousMainStoryInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetChessRogueNousStoryInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChessRogueNousStoryInfoScRsp); i {
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
			RawDescriptor: file_GetChessRogueNousStoryInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetChessRogueNousStoryInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetChessRogueNousStoryInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetChessRogueNousStoryInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetChessRogueNousStoryInfoScRsp_proto = out.File
	file_GetChessRogueNousStoryInfoScRsp_proto_rawDesc = nil
	file_GetChessRogueNousStoryInfoScRsp_proto_goTypes = nil
	file_GetChessRogueNousStoryInfoScRsp_proto_depIdxs = nil
}
