package main

import (
	"context"
	"fmt"
	"time"
)

// Вариант через канал
//func main() {
//	ch := make(chan int, 100)
//	done := make(chan struct{})
//	go func() {
//		for {
//			select {
//			case ch <- doStuff():
//			case <-done:
//				close(ch)
//				return
//			}
//			time.Sleep(100 * time.Millisecond)
//		}
//	}()
//
//	go func() {
//		time.Sleep(3 * time.Second)
//		done <- struct{}{}
//	}()
//
//	for i := range ch {
//		fmt.Println("Получили: ", i)
//	}
//
//	fmt.Println("Заканчиваемся")
//}
//
//func doStuff() int {
//	return 1
//}

// Вариант через контекст
func main() {
	signal := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				signal <- struct{}{}
				return
			default:
				fmt.Println("xd")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-signal
	fmt.Println("Заканчиваемся")
}
