package main

import (
	"fmt"
	"math/big"
)

// В Go операции с большими числами выполняются с помощью пакета math/big
func main() {
	// Инициализируем числа
	a := big.NewInt(18446744e10)
	b := big.NewInt(12352354e10)

	// Создаем промежуточную переменную, с помощью которой осуществляем операции и выводим их на экран
	result := big.NewInt(0)
	fmt.Println("Sum:", result.Add(a, b))
	fmt.Println("Difference:", result.Sub(a, b))
	fmt.Println("Multiplex:", result.Mul(a, b))
	fmt.Println("Subtract:", result.Sub(a, b))
}
