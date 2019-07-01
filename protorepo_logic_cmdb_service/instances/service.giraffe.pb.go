// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package instances

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/easyops-cn/go-proto-giraffe"
	types "github.com/gogo/protobuf/types"
	giraffe_micro "github.com/easyops-cn/giraffe-micro"
	io "io"
	context "context"
	_ "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ = io.EOF
var _ context.Context
var _ giraffe_micro.Client

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = giraffe_micro.SupportPackageIsVersion3 // please upgrade the giraffe_micro package

// Client is the client API for instances service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Client interface {
	GetInstance(ctx context.Context, in *InstanceID) (*types.Struct, error)
	CreateInstance(ctx context.Context, in *Instance) (*InstanceID, error)
	DeleteInstance(ctx context.Context, in *InstanceID) (*types.Empty, error)
}

type client struct {
	c giraffe_micro.Client
}

func NewClient(c giraffe_micro.Client) Client {
	return &client{
		c: c,
	}
}

func (c *client) GetInstance(ctx context.Context, in *InstanceID) (*types.Struct, error) {
	out := new(types.Struct)
	err := c.c.Invoke(ctx, _GetInstanceContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateInstance(ctx context.Context, in *Instance) (*InstanceID, error) {
	out := new(InstanceID)
	err := c.c.Invoke(ctx, _CreateInstanceContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteInstance(ctx context.Context, in *InstanceID) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.c.Invoke(ctx, _DeleteInstanceContract, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service is the server API for instances service.
type Service interface {
	GetInstance(context.Context, *InstanceID) (*types.Struct, error)
	CreateInstance(context.Context, *Instance) (*InstanceID, error)
	DeleteInstance(context.Context, *InstanceID) (*types.Empty, error)
}

func _GetInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetInstance(ctx, req.(*InstanceID))
	}
}

func _CreateInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateInstance(ctx, req.(*Instance))
	}
}

func _DeleteInstanceEndpoint(s Service) giraffe_micro.UnaryEndpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteInstance(ctx, req.(*InstanceID))
	}
}

func RegisterService(s giraffe_micro.Server, srv Service) {
	s.RegisterUnaryEndpoint(_GetInstanceContract, _GetInstanceEndpoint(srv))
	s.RegisterUnaryEndpoint(_CreateInstanceContract, _CreateInstanceEndpoint(srv))
	s.RegisterUnaryEndpoint(_DeleteInstanceContract, _DeleteInstanceEndpoint(srv))
}

// API Contract
var _GetInstanceContract = &getInstanceContract{}

type getInstanceContract struct{}

func (*getInstanceContract) ServiceName() string          { return "instances.rpc" }
func (*getInstanceContract) MethodName() string           { return "GetInstance" }
func (*getInstanceContract) RequestMessage() interface{}  { return new(InstanceID) }
func (*getInstanceContract) ResponseMessage() interface{} { return new(InstanceID) }
func (*getInstanceContract) ContractName() string         { return "/instance.rpc/GetInstance" }
func (*getInstanceContract) ContractVersion() string      { return "V1.0" }
func (*getInstanceContract) Pattern() (string, string) {
	return "GET", "/object/:_object_id/instance/:instance_id"
}
func (*getInstanceContract) Body() string { return "" }

var _CreateInstanceContract = &createInstanceContract{}

type createInstanceContract struct{}

func (*createInstanceContract) ServiceName() string          { return "instances.rpc" }
func (*createInstanceContract) MethodName() string           { return "CreateInstance" }
func (*createInstanceContract) RequestMessage() interface{}  { return new(Instance) }
func (*createInstanceContract) ResponseMessage() interface{} { return new(Instance) }
func (*createInstanceContract) ContractName() string         { return "/instance.rpc/CreateInstance" }
func (*createInstanceContract) ContractVersion() string      { return "V1.0" }
func (*createInstanceContract) Pattern() (string, string) {
	return "POST", "/object/:object_id/instance"
}
func (*createInstanceContract) Body() string { return "data" }

var _DeleteInstanceContract = &deleteInstanceContract{}

type deleteInstanceContract struct{}

func (*deleteInstanceContract) ServiceName() string          { return "instances.rpc" }
func (*deleteInstanceContract) MethodName() string           { return "DeleteInstance" }
func (*deleteInstanceContract) RequestMessage() interface{}  { return new(InstanceID) }
func (*deleteInstanceContract) ResponseMessage() interface{} { return new(InstanceID) }
func (*deleteInstanceContract) ContractName() string         { return "/instance.rpc/DeleteInstance" }
func (*deleteInstanceContract) ContractVersion() string      { return "V1.0" }
func (*deleteInstanceContract) Pattern() (string, string) {
	return "DELETE", "/object/:_object_id/instance/:instance_id"
}
func (*deleteInstanceContract) Body() string { return "" }