package main

import (
	"fmt"
)

type Person struct {
	name string
}
type Student struct {
	Person
	score int
}

func (p *Person) myname() {
	fmt.Println(p.name)
}
func main() {
	s := Student{
		Person{
			name: "li",
		},
		78,
	}
	fmt.Println(s.name)
}
