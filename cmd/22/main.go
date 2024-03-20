package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Задаем значения переменных a и b, которые больше чем 2^20
	a := big.NewInt(2 << 20)
	b := big.NewInt(3 << 20)

	// Умножение
	mult := new(big.Int).Mul(a, b)
	fmt.Println("Умножение a * b:", mult)

	// Деление
	div := new(big.Int).Div(a, b)
	fmt.Println("Деление a / b:", div)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сложение a + b:", sum)

	// Вычитание
	sub := new(big.Int).Sub(a, b)
	fmt.Println("Вычитание a - b:", sub)
}
