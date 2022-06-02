package main

import "fmt"

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}

func isUnique(str string) bool {
	checked := make(map[rune]struct{}, len(str))
	for _, v := range str {
		_, ok := checked[v]
		if ok {
			return false
		}
		checked[v] = struct{}{}
	}

	return true
}
