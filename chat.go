package main

import (
	"fmt"
	"io"
	"main/grpc/chat"
	"main/lamport"
)

func (server *Server) OpenChat(stream chat.ChatService_OpenChatServer) error {
	for {
		msg, err := stream.Recv()

		if err != nil {
			if err != io.EOF {
				log.Printf("Stream receive error: %v", err)
				// TODO: Figure out a way to know what PID to remove.
				return err
			}

			return nil
		}

		log.Printf("Received message '%s'", msg)
		server.events <- NewEvent(Message, msg)
	}
}

func RunChat(server *Server) {
	go func() {
		for event := range server.events {
			msg := event.msg
			peer := server.peers[msg.GetPid()]

			switch event.id {
			case Join:
				fmt.Printf("%s has joined the chat\n", peer.name)
			case Leave:
				fmt.Printf("%s has left the chat\n", peer.name)
			case Message:
				server.ReceiveMessage(msg)

				var peerName string

				// A 'Message' event can be a local log, so peer would be 'nil'
				if peer != nil {
					peerName = peer.name
				} else {
					peerName = *name
				}

				fmt.Printf("[%-3d: %s>] %s\n", msg.GetTime(), peerName, msg.GetContent())
			}
		}
	}()
}

func (server *Server) ReceiveMessage(msg *chat.Message) {
	time := lamport.LamportRecv(server, msg)
	log.Printf("Ticking time (%d -> %d)", server.GetTime(), time)
	server.time = time
}

// Log a message locally (without sending it to the peers)
func (server *Server) LogMessage(msg *chat.Message) {
	server.events <- NewEvent(Message, msg)
}

func (server *Server) SendMessage(content string) {
	time := lamport.LamportSend(server)
	log.Printf("Ticking time (%d -> %d)", server.GetTime(), time)
	server.time = time

	msg := lamport.MakeMessage(server, content)

	for pid, peer := range server.peers {
		err := peer.stream.Send(msg)

		if err != nil {
			log.Printf("Failed to send to message to peer %d: %v", pid, err)
		}
	}

	log.Printf("Sent message '%s'", msg)
	server.LogMessage(msg)
}
