// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FightKickoutType.proto

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

type FightKickoutType int32

const (
	FightKickoutType_FIGHT_KICKOUT_UNKNOWN       FightKickoutType = 0
	FightKickoutType_FIGHT_KICKOUT_BLACK         FightKickoutType = 1
	FightKickoutType_FIGHT_KICKOUT_BY_GM         FightKickoutType = 2
	FightKickoutType_FIGHT_KICKOUT_TIMEOUT       FightKickoutType = 3
	FightKickoutType_FIGHT_KICKOUT_SESSION_RESET FightKickoutType = 4
)

// Enum value maps for FightKickoutType.
var (
	FightKickoutType_name = map[int32]string{
		0: "FIGHT_KICKOUT_UNKNOWN",
		1: "FIGHT_KICKOUT_BLACK",
		2: "FIGHT_KICKOUT_BY_GM",
		3: "FIGHT_KICKOUT_TIMEOUT",
		4: "FIGHT_KICKOUT_SESSION_RESET",
	}
	FightKickoutType_value = map[string]int32{
		"FIGHT_KICKOUT_UNKNOWN":       0,
		"FIGHT_KICKOUT_BLACK":         1,
		"FIGHT_KICKOUT_BY_GM":         2,
		"FIGHT_KICKOUT_TIMEOUT":       3,
		"FIGHT_KICKOUT_SESSION_RESET": 4,
	}
)

func (x FightKickoutType) Enum() *FightKickoutType {
	p := new(FightKickoutType)
	*p = x
	return p
}

func (x FightKickoutType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FightKickoutType) Descriptor() protoreflect.EnumDescriptor {
	return file_FightKickoutType_proto_enumTypes[0].Descriptor()
}

func (FightKickoutType) Type() protoreflect.EnumType {
	return &file_FightKickoutType_proto_enumTypes[0]
}

func (x FightKickoutType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FightKickoutType.Descriptor instead.
func (FightKickoutType) EnumDescriptor() ([]byte, []int) {
	return file_FightKickoutType_proto_rawDescGZIP(), []int{0}
}

var File_FightKickoutType_proto protoreflect.FileDescriptor

var file_FightKickoutType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x46, 0x69, 0x67, 0x68, 0x74, 0x4b, 0x69, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9b, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x4b, 0x69, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x46, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x4f, 0x55, 0x54, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x47, 0x48,
	0x54, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x4f, 0x55, 0x54, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x4f,
	0x55, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x47, 0x4d, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x49,
	0x47, 0x48, 0x54, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x4b,
	0x49, 0x43, 0x4b, 0x4f, 0x55, 0x54, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x53, 0x45, 0x54, 0x10, 0x04, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FightKickoutType_proto_rawDescOnce sync.Once
	file_FightKickoutType_proto_rawDescData = file_FightKickoutType_proto_rawDesc
)

func file_FightKickoutType_proto_rawDescGZIP() []byte {
	file_FightKickoutType_proto_rawDescOnce.Do(func() {
		file_FightKickoutType_proto_rawDescData = protoimpl.X.CompressGZIP(file_FightKickoutType_proto_rawDescData)
	})
	return file_FightKickoutType_proto_rawDescData
}

var file_FightKickoutType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FightKickoutType_proto_goTypes = []interface{}{
	(FightKickoutType)(0), // 0: FightKickoutType
}
var file_FightKickoutType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FightKickoutType_proto_init() }
func file_FightKickoutType_proto_init() {
	if File_FightKickoutType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FightKickoutType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FightKickoutType_proto_goTypes,
		DependencyIndexes: file_FightKickoutType_proto_depIdxs,
		EnumInfos:         file_FightKickoutType_proto_enumTypes,
	}.Build()
	File_FightKickoutType_proto = out.File
	file_FightKickoutType_proto_rawDesc = nil
	file_FightKickoutType_proto_goTypes = nil
	file_FightKickoutType_proto_depIdxs = nil
}