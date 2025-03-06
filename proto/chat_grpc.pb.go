// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/chat.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatService_SendMessage_FullMethodName       = "/proto.ChatService/SendMessage"
	ChatService_GetMessageHistory_FullMethodName = "/proto.ChatService/GetMessageHistory"
	ChatService_SendGroupMessage_FullMethodName  = "/proto.ChatService/SendGroupMessage"
	ChatService_GetGroupMessages_FullMethodName  = "/proto.ChatService/GetGroupMessages"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义聊天服务
type ChatServiceClient interface {
	// 发送私聊消息
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	// 获取用户聊天历史
	GetMessageHistory(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MessageResponse], error)
	// 发送群聊消息
	SendGroupMessage(ctx context.Context, in *GroupMessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	// 获取群聊历史消息
	GetGroupMessages(ctx context.Context, in *GroupHistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MessageResponse], error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, ChatService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetMessageHistory(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MessageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_GetMessageHistory_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HistoryRequest, MessageResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_GetMessageHistoryClient = grpc.ServerStreamingClient[MessageResponse]

func (c *chatServiceClient) SendGroupMessage(ctx context.Context, in *GroupMessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, ChatService_SendGroupMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetGroupMessages(ctx context.Context, in *GroupHistoryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MessageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_GetGroupMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GroupHistoryRequest, MessageResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_GetGroupMessagesClient = grpc.ServerStreamingClient[MessageResponse]

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
//
// 定义聊天服务
type ChatServiceServer interface {
	// 发送私聊消息
	SendMessage(context.Context, *MessageRequest) (*MessageResponse, error)
	// 获取用户聊天历史
	GetMessageHistory(*HistoryRequest, grpc.ServerStreamingServer[MessageResponse]) error
	// 发送群聊消息
	SendGroupMessage(context.Context, *GroupMessageRequest) (*MessageResponse, error)
	// 获取群聊历史消息
	GetGroupMessages(*GroupHistoryRequest, grpc.ServerStreamingServer[MessageResponse]) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) SendMessage(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) GetMessageHistory(*HistoryRequest, grpc.ServerStreamingServer[MessageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetMessageHistory not implemented")
}
func (UnimplementedChatServiceServer) SendGroupMessage(context.Context, *GroupMessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupMessage not implemented")
}
func (UnimplementedChatServiceServer) GetGroupMessages(*GroupHistoryRequest, grpc.ServerStreamingServer[MessageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetGroupMessages not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetMessageHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetMessageHistory(m, &grpc.GenericServerStream[HistoryRequest, MessageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_GetMessageHistoryServer = grpc.ServerStreamingServer[MessageResponse]

func _ChatService_SendGroupMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendGroupMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendGroupMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendGroupMessage(ctx, req.(*GroupMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetGroupMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GroupHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetGroupMessages(m, &grpc.GenericServerStream[GroupHistoryRequest, MessageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_GetGroupMessagesServer = grpc.ServerStreamingServer[MessageResponse]

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
		{
			MethodName: "SendGroupMessage",
			Handler:    _ChatService_SendGroupMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMessageHistory",
			Handler:       _ChatService_GetMessageHistory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetGroupMessages",
			Handler:       _ChatService_GetGroupMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/chat.proto",
}
