package main

import (
	"fmt"
	"math"
)

// Вариант с WaitGroup, которая из себе представляет простейший счетчик для синхронизации
//var wg sync.WaitGroup
//
//func main() {
//	// Инициализируем слайс
//	slice := []float64{2, 4, 6, 8, 10}
//
//	// Запускаем горутину на каждый элемент слайса
//	for _, v := range slice {
//		// Прибавляем счетчик на единичку
//		wg.Add(1)
//		go printPow(v)
//	}
//
//	// Ждем, пока счетчик не обнулится
//	wg.Wait()
//}
//
//// Выводит число типа float64 на экран и уменьшает счетчик WaitGroup на единичку
//func printPow(num float64) {
//	fmt.Println(math.Pow(num, 2))
//	wg.Done()
//}

// Вариант с каналами
func main() {
	slice := []float64{2, 4, 6, 8, 10}
	// Создаем канал
	c := make(chan float64)
	// Запускаем горутины
	for _, v := range slice {
		go pow(c, v)
		// Выводим значения из канала
		fmt.Println(<-c)
	}
}

// pow записывает в канал c квадрат число a
func pow(c chan float64, a float64) {
	c <- math.Pow(a, 2)
}
