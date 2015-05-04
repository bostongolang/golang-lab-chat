# Create the TCP socket (5-10 minutes)

## Goals

The goal of this lesson is to create a TCP server on port 6677 that will
listen for connections.

## Steps


1. Open your `chat.go` file and find the `main` function.

  1. :star2: Use the `net.Listen` function [docs here](http://golang.org/pkg/net/#Listen) 
  to create a listener that binds to TCP port `6677`.  

  1. :star2: Create a new instance of the chatroom using NewChatRoom() and call
  `chatroom.ListenForMessages()`
  
  1. :star2: Write some code that will loop listen for accepted connections on port
  6677, and print out the remote address of the connection using `log`.

  hint: Use the `listener.Accept()` function to accept connections and print
  the remote address when the connection is joined.   You
  can use the `RemoteAddr()` function on `net.Conn` to do this.
  
  [Docs for the net package and example are here](http://golang.org/pkg/net/)
  
[Stuck on any of the steps above? See the solution!](code/02-socket/chat.go)

1. Run the server using `go run chat.go`.  It should wait after the "Chat server starting!" message. 
You can stop the server using CTRL-C.

[Proceed to Lesson 3](03-data-structures.md)
