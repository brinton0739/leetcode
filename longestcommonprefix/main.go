package main

import "fmt"

func longestCommonPrefix(str []string) string {

	if len(str) == 0 {
		return ""
	}
	// we need to find the length of the shortest string in the array.
	minLen := len(str[0])

	for _, strs := range str {
		if len(strs) < minLen {
			minLen = len(strs)
		}
	}
	// we want to next compare the characters of the strings

	for i := 0; i < minLen; i++ {
		char := str[0][i] //get the i-th character of the first string
		for j := 1; j < len(str); j++ {
			if str[j][i] != char {
				return str[0][:i] // return the common prefix found so far.
			}
		}
	}

	return str[0][:minLen] //return the common prefix found in all strings
}

func main() {

	input := []string{"flower", "flour", "flourish"}
	result := longestCommonPrefix(input)

	fmt.Println("Longest common prefix:", result)

}
