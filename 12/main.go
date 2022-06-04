package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	set := getSet(slice)

	fmt.Println(set)
}

// getSet возвращает множество без дубликатов, переданные в strs
func getSet(strs []string) []string {
	set := make(map[string]struct{})
	// Заполняем ключи
	for _, v := range strs {
		set[v] = struct{}{}
	}

	// Превращаем слайс в мапу
	res := sliceFromMap(set)

	return res
}

// sliceFromMap возвращает слайс строк из мапы строка-пустая структура
func sliceFromMap(m map[string]struct{}) []string {
	res := make([]string, 0)
	for k := range m {
		res = append(res, k)
	}

	return res
}
