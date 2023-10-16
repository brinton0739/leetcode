package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	// Initialize variables for result, carry, and indices
	var result strings.Builder
	carry := 0
	i, j := len(a)-1, len(b)-1

	// Loop from the least significant bit to the most significant bit
	for i >= 0 || j >= 0 || carry > 0 {
		// Extract the current bits from a and b
		bitA, bitB := 0, 0
		if i >= 0 {
			bitA = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			bitB = int(b[j] - '0')
			j--
		}

		// Calculate the sum of bits and carry
		sum := bitA + bitB + carry
		result.WriteByte(byte(sum%2 + '0')) // Append the sum bit to the result

		// Calculate the carry for the next iteration
		carry = sum / 2
	}

	// Reverse the result string
	return reverse(result.String())
}

func reverse(s string) string {
	// Helper function to reverse a string
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	a := "11"
	b := "111"
	result := addBinary(a, b)
	fmt.Println(result) // Output: "100"

	a = "1010"
	b = "1011"
	result = addBinary(a, b)
	fmt.Println(result) // Output: "10101"
}
