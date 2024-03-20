package main

import "sync"

func main() {
	m := make(map[string]int)

	numWorkers := 5

	// Создание мьютекса для синхронизации доступа к карте
	var mu sync.Mutex

	// Создание WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m["test"] = 1
			mu.Unlock()
		}()
	}

	wg.Wait()
}
