package main

import (
	"fmt"
	"time"
)

// sleep останавливает выполнение на d времени
func sleep(d time.Duration) {
	// Блокируемся на d секунд, а позже продолжаем выполнение
	<-time.After(d)
}

func main() {
	fmt.Println("Спим")
	sleep(5 * time.Second)
	fmt.Println("Проснулись, завершаемся")
}
