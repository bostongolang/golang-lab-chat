package main

import (
	"log"
	"net"
)

type ChatRoom struct {
	// TODO: populate this
}

// NewChatRoom will create a chatroom
func NewChatRoom() *ChatRoom {
	// TODO: initialize struct members
	return &ChatRoom{}
}

func (cr *ChatRoom) ListenForMessages()     {}
func (cr *ChatRoom) Logout(username string) {}
func (cr *ChatRoom) Join(conn net.Conn)     {}
func (cr *ChatRoom) Broadcast(msg string)   {}

type ChatUser struct {
	// TODO: populate this
}

func NewChatUser(conn net.Conn) *ChatUser {
	// TODO: initialize chat user
	return &ChatUser{}
}

func (cu *ChatUser) ReadIncomingMessages(chatroom *ChatRoom) {
	// TODO: read incoming messages in a loop
}

func (cu *ChatUser) WriteOutgoingMessages(chatroom *ChatRoom) {
	// TODO: wait for outgoing messages in a loop, and write them
}

func (cu *ChatUser) Login(chatroom *ChatRoom) error {
	// TODO: login the user
	return nil
}

func (cu *ChatUser) ReadLine() (string, error) {
	// TODO: read a line from the socket
	return "", nil
}

func (cu *ChatUser) WriteString(msg string) error {
	// TODO: write a line from the socket
	return nil
}

func (cu *ChatUser) Send(msg string) {
	// TODO: put a message on the outgoing messages queue
}

func (cu *ChatUser) Close() {
	// TODO: close the socket
}

//
// main will create a socket, bind to port 6677,
// and loop while waiting for connections.
//
// When it receives a connection it will pass it to
// `chatroom.Join()`.
//
func main() {
	log.Println("Chat server starting!")
	// TODO add other logic
}
