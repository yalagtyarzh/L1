package main

import (
	"context"
	"fmt"
	"time"
)

// Вариант через канал
//func main() {
//	// Создаем буферизированный канал, который будет что-то выполнять
//	ch := make(chan int, 100)
//	// Создаем канал, который будет сигналом для закрытия канала
//	done := make(chan struct{})
//	// запускаем горутину
//	go func() {
//		for {
//			select {
//			case ch <- doStuff():
//			// Останавливаемся, если получили сигнал
//			case <-done:
//				close(ch)
//				return
//			}
//			time.Sleep(100 * time.Millisecond)
//		}
//	}()
//
//	// Ждем сигнал после некоторого времени в еще одной горутине
//	go func() {
//		time.Sleep(3 * time.Second)
//		done <- struct{}{}
//	}()
//
//	for range ch {
//		//fmt.Println("Получили: ", <-ch)
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
	// Создаем канал, который прокидывает сигнал
	signal := make(chan struct{})
	// Создаем контекст для подачи сигнала о его отмеке каналу signal
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем горутину
	go func(ctx context.Context) {
		for {
			select {
			// Останавливаемся, если получили сообщение об отмене
			case <-ctx.Done():
				signal <- struct{}{}
				return
			default:
				fmt.Println("xd")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	// Отменяемся
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	// Как только получили сигнал - продолжаем работу в главной горутине
	<-signal
	fmt.Println("Заканчиваемся")
}
