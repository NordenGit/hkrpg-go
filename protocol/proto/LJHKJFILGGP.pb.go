// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LJHKJFILGGP.proto

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

type LJHKJFILGGP int32

const (
	LJHKJFILGGP_UPDATE_REDDOT_NONE    LJHKJFILGGP = 0
	LJHKJFILGGP_UPDATE_REDDOT_ADD     LJHKJFILGGP = 1
	LJHKJFILGGP_UPDATE_REDDOT_REPLACE LJHKJFILGGP = 2
)

// Enum value maps for LJHKJFILGGP.
var (
	LJHKJFILGGP_name = map[int32]string{
		0: "UPDATE_REDDOT_NONE",
		1: "UPDATE_REDDOT_ADD",
		2: "UPDATE_REDDOT_REPLACE",
	}
	LJHKJFILGGP_value = map[string]int32{
		"UPDATE_REDDOT_NONE":    0,
		"UPDATE_REDDOT_ADD":     1,
		"UPDATE_REDDOT_REPLACE": 2,
	}
)

func (x LJHKJFILGGP) Enum() *LJHKJFILGGP {
	p := new(LJHKJFILGGP)
	*p = x
	return p
}

func (x LJHKJFILGGP) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LJHKJFILGGP) Descriptor() protoreflect.EnumDescriptor {
	return file_LJHKJFILGGP_proto_enumTypes[0].Descriptor()
}

func (LJHKJFILGGP) Type() protoreflect.EnumType {
	return &file_LJHKJFILGGP_proto_enumTypes[0]
}

func (x LJHKJFILGGP) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LJHKJFILGGP.Descriptor instead.
func (LJHKJFILGGP) EnumDescriptor() ([]byte, []int) {
	return file_LJHKJFILGGP_proto_rawDescGZIP(), []int{0}
}

var File_LJHKJFILGGP_proto protoreflect.FileDescriptor

var file_LJHKJFILGGP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x4a, 0x48, 0x4b, 0x4a, 0x46, 0x49, 0x4c, 0x47, 0x47, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x57, 0x0a, 0x0b, 0x4c, 0x4a, 0x48, 0x4b, 0x4a, 0x46, 0x49, 0x4c, 0x47,
	0x47, 0x50, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x44,
	0x44, 0x4f, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x44, 0x44, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x44, 0x44,
	0x4f, 0x54, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LJHKJFILGGP_proto_rawDescOnce sync.Once
	file_LJHKJFILGGP_proto_rawDescData = file_LJHKJFILGGP_proto_rawDesc
)

func file_LJHKJFILGGP_proto_rawDescGZIP() []byte {
	file_LJHKJFILGGP_proto_rawDescOnce.Do(func() {
		file_LJHKJFILGGP_proto_rawDescData = protoimpl.X.CompressGZIP(file_LJHKJFILGGP_proto_rawDescData)
	})
	return file_LJHKJFILGGP_proto_rawDescData
}

var file_LJHKJFILGGP_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LJHKJFILGGP_proto_goTypes = []interface{}{
	(LJHKJFILGGP)(0), // 0: LJHKJFILGGP
}
var file_LJHKJFILGGP_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LJHKJFILGGP_proto_init() }
func file_LJHKJFILGGP_proto_init() {
	if File_LJHKJFILGGP_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LJHKJFILGGP_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LJHKJFILGGP_proto_goTypes,
		DependencyIndexes: file_LJHKJFILGGP_proto_depIdxs,
		EnumInfos:         file_LJHKJFILGGP_proto_enumTypes,
	}.Build()
	File_LJHKJFILGGP_proto = out.File
	file_LJHKJFILGGP_proto_rawDesc = nil
	file_LJHKJFILGGP_proto_goTypes = nil
	file_LJHKJFILGGP_proto_depIdxs = nil
}
