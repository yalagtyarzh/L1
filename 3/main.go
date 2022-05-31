package main

import (
	"fmt"
	"math"
	"sync"
)

//// Вариант с каналом
//func main() {
//	slice := []float64{2, 4, 6, 8, 10}
//	// Создаем канал
//	c := make(chan float64)
//	// Инициализируем переменную, которая будет хранить в себе сумму чисел
//	var sum float64
//	// Запускаем горутины
//	for _, v := range slice {
//		go pow(c, v)
//		// Прибавляем к перемене sum значение из канала
//		sum += <-c
//	}
//
//	// Вывод
//	fmt.Println(sum)
//}
//
//// pow записывает в канал c квадрат число a
//func pow(c chan float64, a float64) {
//	c <- math.Pow(a, 2)
//}

// Вариант с WaitGroup, которая из себе представляет простейший счетчик для синхронизации
var wg sync.WaitGroup

func main() {
	slice := []float64{2, 4, 6, 8, 10}
	// Инициализируем переменную, которая будет хранить в себе сумму чисел
	var sum float64
	// Запускаем горутины
	for _, v := range slice {
		// Прибавляем счетчик на единичку
		wg.Add(1)
		// Полетела наша горутина
		go func(a float64) {
			// Не забываем отнять счетчик не единичку
			defer wg.Done()
			// Прибавляем к sum a^2
			sum += math.Pow(a, 2)
		}(v)
	}

	// Ждем обнуления счетчика
	wg.Wait()
	// Выводит sum
	fmt.Println(sum)
}
