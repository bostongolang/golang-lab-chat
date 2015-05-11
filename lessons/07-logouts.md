# Lesson 7 - Notify When Users Disconnect 

## Goals

We know when users login, we send and receive chatroom messages.. but how do we know when users disconnect?  

Let's write some code to handle disconnects and display a message
`*** [username] disconnected` when a socket disconnects.

How will we accomplish this?  With more channels, of course!

1. Whenever a ChatUser fails to read/write to its connection socket, we will send a message on the `ChatRoom.disconnect` channel to notify the `ListenForMessages` loop
that the socket has been disconnected.
1. When such a message is received by the `ChatRoom`, remove it from the `users` map and broadcast a message to the rest of the clients saying it's disconnected.

## Steps

1. Find the `ChatRoom.Logout` function. Add some code that does the following:
  1. :star2: Send the supplied `username` to the `cr.disconnects` channel.

  [Stuck on any of the steps above? Ask your TA, or see the solution!](code/07-logouts/chat.go#L59-L61)

1. Update the `ListenForMessages` loop to handle messages on `cr.disconnects`:
  1. :star2: Add a `case` statement that reads the `username` from `cr.disconnects`
  1. :star2: if the `username` is in the `cr.users` map:
    1. :star2: call `Close()` on the `ChatUser` object and remove it from the map.
    1. :star2: call `Broadcast` to send a message: _`*** [username] has disconnected`_

  [Stuck on any of the steps above? Ask your TA, or see the solution!](code/07-logouts/chat.go#L45-L54)

1. Now that we have `Logout` implemented, let's call `Logout` whenever the ChatUser read and write loops
have errors reading or writing from the socket. 

  1. Update `ChatUser.ReadIncomingMessages` so that it does the following:

    1. :star2: If cu.ReadLine() returns an error, do not write a message to the `chatroom.incoming` queue. 
    Instead, call `chatroom.Logout` with the current username.
    1. :star2: Add some logic that checks if `cu.disconnect` is set, then exit the loop.
  
    [Stuck on any of the steps above? Ask your TA, or see the solution!](code/07-logouts/chat.go#L115-L121)

  1. Update `ChatUser.WriteOutgoingMessages` so that: 
    
    1. :star2: If cu.ReadLine() returns an error, do not write a message to the queue. Instead,
    call `chatroom.Logout` with the current username.
    1. :star2: Add some logic that checks if `cu.disconnect` is set, then exit the loop.
  
    [Stuck on any of the steps above? Ask your TA, or see the solution!](code/07-logouts/chat.go#L133-L141)

1. Finally, let's implement `ChatUser.Close`.  This will set the `disconnect` variable and close the socket.

  ```go
  func (cu *ChatUser) Close() {
  	// TODO: close the socket
  }
  ```
  1. :star2: In `Close()`, set `ChatUser.disconnect = true`
  1. :star2: Close the `ChatUser.conn`
  
  [Stuck on any of the steps above? Ask your TA, or see the solution!](code/07-logouts/chat.go#L184-L187)

1. Now when a user disconnects, you'll see a disconnect message!
  ```bash
   $ nc localhost 6677
   Welcome to Jen's chat server!
   Please enter your username: funcuddles
   Welcome, funcuddles
   *** funcuddles just joined the chatroom
   *** bob just joined the chatroom
   hello this is dog
   [funcuddles] hello this is dog
   [bob] bye funcuddles, I gotta go!
   *** bob has disconnected
   ```

