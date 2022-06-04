package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	idx := search(slice, 9)
	fmt.Println(idx)
}

// Thanks to Grokking Algorithms
// search ищет значение target в массиве и возвращает индекс target, если он был найден, иначе -1
func search(nums []int, target int) int {
	// Определяем границы поиска в массиве
	left := 0
	right := len(nums) - 1
	// Пока границы не "сомкнлуись" - ищем
	for left <= right {
		// Вычисляемый средний элемент
		mid := (left + right) / 2
		// Если найденный элемент меньше искомого - смещаем левую границу в mid
		// и прибавляем единичку, ведь mid мы уже проверяли
		if nums[mid] < target {
			left = mid + 1

		} else
		// Если найденный элемент больше искомого - смещаем правую границу в mid
		// и прибавляем единичку
		if nums[mid] > target {
			right = mid - 1
		} else
		// В ином случае мы нашли элемент, возвращаем индекс
		{
			return mid
		}
	}

	// Если проверили массив и ничего не нашли - возвращаем -1
	return -1
}
