package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	slice := []int{2, 4, 6, 8, 10}

	wchan := make(chan int)
	rchan := make(chan int)

	wg.Add(1)
	go write(slice, wchan)
	wg.Add(1)
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
