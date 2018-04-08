// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/user-srv/proto/account/account.proto

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/user-srv/proto/account/account.proto

It has these top-level messages:
	User
	Session
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	UpdatePasswordRequest
	UpdatePasswordResponse
	SearchRequest
	SearchResponse
	ReadSessionRequest
	ReadSessionResponse
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
	ReadSession(ctx context.Context, in *ReadSessionRequest, opts ...client.CallOption) (*ReadSessionResponse, error)
}

type accountService struct {
	c           client.Client
	serviceName string
}

func AccountServiceClient(serviceName string, c client.Client) AccountService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "account"
	}
	return &accountService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.UpdatePassword", in)
	out := new(UpdatePasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) ReadSession(ctx context.Context, in *ReadSessionRequest, opts ...client.CallOption) (*ReadSessionResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.ReadSession", in)
	out := new(ReadSessionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *UpdatePasswordResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
	ReadSession(context.Context, *ReadSessionRequest, *ReadSessionResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Account{hdlr}, opts...))
}

type Account struct {
	AccountHandler
}

func (h *Account) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func (h *Account) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AccountHandler.Read(ctx, in, out)
}

func (h *Account) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *Account) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.AccountHandler.Delete(ctx, in, out)
}

func (h *Account) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.AccountHandler.Search(ctx, in, out)
}

func (h *Account) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *UpdatePasswordResponse) error {
	return h.AccountHandler.UpdatePassword(ctx, in, out)
}

func (h *Account) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *Account) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.AccountHandler.Logout(ctx, in, out)
}

func (h *Account) ReadSession(ctx context.Context, in *ReadSessionRequest, out *ReadSessionResponse) error {
	return h.AccountHandler.ReadSession(ctx, in, out)
}
