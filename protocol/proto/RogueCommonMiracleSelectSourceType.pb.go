// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCommonMiracleSelectSourceType.proto

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

type RogueCommonMiracleSelectSourceType int32

const (
	RogueCommonMiracleSelectSourceType_ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_NONE            RogueCommonMiracleSelectSourceType = 0
	RogueCommonMiracleSelectSourceType_ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_DICE_ROLL       RogueCommonMiracleSelectSourceType = 1
	RogueCommonMiracleSelectSourceType_ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_AEON            RogueCommonMiracleSelectSourceType = 2
	RogueCommonMiracleSelectSourceType_ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_BOARD_EVENT     RogueCommonMiracleSelectSourceType = 3
	RogueCommonMiracleSelectSourceType_ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_LEVEL_MECHANISM RogueCommonMiracleSelectSourceType = 4
)

// Enum value maps for RogueCommonMiracleSelectSourceType.
var (
	RogueCommonMiracleSelectSourceType_name = map[int32]string{
		0: "ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_NONE",
		1: "ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_DICE_ROLL",
		2: "ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_AEON",
		3: "ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_BOARD_EVENT",
		4: "ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_LEVEL_MECHANISM",
	}
	RogueCommonMiracleSelectSourceType_value = map[string]int32{
		"ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_NONE":            0,
		"ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_DICE_ROLL":       1,
		"ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_AEON":            2,
		"ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_BOARD_EVENT":     3,
		"ROGUE_COMMON_MIRACLE_SELECT_SOURCE_TYPE_LEVEL_MECHANISM": 4,
	}
)

func (x RogueCommonMiracleSelectSourceType) Enum() *RogueCommonMiracleSelectSourceType {
	p := new(RogueCommonMiracleSelectSourceType)
	*p = x
	return p
}

func (x RogueCommonMiracleSelectSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueCommonMiracleSelectSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueCommonMiracleSelectSourceType_proto_enumTypes[0].Descriptor()
}

func (RogueCommonMiracleSelectSourceType) Type() protoreflect.EnumType {
	return &file_RogueCommonMiracleSelectSourceType_proto_enumTypes[0]
}

func (x RogueCommonMiracleSelectSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueCommonMiracleSelectSourceType.Descriptor instead.
func (RogueCommonMiracleSelectSourceType) EnumDescriptor() ([]byte, []int) {
	return file_RogueCommonMiracleSelectSourceType_proto_rawDescGZIP(), []int{0}
}

var File_RogueCommonMiracleSelectSourceType_proto protoreflect.FileDescriptor

var file_RogueCommonMiracleSelectSourceType_proto_rawDesc = []byte{
	0x0a, 0x28, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb5, 0x02, 0x0a, 0x22, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x30, 0x0a, 0x2c, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f,
	0x4e, 0x5f, 0x4d, 0x49, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54,
	0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x35, 0x0a, 0x31, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45,
	0x43, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x49, 0x43, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x30, 0x0a, 0x2c, 0x52, 0x4f,
	0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x52, 0x41, 0x43,
	0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x45, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x37, 0x0a, 0x33,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x52,
	0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x3b, 0x0a, 0x37, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x45,
	0x4c, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x53, 0x4d,
	0x10, 0x04, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCommonMiracleSelectSourceType_proto_rawDescOnce sync.Once
	file_RogueCommonMiracleSelectSourceType_proto_rawDescData = file_RogueCommonMiracleSelectSourceType_proto_rawDesc
)

func file_RogueCommonMiracleSelectSourceType_proto_rawDescGZIP() []byte {
	file_RogueCommonMiracleSelectSourceType_proto_rawDescOnce.Do(func() {
		file_RogueCommonMiracleSelectSourceType_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCommonMiracleSelectSourceType_proto_rawDescData)
	})
	return file_RogueCommonMiracleSelectSourceType_proto_rawDescData
}

var file_RogueCommonMiracleSelectSourceType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueCommonMiracleSelectSourceType_proto_goTypes = []interface{}{
	(RogueCommonMiracleSelectSourceType)(0), // 0: RogueCommonMiracleSelectSourceType
}
var file_RogueCommonMiracleSelectSourceType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueCommonMiracleSelectSourceType_proto_init() }
func file_RogueCommonMiracleSelectSourceType_proto_init() {
	if File_RogueCommonMiracleSelectSourceType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueCommonMiracleSelectSourceType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCommonMiracleSelectSourceType_proto_goTypes,
		DependencyIndexes: file_RogueCommonMiracleSelectSourceType_proto_depIdxs,
		EnumInfos:         file_RogueCommonMiracleSelectSourceType_proto_enumTypes,
	}.Build()
	File_RogueCommonMiracleSelectSourceType_proto = out.File
	file_RogueCommonMiracleSelectSourceType_proto_rawDesc = nil
	file_RogueCommonMiracleSelectSourceType_proto_goTypes = nil
	file_RogueCommonMiracleSelectSourceType_proto_depIdxs = nil
}