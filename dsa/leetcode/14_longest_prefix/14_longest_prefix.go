package main

import "fmt"

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

*/
func main() {

	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	index := 0
	longestCommonPrefix := strs[index]

	for i := 0; i < len(secondString); i++ {
		if firstString[i] == secondString[i] {
			longestCommonPrefix += string(firstString[i])
		}
	}

}
