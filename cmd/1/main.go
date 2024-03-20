package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

type Action struct {
	Human
}

func (a Action) print() {
	fmt.Println(a.Name, a.Surname)
}

func main() {
	person := Action{
		Human: Human{
			Name:    "john",
			Surname: "doe",
		},
	}
	person.print()
}
