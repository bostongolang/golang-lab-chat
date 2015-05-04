# Populate `ChatRoom` and `ChatUser` (5-10 minutes)

## Goals

In this lesson, we'll setup the ChatRoom and ChatUser structs. 

Let's review the design of the data structures and some of the requirements:

  1. There is one `ChatRoom` in the app.
  1. The `ChatRoom` must know about all of the active connections.
  1. Each connection is tracked in a `ChatUser` object.
  1. The `ChatRoom` must be able to receive messages from a single connection,
  and replay it back to the other connections.  Otherwise, it will not be a very
  good chatroom!
  1. When a new connection is established, the `ChatRoom` must be notified
  of these new connections.

## Steps

1. Find the `ChatRoom` struct in `chat.go`.

  ```go
  type ChatRoom struct {
    // TODO: populate this   
  }
  ```
  
  1. :star2: Add to the `ChatRoom` struct:
    * a private member `users` of type `map[string]*ChatUser`
    * a private member channel `incoming` of type `chan string`
    * a private member channel `joins` of type `chan *ChatUser`
    * a private member channel `disconnects` of type `chan string`
  1. :star2: Initialize all of these data structures in the `NewChatRoom()` constructor.

  [Stuck on any of the steps above? Ask your TA, or see the solution!](code/03-data-structures/chat.go)

1. Find the `ChatUser` struct in `chat.go`.

  ```go
  type ChatUser struct {
    // TODO: populate this   
  }
  ```
  
  1. :star2: Add to the `ChatUser` struct
    * a private member `conn` of type `net.Conn` 
    * a private member `disconnect` of type `bool` 
    * a private member `username` of type `string` 
    * a private member channel `outgoing` of type `chan string` 
    * a private member `reader` of type `*bufio.Reader` 
    * a private member `writer` of type `*bufio.Writer` 
  1. :star2: Initialize all of these data structures in the `NewChatRoom()` constructor.
    * `bufio.NewReader/bufio.NewWriter` should accept the `conn` to create the `reader`
    and `writer` variables.
    * `disconnect` should be initially set to false

  [Stuck on any of the steps above? Ask your TA, or see the solution!](code/03-data-structures/chat.go)

1. Don't worry too much about what all of these variables are for right now.
As we fill out the rest of the code, you will gain a better understanding of what they will be used for, or you can ask your TA for more information.

  [You can also view the comments on the structs here](code/03-data-structures/chat.go).

1. Verify the code still runs when you type `go run chat.go`.

[Next, proceed to Lesson 4 - where we actually read something from the user socket!](04-login.md)
