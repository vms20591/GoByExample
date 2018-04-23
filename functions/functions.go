package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

func variadicFunc(nums ...int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}

func outer() func() int {
	i := 0

	return func() int {
		i++

		return i
	}
}

func main() {
	fmt.Println("plus:", plus(1, 2))

	fmt.Println("plusPlus:", plusPlus(1, 2))

	a, b := vals()
	fmt.Println(a, b)

	nums := []int{1, 2, 3, 4, 5}
	variadicFunc(nums...)

	nextInt := outer()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
