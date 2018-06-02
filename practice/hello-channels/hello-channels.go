package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello goroutine is going to sleep now")

	time.Sleep(4 * time.Second)

	fmt.Println("hello goroutine is awake")

	done <- true
}

func main() {
	done := make(chan bool)

	fmt.Println("main calling hello goroutine")

	go hello(done)

	<-done

	fmt.Println("main resuming")
}
