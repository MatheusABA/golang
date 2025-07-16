package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for range 5 {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

// Change main to main1 to avoid conflict with main function
func main1() {
	// Direct call to the function
	fun("Direct call")

	// Calling the function in a goroutine
	go fun("First goroutine")

	qr := fun

	go qr("Second goroutine")

	// Goroutine as an anonymous function
	go func() {
		fun("Third goroutine")
	}()

	// Ensuring the goroutine has time to run before the main function exits
	time.Sleep(10 * time.Millisecond)

	fmt.Println("Done...")
}
