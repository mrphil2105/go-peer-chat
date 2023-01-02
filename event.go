package main

import "main/grpc/chat"

type Event struct {
	id  EventId
	msg *chat.Message
}

func NewEvent(id EventId, msg *chat.Message) *Event {
	return &Event{
		id:  id,
		msg: msg,
	}
}

type EventId int

const (
	Join EventId = iota
	Leave
	Message
)

var (
	eventName = map[EventId]string{
		Join:    "join",
		Leave:   "leave",
		Message: "msg",
	}
)

func (id EventId) String() string {
	return eventName[id]
}
