package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	fmt.Println("len: ", len(a))

	b[4] = 100
	fmt.Println("set: ", b)
	fmt.Println("get: ", b[4])

	var twoD [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}
