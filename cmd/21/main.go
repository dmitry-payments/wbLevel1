package main

import "fmt"

// Инт который должны привести к нашему коду
type ITest interface {
	hello() string
}

type Test struct {
	name string
}

func (t Test) hello() string {
	return "Hello from " + t.name
}

// Интерфейс к которому нужно привести
type IWork interface {
	getStatus() string
}

type TestAdapter struct {
	test ITest
}

func (ta TestAdapter) getStatus() string {
	return ta.test.hello()
}

func main() {
	t := Test{
		name: "Bob",
	}

	ta := TestAdapter{
		test: t,
	}

	var it IWork = ta

	fmt.Println(it.getStatus())
}
