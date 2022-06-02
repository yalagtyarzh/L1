package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}

	wchan := make(chan int)
	rchan := make(chan int)

	go write(slice, wchan)
	go read(rchan, wchan)

	for v := range rchan {
		fmt.Println(v)
	}
}

func write(slice []int, c chan<- int) {
	for _, v := range slice {
		c <- v
	}

	close(c)
}

func read(rc chan<- int, wc <-chan int) {
	for v := range wc {
		rc <- v * 2
	}

	close(rc)
}
