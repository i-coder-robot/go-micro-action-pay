// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/config/config.proto

package config

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

// Api Endpoints for Configs service

func NewConfigsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Configs service

type ConfigsService interface {
	// 通过token自己更新支付数据
	SelfUpdate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取配置列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取配置
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建配置
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新配置
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除配置
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type configsService struct {
	c    client.Client
	name string
}

func NewConfigsService(name string, c client.Client) ConfigsService {
	return &configsService{
		c:    c,
		name: name,
	}
}

func (c *configsService) SelfUpdate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Configs.SelfUpdate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsService) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Configs.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsService) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Configs.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Configs.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsService) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Configs.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsService) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Configs.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Configs service

type ConfigsHandler interface {
	// 通过token自己更新支付数据
	SelfUpdate(context.Context, *Request, *Response) error
	// 获取配置列表
	List(context.Context, *Request, *Response) error
	// 根据 唯一 获取配置
	Get(context.Context, *Request, *Response) error
	// 创建配置
	Create(context.Context, *Request, *Response) error
	// 更新配置
	Update(context.Context, *Request, *Response) error
	// 删除配置
	Delete(context.Context, *Request, *Response) error
}

func RegisterConfigsHandler(s server.Server, hdlr ConfigsHandler, opts ...server.HandlerOption) error {
	type configs interface {
		SelfUpdate(ctx context.Context, in *Request, out *Response) error
		List(ctx context.Context, in *Request, out *Response) error
		Get(ctx context.Context, in *Request, out *Response) error
		Create(ctx context.Context, in *Request, out *Response) error
		Update(ctx context.Context, in *Request, out *Response) error
		Delete(ctx context.Context, in *Request, out *Response) error
	}
	type Configs struct {
		configs
	}
	h := &configsHandler{hdlr}
	return s.Handle(s.NewHandler(&Configs{h}, opts...))
}

type configsHandler struct {
	ConfigsHandler
}

func (h *configsHandler) SelfUpdate(ctx context.Context, in *Request, out *Response) error {
	return h.ConfigsHandler.SelfUpdate(ctx, in, out)
}

func (h *configsHandler) List(ctx context.Context, in *Request, out *Response) error {
	return h.ConfigsHandler.List(ctx, in, out)
}

func (h *configsHandler) Get(ctx context.Context, in *Request, out *Response) error {
	return h.ConfigsHandler.Get(ctx, in, out)
}

func (h *configsHandler) Create(ctx context.Context, in *Request, out *Response) error {
	return h.ConfigsHandler.Create(ctx, in, out)
}

func (h *configsHandler) Update(ctx context.Context, in *Request, out *Response) error {
	return h.ConfigsHandler.Update(ctx, in, out)
}

func (h *configsHandler) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.ConfigsHandler.Delete(ctx, in, out)
}
