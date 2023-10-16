package main

import "fmt"

func strStr(haystack string, needle string) int {

	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen == 0 {
		return 0
	}

	for i := 0; i <= haystackLen-needleLen; i++ {

		if haystack[i:i+needleLen] == needle {
			return i

		}

	}

	return -1

}

func main() {
	haystack := "Hello, world!"
	needle := "world"

	index := strStr(haystack, needle)

	if index == -1 {
		fmt.Println("Needle not found in haystack.")
	} else {
		fmt.Printf("Needle found at index %d\n", index)
	}
}
