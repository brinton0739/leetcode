package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {

	for i := 0; i < x+1; i++ {
		if i*i == x {
			return i
		} else if i*i > x {
			return i - 1
		}
	}
	return 0
}

func main() {
	fmt.Println(math.Floor(math.Sqrt(8)))
	fmt.Println(mySqrt(6))
}

func mySqrtx(x int) int {
	ans, odd := 0, 1
	for {
		x -= odd
		if x < 0 {
			break
		}
		odd += 2
		ans++
	}
	return ans
}
