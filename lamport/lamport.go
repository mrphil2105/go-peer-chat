package lamport

import "main/grpc/chat"

type Lamport interface {
	GetPid() uint32
	GetTime() uint64
}

func MakeMessage(lamport Lamport, content string) *chat.Message {
	return &chat.Message{
		Time:    lamport.GetTime(),
		Pid:     lamport.GetPid(),
		Content: content,
	}
}

func LamportSend(lamport Lamport) uint64 {
	return lamport.GetTime() + 1
}

func LamportRecv(lamport, other Lamport) uint64 {
	if Compare(lamport, other) > 0 {
		return lamport.GetTime() + 1
	} else {
		return other.GetTime() + 1
	}
}

func Compare(lamport, other Lamport) int {
	// First compare by time, then by pid (lower pid is greater)
	if lamport.GetTime() < other.GetTime() ||
		lamport.GetTime() == other.GetTime() &&
			lamport.GetPid() > other.GetPid() {
		return -1
	} else if lamport.GetPid() == other.GetPid() {
		return 0
	} else {
		return 1
	}
}
