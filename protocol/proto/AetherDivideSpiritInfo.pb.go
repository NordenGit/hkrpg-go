// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AetherDivideSpiritInfo.proto

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

type AetherDivideSpiritInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp          uint32              `protobuf:"varint,15,opt,name=exp,proto3" json:"exp,omitempty"`
	Promotion    uint32              `protobuf:"varint,1,opt,name=promotion,proto3" json:"promotion,omitempty"`
	SpBar        *SpBarInfo          `protobuf:"bytes,3,opt,name=sp_bar,json=spBar,proto3" json:"sp_bar,omitempty"`
	AvatarId     uint32              `protobuf:"varint,2,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	JAAEEMLNALF  uint32              `protobuf:"varint,13,opt,name=JAAEEMLNALF,proto3" json:"JAAEEMLNALF,omitempty"`
	PassiveSkill []*PassiveSkillItem `protobuf:"bytes,10,rep,name=passive_skill,json=passiveSkill,proto3" json:"passive_skill,omitempty"`
}

func (x *AetherDivideSpiritInfo) Reset() {
	*x = AetherDivideSpiritInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AetherDivideSpiritInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AetherDivideSpiritInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AetherDivideSpiritInfo) ProtoMessage() {}

func (x *AetherDivideSpiritInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AetherDivideSpiritInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AetherDivideSpiritInfo.ProtoReflect.Descriptor instead.
func (*AetherDivideSpiritInfo) Descriptor() ([]byte, []int) {
	return file_AetherDivideSpiritInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AetherDivideSpiritInfo) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *AetherDivideSpiritInfo) GetPromotion() uint32 {
	if x != nil {
		return x.Promotion
	}
	return 0
}

func (x *AetherDivideSpiritInfo) GetSpBar() *SpBarInfo {
	if x != nil {
		return x.SpBar
	}
	return nil
}

func (x *AetherDivideSpiritInfo) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *AetherDivideSpiritInfo) GetJAAEEMLNALF() uint32 {
	if x != nil {
		return x.JAAEEMLNALF
	}
	return 0
}

func (x *AetherDivideSpiritInfo) GetPassiveSkill() []*PassiveSkillItem {
	if x != nil {
		return x.PassiveSkill
	}
	return nil
}

var File_AetherDivideSpiritInfo_proto protoreflect.FileDescriptor

var file_AetherDivideSpiritInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x70,
	0x69, 0x72, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x50, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x70, 0x42, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x16, 0x41, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x65, 0x78, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x06, 0x73, 0x70, 0x5f, 0x62, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x70, 0x42, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x73, 0x70, 0x42, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x41, 0x41, 0x45, 0x45, 0x4d, 0x4c, 0x4e, 0x41, 0x4c,
	0x46, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x41, 0x41, 0x45, 0x45, 0x4d, 0x4c,
	0x4e, 0x41, 0x4c, 0x46, 0x12, 0x36, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x69, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c,
	0x70, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AetherDivideSpiritInfo_proto_rawDescOnce sync.Once
	file_AetherDivideSpiritInfo_proto_rawDescData = file_AetherDivideSpiritInfo_proto_rawDesc
)

func file_AetherDivideSpiritInfo_proto_rawDescGZIP() []byte {
	file_AetherDivideSpiritInfo_proto_rawDescOnce.Do(func() {
		file_AetherDivideSpiritInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AetherDivideSpiritInfo_proto_rawDescData)
	})
	return file_AetherDivideSpiritInfo_proto_rawDescData
}

var file_AetherDivideSpiritInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AetherDivideSpiritInfo_proto_goTypes = []interface{}{
	(*AetherDivideSpiritInfo)(nil), // 0: AetherDivideSpiritInfo
	(*SpBarInfo)(nil),              // 1: SpBarInfo
	(*PassiveSkillItem)(nil),       // 2: PassiveSkillItem
}
var file_AetherDivideSpiritInfo_proto_depIdxs = []int32{
	1, // 0: AetherDivideSpiritInfo.sp_bar:type_name -> SpBarInfo
	2, // 1: AetherDivideSpiritInfo.passive_skill:type_name -> PassiveSkillItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AetherDivideSpiritInfo_proto_init() }
func file_AetherDivideSpiritInfo_proto_init() {
	if File_AetherDivideSpiritInfo_proto != nil {
		return
	}
	file_PassiveSkillItem_proto_init()
	file_SpBarInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AetherDivideSpiritInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AetherDivideSpiritInfo); i {
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
			RawDescriptor: file_AetherDivideSpiritInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AetherDivideSpiritInfo_proto_goTypes,
		DependencyIndexes: file_AetherDivideSpiritInfo_proto_depIdxs,
		MessageInfos:      file_AetherDivideSpiritInfo_proto_msgTypes,
	}.Build()
	File_AetherDivideSpiritInfo_proto = out.File
	file_AetherDivideSpiritInfo_proto_rawDesc = nil
	file_AetherDivideSpiritInfo_proto_goTypes = nil
	file_AetherDivideSpiritInfo_proto_depIdxs = nil
}
