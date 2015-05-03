# Create the TCP socket (5-10 minutes)

## Goals

The goal of this lesson is to create a TCP socket that your chat
server will listen for connections

## Steps

1. In your `chat.go` file, use the `net.Listen` function [docs here](http://golang.org/pkg/net/#Listen) 
to bind to TCP port `6677`.  


Ensure the code runs by typing `go run chat.go` .

1. Use the `listener.Accept()` function to accept connections and print
the remote address when the connection is joined.   You
can use the `RemoteAddr()` function on `net.Conn` to do this.

In pseudocode, this looks like this:

```
listener = listen for socket

while (true) {
  connection = accept connection using listener.Accept()
  log("got connection from " - connection.RemoteAddr())
}

```

[Stuck on any of the steps above? See the solution!](code/02-socket/chat-1.go)

