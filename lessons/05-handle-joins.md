# Lesson 5 - Handling Chatroom Joins (15-20 minutes)

## Goals

Let's start playing some more with concurrency! 

In the previous lesson, we implemented `ChatRoom.Join` which placed a newly created
`ChatUser` object on the `joins` channel. 

What we *want* to do is to be able to track all of the users in the `ChatRoom.users`
map, so we can tell the other users when new users have joined, and also make sure
we can broadcast messages correctly.

We will do this by implementing the dispatcher in `chatroom.ListenForMessages` 
which is designed to handle the messages on the various chatroom channels,
and act accordingly (e.g. add or remove users from internal data structures, broadcast
connection, disconnection, and other messages to all users). This dispatcher
will run in its own goroutine constantly looking for message on its queues.

So, let's go!

## Steps


1. In `chat.go`, remember how the `main` function has some code that calls 
`chatroom.ListenForMessages`?

  ```go
  // This binds the socket to TCP port 6677
  // on all interfaces.
  listener, err := net.Listen("tcp", ":6677")

  chatroom.ListenForMessages()
  ```

  Well, the job of `chatroom.ListenForMessages` is to _listen in a loop_ for any messages on the channels in the chatroom
  object, and then handle those messages accordingly.

  Let's write some code in this function that will handle a new `ChatUser` object on the
  `joins` channel!

  1. In `ChatRoom.ListenForMessages`, implement the following
      * :star2: Create a `for/select` loop, and ensure this loop runs in a goroutine.
      * :star2: In the loop, receive a `ChatUser` object on the `cr.joins` channel. Store it in the `users` map by its username.
   
    [Stuck on any of the steps above? Ask your TA, or see the solution!](code/05-handle-joins/chat.go)

1. Great! now the object should be stored in the users map. However: it would
  be nice if the other people in the chatroom knew when new users joined the chatroom server.

  To do this, we need to implement `ChatRoom.Broadcast`.  This function will 
  pass a message on each of the `ChatUser.outgoing` channels (using `ChatUser.Send`).  
  
  The idea is that each `ChatUser` will be listening on this channel for any outgoing messages  via the `ChatUser.WriteOutgoingMessages` 
  function. This function  runs in a loop in a goroutine, and calls `ChatUser.WriteString()` to write that message to the socket when it sees a new message on the `outgoing channel`.

  1. :star2: Implement `ChatUser.Send` to place a message on the `chatuser.outgoing` channel.
    
     [Stuck on any of the steps above? Ask your TA, or see the solution!](code/05-handle-joins/chat.go)

  1. :star2: Implement `ChatUser.WriteOutgoingMessages` by: 

    1. :star2: Creating a loop that constantly reads a msg from the `chatuser.outgoing` channel; 
    1. :star2: Adding a newline to this msg;
    1. :star2: Write the msg to the socket by calling `chatuser.WriteString`.

    [Stuck on any of the steps above? Ask your TA, or see the solution!](code/05-handle-joins/chat.go)
  
  4. :star2: Let's start the outgoing message listener! At the end of `ChatUser.Login`, call `cu.WriteOutgoingMessages` loop and make 
  sure it runs in a goroutine.

    [Stuck on any of the steps above? Ask your TA, or see the solution!](code/05-handle-joins/chat.go)

1. Now that this is all setup, let's broadcast a message whenever the chatroom
sees a new user login.

  1. :star2: Modify `ChatRoom.ListenForMessages` to broadcast a message whenever
  a new user logs in that says `*** [name] joined the chatroom`:  

    ```go
      func (cr *ChatRoom) ListenForMessages() {
      	go func() {
      		for {
      			select {
      			case user := <-cr.joins:
      				cr.users[user.username] = user
      				cr.Broadcast("*** " + user.username + " just joined the chatroom")
      			}
      		}
      	}()
      }
    ```

1. Verify everything works as designed when you connect! You can 
try connecting with multiple simultaneous clients to make sure you see the join messages everywhere.


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
    *** funcuddles just joined the chatroom
    ```

[Finally! Let's chat!  Proceed to Lesson 6](06-broadcast-msgs.md)
