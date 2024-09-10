// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueAeonInfo.proto

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

type ChessRogueAeonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AeonIdList    []uint32                 `protobuf:"varint,5,rep,packed,name=aeon_id_list,json=aeonIdList,proto3" json:"aeon_id_list,omitempty"`
	GameAeonId    uint32                   `protobuf:"varint,6,opt,name=game_aeon_id,json=gameAeonId,proto3" json:"game_aeon_id,omitempty"`
	ChessAeonInfo *ChessRogueQueryAeonInfo `protobuf:"bytes,10,opt,name=chess_aeon_info,json=chessAeonInfo,proto3" json:"chess_aeon_info,omitempty"`
	CCIIMNLADKK   int32                    `protobuf:"varint,3,opt,name=CCIIMNLADKK,proto3" json:"CCIIMNLADKK,omitempty"`
	PLANMPECLAE   *BAKPIDLEIFI             `protobuf:"bytes,2,opt,name=PLANMPECLAE,proto3" json:"PLANMPECLAE,omitempty"`
}

func (x *ChessRogueAeonInfo) Reset() {
	*x = ChessRogueAeonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueAeonInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueAeonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueAeonInfo) ProtoMessage() {}

func (x *ChessRogueAeonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueAeonInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueAeonInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueAeonInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueAeonInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueAeonInfo) GetAeonIdList() []uint32 {
	if x != nil {
		return x.AeonIdList
	}
	return nil
}

func (x *ChessRogueAeonInfo) GetGameAeonId() uint32 {
	if x != nil {
		return x.GameAeonId
	}
	return 0
}

func (x *ChessRogueAeonInfo) GetChessAeonInfo() *ChessRogueQueryAeonInfo {
	if x != nil {
		return x.ChessAeonInfo
	}
	return nil
}

func (x *ChessRogueAeonInfo) GetCCIIMNLADKK() int32 {
	if x != nil {
		return x.CCIIMNLADKK
	}
	return 0
}

func (x *ChessRogueAeonInfo) GetPLANMPECLAE() *BAKPIDLEIFI {
	if x != nil {
		return x.PLANMPECLAE
	}
	return nil
}

var File_ChessRogueAeonInfo_proto protoreflect.FileDescriptor

var file_ChessRogueAeonInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x65, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x41, 0x4b, 0x50, 0x49,
	0x44, 0x4c, 0x45, 0x49, 0x46, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a,
	0x12, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x65, 0x6f, 0x6e, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x65,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x61, 0x6d,
	0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x73, 0x73,
	0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x73,
	0x73, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x43, 0x49,
	0x49, 0x4d, 0x4e, 0x4c, 0x41, 0x44, 0x4b, 0x4b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x43, 0x43, 0x49, 0x49, 0x4d, 0x4e, 0x4c, 0x41, 0x44, 0x4b, 0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x50,
	0x4c, 0x41, 0x4e, 0x4d, 0x50, 0x45, 0x43, 0x4c, 0x41, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x42, 0x41, 0x4b, 0x50, 0x49, 0x44, 0x4c, 0x45, 0x49, 0x46, 0x49, 0x52, 0x0b,
	0x50, 0x4c, 0x41, 0x4e, 0x4d, 0x50, 0x45, 0x43, 0x4c, 0x41, 0x45, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueAeonInfo_proto_rawDescOnce sync.Once
	file_ChessRogueAeonInfo_proto_rawDescData = file_ChessRogueAeonInfo_proto_rawDesc
)

func file_ChessRogueAeonInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueAeonInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueAeonInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueAeonInfo_proto_rawDescData)
	})
	return file_ChessRogueAeonInfo_proto_rawDescData
}

var file_ChessRogueAeonInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueAeonInfo_proto_goTypes = []interface{}{
	(*ChessRogueAeonInfo)(nil),      // 0: ChessRogueAeonInfo
	(*ChessRogueQueryAeonInfo)(nil), // 1: ChessRogueQueryAeonInfo
	(*BAKPIDLEIFI)(nil),             // 2: BAKPIDLEIFI
}
var file_ChessRogueAeonInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueAeonInfo.chess_aeon_info:type_name -> ChessRogueQueryAeonInfo
	2, // 1: ChessRogueAeonInfo.PLANMPECLAE:type_name -> BAKPIDLEIFI
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChessRogueAeonInfo_proto_init() }
func file_ChessRogueAeonInfo_proto_init() {
	if File_ChessRogueAeonInfo_proto != nil {
		return
	}
	file_ChessRogueQueryAeonInfo_proto_init()
	file_BAKPIDLEIFI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueAeonInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueAeonInfo); i {
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
			RawDescriptor: file_ChessRogueAeonInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueAeonInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueAeonInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueAeonInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueAeonInfo_proto = out.File
	file_ChessRogueAeonInfo_proto_rawDesc = nil
	file_ChessRogueAeonInfo_proto_goTypes = nil
	file_ChessRogueAeonInfo_proto_depIdxs = nil
}
