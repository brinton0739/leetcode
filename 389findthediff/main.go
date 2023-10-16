package main

import "fmt"

// func findTheDifference(s string, t string) byte {
// 	result := byte(0)

// 	// Calculate XOR for all characters in s and t
// 	for i := 0; i < len(s); i++ {
// 		result ^= s[i]
// 		result ^= t[i]
// 	}

// 	// XOR the last character in t
// 	result ^= t[len(t)-1]

// 	return result
// }

func main() {
	s := "abcd"
	t := "abcde"
	addedLetter := findTheDifferenceMap(s, t)
	fmt.Println(string(addedLetter)) // Output: "e"
}

func findTheDifferenceMap(s string, t string) byte {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for _, char := range t {
		charCount[char]--
		if charCount[char] < 0 {
			return byte(char)
		}
	}

	return 0 // Default return if no difference is found
}
