package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	writer.WriteString("Please enter your username: ")
	writer.Flush()

	username, _ := reader.ReadString(byte('\n'))
	log.Println("User logged in:", username)

	writer.WriteString("Welcome, " + username)
	writer.Flush()
}

func main() {
	fmt.Println("Chat server starting!")

	// This binds the socket to TCP port 6677
	// on all interfaces.
	listener, err := net.Listen("tcp", ":6677")

	if err != nil {
		log.Fatal("Unable to bind to 6677", err)
	}

	// Accept a connection, and print out the remote
	// address so we know who has connected.
	for {
		conn, _ := listener.Accept()
		log.Println("Connection joined.", conn.RemoteAddr())
		handleConnection(conn)
	}

}
