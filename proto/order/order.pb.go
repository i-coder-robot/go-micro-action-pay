// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/order/order.proto

package order

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Order struct {
	OrderId     int64  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo     string `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId      int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalPrice  int32  `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	PayStatus   int32  `protobuf:"varint,5,opt,name=pay_status,json=payStatus,proto3" json:"pay_status,omitempty"`
	PayType     int32  `protobuf:"varint,6,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"`
	PayTime     string `protobuf:"bytes,7,opt,name=pay_time,json=payTime,proto3" json:"pay_time,omitempty"`
	OrderStatus int32  `protobuf:"varint,8,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
	ExtraInfo   string `protobuf:"bytes,9,opt,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty"`
	UserAddress string `protobuf:"bytes,10,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`
	IsDeleted   int32  `protobuf:"varint,11,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	CreatedAt   string `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Order) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *Order) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Order) GetPayStatus() int32 {
	if m != nil {
		return m.PayStatus
	}
	return 0
}

func (m *Order) GetPayType() int32 {
	if m != nil {
		return m.PayType
	}
	return 0
}

func (m *Order) GetPayTime() string {
	if m != nil {
		return m.PayTime
	}
	return ""
}

func (m *Order) GetOrderStatus() int32 {
	if m != nil {
		return m.OrderStatus
	}
	return 0
}

func (m *Order) GetExtraInfo() string {
	if m != nil {
		return m.ExtraInfo
	}
	return ""
}

func (m *Order) GetUserAddress() string {
	if m != nil {
		return m.UserAddress
	}
	return ""
}

func (m *Order) GetIsDeleted() int32 {
	if m != nil {
		return m.IsDeleted
	}
	return 0
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort  string `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Where string `protobuf:"bytes,4,opt,name=where,proto3" json:"where,omitempty"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{1}
}
func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return m.Size()
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ListQuery) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

