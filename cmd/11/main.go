package main

import "fmt"

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Проходим по каждому элементу первого множества
	for key := range set1 {
		// Если элемент присутствует и во втором множестве, добавляем его в результат
		if set2[key] {
			result[key] = true
		}
	}

	return result
}

func main() {
	// Пример двух неупорядоченных множеств
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Вызов функции пересечения
	result := intersection(set1, set2)

	// Вывод результата
	fmt.Println("Пересечение множеств:", result)
}
