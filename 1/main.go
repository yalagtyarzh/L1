package main

import "fmt"

// Создаем "родительскую"
type Human struct{}

// Создаем структуру, которая встраивает в себя структуру Human
type Action struct {
	Human
}

// Метод "родительской" структуры
func (h *Human) do() {
	fmt.Println("Doing something")
}

func main() {
	a := Action{}
	// Наследуем метод
	a.do()
}
