// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KGMGBCPIMAB.proto

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

type KGMGBCPIMAB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OAPOFDEBJOO []*ABMNFJMIMHO `protobuf:"bytes,6,rep,name=OAPOFDEBJOO,proto3" json:"OAPOFDEBJOO,omitempty"`
	HOOFNKOBIAP IGNPCAPKDFC    `protobuf:"varint,2,opt,name=HOOFNKOBIAP,proto3,enum=IGNPCAPKDFC" json:"HOOFNKOBIAP,omitempty"`
	KDDCEOEHJDN []uint32       `protobuf:"varint,10,rep,packed,name=KDDCEOEHJDN,proto3" json:"KDDCEOEHJDN,omitempty"`
	IIGHCIFJDFG int32          `protobuf:"varint,13,opt,name=IIGHCIFJDFG,proto3" json:"IIGHCIFJDFG,omitempty"`
	HDCFMMNFKEP uint32         `protobuf:"varint,4,opt,name=HDCFMMNFKEP,proto3" json:"HDCFMMNFKEP,omitempty"`
	HMHPJOICOMC uint64         `protobuf:"varint,11,opt,name=HMHPJOICOMC,proto3" json:"HMHPJOICOMC,omitempty"`
	BLOEFMOAKMG *EAPBFAIFBAF   `protobuf:"bytes,9,opt,name=BLOEFMOAKMG,proto3" json:"BLOEFMOAKMG,omitempty"`
}

func (x *KGMGBCPIMAB) Reset() {
	*x = KGMGBCPIMAB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KGMGBCPIMAB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGMGBCPIMAB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGMGBCPIMAB) ProtoMessage() {}

func (x *KGMGBCPIMAB) ProtoReflect() protoreflect.Message {
	mi := &file_KGMGBCPIMAB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGMGBCPIMAB.ProtoReflect.Descriptor instead.
func (*KGMGBCPIMAB) Descriptor() ([]byte, []int) {
	return file_KGMGBCPIMAB_proto_rawDescGZIP(), []int{0}
}

func (x *KGMGBCPIMAB) GetOAPOFDEBJOO() []*ABMNFJMIMHO {
	if x != nil {
		return x.OAPOFDEBJOO
	}
	return nil
}

func (x *KGMGBCPIMAB) GetHOOFNKOBIAP() IGNPCAPKDFC {
	if x != nil {
		return x.HOOFNKOBIAP
	}
	return IGNPCAPKDFC_MATCH3_STATE_IDLE
}

func (x *KGMGBCPIMAB) GetKDDCEOEHJDN() []uint32 {
	if x != nil {
		return x.KDDCEOEHJDN
	}
	return nil
}

func (x *KGMGBCPIMAB) GetIIGHCIFJDFG() int32 {
	if x != nil {
		return x.IIGHCIFJDFG
	}
	return 0
}

func (x *KGMGBCPIMAB) GetHDCFMMNFKEP() uint32 {
	if x != nil {
		return x.HDCFMMNFKEP
	}
	return 0
}

func (x *KGMGBCPIMAB) GetHMHPJOICOMC() uint64 {
	if x != nil {
		return x.HMHPJOICOMC
	}
	return 0
}

func (x *KGMGBCPIMAB) GetBLOEFMOAKMG() *EAPBFAIFBAF {
	if x != nil {
		return x.BLOEFMOAKMG
	}
	return nil
}

var File_KGMGBCPIMAB_proto protoreflect.FileDescriptor

var file_KGMGBCPIMAB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x47, 0x4d, 0x47, 0x42, 0x43, 0x50, 0x49, 0x4d, 0x41, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x41, 0x50, 0x42, 0x46, 0x41, 0x49, 0x46, 0x42, 0x41, 0x46,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x47, 0x4e, 0x50, 0x43, 0x41, 0x50, 0x4b,
	0x44, 0x46, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x42, 0x4d, 0x4e, 0x46,
	0x4a, 0x4d, 0x49, 0x4d, 0x48, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a,
	0x0b, 0x4b, 0x47, 0x4d, 0x47, 0x42, 0x43, 0x50, 0x49, 0x4d, 0x41, 0x42, 0x12, 0x2e, 0x0a, 0x0b,
	0x4f, 0x41, 0x50, 0x4f, 0x46, 0x44, 0x45, 0x42, 0x4a, 0x4f, 0x4f, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x42, 0x4d, 0x4e, 0x46, 0x4a, 0x4d, 0x49, 0x4d, 0x48, 0x4f, 0x52,
	0x0b, 0x4f, 0x41, 0x50, 0x4f, 0x46, 0x44, 0x45, 0x42, 0x4a, 0x4f, 0x4f, 0x12, 0x2e, 0x0a, 0x0b,
	0x48, 0x4f, 0x4f, 0x46, 0x4e, 0x4b, 0x4f, 0x42, 0x49, 0x41, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x49, 0x47, 0x4e, 0x50, 0x43, 0x41, 0x50, 0x4b, 0x44, 0x46, 0x43, 0x52,
	0x0b, 0x48, 0x4f, 0x4f, 0x46, 0x4e, 0x4b, 0x4f, 0x42, 0x49, 0x41, 0x50, 0x12, 0x20, 0x0a, 0x0b,
	0x4b, 0x44, 0x44, 0x43, 0x45, 0x4f, 0x45, 0x48, 0x4a, 0x44, 0x4e, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0b, 0x4b, 0x44, 0x44, 0x43, 0x45, 0x4f, 0x45, 0x48, 0x4a, 0x44, 0x4e, 0x12, 0x20,
	0x0a, 0x0b, 0x49, 0x49, 0x47, 0x48, 0x43, 0x49, 0x46, 0x4a, 0x44, 0x46, 0x47, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x49, 0x49, 0x47, 0x48, 0x43, 0x49, 0x46, 0x4a, 0x44, 0x46, 0x47,
	0x12, 0x20, 0x0a, 0x0b, 0x48, 0x44, 0x43, 0x46, 0x4d, 0x4d, 0x4e, 0x46, 0x4b, 0x45, 0x50, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x44, 0x43, 0x46, 0x4d, 0x4d, 0x4e, 0x46, 0x4b,
	0x45, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4d, 0x48, 0x50, 0x4a, 0x4f, 0x49, 0x43, 0x4f, 0x4d,
	0x43, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x48, 0x4d, 0x48, 0x50, 0x4a, 0x4f, 0x49,
	0x43, 0x4f, 0x4d, 0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x42, 0x4c, 0x4f, 0x45, 0x46, 0x4d, 0x4f, 0x41,
	0x4b, 0x4d, 0x47, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x41, 0x50, 0x42,
	0x46, 0x41, 0x49, 0x46, 0x42, 0x41, 0x46, 0x52, 0x0b, 0x42, 0x4c, 0x4f, 0x45, 0x46, 0x4d, 0x4f,
	0x41, 0x4b, 0x4d, 0x47, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KGMGBCPIMAB_proto_rawDescOnce sync.Once
	file_KGMGBCPIMAB_proto_rawDescData = file_KGMGBCPIMAB_proto_rawDesc
)

