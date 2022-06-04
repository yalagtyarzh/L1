package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := []int{10, 8, 6, 4, 2}
	slice = sort(slice)
	fmt.Println(slice)
}

// Thanks to Grokking Algorithms
// sort сортирует слайс arr по алгоритму quicksort и возвращает отсортированный массив
func sort(slice []int) []int {
	// Базовый случай
	// Если слайс имеет 1 или 0 элементов, то он уже отсортирован, потому возвращаем его
	if len(slice) < 2 {
		return slice
	} else
	// Рекурсивный случай
	{
		// Рандомно выбираем опорный элемент
		rand.Seed(time.Now().Unix())
		pivotIndex := rand.Intn(len(slice))
		pivot := slice[pivotIndex]

		// Удаляем из исходного слайса опорный элемент
		slice = append(slice[:pivotIndex], slice[pivotIndex+1:]...)

		// Формируем два слайса
		// left - все числа, которые меньше опорного элемента
		var left []int
		// right - все числа, которые больше опорного элемента
		var right []int

		// Заполняем слайсы
		for _, item := range slice {
			if item <= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}

		// Рекурсивно сортируем подслайсы и собираем получившиеся результаты воедино
		left = sort(left)
		// Возвращаем опорный элемент в слайс
		left = append(left, pivot)
		right = sort(right)

		// Возвращаем отсортированный слайс, распаковывая все элементы right
		// и добавляя их к left
		return append(left, right...)
	}
}
