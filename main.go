package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	lo "log"
	"main/grpc/connect"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
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

	server := RunServer()
	ReadInput(server)
}

func RunServer() *Server {
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

	return server
}

func ReadInput(server *Server) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		switch input[0] {
		case "/connect":
			port := strconv.Itoa(int(countingPort + ParsePort(input[1])))
			ConnectToPeer(server, port)
		case "/peers":
			PrintPeers(server)
		}
	}
}

func ConnectToPeer(server *Server, peerPort string) {
	conn := ConnectClient(peerPort)
	defer conn.Close() // Close the connection as the 'JoinNetwork' function will connect again properly.
	client := connect.NewConnectServiceClient(conn)

	_, err := client.JoinNetwork(context.Background(), &connect.PeerJoin{
		Pid:  server.GetPid(),
		Port: port,
	})

	if err != nil {
		log.Printf("Unable to connect to peer on port %s", peerPort)
	}
}

func PrintPeers(server *Server) {
	counter := 0

	for pid, _ := range server.peers {
		counter++
		fmt.Printf("%d: %d\n", counter, pid)
	}
}
