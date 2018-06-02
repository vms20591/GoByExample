package main

import (
	"fmt"
)

type I interface {
	walk()
}

type A struct{}
type B struct{}

func (a A) walk() {}
func (b B) walk() {}

func main() {
	var i I
	i = A{}

	vala, oka := i.(A)
	fmt.Println(vala, oka)

	valb, okb := i.(B)
	fmt.Println(valb, okb)
}
