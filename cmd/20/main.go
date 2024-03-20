package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	str := strings.Split(s, " ")
	for i := 0; i < len(str)/2; i++ {

		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	return strings.Join(str, " ")
}

func main() {
	str := "snow dog sun"

	fmt.Println(reverse(str))
}
