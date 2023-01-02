package main

import (
	"context"
	"google.golang.org/grpc"
	"main/grpc/chat"
	"main/grpc/connect"
)

type Peer struct {
	pid        uint32
	name       string
	conn       *grpc.ClientConn
	connClient connect.ConnectServiceClient
	stream     chat.ChatService_OpenChatClient
}

func NewPeer(pid uint32, name string, conn *grpc.ClientConn) *Peer {
	chatClient := chat.NewChatServiceClient(conn)
	stream, err := chatClient.OpenChat(context.Background())

	if err != nil {
		log.Fatalf("Failed to open chat stream: %v", err)
	}

	return &Peer{
		pid:        pid,
		name:       name,
		conn:       conn,
		connClient: connect.NewConnectServiceClient(conn),
		stream:     stream,
	}
}
