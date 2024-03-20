package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(id int, dataChannel <-chan int) {
	for data := range dataChannel {
		fmt.Printf("Воркер %d получил данные: %d\n", id, data)
		time.Sleep(time.Second)
	}
}

func main() {
	numWorkers := 5
	dataChannel := make(chan int, 1000)

	// Запуск воркеров
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, dataChannel)
	}

	// Постоянная запись данных в канал
	// Ожидание сигнала завершения (Ctrl+C) для корректного завершения программы
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)
	for i := 0; ; i++ {
		if len(terminate) == 1 {
			close(dataChannel)
			break
		}
		dataChannel <- i
	}

	fmt.Println("Завершение работы...")
	fmt.Println("Программа завершила работу.")
}
