// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: accbase/srv/auth/proto/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
	Roles(ctx context.Context, in *RoleRequest, opts ...client.CallOption) (*RoleResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.UserInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Roles(ctx context.Context, in *RoleRequest, opts ...client.CallOption) (*RoleResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Roles", in)
	out := new(RoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	UserInfo(context.Context, *UserInfoRequest, *UserInfoResponse) error
	Roles(context.Context, *RoleRequest, *RoleResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		UserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error
		Roles(ctx context.Context, in *RoleRequest, out *RoleResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AuthHandler.Login(ctx, in, out)
}

func (h *authHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.AuthHandler.Register(ctx, in, out)
}

func (h *authHandler) UserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error {
	return h.AuthHandler.UserInfo(ctx, in, out)
}

func (h *authHandler) Roles(ctx context.Context, in *RoleRequest, out *RoleResponse) error {
	return h.AuthHandler.Roles(ctx, in, out)
}
