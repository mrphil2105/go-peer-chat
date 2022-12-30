package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func ConnectClient(port string) *grpc.ClientConn {
	log.Printf("Connecting to on port %s...", port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to peer: %v", err)
	}

	return conn
}
