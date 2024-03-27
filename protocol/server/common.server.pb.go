// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: common.server.proto

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

type GmGive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerUid uint32 `protobuf:"varint,1,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
	ItemId    uint32 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemCount uint32 `protobuf:"varint,3,opt,name=item_count,json=itemCount,proto3" json:"item_count,omitempty"`
	GiveAll   bool   `protobuf:"varint,4,opt,name=give_all,json=giveAll,proto3" json:"give_all,omitempty"`
}

func (x *GmGive) Reset() {
	*x = GmGive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GmGive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GmGive) ProtoMessage() {}

func (x *GmGive) ProtoReflect() protoreflect.Message {
	mi := &file_common_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GmGive.ProtoReflect.Descriptor instead.
func (*GmGive) Descriptor() ([]byte, []int) {
	return file_common_server_proto_rawDescGZIP(), []int{0}
}

func (x *GmGive) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

func (x *GmGive) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GmGive) GetItemCount() uint32 {
	if x != nil {
		return x.ItemCount
	}
	return 0
}

func (x *GmGive) GetGiveAll() bool {
	if x != nil {
		return x.GiveAll
	}
	return false
}

type GmWorldLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerUid  uint32 `protobuf:"varint,1,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
	WorldLevel uint32 `protobuf:"varint,2,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
}

func (x *GmWorldLevel) Reset() {
	*x = GmWorldLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GmWorldLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GmWorldLevel) ProtoMessage() {}

func (x *GmWorldLevel) ProtoReflect() protoreflect.Message {
	mi := &file_common_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GmWorldLevel.ProtoReflect.Descriptor instead.
func (*GmWorldLevel) Descriptor() ([]byte, []int) {
	return file_common_server_proto_rawDescGZIP(), []int{1}
}

func (x *GmWorldLevel) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

func (x *GmWorldLevel) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

var File_common_server_proto protoreflect.FileDescriptor

var file_common_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x06, 0x47, 0x6d, 0x47, 0x69, 0x76, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x74, 0x65,
	0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x69, 0x76, 0x65, 0x5f, 0x61,
	0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x69, 0x76, 0x65, 0x41, 0x6c,
	0x6c, 0x22, 0x4e, 0x0a, 0x0c, 0x47, 0x6d, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_server_proto_rawDescOnce sync.Once
	file_common_server_proto_rawDescData = file_common_server_proto_rawDesc
)

func file_common_server_proto_rawDescGZIP() []byte {
	file_common_server_proto_rawDescOnce.Do(func() {
		file_common_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_server_proto_rawDescData)
	})
	return file_common_server_proto_rawDescData
}

var file_common_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_server_proto_goTypes = []interface{}{
	(*GmGive)(nil),       // 0: GmGive
	(*GmWorldLevel)(nil), // 1: GmWorldLevel
}
var file_common_server_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_server_proto_init() }
func file_common_server_proto_init() {
	if File_common_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GmGive); i {
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
		file_common_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GmWorldLevel); i {
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
			RawDescriptor: file_common_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_server_proto_goTypes,
		DependencyIndexes: file_common_server_proto_depIdxs,
		MessageInfos:      file_common_server_proto_msgTypes,
	}.Build()
	File_common_server_proto = out.File
	file_common_server_proto_rawDesc = nil
	file_common_server_proto_goTypes = nil
	file_common_server_proto_depIdxs = nil
}
