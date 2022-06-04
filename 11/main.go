package main

import "fmt"

func main() {
	set1 := []string{"a", "b", "dog", "d", "f"}
	set2 := []string{"a", "x", "z", "dog", "f"}

	inter := intersect(set1, set2)
	fmt.Println(inter)
}

// intersect формирует множество, представляющее пересечение s1 и s
func intersect(s1, s2 []string) []string {
	// Помечаем все элементы s1 с помощью мапы
	hash := make(map[string]bool)
	for _, v := range s1 {
		hash[v] = true
	}

	// Начинаем сравнивать значения множеств и заполняем мапу пересечения
	inter := make(map[string]struct{})
	for _, v := range s2 {
		// Если пересечение имеется = устанавливаем ключ
		if hash[v] {
			inter[v] = struct{}{}
		}
	}

	// Превращаем мапу в слайс
	res := sliceFromMap(inter)
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
