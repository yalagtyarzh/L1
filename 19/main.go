package main

import (
	"fmt"
)

func main() {
	phrase := "Sera was never an agreeable girl"
	fmt.Println(reverse(phrase))
}

// reverse переворачивает порядок символов в строке
func reverse(str string) string {
	// Преобразовываем строку в массив рун
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		// Меняем символы местами
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
