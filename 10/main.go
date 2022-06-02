package main

import (
	"fmt"
)

func main() {
	nums := []float64{-23.0, -24.0, -21.0, 24.5, 41.0, 13.0, 19.0, 15.5}
	res := groupNums(nums)
	fmt.Println(res)
}

func groupNums(nums []float64) map[int][]float64 {
	res := make(map[int][]float64)
	for _, v := range nums {
		idx := int(v/10) * 10
		res[idx] = append(res[idx], v)
	}

	return res
}
