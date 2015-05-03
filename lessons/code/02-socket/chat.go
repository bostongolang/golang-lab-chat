package main

import (
	"fmt"
	"log"
	"net"
)

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
	}

}
