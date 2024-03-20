package main

import "fmt"

func main() {
	f := 1
	s := 2

	f, s = s, f

	fmt.Println(f, s)
}
