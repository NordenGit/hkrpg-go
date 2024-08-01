// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetVideoVersionKeyScRsp.proto

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

type GetVideoVersionKeyScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                  uint32          `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	VideoKeyInfoList         []*VideoKeyInfo `protobuf:"bytes,13,rep,name=video_key_info_list,json=videoKeyInfoList,proto3" json:"video_key_info_list,omitempty"`                          // 7
	ActivityVideoKeyInfoList []*VideoKeyInfo `protobuf:"bytes,7,rep,name=activity_video_key_info_list,json=activityVideoKeyInfoList,proto3" json:"activity_video_key_info_list,omitempty"` // 13
}

func (x *GetVideoVersionKeyScRsp) Reset() {
	*x = GetVideoVersionKeyScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetVideoVersionKeyScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoVersionKeyScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoVersionKeyScRsp) ProtoMessage() {}

func (x *GetVideoVersionKeyScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetVideoVersionKeyScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoVersionKeyScRsp.ProtoReflect.Descriptor instead.
func (*GetVideoVersionKeyScRsp) Descriptor() ([]byte, []int) {
	return file_GetVideoVersionKeyScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetVideoVersionKeyScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetVideoVersionKeyScRsp) GetVideoKeyInfoList() []*VideoKeyInfo {
	if x != nil {
		return x.VideoKeyInfoList
	}
	return nil
}

func (x *GetVideoVersionKeyScRsp) GetActivityVideoKeyInfoList() []*VideoKeyInfo {
	if x != nil {
		return x.ActivityVideoKeyInfoList
	}
	return nil
}

var File_GetVideoVersionKeyScRsp_proto protoreflect.FileDescriptor

var file_GetVideoVersionKeyScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3c, 0x0a, 0x13, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4b, 0x65,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4b, 0x65, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x1c, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x18, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4b, 0x65, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetVideoVersionKeyScRsp_proto_rawDescOnce sync.Once
	file_GetVideoVersionKeyScRsp_proto_rawDescData = file_GetVideoVersionKeyScRsp_proto_rawDesc
)

func file_GetVideoVersionKeyScRsp_proto_rawDescGZIP() []byte {
	file_GetVideoVersionKeyScRsp_proto_rawDescOnce.Do(func() {
		file_GetVideoVersionKeyScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetVideoVersionKeyScRsp_proto_rawDescData)
	})
	return file_GetVideoVersionKeyScRsp_proto_rawDescData
}

var file_GetVideoVersionKeyScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetVideoVersionKeyScRsp_proto_goTypes = []interface{}{
	(*GetVideoVersionKeyScRsp)(nil), // 0: GetVideoVersionKeyScRsp
	(*VideoKeyInfo)(nil),            // 1: VideoKeyInfo
}
var file_GetVideoVersionKeyScRsp_proto_depIdxs = []int32{
	1, // 0: GetVideoVersionKeyScRsp.video_key_info_list:type_name -> VideoKeyInfo
	1, // 1: GetVideoVersionKeyScRsp.activity_video_key_info_list:type_name -> VideoKeyInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetVideoVersionKeyScRsp_proto_init() }
func file_GetVideoVersionKeyScRsp_proto_init() {
	if File_GetVideoVersionKeyScRsp_proto != nil {
		return
	}
	file_VideoKeyInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetVideoVersionKeyScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoVersionKeyScRsp); i {
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
			RawDescriptor: file_GetVideoVersionKeyScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetVideoVersionKeyScRsp_proto_goTypes,
		DependencyIndexes: file_GetVideoVersionKeyScRsp_proto_depIdxs,
		MessageInfos:      file_GetVideoVersionKeyScRsp_proto_msgTypes,
	}.Build()
	File_GetVideoVersionKeyScRsp_proto = out.File
	file_GetVideoVersionKeyScRsp_proto_rawDesc = nil
	file_GetVideoVersionKeyScRsp_proto_goTypes = nil
	file_GetVideoVersionKeyScRsp_proto_depIdxs = nil
}
