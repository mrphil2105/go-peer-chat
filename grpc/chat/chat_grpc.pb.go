// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: grpc/chat/chat.proto

package chat

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	OpenChat(ctx context.Context, opts ...grpc.CallOption) (ChatService_OpenChatClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) OpenChat(ctx context.Context, opts ...grpc.CallOption) (ChatService_OpenChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/connect.ChatService/OpenChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceOpenChatClient{stream}
	return x, nil
}

type ChatService_OpenChatClient interface {
	Send(*Message) error
	CloseAndRecv() (*Null, error)
	grpc.ClientStream
}

type chatServiceOpenChatClient struct {
	grpc.ClientStream
}

func (x *chatServiceOpenChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceOpenChatClient) CloseAndRecv() (*Null, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Null)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	OpenChat(ChatService_OpenChatServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) OpenChat(ChatService_OpenChatServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenChat not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_OpenChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).OpenChat(&chatServiceOpenChatServer{stream})
}

type ChatService_OpenChatServer interface {
	SendAndClose(*Null) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceOpenChatServer struct {
	grpc.ServerStream
}

func (x *chatServiceOpenChatServer) SendAndClose(m *Null) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceOpenChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connect.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenChat",
			Handler:       _ChatService_OpenChat_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/chat/chat.proto",
}
