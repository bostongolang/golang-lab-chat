# Populate `ChatRoom` and `ChatUser` 

## Goals

In this lesson, we'll setup the ChatRoom and ChatUser structs.

## Steps

1. Find the `ChatRoom` struct in `chat.go`.

  ```go
  type ChatRoom struct {
    // TODO: populate this   
  }
  ```
  
  1. Add to the `ChatRoom` struct:
    * a private member `users` of type `map[string]*ChatUser`
    * a private member channel `incoming` of type `chan string`
    * a private member channel `joins` of type `chan *ChatUser`
    * a private member channel `disconnects` of type `chan string`
  1. Initialize all of these data structures in the `NewChatRoom()` constructor.

[Stuck on any of the steps above? See the solution!](code/03-data-structures/chat.go)

1. Find the `ChatUser` struct in `chat.go`.

  ```go
  type ChatUser struct {
    // TODO: populate this   
  }
  ```
  
  1. Add to the `ChatUser` struct
    * a private member `conn` of type `net.Conn` 
    * a private member `disconnect` of type `bool` 
    * a private member `username` of type `string` 
    * a private member channel `outgoing` of type `chan string` 
    * a private member `reader` of type `*bufio.Reader` 
    * a private member `writer` of type `*bufio.Writer` 
  1. Initialize all of these data structures in the `NewChatRoom()` constructor.
    * `bufio.NewReader/bufio.NewWriter` should accept the `conn` to create the `reader`
    and `writer` variables.
    * `disconnect` should be initially set to false

[Stuck on any of the steps above? See the solution!](code/03-data-structures/chat.go)

1. Review these member variables and ponder their roles.  As we fill out the rest of the code,
you will gain a better understanding of what they will be used for, or you can ask your TA for more information.

  [You can also view the comments on the structs here](code/03-data-structures/chat.go).

1. Verify the code still runs when you type `go run chat.go`.

[Proceed to Lesson 4](04-user-input.md)
