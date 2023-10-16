package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	// this is a common way to solve thing like this with the
	//fibonacci sequence
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	fmt.Println(climbStairs(2)) // Output: 2
	fmt.Println(climbStairs(3)) // Output: 3
	fmt.Println(climbStairs(4)) // Output: 3
	fmt.Println(climbStairs(5)) // Output: 3
}
