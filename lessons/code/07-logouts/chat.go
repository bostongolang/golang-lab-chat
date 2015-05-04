package main

import (
	"bufio"
	"log"
	"net"
)

//
// ChatRoom is the main chatroom data structure.
//
// `users` contains connected ChatUser connections.
// `incoming` receives incoming messages from ChatUser connections.
// `joins` receives incoming new ChatUser connections.
// `disconnects` receives disconnect notifications.
//
type ChatRoom struct {
	users       map[string]*ChatUser
	incoming    chan string
	joins       chan *ChatUser
	disconnects chan string
}

//
// NewChatRoom() will create a new chatroom.
//
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users:       make(map[string]*ChatUser),
		incoming:    make(chan string),
		joins:       make(chan *ChatUser),
		disconnects: make(chan string),
	}
}

func (cr *ChatRoom) ListenForMessages() {
	go func() {
		for {
			select {
			case msg := <-cr.incoming:
				cr.Broadcast(msg)
			case user := <-cr.joins:
				cr.users[user.username] = user
				cr.Broadcast("*** " + user.username + " just joined the chatroom")
			case username := <-cr.disconnects:
				// perform the logout by removing
				// the username from the `users` map if it exists.
				// Also call Close() on the user object.
				if cr.users[username] != nil {
					cr.users[username].Close()
					delete(cr.users, username)
					cr.Broadcast("*** " + username + " has disconnected")
				}
			}
		}
	}()
}

func (cr *ChatRoom) Logout(username string) {
	cr.disconnects <- username
}

func (cr *ChatRoom) Join(conn net.Conn) {
	user := NewChatUser(conn)
	if user.Login(cr) == nil {
		cr.joins <- user
	}
}

func (cr *ChatRoom) Broadcast(msg string) {
	for _, user := range cr.users {
		user.Send(msg)
	}
}

//
// ChatUser contains information for the connected user.
//
// `conn` is the socket.
// `disconnect` indicates whether or not the socket is disconnected.
// `username` is the chat username.
// `outgoing` is a channel with all pending outgoing messages
// to be written to the socket.
// `reader` is the buffered socket read stream.
// `writer` is the buffered socket write stream.
//
type ChatUser struct {
	conn       net.Conn
	disconnect bool
	username   string
	outgoing   chan string
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func NewChatUser(conn net.Conn) *ChatUser {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	cu := &ChatUser{
		conn:       conn,
		disconnect: false,
		reader:     reader,
		writer:     writer,
		outgoing:   make(chan string),
	}

	return cu
}

func (cu *ChatUser) ReadIncomingMessages(chatroom *ChatRoom) {
	go func() {
		for {
			line, err := cu.ReadLine()
			if cu.disconnect {
				break
			}
			if err != nil {
				chatroom.Logout(cu.username)
				break
			}
			if line != "" {
				chatroom.incoming <- ("[" + cu.username + "] " + line)
			}
		}
	}()
}

func (cu *ChatUser) WriteOutgoingMessages(chatroom *ChatRoom) {
	go func() {
		for {
			data := <-cu.outgoing
			if cu.disconnect {
				break
			}
			data = data + "\n"
			err := cu.WriteString(data)
			if err != nil {
				chatroom.Logout(cu.username)
				break
			}
		}
	}()
}

func (cu *ChatUser) Login(chatroom *ChatRoom) error {
	var err error
	cu.WriteString("Welcome to Jen's chat server!\n")
	cu.WriteString("Please enter your username: ")

	cu.username, err = cu.ReadLine()

	if err != nil {
		return err
	}

	log.Println("User logged in:", cu.username)

	cu.WriteString("Welcome, " + cu.username + "\n")

	cu.WriteOutgoingMessages(chatroom)
	cu.ReadIncomingMessages(chatroom)
	return nil
}

func (cu *ChatUser) ReadLine() (string, error) {
	bytes, _, err := cu.reader.ReadLine()
	str := string(bytes)
	return str, err
}

func (cu *ChatUser) WriteString(msg string) error {
	_, err := cu.writer.WriteString(msg)
	if err != nil {
		return err
	}
	return cu.writer.Flush()
}

func (cu *ChatUser) Send(msg string) {
	cu.outgoing <- msg
}

func (cu *ChatUser) Close() {
	cu.disconnect = true
	cu.conn.Close()
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
	chatroom := NewChatRoom()

	// This binds the socket to TCP port 6677
	// on all interfaces.
	listener, err := net.Listen("tcp", ":6677")

	chatroom.ListenForMessages()

	if err != nil {
		log.Fatal("Unable to bind to 6677", err)
	}

	// Accept a connection, and print out the remote
	// address so we know who has connected.
	for {
		conn, _ := listener.Accept()
		log.Println("Connection joined.", conn.RemoteAddr())

		// run this in a goroutine so more than one thing
		// can connect
		go chatroom.Join(conn)
	}
}
