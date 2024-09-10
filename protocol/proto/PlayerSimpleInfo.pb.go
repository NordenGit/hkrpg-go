// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerSimpleInfo.proto

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

type PlayerSimpleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatBubbleId         uint32              `protobuf:"varint,5,opt,name=chat_bubble_id,json=chatBubbleId,proto3" json:"chat_bubble_id,omitempty"`
	Signature            string              `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	AssistSimpleInfoList []*AssistSimpleInfo `protobuf:"bytes,1,rep,name=assist_simple_info_list,json=assistSimpleInfoList,proto3" json:"assist_simple_info_list,omitempty"`
	CDFPLBBIKDI          string              `protobuf:"bytes,10,opt,name=CDFPLBBIKDI,proto3" json:"CDFPLBBIKDI,omitempty"`
	Platform             PlatformType        `protobuf:"varint,3,opt,name=platform,proto3,enum=PlatformType" json:"platform,omitempty"`
	LastActiveTime       int64               `protobuf:"varint,11,opt,name=last_active_time,json=lastActiveTime,proto3" json:"last_active_time,omitempty"`
	HeadIcon             uint32              `protobuf:"varint,8,opt,name=head_icon,json=headIcon,proto3" json:"head_icon,omitempty"`
	Uid                  uint32              `protobuf:"varint,15,opt,name=uid,proto3" json:"uid,omitempty"`
	IsBanned             bool                `protobuf:"varint,4,opt,name=is_banned,json=isBanned,proto3" json:"is_banned,omitempty"`
	Nickname             string              `protobuf:"bytes,7,opt,name=nickname,proto3" json:"nickname,omitempty"`
	OnlineStatus         FriendOnlineStatus  `protobuf:"varint,14,opt,name=online_status,json=onlineStatus,proto3,enum=FriendOnlineStatus" json:"online_status,omitempty"`
	Level                uint32              `protobuf:"varint,12,opt,name=level,proto3" json:"level,omitempty"`
	EMGADLCGDBF          string              `protobuf:"bytes,2,opt,name=EMGADLCGDBF,proto3" json:"EMGADLCGDBF,omitempty"`
}

func (x *PlayerSimpleInfo) Reset() {
	*x = PlayerSimpleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerSimpleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerSimpleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerSimpleInfo) ProtoMessage() {}

func (x *PlayerSimpleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerSimpleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerSimpleInfo.ProtoReflect.Descriptor instead.
func (*PlayerSimpleInfo) Descriptor() ([]byte, []int) {
	return file_PlayerSimpleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerSimpleInfo) GetChatBubbleId() uint32 {
	if x != nil {
		return x.ChatBubbleId
	}
	return 0
}

func (x *PlayerSimpleInfo) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PlayerSimpleInfo) GetAssistSimpleInfoList() []*AssistSimpleInfo {
	if x != nil {
		return x.AssistSimpleInfoList
	}
	return nil
}

func (x *PlayerSimpleInfo) GetCDFPLBBIKDI() string {
	if x != nil {
		return x.CDFPLBBIKDI
	}
	return ""
}

func (x *PlayerSimpleInfo) GetPlatform() PlatformType {
	if x != nil {
		return x.Platform
	}
	return PlatformType_EDITOR
}

func (x *PlayerSimpleInfo) GetLastActiveTime() int64 {
	if x != nil {
		return x.LastActiveTime
	}
	return 0
}

func (x *PlayerSimpleInfo) GetHeadIcon() uint32 {
	if x != nil {
		return x.HeadIcon
	}
	return 0
}

func (x *PlayerSimpleInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlayerSimpleInfo) GetIsBanned() bool {
	if x != nil {
		return x.IsBanned
	}
	return false
}

func (x *PlayerSimpleInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PlayerSimpleInfo) GetOnlineStatus() FriendOnlineStatus {
	if x != nil {
		return x.OnlineStatus
	}
	return FriendOnlineStatus_FRIEND_ONLINE_STATUS_OFFLINE
}

func (x *PlayerSimpleInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *PlayerSimpleInfo) GetEMGADLCGDBF() string {
	if x != nil {
		return x.EMGADLCGDBF
	}
	return ""
}

var File_PlayerSimpleInfo_proto protoreflect.FileDescriptor

var file_PlayerSimpleInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1,
	0x03, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x62, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x61,
	0x74, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x48, 0x0a, 0x17, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x44, 0x46, 0x50, 0x4c, 0x42, 0x42, 0x49, 0x4b, 0x44, 0x49,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x44, 0x46, 0x50, 0x4c, 0x42, 0x42, 0x49,
	0x4b, 0x44, 0x49, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64,
	0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x65, 0x61,
	0x64, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x38, 0x0a, 0x0d, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4d, 0x47, 0x41, 0x44, 0x4c, 0x43, 0x47, 0x44, 0x42, 0x46, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x45, 0x4d, 0x47, 0x41, 0x44, 0x4c, 0x43, 0x47, 0x44,
	0x42, 0x46, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerSimpleInfo_proto_rawDescOnce sync.Once
	file_PlayerSimpleInfo_proto_rawDescData = file_PlayerSimpleInfo_proto_rawDesc
)

func file_PlayerSimpleInfo_proto_rawDescGZIP() []byte {
	file_PlayerSimpleInfo_proto_rawDescOnce.Do(func() {
		file_PlayerSimpleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerSimpleInfo_proto_rawDescData)
	})
	return file_PlayerSimpleInfo_proto_rawDescData
}

var file_PlayerSimpleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerSimpleInfo_proto_goTypes = []interface{}{
	(*PlayerSimpleInfo)(nil), // 0: PlayerSimpleInfo
	(*AssistSimpleInfo)(nil), // 1: AssistSimpleInfo
	(PlatformType)(0),        // 2: PlatformType
	(FriendOnlineStatus)(0),  // 3: FriendOnlineStatus
}
var file_PlayerSimpleInfo_proto_depIdxs = []int32{
	1, // 0: PlayerSimpleInfo.assist_simple_info_list:type_name -> AssistSimpleInfo
	2, // 1: PlayerSimpleInfo.platform:type_name -> PlatformType
	3, // 2: PlayerSimpleInfo.online_status:type_name -> FriendOnlineStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PlayerSimpleInfo_proto_init() }
func file_PlayerSimpleInfo_proto_init() {
	if File_PlayerSimpleInfo_proto != nil {
		return
	}
	file_AssistSimpleInfo_proto_init()
	file_PlatformType_proto_init()
	file_FriendOnlineStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerSimpleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerSimpleInfo); i {
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
			RawDescriptor: file_PlayerSimpleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerSimpleInfo_proto_goTypes,
		DependencyIndexes: file_PlayerSimpleInfo_proto_depIdxs,
		MessageInfos:      file_PlayerSimpleInfo_proto_msgTypes,
	}.Build()
	File_PlayerSimpleInfo_proto = out.File
	file_PlayerSimpleInfo_proto_rawDesc = nil
	file_PlayerSimpleInfo_proto_goTypes = nil
	file_PlayerSimpleInfo_proto_depIdxs = nil
}
