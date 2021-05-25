package protocol

import "errors"

var (
	UnknownCommand = errors.New("Unknown command")
)

// NameCommand used to set client nickname
type NameCommand struct {
	Name string
}

// SendCommand used to send message from client
type SendCommand struct {
	Message string
}

// MessageCommand used to notify about new messages
type MessageCommand struct {
	Name    string
	Message string
}
