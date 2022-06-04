package main

import "fmt"

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}

// isUnique проверяет, все ли символы в строке уникальны
func isUnique(str string) bool {
	// Создаем мапу, ключами которого являются символы строки
	checked := make(map[rune]struct{}, len(str))
	// Пробегаемся по строке
	for _, v := range str {
		_, ok := checked[v]
		// Если в мапе уже есть символ v - вернуть false
		if ok {
			return false
		}
		// Помечаем символ как проверенный
		checked[v] = struct{}{}
	}

	return true
}
