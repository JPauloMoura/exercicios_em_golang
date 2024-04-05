// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: protobuf/user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HandlerUserClient is the client API for HandlerUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandlerUserClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type handlerUserClient struct {
	cc grpc.ClientConnInterface
}

func NewHandlerUserClient(cc grpc.ClientConnInterface) HandlerUserClient {
	return &handlerUserClient{cc}
}

func (c *handlerUserClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/HandlerUser/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandlerUserServer is the server API for HandlerUser service.
// All implementations must embed UnimplementedHandlerUserServer
// for forward compatibility
type HandlerUserServer interface {
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedHandlerUserServer()
}

// UnimplementedHandlerUserServer must be embedded to have forward compatible implementations.
type UnimplementedHandlerUserServer struct {
}

func (UnimplementedHandlerUserServer) Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedHandlerUserServer) mustEmbedUnimplementedHandlerUserServer() {}

// UnsafeHandlerUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandlerUserServer will
// result in compilation errors.
type UnsafeHandlerUserServer interface {
	mustEmbedUnimplementedHandlerUserServer()
}

func RegisterHandlerUserServer(s grpc.ServiceRegistrar, srv HandlerUserServer) {
	s.RegisterService(&HandlerUser_ServiceDesc, srv)
}

func _HandlerUser_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUser/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUserServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HandlerUser_ServiceDesc is the grpc.ServiceDesc for HandlerUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandlerUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HandlerUser",
	HandlerType: (*HandlerUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _HandlerUser_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/user.proto",
}
