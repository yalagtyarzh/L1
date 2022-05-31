package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// Интерфейс для воркера
type Worker interface {
	Do()
}

// Структура воркера со своим ID
type worker struct {
	ID int
}

func main() {
	// Получаем количество воркеров с cli
	var amount int
	fmt.Print("Введите количество воркеров:  ")
	fmt.Scan(&amount)
	if amount < 1 {
		log.Fatal("Количество воркеров не может быть меньше одного, падаю")
	}

	// Инициализируем обработчик, отвечающий за получения сигнала прерывания программы (Ctrl + C)
	//initCloseHandler()

	// Инициализируем воркеры и заставляем их работать (хи-хи ха-ха)
	workers := make([]Worker, amount)
	for i := range workers {
		workers[i] = &worker{i + 1}
		go workers[i].Do()
	}

	// Что-то делаем в бесконечном цикле
	//for {
	//
	//}

	// Если хотим запустить обработчик в горутине main - инициализируем его в конце
	initCloseHandler()
}

// initCloseHandler создает канал, который ловит сигнал os.Interrupt (ctrl+c)
// Если поймали os.Interrupt - выходим из программы
func initCloseHandler() {
	c := make(chan os.Signal)
	// Обозначаем, какие сигналы ОС "ловить"
	signal.Notify(c, os.Interrupt)
	// go
	func() {
		// Получили сигнал - выходим
		<-c
		fmt.Println("\nПрожали Ctrl+c - выходим из программы")
		os.Exit(0)
	}()
}

// Do "заставляет" воркера работать
func (w *worker) Do() {
	for {
		fmt.Printf("Worker %d сделал работу и пошел спать\n", w.ID)
		time.Sleep(time.Second)
	}
}
