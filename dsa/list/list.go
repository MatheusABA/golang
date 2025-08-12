package main

import (
	"fmt"

	"slices"
)

func main() {

	// Creating an array with fixed size
	array := [5]int{}

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	// iterating over the array
	// for i := range array {
	// 	fmt.Println(array[i])
	// }

	// fmt.Printf("Array: %v\n", array)

	// printArray(array[:]) // passing the array as a slice

	slice := []int{} // creating a slice with not fixed size
	// We can use make to create a slice with a specific length and capacity
	// creating a slice with length 0 and capacity 5
	// slice = make([]int, 0, 5)

	// Append only work with slices, not with arrays
	slice = append(slice, 1) // adding an element to the slice
	slice = append(slice, 2) // adding another element to the slice
	slice = append(slice, 5) // adding a third element to the slice

	// iterating over the slice
	// for i := range slice {
	// 	fmt.Println(slice[i])
	// }

	// fmt.Printf("Slice: %v\n", cap(slice))

	printArray(slice)

	slice = removeEvenNumbers(slice) // removing even numbers
	printArray(slice)

	slice = slices.DeleteFunc(slice, func(value int) bool {
		return value == 5
	})

	printArray(slice)

}

func printArray(array []int) {
	for _, value := range array {
		fmt.Printf("%d \t ", value)
	}
	fmt.Println()
}

func removeEvenNumbers(array []int) (notEvenArray []int) {
	for _, value := range array {
		if value%2 != 0 { // value is just like array[i] if the for loop was for i := 0; i < len(array); i++
			notEvenArray = append(notEvenArray, value)
		}
	}

	return notEvenArray
}
