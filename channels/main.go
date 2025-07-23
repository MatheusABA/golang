package main

import "fmt"

// This program demonstrates a simple use of channels in Go.
// It defines two functions: ping and pong, which communicate through channels.
// The ping function sends a message to a channel, and the pong function receives that message
// and sends it to another channel. The main function orchestrates the communication.

// To send a message to a channel, the channel must be defined with a send-only type (chan<-), for example: ch := make(chan<- string)
// To receive a message from a channel, the channel must be defined with a receive-only type (<-chan), for example: ch := make(<-chan string)

// chan<- sends only, <-chan receives only

/*
Receive a send channel<- and insert the string message into it.

*/
func ping(pingChannel chan<- string, msg string) {
	pingChannel <- msg // Send message to pingChannel
}

/*
Receive a send channel<- and a receive channel<-.
Receive a message from the receive channel and send it to the send channel.
*/
func pong(pongChannel chan<- string, pingChannel <-chan string) {
	msg := <-pingChannel // Receive message from pingChannel
	pongChannel <- msg   // Send the same message to pongChannel
}

func main() {
	pingChannel := make(chan string) // Create a channel for ping
	pongChannel := make(chan string) // Create a channel for pong

	go ping(pingChannel, "Hello")     // Call ping with the pingChannel and a message
	go pong(pongChannel, pingChannel) // Call pong with the pongChannel and pingChannel

	// fmt.Println("Ping:", <-pingChannel) // Print the message received from ping -> commented out as it is not used in this example
	fmt.Println("Pong:", <-pongChannel) // Print the message received from pong
}
