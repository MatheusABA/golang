// Given an integer n, return true if it is a power of three. Otherwise, return false.

// An integer n is a power of three, if there exists an integer x such that n == 3x.

package main

import "fmt"

func main() {
	n := 45
	fmt.Printf("Is %d power of three? %t \n", n, isPowerOfThree(n))
}

func isPowerOfThree(n int) bool {
	const maxPowerOfThree = 1162261467
	return n > 0 && maxPowerOfThree%n == 0
}
