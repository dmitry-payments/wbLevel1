package main

import "fmt"

func isUnique(s string) bool {
	seen := make(map[rune]bool)

	for _, char := range s {
		charLower := toLower(char)

		if seen[charLower] {
			return false
		}

		seen[charLower] = true
	}

	return true
}

func toLower(char rune) rune {
	if 'A' <= char && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func main() {
	fmt.Println(isUnique("Hello"))
	fmt.Println(isUnique("world"))
	fmt.Println(isUnique("Gopher"))
	fmt.Println(isUnique("gopher"))
	fmt.Println(isUnique("1234567"))
	fmt.Println(isUnique("112233"))
}
