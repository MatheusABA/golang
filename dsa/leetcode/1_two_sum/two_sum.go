// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

package main

import "fmt"

func main() {

	fmt.Println(findTwoSum([]int{2, 7, 11, 15}, 9))
}

// First try
// func findTwoSum(nums []int, target int) []int {
// 	for i := range nums { // First loop starts in 0
// 		for j := i + 1; j < len(nums); j++ { // Second loop start in 1
// 			if nums[i]+nums[j] == target { // Verify if sum between this 2 positions match target
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

func findTwoSum(nums []int, target int) []int {
	// Create a map to store numbers we've seen so far.
	// Key: number from nums, Value: index of that number.
	numsMap := make(map[int]int)

	// Loop through the array.
	for i, v := range nums {
		// For each number v, calculate its complement (target - v).
		// Check if we've already seen the complement before.
		if index, ok := numsMap[target-v]; ok {
			// If yes, return the indices: [index of complement, current index].
			return []int{index, i}
		}
		// If not, store the current number and its index in the map.
		numsMap[v] = i
	}

	// If no solution found, return nil.
	return nil
}
