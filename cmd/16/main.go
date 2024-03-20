package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 2, 6, 5}

	sort.Ints(a)
	fmt.Println(a)
}
