// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ecommerce.proto

package ecommerce

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Customer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	BirthDate            string   `protobuf:"bytes,4,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba438465af3fc23, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Customer) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Customer) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

type Product struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PriceInCents         int32          `protobuf:"varint,4,opt,name=price_in_cents,json=priceInCents,proto3" json:"price_in_cents,omitempty"`
	DiscountValue        *DiscountValue `protobuf:"bytes,5,opt,name=discount_value,json=discountValue,proto3" json:"discount_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba438465af3fc23, []int{1}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPriceInCents() int32 {
	if m != nil {
		return m.PriceInCents
	}
	return 0
}

func (m *Product) GetDiscountValue() *DiscountValue {
	if m != nil {
		return m.DiscountValue
	}
	return nil
}

type DiscountValue struct {
	Pct                  float32  `protobuf:"fixed32,1,opt,name=pct,proto3" json:"pct,omitempty"`
	ValueInCents         int32    `protobuf:"varint,2,opt,name=value_in_cents,json=valueInCents,proto3" json:"value_in_cents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscountValue) Reset()         { *m = DiscountValue{} }
func (m *DiscountValue) String() string { return proto.CompactTextString(m) }
func (*DiscountValue) ProtoMessage()    {}
func (*DiscountValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba438465af3fc23, []int{2}
}

func (m *DiscountValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscountValue.Unmarshal(m, b)
}
func (m *DiscountValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscountValue.Marshal(b, m, deterministic)
}
func (m *DiscountValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscountValue.Merge(m, src)
}
func (m *DiscountValue) XXX_Size() int {
	return xxx_messageInfo_DiscountValue.Size(m)
}
func (m *DiscountValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscountValue.DiscardUnknown(m)
}

var xxx_messageInfo_DiscountValue proto.InternalMessageInfo

func (m *DiscountValue) GetPct() float32 {
	if m != nil {
		return m.Pct
	}
	return 0
}

func (m *DiscountValue) GetValueInCents() int32 {
	if m != nil {
		return m.ValueInCents
	}
	return 0
}

type DiscountRequest struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	Product              *Product  `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DiscountRequest) Reset()         { *m = DiscountRequest{} }
func (m *DiscountRequest) String() string { return proto.CompactTextString(m) }
func (*DiscountRequest) ProtoMessage()    {}
func (*DiscountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba438465af3fc23, []int{3}
}

func (m *DiscountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscountRequest.Unmarshal(m, b)
}
func (m *DiscountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscountRequest.Marshal(b, m, deterministic)
}
func (m *DiscountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscountRequest.Merge(m, src)
}
func (m *DiscountRequest) XXX_Size() int {
	return xxx_messageInfo_DiscountRequest.Size(m)
}
func (m *DiscountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscountRequest proto.InternalMessageInfo

func (m *DiscountRequest) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *DiscountRequest) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type DiscountResponse struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscountResponse) Reset()         { *m = DiscountResponse{} }
func (m *DiscountResponse) String() string { return proto.CompactTextString(m) }
func (*DiscountResponse) ProtoMessage()    {}
func (*DiscountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba438465af3fc23, []int{4}
}

func (m *DiscountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscountResponse.Unmarshal(m, b)
}
func (m *DiscountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscountResponse.Marshal(b, m, deterministic)
}
func (m *DiscountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscountResponse.Merge(m, src)
}
func (m *DiscountResponse) XXX_Size() int {
	return xxx_messageInfo_DiscountResponse.Size(m)
}
func (m *DiscountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscountResponse proto.InternalMessageInfo

func (m *DiscountResponse) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "ecommerce.Customer")
	proto.RegisterType((*Product)(nil), "ecommerce.Product")
	proto.RegisterType((*DiscountValue)(nil), "ecommerce.DiscountValue")
	proto.RegisterType((*DiscountRequest)(nil), "ecommerce.DiscountRequest")
	proto.RegisterType((*DiscountResponse)(nil), "ecommerce.DiscountResponse")
}

func init() { proto.RegisterFile("ecommerce.proto", fileDescriptor_4ba438465af3fc23) }

