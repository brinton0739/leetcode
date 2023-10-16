package main

import (
	"fmt"
	"strconv"
)

func main() {
	word := "word"
	fmt.Println(Solution(word))
	fmt.Println(isPalindrome(2002))
	fmt.Println(isPalindromeLazy(12345))
}

func isPalindrome(x int) bool {
	// Negative numbers can't be palindromes since the '-' sign makes them different
	if x < 0 {
		return false
	}

	// Initialize variables to store the reversed number and the original number
	reversed := 0
	original := x

	// Reverse the number while keeping track of the result in 'reversed'
	for x != 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return original == reversed
}

func isPalindromeLazy(x int) bool {
	s := strconv.Itoa(x)
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return s == string(rns)

}

func isPalindromeTwo(x int) bool {
	s := strconv.Itoa(x)

	rns := []rune(s)

	var i = 0
	var j = len(rns) - 1

	for i < j {
		rns[i], rns[j] = rns[j], rns[i]
		i++
		j--
	}
	return s == string(rns)
}

func Solution(word string) string {

	rns := []rune(word)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)

}
