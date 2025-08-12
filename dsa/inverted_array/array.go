package main

import (
	"fmt"
	"slices"
)

func main() {

	originalArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	invertedArray := invertArray(originalArray)

	fmt.Printf("Original Array: %v\n", originalArray)
	fmt.Printf("Inverted Array: %v\n", invertedArray)

	// Usings the slices package
	slices.Reverse(originalArray)
	fmt.Printf("Reversed Original Array: %v\n", originalArray)
}

func invertArray(array []int) []int {
	inverted := make([]int, len(array)) // creating a new slice for the inverted array with the same length as the original
	for index, value := range array {
		position := len(array) - 1 - index // calculating the position for the inverted array
		inverted[position] = value
		fmt.Printf("Index: %d, Value: %d\n", position, value)
	}
	return inverted
}
