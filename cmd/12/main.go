package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, el := range s {
		set[el] = true
	}

	fmt.Println(set)
}
