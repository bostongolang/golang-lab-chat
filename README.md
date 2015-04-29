# Golang Boston Lab #1 - Building a Chat Server

Welcome to the first Boston Go language lab!  We are going to build a chat server in Go, then chat with each other!

The goal is to practice implementing some of the concurrency primitives we learned in the previous meetup, and 
get your hands dirty by writing some real Go code. 

## Prerequisites 

1. PLEASE BRING:  A laptop with a Go language environment already set up. Please see 'Vagrant setup' below to get started with one easily.
  * If you are having trouble setting up your Go programming environment, please join the #lab-help channel in the Boston Golang Slack group. You can signup for the Slack group [here](http://bostongolang-slack-invite.herokuapp.com/).

2. PLEASE BRING: A text editor or IDE suitable for writing Go code.
  * For beginners, [Sublime](http://www.sublimetext.com) is a good option. Make sure you install the Go plugin [here](https://github.com/DisposaBoy/GoSublime).
  * For VIM users, there is a pretty nice VIM setup [here](https://github.com/fatih/vim-go).  

3. PLEASE HAVE:  Some basic Go language exposure.  You should be familiar with the Go basics: e.g., Go's types, structs, control flow structures, goroutines, and channels.   
Other programming language experience and concepts (such as networking, etc) will be helpful. A good introduction to the basics of Go for people familiar with 
other programming language is available at: [https://tour.golang.org](https://tour.golang.org). If you can get through this tour, you will be well-prepared for this meetup!


### Vagrant setup

1. Install [Vagrant](http://www.vagrantup.com/downloads) for your platform.
2. Clone this repository into a directory
  
  ```bash
  # open a terminal window and type:
  $ git clone https://github.com/bostongolang/golang-lab-chat.git
  ```

3. Open a terminal window, and cd /path/to/this/repository.

  ```bash
  # open a terminal window and type:
  $ cd golang-lab-chat
  ```

4. From within the `golang-lab-chat` directory, type `vagrant up`. This will create a virtual machine with Ubuntu linux and Go 1.4 installed for the Vagrant user.
5. Type `vagrant ssh` to ssh into the virtual machine.  

Need help? Arrive a little early and we can help you get up-and-running, or join
the Boston Golang Slack group #lab-help channel group by signing up [here](http://bostongolang-slack-invite.herokuapp.com/).

You can also email me directly: [jandre+bostongolang@gmail.com](mailto:jandre+bostongolang@gmail.com).


## Organization

This lab is organized into the following steps -- if you ever need to 'cheat' and see an example of how the code is
implemented, we have provided all of the working code examples.

1. Setting up your Go project
1. Creating a TCP server
1. Writing to the socket: creating your server banner
1. Reading the username and storing in a user struct 
1. Creating a structure for the chatroom
