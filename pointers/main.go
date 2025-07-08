package main

import "fmt"

func main() {

	// Pointers are variables that store the memory address of another variable.
	// & is used to get the address of a variable
	// * is used to declare a pointer in Go.

	// Reminder of how pointers work
	a := 42
	b := &a // b is a pointer to a, so, b holds the address of a. To access the real value of b we need to dereference it using *
	// or we can use new(pointer_type) to create a pointer
	// a := 10
	// b := new(int)
	// *b := &a

	// Printing values
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of b:", b)
	fmt.Println("Value pointed by b:", *b)
	fmt.Println("Address of b:", &b)

	// Dereferencing the pointer to change the value of a
	*b = 100

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of b:", b)
	fmt.Println("Value pointed by b:", *b)
	fmt.Println("Address of b:", &b)
}
