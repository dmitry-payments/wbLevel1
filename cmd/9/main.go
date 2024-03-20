package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	numbers := [5]int{1, 2, 3, 4, 5}

	go func() {
		for _, num := range numbers {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	for num := range ch2 {
		fmt.Println(num)
	}
}
