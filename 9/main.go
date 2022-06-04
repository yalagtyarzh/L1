package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}

	// Создаем каналы записи и чтения
	wchan := make(chan int)
	rchan := make(chan int)

	// Запускаем горутины записи и чтения
	go write(slice, wchan)
	go read(rchan, wchan)

	// Читаем данные из rchan
	for v := range rchan {
		fmt.Println(v)
	}
}

// write получает данные из slice и записывает их в канал с, а по окончании закрывает канал c
func write(slice []int, c chan<- int) {
	for _, v := range slice {
		c <- v
	}

	close(c)
}

// read получает данные из wc и записывает их в канал rc, а по окончании закрывает канал rc
func read(rc chan<- int, wc <-chan int) {
	for v := range wc {
		rc <- v * 2
	}

	close(rc)
}
