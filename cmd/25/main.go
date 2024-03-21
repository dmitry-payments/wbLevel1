package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("1 stage")
	Sleep(60)
	fmt.Println("2 stage")
}
