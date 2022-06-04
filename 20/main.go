package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "Sera was never an agreeable girl"
	fmt.Println(reverse(phrase))
}

// reverse переворачивает порядок слов в строке
func reverse(str string) string {
	// Получаем "словарь" строки
	slice := strings.Split(str, " ")
	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		// Меняем слова местами
		slice[i], slice[j] = slice[j], slice[i]
	}

	return strings.Join(slice, " ")
}
