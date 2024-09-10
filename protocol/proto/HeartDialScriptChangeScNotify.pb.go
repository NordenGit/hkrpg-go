// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HeartDialScriptChangeScNotify.proto

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

type HeartDialScriptChangeScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BECODEANCGJ             []*FLIKGGHIKFG           `protobuf:"bytes,9,rep,name=BECODEANCGJ,proto3" json:"BECODEANCGJ,omitempty"`
	ChangedDialogueInfoList []*HeartDialDialogueInfo `protobuf:"bytes,11,rep,name=changed_dialogue_info_list,json=changedDialogueInfoList,proto3" json:"changed_dialogue_info_list,omitempty"`
	UnlockStatus            HeartDialUnlockStatus    `protobuf:"varint,7,opt,name=unlock_status,json=unlockStatus,proto3,enum=HeartDialUnlockStatus" json:"unlock_status,omitempty"`
	ChangedScriptInfoList   []*HeartDialScriptInfo   `protobuf:"bytes,15,rep,name=changed_script_info_list,json=changedScriptInfoList,proto3" json:"changed_script_info_list,omitempty"`
}

func (x *HeartDialScriptChangeScNotify) Reset() {
	*x = HeartDialScriptChangeScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HeartDialScriptChangeScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartDialScriptChangeScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartDialScriptChangeScNotify) ProtoMessage() {}

func (x *HeartDialScriptChangeScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HeartDialScriptChangeScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartDialScriptChangeScNotify.ProtoReflect.Descriptor instead.
func (*HeartDialScriptChangeScNotify) Descriptor() ([]byte, []int) {
	return file_HeartDialScriptChangeScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HeartDialScriptChangeScNotify) GetBECODEANCGJ() []*FLIKGGHIKFG {
	if x != nil {
		return x.BECODEANCGJ
	}
	return nil
}

func (x *HeartDialScriptChangeScNotify) GetChangedDialogueInfoList() []*HeartDialDialogueInfo {
	if x != nil {
		return x.ChangedDialogueInfoList
	}
	return nil
}

func (x *HeartDialScriptChangeScNotify) GetUnlockStatus() HeartDialUnlockStatus {
	if x != nil {
		return x.UnlockStatus
	}
	return HeartDialUnlockStatus_HEART_DIAL_UNLOCK_STATUS_LOCK
}

func (x *HeartDialScriptChangeScNotify) GetChangedScriptInfoList() []*HeartDialScriptInfo {
	if x != nil {
		return x.ChangedScriptInfoList
	}
	return nil
}

var File_HeartDialScriptChangeScNotify_proto protoreflect.FileDescriptor

var file_HeartDialScriptChangeScNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x4c, 0x49, 0x4b, 0x47, 0x47, 0x48, 0x49, 0x4b,
	0x46, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44,
	0x69, 0x61, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x44, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x02,
	0x0a, 0x1d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x2e, 0x0a, 0x0b, 0x42, 0x45, 0x43, 0x4f, 0x44, 0x45, 0x41, 0x4e, 0x43, 0x47, 0x4a, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x4c, 0x49, 0x4b, 0x47, 0x47, 0x48, 0x49, 0x4b,
	0x46, 0x47, 0x52, 0x0b, 0x42, 0x45, 0x43, 0x4f, 0x44, 0x45, 0x41, 0x4e, 0x43, 0x47, 0x4a, 0x12,
	0x53, 0x0a, 0x1a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x17, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4d, 0x0a, 0x18, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x15, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HeartDialScriptChangeScNotify_proto_rawDescOnce sync.Once
	file_HeartDialScriptChangeScNotify_proto_rawDescData = file_HeartDialScriptChangeScNotify_proto_rawDesc
)

func file_HeartDialScriptChangeScNotify_proto_rawDescGZIP() []byte {
	file_HeartDialScriptChangeScNotify_proto_rawDescOnce.Do(func() {
		file_HeartDialScriptChangeScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HeartDialScriptChangeScNotify_proto_rawDescData)
	})
	return file_HeartDialScriptChangeScNotify_proto_rawDescData
}

var file_HeartDialScriptChangeScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HeartDialScriptChangeScNotify_proto_goTypes = []interface{}{
	(*HeartDialScriptChangeScNotify)(nil), // 0: HeartDialScriptChangeScNotify
	(*FLIKGGHIKFG)(nil),                   // 1: FLIKGGHIKFG
	(*HeartDialDialogueInfo)(nil),         // 2: HeartDialDialogueInfo
	(HeartDialUnlockStatus)(0),            // 3: HeartDialUnlockStatus
	(*HeartDialScriptInfo)(nil),           // 4: HeartDialScriptInfo
}
var file_HeartDialScriptChangeScNotify_proto_depIdxs = []int32{
	1, // 0: HeartDialScriptChangeScNotify.BECODEANCGJ:type_name -> FLIKGGHIKFG
	2, // 1: HeartDialScriptChangeScNotify.changed_dialogue_info_list:type_name -> HeartDialDialogueInfo
	3, // 2: HeartDialScriptChangeScNotify.unlock_status:type_name -> HeartDialUnlockStatus
	4, // 3: HeartDialScriptChangeScNotify.changed_script_info_list:type_name -> HeartDialScriptInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_HeartDialScriptChangeScNotify_proto_init() }
func file_HeartDialScriptChangeScNotify_proto_init() {
	if File_HeartDialScriptChangeScNotify_proto != nil {
		return
	}
	file_FLIKGGHIKFG_proto_init()
	file_HeartDialScriptInfo_proto_init()
	file_HeartDialUnlockStatus_proto_init()
	file_HeartDialDialogueInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HeartDialScriptChangeScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartDialScriptChangeScNotify); i {
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
			RawDescriptor: file_HeartDialScriptChangeScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HeartDialScriptChangeScNotify_proto_goTypes,
		DependencyIndexes: file_HeartDialScriptChangeScNotify_proto_depIdxs,
		MessageInfos:      file_HeartDialScriptChangeScNotify_proto_msgTypes,
	}.Build()
	File_HeartDialScriptChangeScNotify_proto = out.File
	file_HeartDialScriptChangeScNotify_proto_rawDesc = nil
	file_HeartDialScriptChangeScNotify_proto_goTypes = nil
	file_HeartDialScriptChangeScNotify_proto_depIdxs = nil
}
