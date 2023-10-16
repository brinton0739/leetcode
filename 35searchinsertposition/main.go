package main

import "fmt"

func searchInsert(nums []int, target int) int {

	for i, num := range nums {
		if num == target || num > target {
			return i
		}
	}
	return len(nums)

}

func main() {
	nums := []int{1, 2, 5, 6}
	target := 7
	result := searchInsert(nums, target)
	fmt.Println("Index:", result)
}
