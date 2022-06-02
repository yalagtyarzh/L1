package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(18446744e10)
	b := big.NewInt(12352354e10)

	result := big.NewInt(0)
	fmt.Println("Sum:", result.Add(a, b))
	fmt.Println("Difference:", result.Sub(a, b))
	fmt.Println("Multiplex:", result.Mul(a, b))
	fmt.Println("Subtract:", result.Sub(a, b))
}
