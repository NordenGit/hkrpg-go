// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetPunkLordDataScRsp.proto

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

type GetPunkLordDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AEKKNLJAJOM          uint32   `protobuf:"varint,6,opt,name=AEKKNLJAJOM,proto3" json:"AEKKNLJAJOM,omitempty"`
	ADGEIBPOIAN          uint32   `protobuf:"varint,3,opt,name=ADGEIBPOIAN,proto3" json:"ADGEIBPOIAN,omitempty"`
	JKDMALJDFJB          uint32   `protobuf:"varint,12,opt,name=JKDMALJDFJB,proto3" json:"JKDMALJDFJB,omitempty"`
	DLINONOJJFL          uint32   `protobuf:"varint,1,opt,name=DLINONOJJFL,proto3" json:"DLINONOJJFL,omitempty"`
	OEBNIBEIPLC          uint32   `protobuf:"varint,10,opt,name=OEBNIBEIPLC,proto3" json:"OEBNIBEIPLC,omitempty"`
	PLOFECFNGOM          int64    `protobuf:"varint,8,opt,name=PLOFECFNGOM,proto3" json:"PLOFECFNGOM,omitempty"`
	TakenRewardLevelList []uint32 `protobuf:"varint,11,rep,packed,name=taken_reward_level_list,json=takenRewardLevelList,proto3" json:"taken_reward_level_list,omitempty"`
	Retcode              uint32   `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetPunkLordDataScRsp) Reset() {
	*x = GetPunkLordDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPunkLordDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPunkLordDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPunkLordDataScRsp) ProtoMessage() {}

func (x *GetPunkLordDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPunkLordDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPunkLordDataScRsp.ProtoReflect.Descriptor instead.
func (*GetPunkLordDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetPunkLordDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPunkLordDataScRsp) GetAEKKNLJAJOM() uint32 {
	if x != nil {
		return x.AEKKNLJAJOM
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetADGEIBPOIAN() uint32 {
	if x != nil {
		return x.ADGEIBPOIAN
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetJKDMALJDFJB() uint32 {
	if x != nil {
		return x.JKDMALJDFJB
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetDLINONOJJFL() uint32 {
	if x != nil {
		return x.DLINONOJJFL
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetOEBNIBEIPLC() uint32 {
	if x != nil {
		return x.OEBNIBEIPLC
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetPLOFECFNGOM() int64 {
	if x != nil {
		return x.PLOFECFNGOM
	}
	return 0
}

func (x *GetPunkLordDataScRsp) GetTakenRewardLevelList() []uint32 {
	if x != nil {
		return x.TakenRewardLevelList
	}
	return nil
}

func (x *GetPunkLordDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetPunkLordDataScRsp_proto protoreflect.FileDescriptor

var file_GetPunkLordDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x45, 0x4b, 0x4b, 0x4e, 0x4c, 0x4a,
	0x41, 0x4a, 0x4f, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x45, 0x4b, 0x4b,
	0x4e, 0x4c, 0x4a, 0x41, 0x4a, 0x4f, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x44, 0x47, 0x45, 0x49,
	0x42, 0x50, 0x4f, 0x49, 0x41, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x44,
	0x47, 0x45, 0x49, 0x42, 0x50, 0x4f, 0x49, 0x41, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4b, 0x44,
	0x4d, 0x41, 0x4c, 0x4a, 0x44, 0x46, 0x4a, 0x42, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4a, 0x4b, 0x44, 0x4d, 0x41, 0x4c, 0x4a, 0x44, 0x46, 0x4a, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x4c, 0x49, 0x4e, 0x4f, 0x4e, 0x4f, 0x4a, 0x4a, 0x46, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x44, 0x4c, 0x49, 0x4e, 0x4f, 0x4e, 0x4f, 0x4a, 0x4a, 0x46, 0x4c, 0x12, 0x20, 0x0a,
	0x0b, 0x4f, 0x45, 0x42, 0x4e, 0x49, 0x42, 0x45, 0x49, 0x50, 0x4c, 0x43, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x45, 0x42, 0x4e, 0x49, 0x42, 0x45, 0x49, 0x50, 0x4c, 0x43, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x4c, 0x4f, 0x46, 0x45, 0x43, 0x46, 0x4e, 0x47, 0x4f, 0x4d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x4c, 0x4f, 0x46, 0x45, 0x43, 0x46, 0x4e, 0x47, 0x4f,
	0x4d, 0x12, 0x35, 0x0a, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x14, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPunkLordDataScRsp_proto_rawDescOnce sync.Once
	file_GetPunkLordDataScRsp_proto_rawDescData = file_GetPunkLordDataScRsp_proto_rawDesc
)

func file_GetPunkLordDataScRsp_proto_rawDescGZIP() []byte {
	file_GetPunkLordDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetPunkLordDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPunkLordDataScRsp_proto_rawDescData)
	})
	return file_GetPunkLordDataScRsp_proto_rawDescData
}

var file_GetPunkLordDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPunkLordDataScRsp_proto_goTypes = []interface{}{
	(*GetPunkLordDataScRsp)(nil), // 0: GetPunkLordDataScRsp
}
var file_GetPunkLordDataScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetPunkLordDataScRsp_proto_init() }
func file_GetPunkLordDataScRsp_proto_init() {
	if File_GetPunkLordDataScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetPunkLordDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPunkLordDataScRsp); i {
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
			RawDescriptor: file_GetPunkLordDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPunkLordDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetPunkLordDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetPunkLordDataScRsp_proto_msgTypes,
	}.Build()
	File_GetPunkLordDataScRsp_proto = out.File
	file_GetPunkLordDataScRsp_proto_rawDesc = nil
	file_GetPunkLordDataScRsp_proto_goTypes = nil
	file_GetPunkLordDataScRsp_proto_depIdxs = nil
}
