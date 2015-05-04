# Read user input from socket! (Time: 15 minutes)

## Goals

Finally.  Let's start reading and writing some data from the socket.

In this exercise, we will:

 * Print a banner when the user connects. 
 * Implement `ChatUser.Login` (and some related functions) to be able to read
 a username from the user.

Btw: To test your server, you can use `nc localhost 6677` or `telnet localhost 6677`.

## Steps


1. In `chat.go`, view the `main` function and notice how `ChatRoom.Join` is called
on each connection.  Find the section of the code `ChatRoom.Join` 

  ```go
  // This is what we want to modify
  func (cr *ChatRoom) Join(conn net.Conn)     {}
  ```

  1. In `ChatRoom.Join`, write the code that does the following: 
    * :star2: Creates a new `ChatUser` object using `NewChatUser` 
    * :star2: Calls `ChatUser.Login` on this object (and verify there is no error)
    * :star2: Notifies of a new user by putting the newly created `ChatUser` object on the `ChatRoom.joins` channel. 
    (Don't worry about how this is used for now, I'll show you how we consume it later.)

    [Stuck on any of the steps above? See the solution!](code/04-login/chat.go)

1. Great! Now let's start implementing `ChatUser.Login`.  First, let's create a 
helpful banner that says "Welcome to [foo's] server", where `foo` is your name.

  1. :star2: Find `ChatUser.Login` and edit it to call `cu.WriteString` with your banner message. 
    Make sure you also write the newline.

    Here's what it should look like: 
    ```go
    func (cu *ChatUser) Login(chatroom *ChatRoom) error {
    	// TODO: login the user
    	cu.WriteString("Welcome to Jen's chat server!\n")
    	return nil
    }
    ```
  
  1. Find the function `ChatUser.WriteString`

    ```go
    func (cu *ChatUser) WriteString(msg string) error {
      // TODO: write a line to the socket
      return nil
    }
    ```

    1. :star2: Implement the code in `WriteString` that will write the `msg` to the `writer`.
    *Make sure you call `writer.Flush`*.

     [Stuck on any of the steps above? See the solution!](code/04-login/chat.go)

  1. Start the server using `go run chat.go`. Test this using the `telnet` or `nc` tool
  to connect to port `6677`.

    ```bash
    $ telnet localhost 6677                                                                                                                      ~ 1 ↵
    Trying ::1...
    telnet: connect to address ::1: Connection refused
    Trying 127.0.0.1...
    Connected to localhost.
    Escape character is '^]'.
    Welcome to Jen's chat server!
    ```
1. Now we are going to read from the socket.  We want to ask for the person's 
username and store it on the `ChatUser.username` field. 

  1. First, let's implement the `ChatUser.ReadLine` function.

    Find this code:

    ```go
    func (cu *ChatUser) ReadLine() (string, error) {
    	// TODO: read a line from the socket
    	return "", nil
    }
    ```

    1. :star2: Implement the code that calls `cu.reader.ReadLine` and returns the results as a string.

    [Stuck on any of the steps above? See the solution!](code/04-login/chat.go)

  1. Go back to the `ChatUser.Login` function.  After you print the banner, write some code that:

  
    1. :star2: Will print to the socket "Please enter your username:";
    1. :star2: Read the username from the socket using `cu.ReadLine`;
    1. :star2: And write back to the socket "Welcome, [username]".

    [Stuck on any of the steps above? See the solution!](code/04-login/chat.go)

    Here is what is should look like when you connect via  `telnet localhost 6677`: 

       ```bash
      $ telnet localhost 6677                                                                                                                      ~ 1 ↵
      Trying ::1...
      telnet: connect to address ::1: Connection refused
      Trying 127.0.0.1...
      Connected to localhost.
      Escape character is '^]'.
      Welcome to Jen's chat server!
      Please enter your username: funcuddles
      Welcome, funcuddles
      ```

[Proceed to Lesson 5](04-user-struct.md)
