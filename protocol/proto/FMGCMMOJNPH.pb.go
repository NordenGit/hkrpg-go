// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FMGCMMOJNPH.proto

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

type FMGCMMOJNPH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EMEFKPIMKOJ map[uint32]*IPOFINFLFNE `protobuf:"bytes,9,rep,name=EMEFKPIMKOJ,proto3" json:"EMEFKPIMKOJ,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FMGCMMOJNPH) Reset() {
	*x = FMGCMMOJNPH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FMGCMMOJNPH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FMGCMMOJNPH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FMGCMMOJNPH) ProtoMessage() {}

func (x *FMGCMMOJNPH) ProtoReflect() protoreflect.Message {
	mi := &file_FMGCMMOJNPH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FMGCMMOJNPH.ProtoReflect.Descriptor instead.
func (*FMGCMMOJNPH) Descriptor() ([]byte, []int) {
	return file_FMGCMMOJNPH_proto_rawDescGZIP(), []int{0}
}

func (x *FMGCMMOJNPH) GetEMEFKPIMKOJ() map[uint32]*IPOFINFLFNE {
	if x != nil {
		return x.EMEFKPIMKOJ
	}
	return nil
}

var File_FMGCMMOJNPH_proto protoreflect.FileDescriptor

var file_FMGCMMOJNPH_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x4d, 0x47, 0x43, 0x4d, 0x4d, 0x4f, 0x4a, 0x4e, 0x50, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x50, 0x4f, 0x46, 0x49, 0x4e, 0x46, 0x4c, 0x46, 0x4e, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x46, 0x4d, 0x47, 0x43, 0x4d,
	0x4d, 0x4f, 0x4a, 0x4e, 0x50, 0x48, 0x12, 0x3f, 0x0a, 0x0b, 0x45, 0x4d, 0x45, 0x46, 0x4b, 0x50,
	0x49, 0x4d, 0x4b, 0x4f, 0x4a, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x4d,
	0x47, 0x43, 0x4d, 0x4d, 0x4f, 0x4a, 0x4e, 0x50, 0x48, 0x2e, 0x45, 0x4d, 0x45, 0x46, 0x4b, 0x50,
	0x49, 0x4d, 0x4b, 0x4f, 0x4a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x45, 0x4d, 0x45, 0x46,
	0x4b, 0x50, 0x49, 0x4d, 0x4b, 0x4f, 0x4a, 0x1a, 0x4c, 0x0a, 0x10, 0x45, 0x4d, 0x45, 0x46, 0x4b,
	0x50, 0x49, 0x4d, 0x4b, 0x4f, 0x4a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49,
	0x50, 0x4f, 0x46, 0x49, 0x4e, 0x46, 0x4c, 0x46, 0x4e, 0x45, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FMGCMMOJNPH_proto_rawDescOnce sync.Once
	file_FMGCMMOJNPH_proto_rawDescData = file_FMGCMMOJNPH_proto_rawDesc
)

func file_FMGCMMOJNPH_proto_rawDescGZIP() []byte {
	file_FMGCMMOJNPH_proto_rawDescOnce.Do(func() {
		file_FMGCMMOJNPH_proto_rawDescData = protoimpl.X.CompressGZIP(file_FMGCMMOJNPH_proto_rawDescData)
	})
	return file_FMGCMMOJNPH_proto_rawDescData
}

var file_FMGCMMOJNPH_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FMGCMMOJNPH_proto_goTypes = []interface{}{
	(*FMGCMMOJNPH)(nil), // 0: FMGCMMOJNPH
	nil,                 // 1: FMGCMMOJNPH.EMEFKPIMKOJEntry
	(*IPOFINFLFNE)(nil), // 2: IPOFINFLFNE
}
var file_FMGCMMOJNPH_proto_depIdxs = []int32{
	1, // 0: FMGCMMOJNPH.EMEFKPIMKOJ:type_name -> FMGCMMOJNPH.EMEFKPIMKOJEntry
	2, // 1: FMGCMMOJNPH.EMEFKPIMKOJEntry.value:type_name -> IPOFINFLFNE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FMGCMMOJNPH_proto_init() }
func file_FMGCMMOJNPH_proto_init() {
	if File_FMGCMMOJNPH_proto != nil {
		return
	}
	file_IPOFINFLFNE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FMGCMMOJNPH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FMGCMMOJNPH); i {
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
			RawDescriptor: file_FMGCMMOJNPH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FMGCMMOJNPH_proto_goTypes,
		DependencyIndexes: file_FMGCMMOJNPH_proto_depIdxs,
		MessageInfos:      file_FMGCMMOJNPH_proto_msgTypes,
	}.Build()
	File_FMGCMMOJNPH_proto = out.File
	file_FMGCMMOJNPH_proto_rawDesc = nil
	file_FMGCMMOJNPH_proto_goTypes = nil
	file_FMGCMMOJNPH_proto_depIdxs = nil
}