// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DILKCHLPOBA.proto

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

type DILKCHLPOBA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         uint32       `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Level       uint32       `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Nickname    string       `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	JOHLDPCECHI uint32       `protobuf:"varint,4,opt,name=JOHLDPCECHI,proto3" json:"JOHLDPCECHI,omitempty"`
	Platform    PlatformType `protobuf:"varint,5,opt,name=platform,proto3,enum=PlatformType" json:"platform,omitempty"`
	CDFPLBBIKDI string       `protobuf:"bytes,6,opt,name=CDFPLBBIKDI,proto3" json:"CDFPLBBIKDI,omitempty"`
	COJFJAKCLKJ string       `protobuf:"bytes,7,opt,name=COJFJAKCLKJ,proto3" json:"COJFJAKCLKJ,omitempty"`
	Version     uint64       `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DILKCHLPOBA) Reset() {
	*x = DILKCHLPOBA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DILKCHLPOBA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DILKCHLPOBA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DILKCHLPOBA) ProtoMessage() {}

func (x *DILKCHLPOBA) ProtoReflect() protoreflect.Message {
	mi := &file_DILKCHLPOBA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DILKCHLPOBA.ProtoReflect.Descriptor instead.
func (*DILKCHLPOBA) Descriptor() ([]byte, []int) {
	return file_DILKCHLPOBA_proto_rawDescGZIP(), []int{0}
}

func (x *DILKCHLPOBA) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DILKCHLPOBA) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *DILKCHLPOBA) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *DILKCHLPOBA) GetJOHLDPCECHI() uint32 {
	if x != nil {
		return x.JOHLDPCECHI
	}
	return 0
}

func (x *DILKCHLPOBA) GetPlatform() PlatformType {
	if x != nil {
		return x.Platform
	}
	return PlatformType_EDITOR
}

func (x *DILKCHLPOBA) GetCDFPLBBIKDI() string {
	if x != nil {
		return x.CDFPLBBIKDI
	}
	return ""
}

func (x *DILKCHLPOBA) GetCOJFJAKCLKJ() string {
	if x != nil {
		return x.COJFJAKCLKJ
	}
	return ""
}

func (x *DILKCHLPOBA) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_DILKCHLPOBA_proto protoreflect.FileDescriptor

var file_DILKCHLPOBA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x49, 0x4c, 0x4b, 0x43, 0x48, 0x4c, 0x50, 0x4f, 0x42, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x0b, 0x44, 0x49, 0x4c, 0x4b,
	0x43, 0x48, 0x4c, 0x50, 0x4f, 0x42, 0x41, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x4f, 0x48, 0x4c, 0x44, 0x50, 0x43, 0x45, 0x43, 0x48, 0x49, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x4f, 0x48, 0x4c, 0x44, 0x50, 0x43, 0x45, 0x43, 0x48, 0x49, 0x12, 0x29, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x44, 0x46, 0x50,
	0x4c, 0x42, 0x42, 0x49, 0x4b, 0x44, 0x49, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43,
	0x44, 0x46, 0x50, 0x4c, 0x42, 0x42, 0x49, 0x4b, 0x44, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4f,
	0x4a, 0x46, 0x4a, 0x41, 0x4b, 0x43, 0x4c, 0x4b, 0x4a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x43, 0x4f, 0x4a, 0x46, 0x4a, 0x41, 0x4b, 0x43, 0x4c, 0x4b, 0x4a, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DILKCHLPOBA_proto_rawDescOnce sync.Once
	file_DILKCHLPOBA_proto_rawDescData = file_DILKCHLPOBA_proto_rawDesc
)

func file_DILKCHLPOBA_proto_rawDescGZIP() []byte {
	file_DILKCHLPOBA_proto_rawDescOnce.Do(func() {
		file_DILKCHLPOBA_proto_rawDescData = protoimpl.X.CompressGZIP(file_DILKCHLPOBA_proto_rawDescData)
	})
	return file_DILKCHLPOBA_proto_rawDescData
}

var file_DILKCHLPOBA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DILKCHLPOBA_proto_goTypes = []interface{}{
	(*DILKCHLPOBA)(nil), // 0: DILKCHLPOBA
	(PlatformType)(0),   // 1: PlatformType
}
var file_DILKCHLPOBA_proto_depIdxs = []int32{
	1, // 0: DILKCHLPOBA.platform:type_name -> PlatformType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DILKCHLPOBA_proto_init() }
func file_DILKCHLPOBA_proto_init() {
	if File_DILKCHLPOBA_proto != nil {
		return
	}
	file_PlatformType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DILKCHLPOBA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DILKCHLPOBA); i {
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
			RawDescriptor: file_DILKCHLPOBA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DILKCHLPOBA_proto_goTypes,
		DependencyIndexes: file_DILKCHLPOBA_proto_depIdxs,
		MessageInfos:      file_DILKCHLPOBA_proto_msgTypes,
	}.Build()
	File_DILKCHLPOBA_proto = out.File
	file_DILKCHLPOBA_proto_rawDesc = nil
	file_DILKCHLPOBA_proto_goTypes = nil
	file_DILKCHLPOBA_proto_depIdxs = nil
}