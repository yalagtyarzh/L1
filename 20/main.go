package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "Sera was never an agreeable girl"
	fmt.Println(reverse(phrase))
}

func reverse(str string) string {
	slice := strings.Split(str, " ")
	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return strings.Join(slice, " ")
}
