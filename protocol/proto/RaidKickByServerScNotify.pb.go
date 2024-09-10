// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RaidKickByServerScNotify.proto

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

type RaidKickByServerScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason     EHHJOCAAABA `protobuf:"varint,1,opt,name=reason,proto3,enum=EHHJOCAAABA" json:"reason,omitempty"`
	Lineup     *LineupInfo `protobuf:"bytes,5,opt,name=lineup,proto3" json:"lineup,omitempty"`
	WorldLevel uint32      `protobuf:"varint,12,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	RaidId     uint32      `protobuf:"varint,6,opt,name=raid_id,json=raidId,proto3" json:"raid_id,omitempty"`
	Scene      *SceneInfo  `protobuf:"bytes,11,opt,name=scene,proto3" json:"scene,omitempty"`
}

func (x *RaidKickByServerScNotify) Reset() {
	*x = RaidKickByServerScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RaidKickByServerScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaidKickByServerScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaidKickByServerScNotify) ProtoMessage() {}

func (x *RaidKickByServerScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RaidKickByServerScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaidKickByServerScNotify.ProtoReflect.Descriptor instead.
func (*RaidKickByServerScNotify) Descriptor() ([]byte, []int) {
	return file_RaidKickByServerScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RaidKickByServerScNotify) GetReason() EHHJOCAAABA {
	if x != nil {
		return x.Reason
	}
	return EHHJOCAAABA_RAID_KICK_REASON_NONE
}

func (x *RaidKickByServerScNotify) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *RaidKickByServerScNotify) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *RaidKickByServerScNotify) GetRaidId() uint32 {
	if x != nil {
		return x.RaidId
	}
	return 0
}

func (x *RaidKickByServerScNotify) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

var File_RaidKickByServerScNotify_proto protoreflect.FileDescriptor

var file_RaidKickByServerScNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x52, 0x61, 0x69, 0x64, 0x4b, 0x69, 0x63, 0x6b, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x48, 0x48, 0x4a, 0x4f, 0x43, 0x41, 0x41, 0x41, 0x42, 0x41,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x18, 0x52, 0x61, 0x69, 0x64, 0x4b,
	0x69, 0x63, 0x6b, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x63, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x48, 0x48, 0x4a, 0x4f, 0x43, 0x41, 0x41, 0x41, 0x42,
	0x41, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x06, 0x6c, 0x69, 0x6e,
	0x65, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x72, 0x61, 0x69, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_RaidKickByServerScNotify_proto_rawDescOnce sync.Once
	file_RaidKickByServerScNotify_proto_rawDescData = file_RaidKickByServerScNotify_proto_rawDesc
)

func file_RaidKickByServerScNotify_proto_rawDescGZIP() []byte {
	file_RaidKickByServerScNotify_proto_rawDescOnce.Do(func() {
		file_RaidKickByServerScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RaidKickByServerScNotify_proto_rawDescData)
	})
	return file_RaidKickByServerScNotify_proto_rawDescData
}

var file_RaidKickByServerScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RaidKickByServerScNotify_proto_goTypes = []interface{}{
	(*RaidKickByServerScNotify)(nil), // 0: RaidKickByServerScNotify
	(EHHJOCAAABA)(0),                 // 1: EHHJOCAAABA
	(*LineupInfo)(nil),               // 2: LineupInfo
	(*SceneInfo)(nil),                // 3: SceneInfo
}
var file_RaidKickByServerScNotify_proto_depIdxs = []int32{
	1, // 0: RaidKickByServerScNotify.reason:type_name -> EHHJOCAAABA
	2, // 1: RaidKickByServerScNotify.lineup:type_name -> LineupInfo
	3, // 2: RaidKickByServerScNotify.scene:type_name -> SceneInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RaidKickByServerScNotify_proto_init() }
func file_RaidKickByServerScNotify_proto_init() {
	if File_RaidKickByServerScNotify_proto != nil {
		return
	}
	file_SceneInfo_proto_init()
	file_LineupInfo_proto_init()
	file_EHHJOCAAABA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RaidKickByServerScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaidKickByServerScNotify); i {
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
			RawDescriptor: file_RaidKickByServerScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RaidKickByServerScNotify_proto_goTypes,
		DependencyIndexes: file_RaidKickByServerScNotify_proto_depIdxs,
		MessageInfos:      file_RaidKickByServerScNotify_proto_msgTypes,
	}.Build()
	File_RaidKickByServerScNotify_proto = out.File
	file_RaidKickByServerScNotify_proto_rawDesc = nil
	file_RaidKickByServerScNotify_proto_goTypes = nil
	file_RaidKickByServerScNotify_proto_depIdxs = nil
}
