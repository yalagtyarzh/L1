package main

import "fmt"

func main() {
	a, b := 10, 20
	a, b = b, a
	fmt.Printf("a: %d, b: %d", a, b)
}
