package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, found := numMap[complement]; found {

			return []int{nums[idx], num}
		}

		numMap[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{1, 3, 5}
	target := 8
	result := twoSum(nums, target)
	fmt.Println(result)
}
