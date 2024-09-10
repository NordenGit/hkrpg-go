// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RelicFilterPlanSettings.proto

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

type RelicFilterPlanSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RarityBitset               uint32   `protobuf:"varint,1,opt,name=rarity_bitset,json=rarityBitset,proto3" json:"rarity_bitset,omitempty"`
	RelicSetList               []uint32 `protobuf:"varint,2,rep,packed,name=relic_set_list,json=relicSetList,proto3" json:"relic_set_list,omitempty"`
	BodyMainPropertyList       []uint32 `protobuf:"varint,3,rep,packed,name=body_main_property_list,json=bodyMainPropertyList,proto3" json:"body_main_property_list,omitempty"`
	FootMainPropertyList       []uint32 `protobuf:"varint,4,rep,packed,name=foot_main_property_list,json=footMainPropertyList,proto3" json:"foot_main_property_list,omitempty"`
	SphereMainPropertyList     []uint32 `protobuf:"varint,5,rep,packed,name=sphere_main_property_list,json=sphereMainPropertyList,proto3" json:"sphere_main_property_list,omitempty"`
	RopeMainPropertyList       []uint32 `protobuf:"varint,6,rep,packed,name=rope_main_property_list,json=ropeMainPropertyList,proto3" json:"rope_main_property_list,omitempty"`
	IsIncludeFilterSubProperty bool     `protobuf:"varint,7,opt,name=is_include_filter_sub_property,json=isIncludeFilterSubProperty,proto3" json:"is_include_filter_sub_property,omitempty"`
	SubPropertyNum             uint32   `protobuf:"varint,8,opt,name=sub_property_num,json=subPropertyNum,proto3" json:"sub_property_num,omitempty"`
	SubPropertyList            []uint32 `protobuf:"varint,9,rep,packed,name=sub_property_list,json=subPropertyList,proto3" json:"sub_property_list,omitempty"`
}

func (x *RelicFilterPlanSettings) Reset() {
	*x = RelicFilterPlanSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RelicFilterPlanSettings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelicFilterPlanSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelicFilterPlanSettings) ProtoMessage() {}

func (x *RelicFilterPlanSettings) ProtoReflect() protoreflect.Message {
	mi := &file_RelicFilterPlanSettings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelicFilterPlanSettings.ProtoReflect.Descriptor instead.
func (*RelicFilterPlanSettings) Descriptor() ([]byte, []int) {
	return file_RelicFilterPlanSettings_proto_rawDescGZIP(), []int{0}
}

func (x *RelicFilterPlanSettings) GetRarityBitset() uint32 {
	if x != nil {
		return x.RarityBitset
	}
	return 0
}

func (x *RelicFilterPlanSettings) GetRelicSetList() []uint32 {
	if x != nil {
		return x.RelicSetList
	}
	return nil
}

func (x *RelicFilterPlanSettings) GetBodyMainPropertyList() []uint32 {
	if x != nil {
		return x.BodyMainPropertyList
	}
	return nil
}

func (x *RelicFilterPlanSettings) GetFootMainPropertyList() []uint32 {
	if x != nil {
		return x.FootMainPropertyList
	}
	return nil
}

func (x *RelicFilterPlanSettings) GetSphereMainPropertyList() []uint32 {
	if x != nil {
		return x.SphereMainPropertyList
	}
	return nil
}

func (x *RelicFilterPlanSettings) GetRopeMainPropertyList() []uint32 {
	if x != nil {
		return x.RopeMainPropertyList
	}
	return nil
}

func (x *RelicFilterPlanSettings) GetIsIncludeFilterSubProperty() bool {
	if x != nil {
		return x.IsIncludeFilterSubProperty
	}
	return false
}

func (x *RelicFilterPlanSettings) GetSubPropertyNum() uint32 {
	if x != nil {
		return x.SubPropertyNum
	}
	return 0
}

func (x *RelicFilterPlanSettings) GetSubPropertyList() []uint32 {
	if x != nil {
		return x.SubPropertyList
	}
	return nil
}

var File_RelicFilterPlanSettings_proto protoreflect.FileDescriptor

var file_RelicFilterPlanSettings_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xde, 0x03, 0x0a, 0x17, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x42, 0x69, 0x74, 0x73, 0x65, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x53,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x14, 0x62, 0x6f, 0x64, 0x79, 0x4d, 0x61, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a,
	0x17, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x14,
	0x66, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x16, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x4d,
	0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x17, 0x72, 0x6f, 0x70, 0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x14, 0x72, 0x6f, 0x70, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x1e, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a,
	0x69, 0x73, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x75,
	0x62, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0f, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RelicFilterPlanSettings_proto_rawDescOnce sync.Once
	file_RelicFilterPlanSettings_proto_rawDescData = file_RelicFilterPlanSettings_proto_rawDesc
)

func file_RelicFilterPlanSettings_proto_rawDescGZIP() []byte {
	file_RelicFilterPlanSettings_proto_rawDescOnce.Do(func() {
		file_RelicFilterPlanSettings_proto_rawDescData = protoimpl.X.CompressGZIP(file_RelicFilterPlanSettings_proto_rawDescData)
	})
	return file_RelicFilterPlanSettings_proto_rawDescData
}

var file_RelicFilterPlanSettings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RelicFilterPlanSettings_proto_goTypes = []interface{}{
	(*RelicFilterPlanSettings)(nil), // 0: RelicFilterPlanSettings
}
var file_RelicFilterPlanSettings_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RelicFilterPlanSettings_proto_init() }
func file_RelicFilterPlanSettings_proto_init() {
	if File_RelicFilterPlanSettings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RelicFilterPlanSettings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelicFilterPlanSettings); i {
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
			RawDescriptor: file_RelicFilterPlanSettings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RelicFilterPlanSettings_proto_goTypes,
		DependencyIndexes: file_RelicFilterPlanSettings_proto_depIdxs,
		MessageInfos:      file_RelicFilterPlanSettings_proto_msgTypes,
	}.Build()
	File_RelicFilterPlanSettings_proto = out.File
	file_RelicFilterPlanSettings_proto_rawDesc = nil
	file_RelicFilterPlanSettings_proto_goTypes = nil
	file_RelicFilterPlanSettings_proto_depIdxs = nil
}
