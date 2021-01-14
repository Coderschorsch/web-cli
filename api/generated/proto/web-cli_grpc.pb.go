// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package web_cli

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NewsFeedClient is the client API for NewsFeed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsFeedClient interface {
	ListFeeds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListFeedsResponse, error)
}

type newsFeedClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsFeedClient(cc grpc.ClientConnInterface) NewsFeedClient {
	return &newsFeedClient{cc}
}

func (c *newsFeedClient) ListFeeds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListFeedsResponse, error) {
	out := new(ListFeedsResponse)
	err := c.cc.Invoke(ctx, "/webcli.service.v1.NewsFeed/ListFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsFeedServer is the server API for NewsFeed service.
// All implementations must embed UnimplementedNewsFeedServer
// for forward compatibility
type NewsFeedServer interface {
	ListFeeds(context.Context, *emptypb.Empty) (*ListFeedsResponse, error)
	mustEmbedUnimplementedNewsFeedServer()
}

// UnimplementedNewsFeedServer must be embedded to have forward compatible implementations.
type UnimplementedNewsFeedServer struct {
}

func (UnimplementedNewsFeedServer) ListFeeds(context.Context, *emptypb.Empty) (*ListFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeeds not implemented")
}
func (UnimplementedNewsFeedServer) mustEmbedUnimplementedNewsFeedServer() {}

// UnsafeNewsFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsFeedServer will
// result in compilation errors.
type UnsafeNewsFeedServer interface {
	mustEmbedUnimplementedNewsFeedServer()
}

func RegisterNewsFeedServer(s grpc.ServiceRegistrar, srv NewsFeedServer) {
	s.RegisterService(&_NewsFeed_serviceDesc, srv)
}

func _NewsFeed_ListFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsFeedServer).ListFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webcli.service.v1.NewsFeed/ListFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsFeedServer).ListFeeds(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NewsFeed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "webcli.service.v1.NewsFeed",
	HandlerType: (*NewsFeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFeeds",
			Handler:    _NewsFeed_ListFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/web-cli.proto",
}