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

  tip: If you are using Vagrant, this will 'appear' in `/opt/golang-lab/chat-server`.
  The /opt/golang-lab folder is shared from the repository folder in your host operating
  system, so any changes you make will be reflected.

1. Under this folder, Use your code editor to create a file `chat.go` that will contain your `main` function.  Write some
code that will print "Chat server starting!".

[See the solution!](code/01-setup/chat.go)

1. Ensure the code runs by typing `go run chat.go` 

```bash
cd /opt/golang-lab # if using vagrant
go run chat.go
```

