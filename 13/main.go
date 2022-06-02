package main

import "fmt"

// Фича Go
func main() {
	a, b := 10, 20
	a, b = b, a
	fmt.Printf("a: %d, b: %d", a, b)
}

//// Вариант костыльный
//func main() {
//	a, b := 23452345, 2087654
//	a += b
//	b = a - b
//	a -= b
//	fmt.Println(a, b)
//}
