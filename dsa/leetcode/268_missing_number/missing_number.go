package main

// Leetcode Number 268
// Given an array nums containg n distinc umber in the range [0, n], return the only number in the range that is missing from the array
func main() {
	var nums = []int{5, 2, 1, 0, 3}

	findMissingNumber(nums)
}

func findMissingNumber(array []int) int {
	var n = len(array)                // array length
	var expectedSum = (n + 1) * n / 2 // expected sum from 0 to n
	var actualSum = 0
	for _, v := range array {
		actualSum += v // Sum all values inside the array
	}

	return expectedSum - actualSum // Return the diference between the expeted result and sum. There it is the missing number
}
