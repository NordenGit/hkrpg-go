// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneSummonUnitInfo.proto

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

type SceneSummonUnitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CasterEntityId  uint32   `protobuf:"varint,7,opt,name=caster_entity_id,json=casterEntityId,proto3" json:"caster_entity_id,omitempty"`
	SummonUnitId    uint32   `protobuf:"varint,4,opt,name=summon_unit_id,json=summonUnitId,proto3" json:"summon_unit_id,omitempty"`
	TriggerNameList []string `protobuf:"bytes,6,rep,name=trigger_name_list,json=triggerNameList,proto3" json:"trigger_name_list,omitempty"`
	LifeTimeMs      int32    `protobuf:"varint,5,opt,name=life_time_ms,json=lifeTimeMs,proto3" json:"life_time_ms,omitempty"`
	CreateTimeMs    uint64   `protobuf:"varint,11,opt,name=create_time_ms,json=createTimeMs,proto3" json:"create_time_ms,omitempty"`
	AttachEntityId  uint32   `protobuf:"varint,9,opt,name=attach_entity_id,json=attachEntityId,proto3" json:"attach_entity_id,omitempty"`
}

func (x *SceneSummonUnitInfo) Reset() {
	*x = SceneSummonUnitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneSummonUnitInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneSummonUnitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneSummonUnitInfo) ProtoMessage() {}

func (x *SceneSummonUnitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneSummonUnitInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneSummonUnitInfo.ProtoReflect.Descriptor instead.
func (*SceneSummonUnitInfo) Descriptor() ([]byte, []int) {
	return file_SceneSummonUnitInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneSummonUnitInfo) GetCasterEntityId() uint32 {
	if x != nil {
		return x.CasterEntityId
	}
	return 0
}

func (x *SceneSummonUnitInfo) GetSummonUnitId() uint32 {
	if x != nil {
		return x.SummonUnitId
	}
	return 0
}

func (x *SceneSummonUnitInfo) GetTriggerNameList() []string {
	if x != nil {
		return x.TriggerNameList
	}
	return nil
}

func (x *SceneSummonUnitInfo) GetLifeTimeMs() int32 {
	if x != nil {
		return x.LifeTimeMs
	}
	return 0
}

func (x *SceneSummonUnitInfo) GetCreateTimeMs() uint64 {
	if x != nil {
		return x.CreateTimeMs
	}
	return 0
}

func (x *SceneSummonUnitInfo) GetAttachEntityId() uint32 {
	if x != nil {
		return x.AttachEntityId
	}
	return 0
}

var File_SceneSummonUnitInfo_proto protoreflect.FileDescriptor

var file_SceneSummonUnitInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x13,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f,
	0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0c, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4d,
	0x73, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x6d, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneSummonUnitInfo_proto_rawDescOnce sync.Once
	file_SceneSummonUnitInfo_proto_rawDescData = file_SceneSummonUnitInfo_proto_rawDesc
)

func file_SceneSummonUnitInfo_proto_rawDescGZIP() []byte {
	file_SceneSummonUnitInfo_proto_rawDescOnce.Do(func() {
		file_SceneSummonUnitInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneSummonUnitInfo_proto_rawDescData)
	})
	return file_SceneSummonUnitInfo_proto_rawDescData
}

var file_SceneSummonUnitInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneSummonUnitInfo_proto_goTypes = []interface{}{
	(*SceneSummonUnitInfo)(nil), // 0: SceneSummonUnitInfo
}
var file_SceneSummonUnitInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneSummonUnitInfo_proto_init() }
func file_SceneSummonUnitInfo_proto_init() {
	if File_SceneSummonUnitInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneSummonUnitInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneSummonUnitInfo); i {
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
			RawDescriptor: file_SceneSummonUnitInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneSummonUnitInfo_proto_goTypes,
		DependencyIndexes: file_SceneSummonUnitInfo_proto_depIdxs,
		MessageInfos:      file_SceneSummonUnitInfo_proto_msgTypes,
	}.Build()
	File_SceneSummonUnitInfo_proto = out.File
	file_SceneSummonUnitInfo_proto_rawDesc = nil
	file_SceneSummonUnitInfo_proto_goTypes = nil
	file_SceneSummonUnitInfo_proto_depIdxs = nil
}
