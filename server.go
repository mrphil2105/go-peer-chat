package main

import (
	"context"
	"log"
	"main/grpc/connect"
	"os"
)

type Server struct {
	connect.UnimplementedConnectServiceServer
	pid   uint32
	peers map[uint32]*Peer
}

func NewServer() *Server {
	return &Server{
		pid:   uint32(os.Getpid()),
		peers: make(map[uint32]*Peer),
	}
}

func (server *Server) JoinNetwork(ctx context.Context, peerJoin *connect.PeerJoin) (*connect.ConnectedTo, error) {
	if _, exists := server.peers[peerJoin.GetPid()]; !exists {
		conn := ConnectClient(peerJoin.GetPort())
		client := connect.NewConnectServiceClient(conn)
		server.AddPeer(NewPeer(peerJoin.GetPid(), conn, client))

		log.Printf("Connected to peer %d", peerJoin.GetPid())

		// Tell the rest of the network about the new peer
		for pid, peer := range server.peers {
			var err error

			if pid != peerJoin.GetPid() {
				_, err = peer.connClient.JoinNetwork(ctx, peerJoin)
			} else {
				// Tell the new peer about us
				_, err = peer.connClient.JoinNetwork(ctx, &connect.PeerJoin{
					Pid:  server.GetPid(),
					Port: port,
				})
			}

			if err != nil {
				log.Fatalf("Failed to propagate peer join: %v", err)
			}
		}
	}

	return &connect.ConnectedTo{Pid: server.GetPid()}, nil
}

func (server *Server) GetPid() uint32 {
	return server.pid
}

func (server *Server) AddPeer(peer *Peer) {
	server.peers[peer.pid] = peer
}