type Request struct {
	ListQuery *ListQuery `protobuf:"bytes,1,opt,name=list_query,json=listQuery,proto3" json:"list_query,omitempty"`
	Order     *Order     `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetListQuery() *ListQuery {
	if m != nil {
		return m.ListQuery
	}
	return nil
}

func (m *Request) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Response struct {
	Valid     bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Total     int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Order     *Order   `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	OrderList []*Order `protobuf:"bytes,4,rep,name=order_list,json=orderList,proto3" json:"order_list,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Response) GetOrderList() []*Order {
	if m != nil {
		return m.OrderList
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "order.Order")
	proto.RegisterType((*ListQuery)(nil), "order.ListQuery")
	proto.RegisterType((*Request)(nil), "order.Request")
	proto.RegisterType((*Response)(nil), "order.Response")
}

func init() { proto.RegisterFile("proto/order/order.proto", fileDescriptor_986e030a471601a2) }

var fileDescriptor_986e030a471601a2 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x8b, 0x13, 0x4f,
	0x10, 0xcd, 0xec, 0x24, 0x93, 0x4c, 0x4d, 0x7e, 0x3f, 0xa5, 0x11, 0xb6, 0x15, 0x76, 0x8c, 0x39,
	0x45, 0x02, 0xbb, 0x10, 0x3f, 0x41, 0x54, 0x90, 0x80, 0xf8, 0x67, 0xd4, 0xab, 0x43, 0xbb, 0x53,
	0xab, 0x0d, 0x93, 0xf4, 0x6c, 0x77, 0x8f, 0x3a, 0x9f, 0xc0, 0x9b, 0xe8, 0xb7, 0xf2, 0xb8, 0x47,
	0x8f, 0x92, 0x7c, 0x11, 0xe9, 0xea, 0x36, 0x7b, 0x10, 0x21, 0x97, 0xd0, 0xf5, 0xde, 0xab, 0xaa,
	0x37, 0xdd, 0x8f, 0xc0, 0x71, 0xa3, 0x95, 0x55, 0x67, 0x4a, 0x57, 0xa8, 0xfd, 0xef, 0x29, 0x21,
	0x6c, 0x40, 0xc5, 0xf4, 0x6b, 0x0c, 0x83, 0xe7, 0xee, 0xc4, 0x6e, 0xc3, 0x88, 0xa0, 0x52, 0x56,
	0x3c, 0x9a, 0x44, 0xb3, 0xb8, 0x18, 0x52, 0xbd, 0xaa, 0xae, 0xa9, 0x8d, 0xe2, 0x47, 0x93, 0x68,
	0x96, 0x06, 0xea, 0x99, 0x62, 0xc7, 0x30, 0x6c, 0x8d, 0x6f, 0x8a, 0xa9, 0x29, 0x71, 0xe5, 0xaa,
	0x62, 0x77, 0x21, 0xb3, 0xca, 0x8a, 0xba, 0x6c, 0xb4, 0x3c, 0x47, 0xde, 0x9f, 0x44, 0xb3, 0x41,
	0x01, 0x04, 0xbd, 0x70, 0x08, 0x3b, 0x01, 0x68, 0x44, 0x57, 0x1a, 0x2b, 0x6c, 0x6b, 0xf8, 0x80,
	0xf8, 0xb4, 0x11, 0xdd, 0x2b, 0x02, 0xdc, 0x4e, 0x47, 0xdb, 0xae, 0x41, 0x9e, 0x10, 0x39, 0x6c,
	0x44, 0xf7, 0xba, 0x6b, 0x70, 0x4f, 0xc9, 0x35, 0xf2, 0xa1, 0xb7, 0xe3, 0x28, 0xb9, 0x46, 0x76,
	0x0f, 0xc6, 0xde, 0x69, 0x18, 0x3b, 0xa2, 0xce, 0x8c, 0xb0, 0x30, 0xf8, 0x04, 0x00, 0x3f, 0x5b,
	0x2d, 0x4a, 0xb9, 0xb9, 0x50, 0x3c, 0xa5, 0xfe, 0x94, 0x90, 0xd5, 0xe6, 0x42, 0xb9, 0x09, 0xf4,
	0x41, 0xa2, 0xaa, 0x34, 0x1a, 0xc3, 0x81, 0x04, 0x99, 0xc3, 0x96, 0x1e, 0x72, 0x13, 0xa4, 0x29,
	0x2b, 0xac, 0xd1, 0x62, 0xc5, 0x33, 0xef, 0x5c, 0x9a, 0xc7, 0x1e, 0x70, 0xf4, 0xb9, 0x46, 0x61,
	0xb1, 0x2a, 0x85, 0xe5, 0x63, 0xbf, 0x20, 0x20, 0x4b, 0xeb, 0xe8, 0xb6, 0xa9, 0xfe, 0xd0, 0xff,
	0x79, 0x3a, 0x20, 0x4b, 0x3b, 0x2d, 0x21, 0x7d, 0x2a, 0x8d, 0x7d, 0xd9, 0xa2, 0xee, 0xd8, 0x2d,
	0x18, 0xd4, 0x72, 0x2d, 0x6d, 0x78, 0x10, 0x5f, 0x30, 0x06, 0xfd, 0x46, 0xbc, 0x47, 0x7a, 0x8a,
	0xb8, 0xa0, 0xb3, 0xc3, 0x8c, 0xd2, 0x96, 0x1e, 0x21, 0x2d, 0xe8, 0xec, 0xba, 0x3f, 0x7d, 0x40,
	0xed, 0x2f, 0x3f, 0x2d, 0x7c, 0x31, 0x7d, 0x0b, 0xc3, 0x02, 0x2f, 0x5b, 0x34, 0x96, 0x9d, 0x01,
	0xd4, 0xd2, 0xd8, 0xf2, 0xd2, 0x2d, 0xa3, 0x1d, 0xd9, 0xe2, 0xe6, 0xa9, 0x4f, 0xc9, 0xde, 0x44,
	0x91, 0xd6, 0x7b, 0x3f, 0x53, 0xf0, 0xb1, 0xa1, 0xd5, 0xd9, 0x62, 0x1c, 0xb4, 0x14, 0xa0, 0x22,
	0x24, 0xea, 0x4b, 0x04, 0xa3, 0x02, 0x4d, 0xa3, 0x36, 0x06, 0x9d, 0x85, 0x8f, 0xa2, 0x0e, 0x89,
	0x1a, 0x15, 0xbe, 0x70, 0x28, 0x05, 0x21, 0x7c, 0x81, 0x2f, 0xae, 0x87, 0xc7, 0xff, 0x1c, 0xce,
	0xe6, 0x00, 0xfe, 0x7d, 0x9d, 0x27, 0xde, 0x9f, 0xc4, 0x7f, 0x09, 0x53, 0x2a, 0x9c, 0xfb, 0xc5,
	0xf7, 0x23, 0x48, 0x08, 0x34, 0x6c, 0x0e, 0xc9, 0x72, 0xad, 0xda, 0x8d, 0x65, 0xff, 0x07, 0x75,
	0xb8, 0x83, 0x3b, 0x37, 0xf6, 0xb5, 0xb7, 0x3c, 0xed, 0xb1, 0xfb, 0xd0, 0x77, 0xfd, 0x87, 0x48,
	0x67, 0x10, 0x3f, 0xc1, 0x83, 0x94, 0x73, 0x48, 0x1e, 0x51, 0x06, 0x0e, 0x14, 0xbf, 0xa1, 0x44,
	0x1c, 0x28, 0xf6, 0xd1, 0x3b, 0x40, 0xfc, 0x90, 0xff, 0xd8, 0xe6, 0xd1, 0xd5, 0x36, 0x8f, 0x7e,
	0x6d, 0xf3, 0xe8, 0xdb, 0x2e, 0xef, 0x5d, 0xed, 0xf2, 0xde, 0xcf, 0x5d, 0xde, 0x7b, 0x97, 0xd0,
	0xff, 0xc2, 0x83, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0x74, 0xbe, 0xb4, 0x32, 0x04, 0x00,
	0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UpdatedAt) > 0 {
		i -= len(m.UpdatedAt)
		copy(dAtA[i:], m.UpdatedAt)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.UpdatedAt)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x62
	}
	if m.IsDeleted != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.IsDeleted))
		i--
		dAtA[i] = 0x58
	}
	if len(m.UserAddress) > 0 {
		i -= len(m.UserAddress)
		copy(dAtA[i:], m.UserAddress)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.UserAddress)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ExtraInfo) > 0 {
		i -= len(m.ExtraInfo)
		copy(dAtA[i:], m.ExtraInfo)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.ExtraInfo)))
		i--
		dAtA[i] = 0x4a
	}
	if m.OrderStatus != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.OrderStatus))
		i--
		dAtA[i] = 0x40
	}
	if len(m.PayTime) > 0 {
		i -= len(m.PayTime)
		copy(dAtA[i:], m.PayTime)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.PayTime)))
		i--
		dAtA[i] = 0x3a
	}
	if m.PayType != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.PayType))
		i--
		dAtA[i] = 0x30
	}
	if m.PayStatus != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.PayStatus))
		i--
		dAtA[i] = 0x28
	}
	if m.TotalPrice != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.TotalPrice))
		i--
		dAtA[i] = 0x20
	}
	if m.UserId != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OrderNo) > 0 {
		i -= len(m.OrderNo)
		copy(dAtA[i:], m.OrderNo)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.OrderNo)))
		i--
		dAtA[i] = 0x12
	}
	if m.OrderId != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.OrderId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Where) > 0 {
		i -= len(m.Where)
		copy(dAtA[i:], m.Where)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Where)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sort) > 0 {
		i -= len(m.Sort)
		copy(dAtA[i:], m.Sort)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Sort)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Page != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x10
	}
	if m.Limit != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Order != nil {
		{
			size, err := m.Order.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ListQuery != nil {
		{
			size, err := m.ListQuery.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OrderList) > 0 {
		for iNdEx := len(m.OrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOrder(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Order != nil {
		{
			size, err := m.Order.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Total != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderId != 0 {
		n += 1 + sovOrder(uint64(m.OrderId))
	}
	l = len(m.OrderNo)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.UserId != 0 {
		n += 1 + sovOrder(uint64(m.UserId))
	}
	if m.TotalPrice != 0 {
		n += 1 + sovOrder(uint64(m.TotalPrice))
	}
	if m.PayStatus != 0 {
		n += 1 + sovOrder(uint64(m.PayStatus))
	}
	if m.PayType != 0 {
		n += 1 + sovOrder(uint64(m.PayType))
	}
	l = len(m.PayTime)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.OrderStatus != 0 {
		n += 1 + sovOrder(uint64(m.OrderStatus))
	}
	l = len(m.ExtraInfo)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.UserAddress)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.IsDeleted != 0 {
		n += 1 + sovOrder(uint64(m.IsDeleted))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func (m *ListQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Limit != 0 {
		n += 1 + sovOrder(uint64(m.Limit))
	}
	if m.Page != 0 {
		n += 1 + sovOrder(uint64(m.Page))
	}
	l = len(m.Sort)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Where)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListQuery != nil {
		l = m.ListQuery.Size()
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Order != nil {
		l = m.Order.Size()
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valid {
		n += 2
	}
	if m.Total != 0 {
		n += 1 + sovOrder(uint64(m.Total))
	}
	if m.Order != nil {
		l = m.Order.Size()
		n += 1 + l + sovOrder(uint64(l))
	}
	if len(m.OrderList) > 0 {
		for _, e := range m.OrderList {
			l = e.Size()
			n += 1 + l + sovOrder(uint64(l))
		}
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			m.OrderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderNo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderNo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalPrice", wireType)
			}
			m.TotalPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalPrice |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayStatus", wireType)
			}
			m.PayStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PayStatus |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayType", wireType)
			}
			m.PayType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PayType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderStatus", wireType)
			}
			m.OrderStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderStatus |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsDeleted", wireType)
			}
			m.IsDeleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsDeleted |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Where", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Where = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListQuery", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ListQuery == nil {
				m.ListQuery = &ListQuery{}
			}
			if err := m.ListQuery.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Order == nil {
				m.Order = &Order{}
			}
			if err := m.Order.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Order == nil {
				m.Order = &Order{}
			}
			if err := m.Order.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderList = append(m.OrderList, &Order{})
			if err := m.OrderList[len(m.OrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)