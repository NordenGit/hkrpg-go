// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MissionStatusBySceneInfo.proto

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

type MissionStatusBySceneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainMissionMcvList          []*MainMissionCustomValue `protobuf:"bytes,10,rep,name=main_mission_mcv_list,json=mainMissionMcvList,proto3" json:"main_mission_mcv_list,omitempty"`
	DisabledMainMissionIdList   []uint32                  `protobuf:"varint,1,rep,packed,name=disabled_main_mission_id_list,json=disabledMainMissionIdList,proto3" json:"disabled_main_mission_id_list,omitempty"` // 7
	SubMissionStatusList        []*Mission                `protobuf:"bytes,5,rep,name=sub_mission_status_list,json=subMissionStatusList,proto3" json:"sub_mission_status_list,omitempty"`
	UnfinishedMainMissionIdList []uint32                  `protobuf:"varint,7,rep,packed,name=unfinished_main_mission_id_list,json=unfinishedMainMissionIdList,proto3" json:"unfinished_main_mission_id_list,omitempty"` // 1
	FinishedMainMissionIdList   []uint32                  `protobuf:"varint,8,rep,packed,name=finished_main_mission_id_list,json=finishedMainMissionIdList,proto3" json:"finished_main_mission_id_list,omitempty"`
}

func (x *MissionStatusBySceneInfo) Reset() {
	*x = MissionStatusBySceneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MissionStatusBySceneInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissionStatusBySceneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissionStatusBySceneInfo) ProtoMessage() {}

func (x *MissionStatusBySceneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MissionStatusBySceneInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissionStatusBySceneInfo.ProtoReflect.Descriptor instead.
func (*MissionStatusBySceneInfo) Descriptor() ([]byte, []int) {
	return file_MissionStatusBySceneInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MissionStatusBySceneInfo) GetMainMissionMcvList() []*MainMissionCustomValue {
	if x != nil {
		return x.MainMissionMcvList
	}
	return nil
}

func (x *MissionStatusBySceneInfo) GetDisabledMainMissionIdList() []uint32 {
	if x != nil {
		return x.DisabledMainMissionIdList
	}
	return nil
}

func (x *MissionStatusBySceneInfo) GetSubMissionStatusList() []*Mission {
	if x != nil {
		return x.SubMissionStatusList
	}
	return nil
}

func (x *MissionStatusBySceneInfo) GetUnfinishedMainMissionIdList() []uint32 {
	if x != nil {
		return x.UnfinishedMainMissionIdList
	}
	return nil
}

func (x *MissionStatusBySceneInfo) GetFinishedMainMissionIdList() []uint32 {
	if x != nil {
		return x.FinishedMainMissionIdList
	}
	return nil
}

var File_MissionStatusBySceneInfo_proto protoreflect.FileDescriptor

var file_MissionStatusBySceneInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x79, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x02,
	0x0a, 0x18, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x79, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x15, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x63, 0x76, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x4d, 0x61, 0x69, 0x6e,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d,
	0x63, 0x76, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x1d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x19, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x17, 0x73, 0x75, 0x62, 0x5f,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x14, 0x73, 0x75, 0x62, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x1f, 0x75, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x1b, 0x75, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4d, 0x61,
	0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x40, 0x0a, 0x1d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x19, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MissionStatusBySceneInfo_proto_rawDescOnce sync.Once
	file_MissionStatusBySceneInfo_proto_rawDescData = file_MissionStatusBySceneInfo_proto_rawDesc
)

func file_MissionStatusBySceneInfo_proto_rawDescGZIP() []byte {
	file_MissionStatusBySceneInfo_proto_rawDescOnce.Do(func() {
		file_MissionStatusBySceneInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MissionStatusBySceneInfo_proto_rawDescData)
	})
	return file_MissionStatusBySceneInfo_proto_rawDescData
}

var file_MissionStatusBySceneInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MissionStatusBySceneInfo_proto_goTypes = []interface{}{
	(*MissionStatusBySceneInfo)(nil), // 0: MissionStatusBySceneInfo
	(*MainMissionCustomValue)(nil),   // 1: MainMissionCustomValue
	(*Mission)(nil),                  // 2: Mission
}
var file_MissionStatusBySceneInfo_proto_depIdxs = []int32{
	1, // 0: MissionStatusBySceneInfo.main_mission_mcv_list:type_name -> MainMissionCustomValue
	2, // 1: MissionStatusBySceneInfo.sub_mission_status_list:type_name -> Mission
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MissionStatusBySceneInfo_proto_init() }
func file_MissionStatusBySceneInfo_proto_init() {
	if File_MissionStatusBySceneInfo_proto != nil {
		return
	}
	file_Mission_proto_init()
	file_MainMissionCustomValue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MissionStatusBySceneInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissionStatusBySceneInfo); i {
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
			RawDescriptor: file_MissionStatusBySceneInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MissionStatusBySceneInfo_proto_goTypes,
		DependencyIndexes: file_MissionStatusBySceneInfo_proto_depIdxs,
		MessageInfos:      file_MissionStatusBySceneInfo_proto_msgTypes,
	}.Build()
	File_MissionStatusBySceneInfo_proto = out.File
	file_MissionStatusBySceneInfo_proto_rawDesc = nil
	file_MissionStatusBySceneInfo_proto_goTypes = nil
	file_MissionStatusBySceneInfo_proto_depIdxs = nil
}
