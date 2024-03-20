package main

import "fmt"

func reverse(s string) string {
	rns := []rune(s)
	for i := 0; i < len(rns)/2; i++ {

		rns[i], rns[len(rns)-1-i] = rns[len(rns)-1-i], rns[i]
	}

	return string(rns)
}

func main() {
	str := "главрыба"
	fmt.Println(reverse(str))
}
