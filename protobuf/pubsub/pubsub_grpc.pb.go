// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: protobuf/pubsub/pubsub.proto

package pubsub

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

const (
	PubSub_Subscribe_FullMethodName = "/task.PubSub/Subscribe"
	PubSub_Publish_FullMethodName   = "/task.PubSub/Publish"
)

// PubSubClient is the client API for PubSub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSubClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSub_SubscribeClient, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
}

type pubSubClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubClient(cc grpc.ClientConnInterface) PubSubClient {
	return &pubSubClient{cc}
}

func (c *pubSubClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSub_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubSub_ServiceDesc.Streams[0], PubSub_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSub_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type pubSubSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, PubSub_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSubServer is the server API for PubSub service.
// All implementations must embed UnimplementedPubSubServer
// for forward compatibility
type PubSubServer interface {
	Subscribe(*SubscribeRequest, PubSub_SubscribeServer) error
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	mustEmbedUnimplementedPubSubServer()
}

// UnimplementedPubSubServer must be embedded to have forward compatible implementations.
type UnimplementedPubSubServer struct {
}

func (UnimplementedPubSubServer) Subscribe(*SubscribeRequest, PubSub_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedPubSubServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPubSubServer) mustEmbedUnimplementedPubSubServer() {}

// UnsafePubSubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSubServer will
// result in compilation errors.
type UnsafePubSubServer interface {
	mustEmbedUnimplementedPubSubServer()
}

func RegisterPubSubServer(s grpc.ServiceRegistrar, srv PubSubServer) {
	s.RegisterService(&PubSub_ServiceDesc, srv)
}

func _PubSub_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServer).Subscribe(m, &pubSubSubscribeServer{stream})
}

type PubSub_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type pubSubSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PubSub_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSub_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PubSub_ServiceDesc is the grpc.ServiceDesc for PubSub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.PubSub",
	HandlerType: (*PubSubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubSub_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSub_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/pubsub/pubsub.proto",
}
