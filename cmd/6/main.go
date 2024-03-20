package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(done chan bool) {
	// Работа горутины
	// Когда работа завершена, отправляем сигнал в канал
	done <- true
}

func workerCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Получен сигнал об отмене контекста
			return
		default:
			// Работа горутины
		}
	}
}

func main() {
	// Создаем канал для уведомления о завершении работы горутины
	done := make(chan bool)

	// Запускаем горутину
	go worker(done)

	// Ждем, пока горутина завершится
	<-done

	//**********************//

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Отложенный вызов cancel() для освобождения ресурсов

	// Запускаем горутину
	go workerCtx(ctx)

	// После некоторого времени отменяем выполнение
	time.Sleep(1 * time.Second)
	cancel()

	//**********************//

	// Создаем канал для прослушивания сигналов ОС
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Ждем получения сигнала
	<-sig
	// Здесь можно добавить логику завершения программы
}
