// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order/order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Orders service

func NewOrdersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Orders service

type OrdersService interface {
	// 查询总价
	Amount(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取订单列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取唯一订单
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建订单
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新订单
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除订单
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type ordersService struct {
	c    client.Client
	name string
}

func NewOrdersService(name string, c client.Client) OrdersService {
	return &ordersService{
		c:    c,
		name: name,
	}
}

func (c *ordersService) Amount(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.Amount", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersService) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Orders.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Orders service

type OrdersHandler interface {
	// 查询总价
	Amount(context.Context, *Request, *Response) error
	// 获取订单列表
	List(context.Context, *Request, *Response) error
	// 获取唯一订单
	Get(context.Context, *Request, *Response) error
	// 创建订单
	Create(context.Context, *Request, *Response) error
	// 更新订单
	Update(context.Context, *Request, *Response) error
	// 删除订单
	Delete(context.Context, *Request, *Response) error
}

func RegisterOrdersHandler(s server.Server, hdlr OrdersHandler, opts ...server.HandlerOption) error {
	type orders interface {
		Amount(ctx context.Context, in *Request, out *Response) error
		List(ctx context.Context, in *Request, out *Response) error
		Get(ctx context.Context, in *Request, out *Response) error
		Create(ctx context.Context, in *Request, out *Response) error
		Update(ctx context.Context, in *Request, out *Response) error
		Delete(ctx context.Context, in *Request, out *Response) error
	}
	type Orders struct {
		orders
	}
	h := &ordersHandler{hdlr}
	return s.Handle(s.NewHandler(&Orders{h}, opts...))
}

type ordersHandler struct {
	OrdersHandler
}

func (h *ordersHandler) Amount(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.Amount(ctx, in, out)
}

func (h *ordersHandler) List(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.List(ctx, in, out)
}

func (h *ordersHandler) Get(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.Get(ctx, in, out)
}

func (h *ordersHandler) Create(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.Create(ctx, in, out)
}

func (h *ordersHandler) Update(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.Update(ctx, in, out)
}

func (h *ordersHandler) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.OrdersHandler.Delete(ctx, in, out)
}
