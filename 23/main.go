package main

import (
	"fmt"
	"log"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	slice, err := remove(slice, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slice)
}

func remove(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return nil, fmt.Errorf("invalid index")
	}

	return append(s[:i], s[i+1:]...), nil
}
