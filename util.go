package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strconv"
)

func ConnectClient(port string) *grpc.ClientConn {
	log.Printf("Connecting to peer on port %s...", port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to peer: %v", err)
	}

	return conn
}

func ParsePort(portStr string) uint16 {
	parsed, err := strconv.ParseUint(portStr, 10, 16)

	if err != nil {
		log.Fatalf("Failed to parse '%s' as port: %v", portStr, err)
	}

	return uint16(parsed)
}
