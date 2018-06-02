package main

import "fmt"

func digits(num int, dchn chan int) {
	for num != 0 {
		digit := num % 10
		num /= 10
		dchn <- digit
	}

	close(dchn)
}

func calcSquares(num int, squareC chan int) {
	sum := 0
	dchn := make(chan int)

	go digits(num, dchn)

	for digit := range dchn {
		sum += (digit * digit)
	}

	squareC <- sum
}

func calcCubes(num int, cubeC chan int) {
	sum := 0
	dchn := make(chan int)

	go digits(num, dchn)

	for digit := range dchn {
		sum += (digit * digit * digit)
	}

	cubeC <- sum
}

func main() {
	squareC := make(chan int)
	cubeC := make(chan int)
	num := 12341

	go calcSquares(num, squareC)
	go calcCubes(num, cubeC)

	fmt.Println(<-squareC)
	fmt.Println(<-cubeC)
}
