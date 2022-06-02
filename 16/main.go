package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := []int{10, 8, 6, 4, 2}
	slice = sort(slice)
	fmt.Println(slice)
}

// Thanks to Grokking Algorithms
func sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		rand.Seed(time.Now().Unix())
		pivotIndex := rand.Intn(len(arr))
		pivot := arr[pivotIndex]

		arr = append(arr[:pivotIndex], arr[pivotIndex+1:]...)

		var left []int
		var right []int

		for _, item := range arr {
			if item <= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}

		left = sort(left)
		left = append(left, pivot)
		right = sort(right)

		return append(left, right...)
	}
}
