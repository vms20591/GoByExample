package main

import (
	"fmt"
)

func ping(pingChn chan<- string, message string) {
	pingChn <- message
}

func pong(pingChn <-chan string, pongChn chan<- string) {
	message := <-pingChn

	pongChn <- message
}

func main() {
	pingChn := make(chan string, 1)
	pongChn := make(chan string, 1)

	ping(pingChn, "pinged message")
	pong(pingChn, pongChn)

	fmt.Println(<-pongChn)
}
