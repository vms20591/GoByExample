package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"vms", 26})
	fmt.Println(person{name: "vms", age: 26})
	fmt.Println(person{name: "vms"})
	fmt.Println(&person{name: "vms", age: 26})

	s := person{"vms", 26}

	sp := &s
	sp.age = 27

	fmt.Println("age: ", sp.age)
}
