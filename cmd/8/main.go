package main

import (
	"fmt"
)

// Функция для установки i-го бита в значении n в 1 или 0
func setBit(n int64, i uint, bitValue uint) int64 {
	// Создаем маску, установив i-й бит в 1
	mask := int64(1 << i)

	// Если bitValue равен 1, устанавливаем i-й бит в 1 путем выполнения побитового ИЛИ с маской
	if bitValue == 1 {
		return n | mask
	}

	// Если bitValue равен 0, устанавливаем i-й бит в 0 путем выполнения побитового И с инвертированной маской
	return n &^ mask
}

func main() {
	// Исходное значение
	var n int64 = 42
	fmt.Printf("Исходное значение: %d\n", n)

	// Устанавливаем 3-й бит в 1
	n = setBit(n, 3, 1)
	fmt.Printf("Установлен 3-й бит в 1: %d\n", n)

	// Устанавливаем 5-й бит в 0
	n = setBit(n, 5, 0)
	fmt.Printf("Установлен 5-й бит в 0: %d\n", n)
}
