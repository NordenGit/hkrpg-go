// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeBossEquipmentInfo.proto

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

type ChallengeBossEquipmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid       uint32 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Level     uint32 `protobuf:"varint,8,opt,name=level,proto3" json:"level,omitempty"`
	Rank      uint32 `protobuf:"varint,7,opt,name=rank,proto3" json:"rank,omitempty"`
	UniqueId  uint32 `protobuf:"varint,9,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	Promotion uint32 `protobuf:"varint,4,opt,name=promotion,proto3" json:"promotion,omitempty"`
}

func (x *ChallengeBossEquipmentInfo) Reset() {
	*x = ChallengeBossEquipmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeBossEquipmentInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeBossEquipmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeBossEquipmentInfo) ProtoMessage() {}

func (x *ChallengeBossEquipmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeBossEquipmentInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeBossEquipmentInfo.ProtoReflect.Descriptor instead.
func (*ChallengeBossEquipmentInfo) Descriptor() ([]byte, []int) {
	return file_ChallengeBossEquipmentInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeBossEquipmentInfo) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *ChallengeBossEquipmentInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ChallengeBossEquipmentInfo) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *ChallengeBossEquipmentInfo) GetUniqueId() uint32 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

func (x *ChallengeBossEquipmentInfo) GetPromotion() uint32 {
	if x != nil {
		return x.Promotion
	}
	return 0
}

var File_ChallengeBossEquipmentInfo_proto protoreflect.FileDescriptor

var file_ChallengeBossEquipmentInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x1a, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x42, 0x6f, 0x73, 0x73, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x74, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeBossEquipmentInfo_proto_rawDescOnce sync.Once
	file_ChallengeBossEquipmentInfo_proto_rawDescData = file_ChallengeBossEquipmentInfo_proto_rawDesc
)

func file_ChallengeBossEquipmentInfo_proto_rawDescGZIP() []byte {
	file_ChallengeBossEquipmentInfo_proto_rawDescOnce.Do(func() {
		file_ChallengeBossEquipmentInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeBossEquipmentInfo_proto_rawDescData)
	})
	return file_ChallengeBossEquipmentInfo_proto_rawDescData
}

var file_ChallengeBossEquipmentInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeBossEquipmentInfo_proto_goTypes = []interface{}{
	(*ChallengeBossEquipmentInfo)(nil), // 0: ChallengeBossEquipmentInfo
}
var file_ChallengeBossEquipmentInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChallengeBossEquipmentInfo_proto_init() }
func file_ChallengeBossEquipmentInfo_proto_init() {
	if File_ChallengeBossEquipmentInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChallengeBossEquipmentInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeBossEquipmentInfo); i {
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
			RawDescriptor: file_ChallengeBossEquipmentInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeBossEquipmentInfo_proto_goTypes,
		DependencyIndexes: file_ChallengeBossEquipmentInfo_proto_depIdxs,
		MessageInfos:      file_ChallengeBossEquipmentInfo_proto_msgTypes,
	}.Build()
	File_ChallengeBossEquipmentInfo_proto = out.File
	file_ChallengeBossEquipmentInfo_proto_rawDesc = nil
	file_ChallengeBossEquipmentInfo_proto_goTypes = nil
	file_ChallengeBossEquipmentInfo_proto_depIdxs = nil
}