func file_KGMGBCPIMAB_proto_rawDescGZIP() []byte {
	file_KGMGBCPIMAB_proto_rawDescOnce.Do(func() {
		file_KGMGBCPIMAB_proto_rawDescData = protoimpl.X.CompressGZIP(file_KGMGBCPIMAB_proto_rawDescData)
	})
	return file_KGMGBCPIMAB_proto_rawDescData
}

var file_KGMGBCPIMAB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KGMGBCPIMAB_proto_goTypes = []interface{}{
	(*KGMGBCPIMAB)(nil), // 0: KGMGBCPIMAB
	(*ABMNFJMIMHO)(nil), // 1: ABMNFJMIMHO
	(IGNPCAPKDFC)(0),    // 2: IGNPCAPKDFC
	(*EAPBFAIFBAF)(nil), // 3: EAPBFAIFBAF
}
var file_KGMGBCPIMAB_proto_depIdxs = []int32{
	1, // 0: KGMGBCPIMAB.OAPOFDEBJOO:type_name -> ABMNFJMIMHO
	2, // 1: KGMGBCPIMAB.HOOFNKOBIAP:type_name -> IGNPCAPKDFC
	3, // 2: KGMGBCPIMAB.BLOEFMOAKMG:type_name -> EAPBFAIFBAF
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_KGMGBCPIMAB_proto_init() }
func file_KGMGBCPIMAB_proto_init() {
	if File_KGMGBCPIMAB_proto != nil {
		return
	}
	file_EAPBFAIFBAF_proto_init()
	file_IGNPCAPKDFC_proto_init()
	file_ABMNFJMIMHO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KGMGBCPIMAB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGMGBCPIMAB); i {
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
			RawDescriptor: file_KGMGBCPIMAB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KGMGBCPIMAB_proto_goTypes,
		DependencyIndexes: file_KGMGBCPIMAB_proto_depIdxs,
		MessageInfos:      file_KGMGBCPIMAB_proto_msgTypes,
	}.Build()
	File_KGMGBCPIMAB_proto = out.File
	file_KGMGBCPIMAB_proto_rawDesc = nil
	file_KGMGBCPIMAB_proto_goTypes = nil
	file_KGMGBCPIMAB_proto_depIdxs = nil
}
