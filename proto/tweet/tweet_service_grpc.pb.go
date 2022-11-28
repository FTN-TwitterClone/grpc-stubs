// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: tweet_service.proto

package tweet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TweetServiceClient is the client API for TweetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TweetServiceClient interface {
	UpdateFeed(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tweetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTweetServiceClient(cc grpc.ClientConnInterface) TweetServiceClient {
	return &tweetServiceClient{cc}
}

func (c *tweetServiceClient) UpdateFeed(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/tweet.TweetService/UpdateFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetServiceServer is the server API for TweetService service.
// All implementations must embed UnimplementedTweetServiceServer
// for forward compatibility
type TweetServiceServer interface {
	UpdateFeed(context.Context, *User) (*emptypb.Empty, error)
	mustEmbedUnimplementedTweetServiceServer()
}

// UnimplementedTweetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTweetServiceServer struct {
}

func (UnimplementedTweetServiceServer) UpdateFeed(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeed not implemented")
}
func (UnimplementedTweetServiceServer) mustEmbedUnimplementedTweetServiceServer() {}

// UnsafeTweetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TweetServiceServer will
// result in compilation errors.
type UnsafeTweetServiceServer interface {
	mustEmbedUnimplementedTweetServiceServer()
}

func RegisterTweetServiceServer(s grpc.ServiceRegistrar, srv TweetServiceServer) {
	s.RegisterService(&TweetService_ServiceDesc, srv)
}

func _TweetService_UpdateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).UpdateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tweet.TweetService/UpdateFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).UpdateFeed(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// TweetService_ServiceDesc is the grpc.ServiceDesc for TweetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TweetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tweet.TweetService",
	HandlerType: (*TweetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateFeed",
			Handler:    _TweetService_UpdateFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tweet_service.proto",
}