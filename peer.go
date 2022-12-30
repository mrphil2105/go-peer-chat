package main

import (
	"google.golang.org/grpc"
	"main/grpc/connect"
)

type Peer struct {
	pid        uint32
	conn       *grpc.ClientConn
	connClient connect.ConnectServiceClient
}

func NewPeer(pid uint32, conn *grpc.ClientConn, connClient connect.ConnectServiceClient) *Peer {
	return &Peer{
		pid:        pid,
		conn:       conn,
		connClient: connClient,
	}
}
