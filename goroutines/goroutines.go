package main

import "fmt"

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("still going")

	fmt.Scanln()
	fmt.Println("done")
}
