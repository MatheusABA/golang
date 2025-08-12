package main

import "fmt"

/*
We can pass the capacity of a channel as an argument to the make function.
This allows us to create a buffered channel, which can hold a certain number of values before it blocks.
Buffered channels are useful when we want to allow some amount of data to be stored in the channel without requiring immediate processing.
Buffered channels can also help to choose how much goroutines can run concurrently without blocking.

Example:
func main() {
	ch := make(chan int, 5) // Create a buffered channel with a capacity of 5
	ch <- 1                  // Send value to the channel
	ch <- 2                  // Send another value
	fmt.Println(<-ch)        // Receive value from the channel
	fmt.Println(<-ch)        // Receive another value
}

In the other ways, we can create an unbuffered channel by not specifying a capacity, or by using a capacity of 0.
Example:
func main() {
	ch := make(chan int) // Create an unbuffered channel
	go func() {
		ch <- 1 // Send value to the channel
	}()
	fmt.Println(<-ch) // Receive value from the channel
}

*/
func main() {
	fmt.Println("Hello, World!")
}
