package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Получаем количество секунд для таймера с cli
	var sec int
	fmt.Print("Введите количество секунд:  ")
	fmt.Scan(&sec)
	if sec < 0 {
		log.Fatal("Количество секунд не может быть меньше нуля, падаю")
	}

	// Создаем канал для чтения и записи данных
	c := make(chan int)
	// Начинаем записывать в канал и читать из него
	go write(c)
	go read(c)

	// В main горутине спим пока не истечет таймер
	time.Sleep(time.Duration(sec) * time.Second)
}

// write записывает данные в канал и засыпает
func write(c chan int) {
	for {
		c <- 79
		time.Sleep(time.Second)
	}
}

// read читает данные из канала
func read(c chan int) {
	for {
		fmt.Println(<-c)
	}
}
