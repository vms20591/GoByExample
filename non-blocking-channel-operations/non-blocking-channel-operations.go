package main

import (
	"fmt"
)

func main() {
	messages := make(chan int)
	signals := make(chan int)

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("no message received")
	}

	msg := 1

	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received Signal:", sig)
	default:
		fmt.Println("no activity")
	}
}
