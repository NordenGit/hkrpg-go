// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BuyGoodsCsReq.proto

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

type BuyGoodsCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId    uint32 `protobuf:"varint,8,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	ItemId    uint32 `protobuf:"varint,15,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	GoodsNum  uint32 `protobuf:"varint,13,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num,omitempty"`
	GoodsId   uint32 `protobuf:"varint,3,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	MonsterId uint32 `protobuf:"varint,1,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
}

func (x *BuyGoodsCsReq) Reset() {
	*x = BuyGoodsCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BuyGoodsCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyGoodsCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyGoodsCsReq) ProtoMessage() {}

func (x *BuyGoodsCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_BuyGoodsCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyGoodsCsReq.ProtoReflect.Descriptor instead.
func (*BuyGoodsCsReq) Descriptor() ([]byte, []int) {
	return file_BuyGoodsCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *BuyGoodsCsReq) GetShopId() uint32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *BuyGoodsCsReq) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *BuyGoodsCsReq) GetGoodsNum() uint32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *BuyGoodsCsReq) GetGoodsId() uint32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *BuyGoodsCsReq) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

var File_BuyGoodsCsReq_proto protoreflect.FileDescriptor

var file_BuyGoodsCsReq_proto_rawDesc = []byte{
	0x0a, 0x13, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0d, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BuyGoodsCsReq_proto_rawDescOnce sync.Once
	file_BuyGoodsCsReq_proto_rawDescData = file_BuyGoodsCsReq_proto_rawDesc
)

func file_BuyGoodsCsReq_proto_rawDescGZIP() []byte {
	file_BuyGoodsCsReq_proto_rawDescOnce.Do(func() {
		file_BuyGoodsCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_BuyGoodsCsReq_proto_rawDescData)
	})
	return file_BuyGoodsCsReq_proto_rawDescData
}

var file_BuyGoodsCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BuyGoodsCsReq_proto_goTypes = []interface{}{
	(*BuyGoodsCsReq)(nil), // 0: BuyGoodsCsReq
}
var file_BuyGoodsCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BuyGoodsCsReq_proto_init() }
func file_BuyGoodsCsReq_proto_init() {
	if File_BuyGoodsCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BuyGoodsCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyGoodsCsReq); i {
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
			RawDescriptor: file_BuyGoodsCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BuyGoodsCsReq_proto_goTypes,
		DependencyIndexes: file_BuyGoodsCsReq_proto_depIdxs,
		MessageInfos:      file_BuyGoodsCsReq_proto_msgTypes,
	}.Build()
	File_BuyGoodsCsReq_proto = out.File
	file_BuyGoodsCsReq_proto_rawDesc = nil
	file_BuyGoodsCsReq_proto_goTypes = nil
	file_BuyGoodsCsReq_proto_depIdxs = nil
}
