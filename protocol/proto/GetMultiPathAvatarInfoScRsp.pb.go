// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMultiPathAvatarInfoScRsp.proto

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

type GetMultiPathAvatarInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                 uint32                         `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MultiPathAvatarInfoList []*MultiPathAvatarInfo         `protobuf:"bytes,2,rep,name=multi_path_avatar_info_list,json=multiPathAvatarInfoList,proto3" json:"multi_path_avatar_info_list,omitempty"`
	BasicTypeIdList         []uint32                       `protobuf:"varint,4,rep,packed,name=basic_type_id_list,json=basicTypeIdList,proto3" json:"basic_type_id_list,omitempty"`
	CurAvatarPath           map[uint32]MultiPathAvatarType `protobuf:"bytes,13,rep,name=cur_avatar_path,json=curAvatarPath,proto3" json:"cur_avatar_path,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=MultiPathAvatarType"`
}

func (x *GetMultiPathAvatarInfoScRsp) Reset() {
	*x = GetMultiPathAvatarInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMultiPathAvatarInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultiPathAvatarInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultiPathAvatarInfoScRsp) ProtoMessage() {}

func (x *GetMultiPathAvatarInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMultiPathAvatarInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultiPathAvatarInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetMultiPathAvatarInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetMultiPathAvatarInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMultiPathAvatarInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMultiPathAvatarInfoScRsp) GetMultiPathAvatarInfoList() []*MultiPathAvatarInfo {
	if x != nil {
		return x.MultiPathAvatarInfoList
	}
	return nil
}

func (x *GetMultiPathAvatarInfoScRsp) GetBasicTypeIdList() []uint32 {
	if x != nil {
		return x.BasicTypeIdList
	}
	return nil
}

func (x *GetMultiPathAvatarInfoScRsp) GetCurAvatarPath() map[uint32]MultiPathAvatarType {
	if x != nil {
		return x.CurAvatarPath
	}
	return nil
}

var File_GetMultiPathAvatarInfoScRsp_proto protoreflect.FileDescriptor

var file_GetMultiPathAvatarInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x1b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x50, 0x61, 0x74, 0x68, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x17,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x43, 0x75, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x63, 0x75, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x56, 0x0a,
	0x12, 0x43, 0x75, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMultiPathAvatarInfoScRsp_proto_rawDescOnce sync.Once
	file_GetMultiPathAvatarInfoScRsp_proto_rawDescData = file_GetMultiPathAvatarInfoScRsp_proto_rawDesc
)

func file_GetMultiPathAvatarInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetMultiPathAvatarInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetMultiPathAvatarInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMultiPathAvatarInfoScRsp_proto_rawDescData)
	})
	return file_GetMultiPathAvatarInfoScRsp_proto_rawDescData
}

var file_GetMultiPathAvatarInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetMultiPathAvatarInfoScRsp_proto_goTypes = []interface{}{
	(*GetMultiPathAvatarInfoScRsp)(nil), // 0: GetMultiPathAvatarInfoScRsp
	nil,                                 // 1: GetMultiPathAvatarInfoScRsp.CurAvatarPathEntry
	(*MultiPathAvatarInfo)(nil),         // 2: MultiPathAvatarInfo
	(MultiPathAvatarType)(0),            // 3: MultiPathAvatarType
}
var file_GetMultiPathAvatarInfoScRsp_proto_depIdxs = []int32{
	2, // 0: GetMultiPathAvatarInfoScRsp.multi_path_avatar_info_list:type_name -> MultiPathAvatarInfo
	1, // 1: GetMultiPathAvatarInfoScRsp.cur_avatar_path:type_name -> GetMultiPathAvatarInfoScRsp.CurAvatarPathEntry
	3, // 2: GetMultiPathAvatarInfoScRsp.CurAvatarPathEntry.value:type_name -> MultiPathAvatarType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetMultiPathAvatarInfoScRsp_proto_init() }
func file_GetMultiPathAvatarInfoScRsp_proto_init() {
	if File_GetMultiPathAvatarInfoScRsp_proto != nil {
		return
	}
	file_MultiPathAvatarInfo_proto_init()
	file_MultiPathAvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMultiPathAvatarInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultiPathAvatarInfoScRsp); i {
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
			RawDescriptor: file_GetMultiPathAvatarInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMultiPathAvatarInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetMultiPathAvatarInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetMultiPathAvatarInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetMultiPathAvatarInfoScRsp_proto = out.File
	file_GetMultiPathAvatarInfoScRsp_proto_rawDesc = nil
	file_GetMultiPathAvatarInfoScRsp_proto_goTypes = nil
	file_GetMultiPathAvatarInfoScRsp_proto_depIdxs = nil
}
