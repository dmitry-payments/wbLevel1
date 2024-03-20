package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5 // Время в секундах

	// Создаем канал для обмена данными
	ch := make(chan int)

	// Горутина для отправки данных в канал
	go func() {
		for i := 1; ; i++ {
			ch <- i                 // Отправляем значение в канал
			time.Sleep(time.Second) // Ждем 1 секунду перед отправкой следующего значения
		}
	}()

	// Чтение значений из канала и вывод их в стандартный вывод
	for {
		select {
		case val := <-ch:
			fmt.Println("Принято значение из канала:", val)
		case <-time.After(time.Duration(N) * time.Second): // Ждем N секунд перед завершением программы
			fmt.Println("Программа завершает работу...")
			return
		}
	}
}
