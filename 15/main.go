package main

import "fmt"

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10)
	// Нарезаем слайс БАЙТОВ, а не рун, потому при работе с символами, занимающими
	// 1 байта памяти, мы можем столкнуться с некоторыми проблемами
	// В данном случае мы выводим меньшее количество символов, чем выбираем из строки
	justString = v[:100]

	// КОРРЕКТНЫЙ ВАРИАНТ
	// Приводим строку к слайсу рун, а после вытаскиваем из слайса 100 рун, приведя их обратно к строке
	runes := []rune(v)
	trueString := string(runes[:100])

	fmt.Println(justString)
	fmt.Println(trueString)
}

func createHugeString(size int) string {
	var s string
	for i := 0; i < size; i++ {
		s += "ゴ"
	}

	return s
}
