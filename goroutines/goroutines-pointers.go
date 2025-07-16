package main

import (
	"fmt"
	"time"
)

func goroutinePointer(value *string) {
	for {
		fmt.Println("Ptr pointed", value)  // will print address that value points to
		fmt.Println("Ptr value", *value)   // will print value
		fmt.Println("Ptr address", &value) // will print address of pointer
		time.Sleep(4 * time.Millisecond)
	}
}

func main() {
	value := "Goroutine Pointer" // Initialize a string variable
	ptr := &value                // Create a pointer to the string address
	var ptr2 *string = &value    // Another pointer to the same address

	go goroutinePointer(ptr)

	time.Sleep(time.Second) // Ensuring the goroutine has time to start

	*ptr2 = "Updated Value" // Update the value through the pointer

	time.Sleep(time.Second) // Ensure the goroutine has time to run
}
