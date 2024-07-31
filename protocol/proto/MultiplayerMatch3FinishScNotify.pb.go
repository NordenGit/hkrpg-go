// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MultiplayerMatch3FinishScNotify.proto

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

type MultiplayerMatch3FinishScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DJPAJNOCHHG uint32             `protobuf:"varint,7,opt,name=DJPAJNOCHHG,proto3" json:"DJPAJNOCHHG,omitempty"`
	FJAFGJLFJGD uint32             `protobuf:"varint,13,opt,name=FJAFGJLFJGD,proto3" json:"FJAFGJLFJGD,omitempty"`
	Reason      Match3FinishReason `protobuf:"varint,10,opt,name=reason,proto3,enum=Match3FinishReason" json:"reason,omitempty"`
	FANMEJFOCIL *HLEMBLOIKAM       `protobuf:"bytes,11,opt,name=FANMEJFOCIL,proto3" json:"FANMEJFOCIL,omitempty"`
}

func (x *MultiplayerMatch3FinishScNotify) Reset() {
	*x = MultiplayerMatch3FinishScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultiplayerMatch3FinishScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplayerMatch3FinishScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplayerMatch3FinishScNotify) ProtoMessage() {}

func (x *MultiplayerMatch3FinishScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MultiplayerMatch3FinishScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplayerMatch3FinishScNotify.ProtoReflect.Descriptor instead.
func (*MultiplayerMatch3FinishScNotify) Descriptor() ([]byte, []int) {
	return file_MultiplayerMatch3FinishScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MultiplayerMatch3FinishScNotify) GetDJPAJNOCHHG() uint32 {
	if x != nil {
		return x.DJPAJNOCHHG
	}
	return 0
}

func (x *MultiplayerMatch3FinishScNotify) GetFJAFGJLFJGD() uint32 {
	if x != nil {
		return x.FJAFGJLFJGD
	}
	return 0
}

func (x *MultiplayerMatch3FinishScNotify) GetReason() Match3FinishReason {
	if x != nil {
		return x.Reason
	}
	return Match3FinishReason_MATCH3_FINISH_REASON_DEFAULT
}

func (x *MultiplayerMatch3FinishScNotify) GetFANMEJFOCIL() *HLEMBLOIKAM {
	if x != nil {
		return x.FANMEJFOCIL
	}
	return nil
}

var File_MultiplayerMatch3FinishScNotify_proto protoreflect.FileDescriptor

var file_MultiplayerMatch3FinishScNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x33, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x33, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x48, 0x4c, 0x45, 0x4d, 0x42, 0x4c, 0x4f, 0x49, 0x4b, 0x41, 0x4d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x1f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x33, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4a, 0x50, 0x41,
	0x4a, 0x4e, 0x4f, 0x43, 0x48, 0x48, 0x47, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44,
	0x4a, 0x50, 0x41, 0x4a, 0x4e, 0x4f, 0x43, 0x48, 0x48, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4a,
	0x41, 0x46, 0x47, 0x4a, 0x4c, 0x46, 0x4a, 0x47, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x46, 0x4a, 0x41, 0x46, 0x47, 0x4a, 0x4c, 0x46, 0x4a, 0x47, 0x44, 0x12, 0x2b, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x33, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x41, 0x4e,
	0x4d, 0x45, 0x4a, 0x46, 0x4f, 0x43, 0x49, 0x4c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x48, 0x4c, 0x45, 0x4d, 0x42, 0x4c, 0x4f, 0x49, 0x4b, 0x41, 0x4d, 0x52, 0x0b, 0x46, 0x41,
	0x4e, 0x4d, 0x45, 0x4a, 0x46, 0x4f, 0x43, 0x49, 0x4c, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MultiplayerMatch3FinishScNotify_proto_rawDescOnce sync.Once
	file_MultiplayerMatch3FinishScNotify_proto_rawDescData = file_MultiplayerMatch3FinishScNotify_proto_rawDesc
)

func file_MultiplayerMatch3FinishScNotify_proto_rawDescGZIP() []byte {
	file_MultiplayerMatch3FinishScNotify_proto_rawDescOnce.Do(func() {
		file_MultiplayerMatch3FinishScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MultiplayerMatch3FinishScNotify_proto_rawDescData)
	})
	return file_MultiplayerMatch3FinishScNotify_proto_rawDescData
}

var file_MultiplayerMatch3FinishScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MultiplayerMatch3FinishScNotify_proto_goTypes = []interface{}{
	(*MultiplayerMatch3FinishScNotify)(nil), // 0: MultiplayerMatch3FinishScNotify
	(Match3FinishReason)(0),                 // 1: Match3FinishReason
	(*HLEMBLOIKAM)(nil),                     // 2: HLEMBLOIKAM
}
var file_MultiplayerMatch3FinishScNotify_proto_depIdxs = []int32{
	1, // 0: MultiplayerMatch3FinishScNotify.reason:type_name -> Match3FinishReason
	2, // 1: MultiplayerMatch3FinishScNotify.FANMEJFOCIL:type_name -> HLEMBLOIKAM
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MultiplayerMatch3FinishScNotify_proto_init() }
func file_MultiplayerMatch3FinishScNotify_proto_init() {
	if File_MultiplayerMatch3FinishScNotify_proto != nil {
		return
	}
	file_Match3FinishReason_proto_init()
	file_HLEMBLOIKAM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MultiplayerMatch3FinishScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplayerMatch3FinishScNotify); i {
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
			RawDescriptor: file_MultiplayerMatch3FinishScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MultiplayerMatch3FinishScNotify_proto_goTypes,
		DependencyIndexes: file_MultiplayerMatch3FinishScNotify_proto_depIdxs,
		MessageInfos:      file_MultiplayerMatch3FinishScNotify_proto_msgTypes,
	}.Build()
	File_MultiplayerMatch3FinishScNotify_proto = out.File
	file_MultiplayerMatch3FinishScNotify_proto_rawDesc = nil
	file_MultiplayerMatch3FinishScNotify_proto_goTypes = nil
	file_MultiplayerMatch3FinishScNotify_proto_depIdxs = nil
}