// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: social_graph_service.proto

package social_graph

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	users "proto/users"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SocialGraphServiceClient is the client API for SocialGraphService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialGraphServiceClient interface {
	RegisterUser(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*empty.Empty, error)
	RegisterBusinessUser(ctx context.Context, in *users.BusinessUser, opts ...grpc.CallOption) (*empty.Empty, error)
}

type socialGraphServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialGraphServiceClient(cc grpc.ClientConnInterface) SocialGraphServiceClient {
	return &socialGraphServiceClient{cc}
}

func (c *socialGraphServiceClient) RegisterUser(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/social_graph.SocialGraphService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialGraphServiceClient) RegisterBusinessUser(ctx context.Context, in *users.BusinessUser, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/social_graph.SocialGraphService/RegisterBusinessUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialGraphServiceServer is the server API for SocialGraphService service.
// All implementations must embed UnimplementedSocialGraphServiceServer
// for forward compatibility
type SocialGraphServiceServer interface {
	RegisterUser(context.Context, *users.User) (*empty.Empty, error)
	RegisterBusinessUser(context.Context, *users.BusinessUser) (*empty.Empty, error)
	mustEmbedUnimplementedSocialGraphServiceServer()
}

// UnimplementedSocialGraphServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSocialGraphServiceServer struct {
}

func (UnimplementedSocialGraphServiceServer) RegisterUser(context.Context, *users.User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedSocialGraphServiceServer) RegisterBusinessUser(context.Context, *users.BusinessUser) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterBusinessUser not implemented")
}
func (UnimplementedSocialGraphServiceServer) mustEmbedUnimplementedSocialGraphServiceServer() {}

// UnsafeSocialGraphServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialGraphServiceServer will
// result in compilation errors.
type UnsafeSocialGraphServiceServer interface {
	mustEmbedUnimplementedSocialGraphServiceServer()
}

func RegisterSocialGraphServiceServer(s grpc.ServiceRegistrar, srv SocialGraphServiceServer) {
	s.RegisterService(&SocialGraphService_ServiceDesc, srv)
}

func _SocialGraphService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(users.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialGraphServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social_graph.SocialGraphService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialGraphServiceServer).RegisterUser(ctx, req.(*users.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialGraphService_RegisterBusinessUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(users.BusinessUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialGraphServiceServer).RegisterBusinessUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social_graph.SocialGraphService/RegisterBusinessUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialGraphServiceServer).RegisterBusinessUser(ctx, req.(*users.BusinessUser))
	}
	return interceptor(ctx, in, info, handler)
}

// SocialGraphService_ServiceDesc is the grpc.ServiceDesc for SocialGraphService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocialGraphService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "social_graph.SocialGraphService",
	HandlerType: (*SocialGraphServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _SocialGraphService_RegisterUser_Handler,
		},
		{
			MethodName: "RegisterBusinessUser",
			Handler:    _SocialGraphService_RegisterBusinessUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "social_graph_service.proto",
}