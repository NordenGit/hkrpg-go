// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournSettleReason.proto

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

type RogueTournSettleReason int32

const (
	RogueTournSettleReason_ROGUE_TOURN_SETTLE_REASON_NONE      RogueTournSettleReason = 0
	RogueTournSettleReason_ROGUE_TOURN_SETTLE_REASON_WIN       RogueTournSettleReason = 1
	RogueTournSettleReason_ROGUE_TOURN_SETTLE_REASON_FAIL      RogueTournSettleReason = 2
	RogueTournSettleReason_ROGUE_TOURN_SETTLE_REASON_INTERRUPT RogueTournSettleReason = 3
)

// Enum value maps for RogueTournSettleReason.
var (
	RogueTournSettleReason_name = map[int32]string{
		0: "ROGUE_TOURN_SETTLE_REASON_NONE",
		1: "ROGUE_TOURN_SETTLE_REASON_WIN",
		2: "ROGUE_TOURN_SETTLE_REASON_FAIL",
		3: "ROGUE_TOURN_SETTLE_REASON_INTERRUPT",
	}
	RogueTournSettleReason_value = map[string]int32{
		"ROGUE_TOURN_SETTLE_REASON_NONE":      0,
		"ROGUE_TOURN_SETTLE_REASON_WIN":       1,
		"ROGUE_TOURN_SETTLE_REASON_FAIL":      2,
		"ROGUE_TOURN_SETTLE_REASON_INTERRUPT": 3,
	}
)

func (x RogueTournSettleReason) Enum() *RogueTournSettleReason {
	p := new(RogueTournSettleReason)
	*p = x
	return p
}

func (x RogueTournSettleReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueTournSettleReason) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueTournSettleReason_proto_enumTypes[0].Descriptor()
}

func (RogueTournSettleReason) Type() protoreflect.EnumType {
	return &file_RogueTournSettleReason_proto_enumTypes[0]
}

func (x RogueTournSettleReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueTournSettleReason.Descriptor instead.
func (RogueTournSettleReason) EnumDescriptor() ([]byte, []int) {
	return file_RogueTournSettleReason_proto_rawDescGZIP(), []int{0}
}

var File_RogueTournSettleReason_proto protoreflect.FileDescriptor

var file_RogueTournSettleReason_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xac,
	0x01, 0x0a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x4f, 0x47,
	0x55, 0x45, 0x5f, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a,
	0x1d, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x45, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x57, 0x49, 0x4e, 0x10, 0x01,
	0x12, 0x22, 0x0a, 0x1e, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x54, 0x4f,
	0x55, 0x52, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x10, 0x03, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournSettleReason_proto_rawDescOnce sync.Once
	file_RogueTournSettleReason_proto_rawDescData = file_RogueTournSettleReason_proto_rawDesc
)

func file_RogueTournSettleReason_proto_rawDescGZIP() []byte {
	file_RogueTournSettleReason_proto_rawDescOnce.Do(func() {
		file_RogueTournSettleReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournSettleReason_proto_rawDescData)
	})
	return file_RogueTournSettleReason_proto_rawDescData
}

var file_RogueTournSettleReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueTournSettleReason_proto_goTypes = []interface{}{
	(RogueTournSettleReason)(0), // 0: RogueTournSettleReason
}
var file_RogueTournSettleReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueTournSettleReason_proto_init() }
func file_RogueTournSettleReason_proto_init() {
	if File_RogueTournSettleReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueTournSettleReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournSettleReason_proto_goTypes,
		DependencyIndexes: file_RogueTournSettleReason_proto_depIdxs,
		EnumInfos:         file_RogueTournSettleReason_proto_enumTypes,
	}.Build()
	File_RogueTournSettleReason_proto = out.File
	file_RogueTournSettleReason_proto_rawDesc = nil
	file_RogueTournSettleReason_proto_goTypes = nil
	file_RogueTournSettleReason_proto_depIdxs = nil
}