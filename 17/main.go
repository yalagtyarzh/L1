package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	idx := search(slice, 9)
	fmt.Println(idx)
}

// Thanks to Grokking Algorithms
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
