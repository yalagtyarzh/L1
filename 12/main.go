package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	set := getSet(slice)

	for k := range set {
		fmt.Printf("%s ", k)
	}
}

func getSet(strs []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, v := range strs {
		set[v] = struct{}{}
	}

	return set
}
