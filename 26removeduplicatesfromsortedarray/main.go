package main

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	pos := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			pos++
			nums[pos] = nums[i]
		}
	}

	return pos + 1

}

func main() {
	nums := []int{1, 1, 2, 2, 2, 3, 4, 4, 5}

	// Call the removeDuplicates function
	uniqueCount := removeDuplicates(nums)

	// Print the modified slice and the count of unique elements
	fmt.Printf("Modified Slice: %v\n", nums[:uniqueCount])
	fmt.Printf("Count of Unique Elements: %d\n", uniqueCount)
}

// Given an integer array nums sorted in non-decreasing order, remove the duplicates
// in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same.
// Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted,
// you need to do the following things:

// Change the array nums such that the first k elements of nums contain the
// unique elements in the order they were present in nums initially.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.
