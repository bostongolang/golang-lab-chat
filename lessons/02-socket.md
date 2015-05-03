# Create the TCP socket (5-10 minutes)

## Goals

The goal of this lesson is to create a TCP socket that your chat
server will listen for connections.

## Steps

1. Add some code to chat.go that creates a socket to bind all interfaces to TCP port 6677.

  In your `chat.go` file, use the `net.Listen` function [docs here](http://golang.org/pkg/net/#Listen) 
  to createa listener that binds to TCP port `6677`.  
  
  Ensure the code runs by typing `go run chat.go` .

1. Write some code that will listen for accepted connections on port
6677, and print out the remote address of the connection.

  Use the `listener.Accept()` function to accept connections and print
  the remote address when the connection is joined.   You
  can use the `RemoteAddr()` function on `net.Conn` to do this.
  
  [Docs for the net package and example are here](http://golang.org/pkg/net/)
  
  In pseudocode, this looks like this:
  
  ```
  listener = listen for socket
  
  while (true) {
    connection = accept connection using listener.Accept()
    log("got connection from " - connection.RemoteAddr())
  }
  
  ```

  [Stuck on any of the steps above? See the solution!](code/02-socket/chat.go)

  Restart the server using `go run chat.go`.


[Proceed to Lesson 3](03-user-input.md)
