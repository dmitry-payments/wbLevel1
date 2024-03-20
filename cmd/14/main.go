package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 4
	var b interface{} = ""
	var c interface{} = true
	var d interface{} = make(chan int)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
}
