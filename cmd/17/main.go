package main

import "fmt"

func binarySearch(arr []int, target int) int {
	// Определяем нижнюю и верхнюю границу массива
	low, high := 0, len(arr)-1

	for low <= high {
		// Вычисляем средний индекс
		mid := low + (high-low)/2
		// Провряем не нашли ли мы элемент
		if arr[mid] == target {
			return mid
		}

		// Проверяем число слева или справа от середины
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	target := binarySearch(a, 4)
	fmt.Println(target)
}

// Начиная с версии SDK 1.21 будет работать
//import (
//	"fmt"
//	"slices"
//)
//
//func main() {
//	a := []int{1, 2, 3, 4, 5, 6, 7}
//
//	n, found := slices.BinarySearch(a, 4)
//	fmt.Println("4:", n, found)
//}
