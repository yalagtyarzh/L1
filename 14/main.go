package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = make(chan interface{})
	t := checkType(a)
	fmt.Println(t)
}

// checkType возвращает тип val
func checkType(val interface{}) reflect.Type {
	return reflect.TypeOf(val)
}