var fileDescriptor_4ba438465af3fc23 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x6d, 0x67, 0x5d, 0xfb, 0xd6, 0xfd, 0x21, 0x7a, 0x28, 0x1b, 0x83, 0x52, 0x3c, 0xec,
	0x20, 0x13, 0xea, 0x07, 0x50, 0xd9, 0x40, 0xf4, 0x20, 0x92, 0xc3, 0xae, 0xa5, 0x4b, 0x23, 0x06,
	0xda, 0x26, 0x26, 0xa9, 0xe0, 0x37, 0xf3, 0xe3, 0x49, 0xd3, 0x76, 0xab, 0x38, 0xf0, 0xb6, 0xfc,
	0x9e, 0x77, 0xcf, 0xf3, 0xe4, 0x4d, 0x61, 0x42, 0x09, 0x2f, 0x0a, 0x2a, 0x09, 0x5d, 0x09, 0xc9,
	0x35, 0x47, 0xde, 0x1e, 0x44, 0x15, 0xb8, 0xeb, 0x4a, 0x69, 0x5e, 0x50, 0x89, 0xc6, 0x60, 0xb3,
	0x2c, 0xb0, 0x42, 0x6b, 0xe9, 0x60, 0x9b, 0x65, 0x68, 0x01, 0xf0, 0xc6, 0xa4, 0xd2, 0x49, 0x99,
	0x16, 0x34, 0xb0, 0x43, 0x6b, 0xe9, 0x61, 0xcf, 0x90, 0x97, 0xb4, 0xa0, 0x68, 0x0e, 0x5e, 0x9e,
	0x76, 0xea, 0xc0, 0xa8, 0x6e, 0x0d, 0x8c, 0xb8, 0x00, 0xd8, 0x31, 0xa9, 0xdf, 0x93, 0x2c, 0xd5,
	0x34, 0x38, 0x6d, 0xfe, 0x6b, 0xc8, 0x26, 0xd5, 0x34, 0xfa, 0xb6, 0x60, 0xf8, 0x2a, 0x79, 0x56,
	0x11, 0xfd, 0x27, 0xf6, 0x12, 0x1c, 0xcd, 0x74, 0xde, 0x25, 0x36, 0x07, 0x14, 0x82, 0x9f, 0x51,
	0x45, 0x24, 0x13, 0x9a, 0xf1, 0xb2, 0xcd, 0xeb, 0x23, 0x74, 0x05, 0x63, 0x21, 0x19, 0xa1, 0x09,
	0x2b, 0x13, 0x42, 0x4b, 0xad, 0x4c, 0xac, 0x83, 0xcf, 0x0d, 0x7d, 0x2a, 0xd7, 0x35, 0x43, 0x77,
	0x30, 0xce, 0x98, 0x22, 0xbc, 0x2a, 0x75, 0xf2, 0x99, 0xe6, 0x15, 0x0d, 0x9c, 0xd0, 0x5a, 0xfa,
	0x71, 0xb0, 0x3a, 0x6c, 0x69, 0xd3, 0x0e, 0x6c, 0x6b, 0x1d, 0x8f, 0xb2, 0xfe, 0x31, 0x7a, 0x84,
	0xd1, 0x2f, 0x1d, 0x4d, 0x61, 0x20, 0x88, 0x36, 0x17, 0xb0, 0x71, 0xfd, 0xb3, 0x6e, 0x62, 0xac,
	0x0f, 0x4d, 0xec, 0xa6, 0x89, 0xa1, 0x6d, 0x93, 0x48, 0xc0, 0xa4, 0x33, 0xc2, 0xf4, 0xa3, 0xa2,
	0x4a, 0xa3, 0x1b, 0x70, 0x49, 0xfb, 0x1a, 0xc6, 0xcf, 0x8f, 0x2f, 0x7a, 0xb5, 0xba, 0x87, 0xc2,
	0xfb, 0x21, 0x74, 0x0d, 0x43, 0xd1, 0xac, 0xd1, 0x44, 0xf8, 0x31, 0xea, 0xcd, 0xb7, 0x0b, 0xc6,
	0xdd, 0x48, 0x74, 0x0f, 0xd3, 0x43, 0xa2, 0x12, 0xbc, 0x54, 0xb4, 0xef, 0x60, 0xfd, 0xeb, 0x10,
	0x6f, 0xc1, 0xed, 0x1c, 0xd0, 0x33, 0x8c, 0x1e, 0x84, 0xc8, 0xbf, 0xf6, 0x60, 0x76, 0x64, 0x85,
	0xed, 0xcd, 0x66, 0xf3, 0xa3, 0x5a, 0xd3, 0x21, 0x3a, 0xd9, 0x9d, 0x99, 0x0f, 0xf3, 0xf6, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x81, 0x96, 0x0e, 0x8b, 0xab, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscountClient is the client API for Discount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscountClient interface {
	ApplyDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error)
}

type discountClient struct {
	cc *grpc.ClientConn
}

func NewDiscountClient(cc *grpc.ClientConn) DiscountClient {
	return &discountClient{cc}
}

func (c *discountClient) ApplyDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error) {
	out := new(DiscountResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.Discount/ApplyDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountServer is the server API for Discount service.
type DiscountServer interface {
	ApplyDiscount(context.Context, *DiscountRequest) (*DiscountResponse, error)
}

func RegisterDiscountServer(s *grpc.Server, srv DiscountServer) {
	s.RegisterService(&_Discount_serviceDesc, srv)
}

func _Discount_ApplyDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServer).ApplyDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.Discount/ApplyDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServer).ApplyDiscount(ctx, req.(*DiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discount_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.Discount",
	HandlerType: (*DiscountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyDiscount",
			Handler:    _Discount_ApplyDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecommerce.proto",
}
