# Setup the chat project (5-10 minutes)

## Goals

The goal of the portion is to setup project directory `chat-server` which will
hold the code for your chat project. 

## Steps

1. Create a directory `chat-server` for your code underneath the main folder for 
this repository.

  E.g. under `golang-lab-chat/chat-server`
  
    ```bash
    mkdir -p chat-server
    ```

  If you are using Vagrant, this will 'appear' in `/opt/golang-lab/chat-server`.

1. Under this folder, create a file `chat.go` that will contain your `main` function.  Write some
code that will print "Chat server starting!" 

>!  ```go
>!  package main
>!  
>!  func main() {
>!    fmt.Println("Chat server starting!")    
>!  }
>!  
>!  ```

1. Ensure the code runs by typing `go run chat.go` 

