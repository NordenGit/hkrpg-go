// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueEndlessAreaData.proto

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

type RogueEndlessAreaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PJGAAONEPAH *RogueEndlessLayerInfo `protobuf:"bytes,11,opt,name=PJGAAONEPAH,proto3" json:"PJGAAONEPAH,omitempty"`
	AreaId      uint32                 `protobuf:"varint,6,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	PanelId     uint32                 `protobuf:"varint,4,opt,name=panel_id,json=panelId,proto3" json:"panel_id,omitempty"`
	GNBPBPFMGIO *LMMKPEAJFKA           `protobuf:"bytes,9,opt,name=GNBPBPFMGIO,proto3" json:"GNBPBPFMGIO,omitempty"`
}

func (x *RogueEndlessAreaData) Reset() {
	*x = RogueEndlessAreaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueEndlessAreaData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueEndlessAreaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueEndlessAreaData) ProtoMessage() {}

func (x *RogueEndlessAreaData) ProtoReflect() protoreflect.Message {
	mi := &file_RogueEndlessAreaData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueEndlessAreaData.ProtoReflect.Descriptor instead.
func (*RogueEndlessAreaData) Descriptor() ([]byte, []int) {
	return file_RogueEndlessAreaData_proto_rawDescGZIP(), []int{0}
}

func (x *RogueEndlessAreaData) GetPJGAAONEPAH() *RogueEndlessLayerInfo {
	if x != nil {
		return x.PJGAAONEPAH
	}
	return nil
}

func (x *RogueEndlessAreaData) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *RogueEndlessAreaData) GetPanelId() uint32 {
	if x != nil {
		return x.PanelId
	}
	return 0
}

func (x *RogueEndlessAreaData) GetGNBPBPFMGIO() *LMMKPEAJFKA {
	if x != nil {
		return x.GNBPBPFMGIO
	}
	return nil
}

var File_RogueEndlessAreaData_proto protoreflect.FileDescriptor

var file_RogueEndlessAreaData_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x41, 0x72,
	0x65, 0x61, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x4d,
	0x4d, 0x4b, 0x50, 0x45, 0x41, 0x4a, 0x46, 0x4b, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a,
	0x14, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x41, 0x72, 0x65,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0b, 0x50, 0x4a, 0x47, 0x41, 0x41, 0x4f, 0x4e,
	0x45, 0x50, 0x41, 0x48, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0b, 0x50, 0x4a, 0x47, 0x41, 0x41, 0x4f, 0x4e, 0x45, 0x50, 0x41, 0x48, 0x12,
	0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x4e, 0x42, 0x50, 0x42, 0x50, 0x46, 0x4d, 0x47,
	0x49, 0x4f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x4d, 0x4d, 0x4b, 0x50,
	0x45, 0x41, 0x4a, 0x46, 0x4b, 0x41, 0x52, 0x0b, 0x47, 0x4e, 0x42, 0x50, 0x42, 0x50, 0x46, 0x4d,
	0x47, 0x49, 0x4f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueEndlessAreaData_proto_rawDescOnce sync.Once
	file_RogueEndlessAreaData_proto_rawDescData = file_RogueEndlessAreaData_proto_rawDesc
)

func file_RogueEndlessAreaData_proto_rawDescGZIP() []byte {
	file_RogueEndlessAreaData_proto_rawDescOnce.Do(func() {
		file_RogueEndlessAreaData_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueEndlessAreaData_proto_rawDescData)
	})
	return file_RogueEndlessAreaData_proto_rawDescData
}

var file_RogueEndlessAreaData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueEndlessAreaData_proto_goTypes = []interface{}{
	(*RogueEndlessAreaData)(nil),  // 0: RogueEndlessAreaData
	(*RogueEndlessLayerInfo)(nil), // 1: RogueEndlessLayerInfo
	(*LMMKPEAJFKA)(nil),           // 2: LMMKPEAJFKA
}
var file_RogueEndlessAreaData_proto_depIdxs = []int32{
	1, // 0: RogueEndlessAreaData.PJGAAONEPAH:type_name -> RogueEndlessLayerInfo
	2, // 1: RogueEndlessAreaData.GNBPBPFMGIO:type_name -> LMMKPEAJFKA
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueEndlessAreaData_proto_init() }
func file_RogueEndlessAreaData_proto_init() {
	if File_RogueEndlessAreaData_proto != nil {
		return
	}
	file_LMMKPEAJFKA_proto_init()
	file_RogueEndlessLayerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueEndlessAreaData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueEndlessAreaData); i {
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
			RawDescriptor: file_RogueEndlessAreaData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueEndlessAreaData_proto_goTypes,
		DependencyIndexes: file_RogueEndlessAreaData_proto_depIdxs,
		MessageInfos:      file_RogueEndlessAreaData_proto_msgTypes,
	}.Build()
	File_RogueEndlessAreaData_proto = out.File
	file_RogueEndlessAreaData_proto_rawDesc = nil
	file_RogueEndlessAreaData_proto_goTypes = nil
	file_RogueEndlessAreaData_proto_depIdxs = nil
}
