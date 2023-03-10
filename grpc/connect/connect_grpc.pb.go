// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: grpc/connect/connect.proto

package connect

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

// ConnectServiceClient is the client API for ConnectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectServiceClient interface {
	JoinNetwork(ctx context.Context, in *PeerJoin, opts ...grpc.CallOption) (*ConnectedTo, error)
	LeaveNetwork(ctx context.Context, in *PeerLeave, opts ...grpc.CallOption) (*Void, error)
}

type connectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectServiceClient(cc grpc.ClientConnInterface) ConnectServiceClient {
	return &connectServiceClient{cc}
}

func (c *connectServiceClient) JoinNetwork(ctx context.Context, in *PeerJoin, opts ...grpc.CallOption) (*ConnectedTo, error) {
	out := new(ConnectedTo)
	err := c.cc.Invoke(ctx, "/connect.ConnectService/JoinNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectServiceClient) LeaveNetwork(ctx context.Context, in *PeerLeave, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/connect.ConnectService/LeaveNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServiceServer is the server API for ConnectService service.
// All implementations must embed UnimplementedConnectServiceServer
// for forward compatibility
type ConnectServiceServer interface {
	JoinNetwork(context.Context, *PeerJoin) (*ConnectedTo, error)
	LeaveNetwork(context.Context, *PeerLeave) (*Void, error)
	mustEmbedUnimplementedConnectServiceServer()
}

// UnimplementedConnectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectServiceServer struct {
}

func (UnimplementedConnectServiceServer) JoinNetwork(context.Context, *PeerJoin) (*ConnectedTo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinNetwork not implemented")
}
func (UnimplementedConnectServiceServer) LeaveNetwork(context.Context, *PeerLeave) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveNetwork not implemented")
}
func (UnimplementedConnectServiceServer) mustEmbedUnimplementedConnectServiceServer() {}

// UnsafeConnectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectServiceServer will
// result in compilation errors.
type UnsafeConnectServiceServer interface {
	mustEmbedUnimplementedConnectServiceServer()
}

func RegisterConnectServiceServer(s grpc.ServiceRegistrar, srv ConnectServiceServer) {
	s.RegisterService(&ConnectService_ServiceDesc, srv)
}

func _ConnectService_JoinNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerJoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).JoinNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.ConnectService/JoinNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).JoinNetwork(ctx, req.(*PeerJoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectService_LeaveNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerLeave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServiceServer).LeaveNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.ConnectService/LeaveNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServiceServer).LeaveNetwork(ctx, req.(*PeerLeave))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectService_ServiceDesc is the grpc.ServiceDesc for ConnectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connect.ConnectService",
	HandlerType: (*ConnectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinNetwork",
			Handler:    _ConnectService_JoinNetwork_Handler,
		},
		{
			MethodName: "LeaveNetwork",
			Handler:    _ConnectService_LeaveNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/connect/connect.proto",
}
