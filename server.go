package main

import (
	"context"
	"fmt"
	"main/grpc/chat"
	"main/grpc/connect"
	"os"
)

type Server struct {
	connect.UnimplementedConnectServiceServer
	chat.UnimplementedChatServiceServer
	pid    uint32
	time   uint64
	peers  map[uint32]*Peer
	events chan *Event
}

func NewServer() *Server {
	return &Server{
		pid:    uint32(os.Getpid()),
		peers:  make(map[uint32]*Peer),
		events: make(chan *Event, 1),
	}
}

func (server *Server) JoinNetwork(ctx context.Context, peerJoin *connect.PeerJoin) (*connect.ConnectedTo, error) {
	if _, exists := server.peers[peerJoin.GetPid()]; !exists {
		conn := ConnectClient(peerJoin.GetPort())
		server.AddPeer(NewPeer(peerJoin.GetPid(), peerJoin.GetName(), conn))

		server.events <- NewEvent(Join, &chat.Message{Pid: peerJoin.GetPid()})
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
					Name: *name,
				})
			}

			if err != nil {
				log.Fatalf("Failed to propagate peer join: %v", err)
			}
		}
	}

	return &connect.ConnectedTo{
		Pid:  server.GetPid(),
		Time: server.GetTime(),
	}, nil
}

func (server *Server) LeaveNetwork(ctx context.Context, peerLeave *connect.PeerLeave) (*connect.Void, error) {
	if peer, exists := server.peers[peerLeave.GetPid()]; exists {
		// We raise event before removing peer to still allow reading from the map
		server.events <- NewEvent(Leave, &chat.Message{Pid: peerLeave.GetPid()})

		err := peer.stream.CloseSend()

		if err != nil {
			log.Printf("Unable to close stream to peer %d: %v", peer.pid, err)
		}

		err = peer.conn.Close()

		if err != nil {
			log.Printf("Unable to close connection to peer %d: %v", peer.pid, err)
		}

		server.RemovePeer(peer.pid)

		log.Printf("Peer %d has left the network", peer.pid)

		return &connect.Void{}, nil
	}

	return nil, fmt.Errorf("peer %d was not found", peerLeave.GetPid())
}

func (server *Server) GetPid() uint32 {
	return server.pid
}

func (server *Server) GetTime() uint64 {
	return server.time
}

func (server *Server) AddPeer(peer *Peer) {
	server.peers[peer.pid] = peer
}

func (server *Server) RemovePeer(pid uint32) {
	delete(server.peers, pid)
}
