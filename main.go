package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	lo "log"
	"main/grpc/connect"
	"net"
	"os"
	"path"
	"strconv"
)

var (
	id           = flag.String("id", "0", "A number to assign the peer. Used to determine port")
	name         = flag.String("name", "NoName", "Name of the peer")
	countingPort = uint16(50050)
	port         = ""
	log          = lo.Default()
)

func main() {
	flag.Parse()
	port = strconv.Itoa(int(countingPort + ParsePort(*id)))

	prefix := fmt.Sprintf("%-8s: ", *name)
	logFileName := path.Join("logs", fmt.Sprintf("%s.log", *name))
	_ = os.Mkdir("logs", 0664)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatalf("Unable to open or create file %s: %v", logFileName, err)
	}

	log = lo.New(logFile, prefix, lo.Ltime)

	RunServer()
}

func RunServer() {
	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", port))

	if err != nil {
		log.Fatalf("Unable to listen on port %s: %v", port, err)
	}

	server := NewServer()

	go func() {
		grpcServer := grpc.NewServer()
		connect.RegisterConnectServiceServer(grpcServer, server)
		log.Printf("Created gRPC server on port %s", port)

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Stopped serving due to error: %v", err)
		}
	}()
}
