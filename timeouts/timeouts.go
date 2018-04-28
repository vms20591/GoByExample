package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)

		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Received:", msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	select {
	case msg1 := <-c2:
		fmt.Println("Received:", msg1)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}
