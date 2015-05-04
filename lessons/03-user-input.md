# Interact with the user (10-15 minutes)

## Goals

The goal of this lesson is to:

1. Verify you can accept connections on your socket.
1. Read the username of the user connected.
1. Echo a response back to the user on the socket.

## Steps

In `chat.go`, you should now have some logic that looks like this to 
accept connections:

```go
  // Accept a connection, and print out the remote
	// address so we know who has connected.
	for {
		conn, _ := listener.Accept()
		log.Println("Connection joined.", conn.RemoteAddr())
	}
```

1. Above the `main` function, create an empty function `func handleConnection(conn net.Conn)` that
will handle the connection received by `listener.Accept`.
  
  Right after you accept the connection and print the 
  'Connection joined' message, call your new function `handleConnection(conn)`.

1. In the body of `handleConnection`, let's create a `bufio.Reader` and a
`bufio.Writer` object that will allow us to read and write to the socket.

  You can use `bufio.NewReader(conn)` and `bufio.NewWriter(conn)`.  
  [Docs for the bufio package are here](http://golang.org/pkg/bufio/)
  
  ```go
    func handleConnection(conn net.Conn) {
    	writer := bufio.NewWriter(conn)
    	reader := bufio.NewReader(conn)
      ...
    }
  ```

1. Using the reader and writer we just created, write some code in `handleConnection` that does the following:

  1. Prints "Enter your username: " on the socket.
  2. Reads the username from the user connected to the socket.
  3. Echos the username back on the socket to the user by printing "Welcome, [username]!"

  Hint: make sure you call `Flush()` to the `writer` socket after writing to it.

You should be able to verify everything is working correctly by `telnet localhost 6677`.

You'll see something like this!

```bash
$ telnet localhost 6677                                                                                                                      ~ 1 ↵
Trying ::1...
telnet: connect to address ::1: Connection refused
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Please enter your username: jen
Welcome, jen
```

[Stuck on any of the steps above? See the solution!](code/03-user-input/chat.go)


[Proceed to Lesson 4](04-data-structures.md)
