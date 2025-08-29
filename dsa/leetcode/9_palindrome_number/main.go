package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 544445
	fmt.Printf("Is palidrone? %t \n", isPalindrome(n))
}

func isPalindrome(x int) bool {

	// if x < 0 {
	// 	return false
	// }

	// original := x
	// reversed := 0
	// for x > 0 {
	// 	reversed = reversed*10 + x%10 // x%10 -> get last digit from x     reversed*10 -> push inverted digits to left
	// 	x /= 10                       // Update x to remove last digit
	// }
	// return original == reversed

	// Simple way
	// Convert int to string
	intConvertion := strconv.Itoa(x)
	start := 0
	end := len(intConvertion) - 1

	for start < end {
		numberStart := intConvertion[start]
		numberEnd := intConvertion[end]

		if numberStart != numberEnd {
			return false
		}

		start++
		end--
	}

	return true
}
