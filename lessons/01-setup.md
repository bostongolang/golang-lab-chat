# Setup the chat project (5-10 minutes)

## Goals

The goal of the portion is to setup project `chat-server` which will
hold the code for your chat project. 

In the interest of time, we will all organize our chat server around the following
data structures which we have scaffolded for you.

 * a `ChatUser` struct to handle the individual chat connection for the user.
 * a `ChatRoom` struct to handle all of the connections, chatroom messages, and disconnects.
 * a `main` function that will be the entry point of your code.

## Steps

1. Create a directory `chat-server` for your code underneath the main folder for 
this repository, and a file `chat.go` that will contain the code.

  E.g. under `golang-lab-chat/chat-server`
  
    ```bash
    mkdir -p chat-server
    cd chat-server
    touch chat.go
    ```

  tip: If you are using Vagrant, this will 'appear' in `/opt/golang-lab-chat/chat-server`.
  The /opt/golang-lab-chat folder is shared from the repository folder in your host operating
  system, so any changes you make will be reflected.


1. Download the basic scaffold code [chat.go](code/01-setup/chat.go), and put it in the `chat-server` folder.  Open `chat.go` with your favorite code editor.

  1. Review the struct `ChatRoom`.  This struct will handle 
    
    * users joining;
    * users disconnecting;
    * receiving individual messages from users and broadcasting (relaying) them to other users.

    ```go
    type ChatRoom struct {
      // TODO: populate this  
    }

    // NewChatRoom will create a chatroom
    func NewChatRoom() *ChatRoom {
      // TODO: return a chatroom
      return nil 
    }

    
    func (cr *ChatRoom) ListenForMessages() {}
    func (cr *ChatRoom) Logout(username string) {}
    func (cr *ChatRoom) Join(conn net.Conn) {}
    func (cr *ChatRoom) Broadcast(msg string) {}
    ```

  1. Review the struct 

1. Ensure the code runs by typing `go run chat.go` 

  ```bash
  cd /opt/golang-lab-chat # if using vagrant
  go run chat.go
  ```

[See the solution!](code/01-setup/chat.go)

[Proceed to Lesson 2](02-socket.md)

