// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HandleFriendScRsp.proto

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

type HandleFriendScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    uint32            `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IsAccept   bool              `protobuf:"varint,11,opt,name=is_accept,json=isAccept,proto3" json:"is_accept,omitempty"`
	FriendInfo *FriendSimpleInfo `protobuf:"bytes,6,opt,name=friend_info,json=friendInfo,proto3" json:"friend_info,omitempty"`
	Uid        uint32            `protobuf:"varint,7,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *HandleFriendScRsp) Reset() {
	*x = HandleFriendScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HandleFriendScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleFriendScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleFriendScRsp) ProtoMessage() {}

func (x *HandleFriendScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_HandleFriendScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleFriendScRsp.ProtoReflect.Descriptor instead.
func (*HandleFriendScRsp) Descriptor() ([]byte, []int) {
	return file_HandleFriendScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *HandleFriendScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *HandleFriendScRsp) GetIsAccept() bool {
	if x != nil {
		return x.IsAccept
	}
	return false
}

func (x *HandleFriendScRsp) GetFriendInfo() *FriendSimpleInfo {
	if x != nil {
		return x.FriendInfo
	}
	return nil
}

func (x *HandleFriendScRsp) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_HandleFriendScRsp_proto protoreflect.FileDescriptor

var file_HandleFriendScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x32,
	0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HandleFriendScRsp_proto_rawDescOnce sync.Once
	file_HandleFriendScRsp_proto_rawDescData = file_HandleFriendScRsp_proto_rawDesc
)

func file_HandleFriendScRsp_proto_rawDescGZIP() []byte {
	file_HandleFriendScRsp_proto_rawDescOnce.Do(func() {
		file_HandleFriendScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_HandleFriendScRsp_proto_rawDescData)
	})
	return file_HandleFriendScRsp_proto_rawDescData
}

var file_HandleFriendScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HandleFriendScRsp_proto_goTypes = []interface{}{
	(*HandleFriendScRsp)(nil), // 0: HandleFriendScRsp
	(*FriendSimpleInfo)(nil),  // 1: FriendSimpleInfo
}
var file_HandleFriendScRsp_proto_depIdxs = []int32{
	1, // 0: HandleFriendScRsp.friend_info:type_name -> FriendSimpleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HandleFriendScRsp_proto_init() }
func file_HandleFriendScRsp_proto_init() {
	if File_HandleFriendScRsp_proto != nil {
		return
	}
	file_FriendSimpleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HandleFriendScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleFriendScRsp); i {
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
			RawDescriptor: file_HandleFriendScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HandleFriendScRsp_proto_goTypes,
		DependencyIndexes: file_HandleFriendScRsp_proto_depIdxs,
		MessageInfos:      file_HandleFriendScRsp_proto_msgTypes,
	}.Build()
	File_HandleFriendScRsp_proto = out.File
	file_HandleFriendScRsp_proto_rawDesc = nil
	file_HandleFriendScRsp_proto_goTypes = nil
	file_HandleFriendScRsp_proto_depIdxs = nil
}
