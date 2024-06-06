// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FriendSimpleInfo.proto

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

type FriendSimpleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerInfo  *PlayerSimpleInfo `protobuf:"bytes,7,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	RemarkName  string            `protobuf:"bytes,10,opt,name=remark_name,json=remarkName,proto3" json:"remark_name,omitempty"`
	PlayerState PlayerStateType   `protobuf:"varint,5,opt,name=player_state,json=playerState,proto3,enum=PlayerStateType" json:"player_state,omitempty"`
	CFMIKLHJMLE *ALCOEANIKIL      `protobuf:"bytes,6,opt,name=CFMIKLHJMLE,proto3" json:"CFMIKLHJMLE,omitempty"`
	IsMarked    bool              `protobuf:"varint,9,opt,name=is_marked,json=isMarked,proto3" json:"is_marked,omitempty"`
}

func (x *FriendSimpleInfo) Reset() {
	*x = FriendSimpleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FriendSimpleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendSimpleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendSimpleInfo) ProtoMessage() {}

func (x *FriendSimpleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FriendSimpleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendSimpleInfo.ProtoReflect.Descriptor instead.
func (*FriendSimpleInfo) Descriptor() ([]byte, []int) {
	return file_FriendSimpleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FriendSimpleInfo) GetPlayerInfo() *PlayerSimpleInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *FriendSimpleInfo) GetRemarkName() string {
	if x != nil {
		return x.RemarkName
	}
	return ""
}

func (x *FriendSimpleInfo) GetPlayerState() PlayerStateType {
	if x != nil {
		return x.PlayerState
	}
	return PlayerStateType_PLAYING_STATE_NONE
}

func (x *FriendSimpleInfo) GetCFMIKLHJMLE() *ALCOEANIKIL {
	if x != nil {
		return x.CFMIKLHJMLE
	}
	return nil
}

func (x *FriendSimpleInfo) GetIsMarked() bool {
	if x != nil {
		return x.IsMarked
	}
	return false
}

var File_FriendSimpleInfo_proto protoreflect.FileDescriptor

var file_FriendSimpleInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4c, 0x43, 0x4f, 0x45, 0x41, 0x4e,
	0x49, 0x4b, 0x49, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x10, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x32, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x46, 0x4d,
	0x49, 0x4b, 0x4c, 0x48, 0x4a, 0x4d, 0x4c, 0x45, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x41, 0x4c, 0x43, 0x4f, 0x45, 0x41, 0x4e, 0x49, 0x4b, 0x49, 0x4c, 0x52, 0x0b, 0x43, 0x46,
	0x4d, 0x49, 0x4b, 0x4c, 0x48, 0x4a, 0x4d, 0x4c, 0x45, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FriendSimpleInfo_proto_rawDescOnce sync.Once
	file_FriendSimpleInfo_proto_rawDescData = file_FriendSimpleInfo_proto_rawDesc
)

func file_FriendSimpleInfo_proto_rawDescGZIP() []byte {
	file_FriendSimpleInfo_proto_rawDescOnce.Do(func() {
		file_FriendSimpleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FriendSimpleInfo_proto_rawDescData)
	})
	return file_FriendSimpleInfo_proto_rawDescData
}

var file_FriendSimpleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FriendSimpleInfo_proto_goTypes = []interface{}{
	(*FriendSimpleInfo)(nil), // 0: FriendSimpleInfo
	(*PlayerSimpleInfo)(nil), // 1: PlayerSimpleInfo
	(PlayerStateType)(0),     // 2: PlayerStateType
	(*ALCOEANIKIL)(nil),      // 3: ALCOEANIKIL
}
var file_FriendSimpleInfo_proto_depIdxs = []int32{
	1, // 0: FriendSimpleInfo.player_info:type_name -> PlayerSimpleInfo
	2, // 1: FriendSimpleInfo.player_state:type_name -> PlayerStateType
	3, // 2: FriendSimpleInfo.CFMIKLHJMLE:type_name -> ALCOEANIKIL
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_FriendSimpleInfo_proto_init() }
func file_FriendSimpleInfo_proto_init() {
	if File_FriendSimpleInfo_proto != nil {
		return
	}
	file_PlayerSimpleInfo_proto_init()
	file_PlayerStateType_proto_init()
	file_ALCOEANIKIL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FriendSimpleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendSimpleInfo); i {
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
			RawDescriptor: file_FriendSimpleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FriendSimpleInfo_proto_goTypes,
		DependencyIndexes: file_FriendSimpleInfo_proto_depIdxs,
		MessageInfos:      file_FriendSimpleInfo_proto_msgTypes,
	}.Build()
	File_FriendSimpleInfo_proto = out.File
	file_FriendSimpleInfo_proto_rawDesc = nil
	file_FriendSimpleInfo_proto_goTypes = nil
	file_FriendSimpleInfo_proto_depIdxs = nil
}