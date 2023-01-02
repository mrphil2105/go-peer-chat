package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	lo "log"
	"main/grpc/chat"
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
	RunChat(server)
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
		chat.RegisterChatServiceServer(grpcServer, server)
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
		line := scanner.Text()

		if !strings.HasPrefix(line, "/") {
			server.SendMessage(line)
			continue
		}

		input := strings.Split(line, " ")

		switch input[0] {
		case "/connect":
			port := strconv.Itoa(int(countingPort + ParsePort(input[1])))
			ConnectToPeer(server, port)
		case "/leave":
			Disconnect(server)
		case "/peers":
			PrintPeers(server)
		default:
			fmt.Printf("Unknown command '%s'\n", input[0][1:])
		}
	}
}

func ConnectToPeer(server *Server, peerPort string) {
	conn := ConnectClient(peerPort)
	defer conn.Close() // Close the connection as the 'JoinNetwork' function will connect again properly.
	client := connect.NewConnectServiceClient(conn)

	connectedTo, err := client.JoinNetwork(context.Background(), &connect.PeerJoin{
		Pid:  server.GetPid(),
		Port: port,
		Name: *name,
	})

	if err != nil {
		log.Printf("Unable to connect to peer on port %s", peerPort)
	}

	// Sync our time with the first peer we connected to
	server.time = connectedTo.GetTime()
}

func Disconnect(server *Server) {
	for pid, peer := range server.peers {
		log.Printf("Sending PeerLeave to peer %d", pid)
		_, err := peer.connClient.LeaveNetwork(context.Background(), &connect.PeerLeave{Pid: server.GetPid()})

		if err != nil {
			log.Printf("Failed to send PeerLeave to peer %d: %v", pid, err)
		}

		err = peer.stream.CloseSend()

		if err != nil {
			log.Printf("Unable to close stream to peer %d: %v", peer.pid, err)
		}

		err = peer.conn.Close()

		if err != nil {
			log.Printf("Unable to close connection to peer %d: %v", peer.pid, err)
		}
	}

	server.peers = make(map[uint32]*Peer)
}

func PrintPeers(server *Server) {
	counter := 0

	for pid, _ := range server.peers {
		counter++
		fmt.Printf("%d: %d\n", counter, pid)
	}
}
