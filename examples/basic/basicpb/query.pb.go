// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: query.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

type Cashier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Cashier) Reset() {
	*x = Cashier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cashier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cashier) ProtoMessage() {}

func (x *Cashier) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cashier.ProtoReflect.Descriptor instead.
func (*Cashier) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{2}
}

func (x *Cashier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cashier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Products  []*Product             `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	Cashier   *Cashier               `protobuf:"bytes,5,opt,name=cashier,proto3" json:"cashier,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Order) GetCashier() *Cashier {
	if x != nil {
		return x.Cashier
	}
	return nil
}

var File_query_proto protoreflect.FileDescriptor

var file_query_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x66, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xc2, 0x3e, 0x04, 0x22, 0x02, 0x08, 0x1a, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xc2, 0x3e, 0x03, 0x0a, 0x01, 0x01, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0xc2, 0x3e, 0x1d, 0x1a, 0x1b, 0x50, 0x52, 0x49,
	0x4d, 0x41, 0x52, 0x59, 0x20, 0x4b, 0x45, 0x59, 0x20, 0x28, 0x69, 0x64, 0x2c, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x22, 0x4d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xc2, 0x3e, 0x04, 0x22, 0x02, 0x08, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xc2, 0x3e, 0x03, 0x0a, 0x01, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x07, 0x43, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xc2, 0x3e, 0x04, 0x22, 0x02, 0x08, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xe0, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xc2, 0x3e, 0x04, 0x22, 0x02, 0x08, 0x1a, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xc2, 0x3e, 0x07, 0x0a, 0x01, 0x00, 0x22, 0x02, 0x08, 0x1a, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xc2, 0x3e, 0x03, 0x0a, 0x01, 0x00, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x05, 0xc2, 0x3e, 0x02, 0x10, 0x03, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x42, 0x05, 0xc2, 0x3e, 0x02, 0x10, 0x02, 0x52, 0x07, 0x63, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x32, 0xcf, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x0c, 0x2e, 0x67, 0x6f, 0x70, 0x67, 0x78, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0d,
	0x2e, 0x67, 0x6f, 0x70, 0x67, 0x78, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1a, 0xc2,
	0x3e, 0x17, 0x0a, 0x15, 0x0a, 0x13, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20, 0x2a, 0x20, 0x46,
	0x52, 0x4f, 0x4d, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x79, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0c, 0x2e, 0x67, 0x6f,
	0x70, 0x67, 0x78, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x67, 0x6f, 0x70, 0x67,
	0x78, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4a, 0xc2, 0x3e, 0x47, 0x12, 0x45, 0x0a,
	0x10, 0x12, 0x0e, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x69,
	0x64, 0x0a, 0x2f, 0x0a, 0x2d, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20, 0x2a, 0x20, 0x46, 0x52,
	0x4f, 0x4d, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x20, 0x57, 0x48, 0x45, 0x52, 0x45, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x20, 0x3d, 0x20, 0x24, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x10, 0x01, 0x42, 0x89, 0x01, 0xc2, 0x3e, 0x7d, 0xc2, 0x3e, 0x6a, 0x0a, 0x42, 0x0a,
	0x0e, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12,
	0x30, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20, 0x2a, 0x20, 0x46, 0x52, 0x4f, 0x4d, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x20, 0x57, 0x48, 0x45, 0x52, 0x45, 0x20, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x24, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x0a, 0x24, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x13, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20, 0x2a, 0x20, 0x46, 0x52, 0x4f,
	0x4d, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0xca, 0x3e, 0x06, 0x76, 0x31, 0x2e, 0x30, 0x2e, 0x30,
	0xd2, 0x3e, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_proto_rawDescOnce sync.Once
	file_query_proto_rawDescData = file_query_proto_rawDesc
)

func file_query_proto_rawDescGZIP() []byte {
	file_query_proto_rawDescOnce.Do(func() {
		file_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_proto_rawDescData)
	})
	return file_query_proto_rawDescData
}

var file_query_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_query_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: User
	(*Product)(nil),               // 1: Product
	(*Cashier)(nil),               // 2: Cashier
	(*Order)(nil),                 // 3: Order
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*Query)(nil),                 // 5: gopgx.Query
	(*Result)(nil),                // 6: gopgx.Result
}
var file_query_proto_depIdxs = []int32{
	4, // 0: Order.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: Order.products:type_name -> Product
	2, // 2: Order.cashier:type_name -> Cashier
	5, // 3: QueryService.ListUsers:input_type -> gopgx.Query
	5, // 4: QueryService.ListUserOrders:input_type -> gopgx.Query
	6, // 5: QueryService.ListUsers:output_type -> gopgx.Result
	6, // 6: QueryService.ListUserOrders:output_type -> gopgx.Result
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_query_proto_init() }
func file_query_proto_init() {
	if File_query_proto != nil {
		return
	}
	file_descriptor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cashier); i {
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
		file_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_proto_goTypes,
		DependencyIndexes: file_query_proto_depIdxs,
		MessageInfos:      file_query_proto_msgTypes,
	}.Build()
	File_query_proto = out.File
	file_query_proto_rawDesc = nil
	file_query_proto_goTypes = nil
	file_query_proto_depIdxs = nil
}