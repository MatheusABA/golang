package main

import "fmt"

// Implementing a module to move all zeros to the end of an array
func main() {
	originalArray := []int{0, 1, 0, 3, 12}
	fmt.Printf("Original Array: %v\n", originalArray)
	movedArray := moveZerosToEnd(originalArray)
	fmt.Printf("Moved Array: %v\n", movedArray)
}

func moveZerosToEnd(nums []int) []int {

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	return nums
}
