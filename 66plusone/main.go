package main

import "fmt"

func plusOne(digits []int) []int {

	n := len(digits)

	digits[n-1]++

	for i := n - 1; i > 0 && digits[i] == 10; i-- {

		digits[i] = 0
		digits[i-1]++
	}

	if digits[0] == 10 {
		digits[0] = 1
		digits = append(digits, 0)
	}

	return digits

}

func main() {

	digits1 := []int{1, 2, 3}
	digits2 := []int{4, 3, 2, 1}
	digits3 := []int{9, 9, 8, 9}

	fmt.Println(plusOne(digits1))
	fmt.Println(plusOne(digits2))
	fmt.Println(plusOne(digits3))
}
