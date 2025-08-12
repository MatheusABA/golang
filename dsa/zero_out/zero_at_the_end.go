package main

import "fmt"

// Implementing a module to move all zeros to the end of an array
func main() {
	originalArray := []int{1, 0, 2, 0, 3, 4, 0, 5}
	fmt.Printf("Original Array: %v\n", originalArray)
	movedArray := moveZerosToEnd(originalArray)
	fmt.Printf("Moved Array: %v\n", movedArray)
}

func moveZerosToEnd(array []int) []int {
	arrayWithZeros := make([]int, 0, len(array)) // slice to hold non-zero elements
	zeroCounts := 0                              // variable to help counting how much zeros we have

	for _, value := range array {
		if value != 0 { // iterate over the original array to create a new array with non zero values at the start
			arrayWithZeros = append(arrayWithZeros, value) // add non-zero elements to the new slice
		} else {
			zeroCounts++ // count how much zeros we need to add at the end of the new array
		}
	}

	for i := 0; i < zeroCounts; i++ { // iterate over the zeroCounts to append zeros at the end
		arrayWithZeros = append(arrayWithZeros, 0) // append zeros at the end
	}

	return arrayWithZeros // return the modified array
}
