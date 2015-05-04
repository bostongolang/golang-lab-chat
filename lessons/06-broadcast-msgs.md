# Lesson 6 - Broadcasting Messages (15-20 minutes)

## Goals

Finally, to the heart of the matter! Chatting!

To chat with each other, we need to be able to read messages from each 
individual ChatUser's socket, then broadcast them back to all of the other user
sockets.  

Here's how it will work:

  * We have a channel `ChatRoom.incoming` where we will put all of the messages
  that come in from all of the user sockets.
  * Each `ChatUser` object will be modified to have a loop that looks for any new lines
  read in from the client, and sends it to `ChatRoom.incoming` channel.
  * In our `ChatRoom.ListenForMessages`, we will look for any new messages
  on `ChatRoom.incoming` and call `Broadcast` to broadcast it out to all of
  the other connected sockets.

Got it?? Not sure? No worries, Let's go through the steps.

## Steps


1. Look for the empty `ChatUser.ReadIncomingMessages` function. Let's implement it.

  ```go
    func (cu *ChatUser) ReadIncomingMessages(chatroom *ChatRoom) {
    	// TODO: read incoming messages in a loop
    }
  ```
  1. In `ChatUser.ReadIncomingMessages`, implement the following:

    1. :star2: Create a `for/select` loop, and ensure this loop runs in a goroutine.
    1. :star2: In the loop, call `cu.ReadLine` to read an incoming msg.
    1. :star2: Prepend the incoming msg with the username surrounded by brackets, e.g.: `msg = "[" + cu.username "]" + msg`
    1. :star2: Place this modified message on the `chatroom.incoming` queue.
   
    [Stuck on any of the steps above? Ask your TA, or see the solution!](code/06-broadcast-msgs/chat.go)

  1. Now, you need to make sure that the `ReadIncomingMessages` goroutine is started
  in `ChatUser.Login`.  
    
    1. :star2: Modify the `Login` function to start the goroutine for `ReadIncomingMessages` after `WriteOutgoingMessages`.

    [Stuck on any of the steps above? Ask your TA, or see the solution!](code/06-broadcast-msgs/chat.go)

1. Cool. Remember `ChatRoom.ListenForMessages()`?  Let's modify it so it sees
  the new messages on the `chatroom.incoming` queue.

  1. :star2: Add a 'case' within the `for/select` loop that reads a msg string
  1. :star2: Call `ChatRoom.Broadcast` to broadcast the read msg.
    
  [Stuck on any of the steps above? Ask your TA, or see the solution!](code/06-broadcast-msgs/chat.go)

  At the end it will look like:

  ```go
  func (cr *ChatRoom) ListenForMessages() {
	  go func() {
	  	for {
	  		select {
	  		case msg := <-cr.incoming:
	  			cr.Broadcast(msg)
	  		case user := <-cr.joins:
	  			cr.users[user.username] = user
	  			cr.Broadcast("*** " + user.username + " just joined the chatroom")
	  		}
	  	}
	  }()
  }
  ```
  
1. Verify everything works as designed.  Connect to port 6677, then type a message and
hit enter to send a chat msg to everyone. You can try connecting with 
multiple simultaneous clients to test how messages look to other users.

  ```bash
   $ nc localhost 6677
   Welcome to Jen's chat server!
   Please enter your username: funcuddles
   Welcome, funcuddles
   *** funcuddles just joined the chatroom
   hello this is dog
   [funcuddles] hello this is dog
   ```

[Next lesson - Handle user disconnects](07-disconnects.md)
