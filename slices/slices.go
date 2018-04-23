package main

import "fmt"

func main() {
	var a []int
	fmt.Println("emp: ", a)

	b := []int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	fmt.Println("len: ", len(b))

	b[4] = 100
	fmt.Println("set: ", b)
	fmt.Println("get: ", b[4])

	c := make([]string, 5)
	c = append(c, "a")
	c = append(c, "b")
	c[2] = "c"
	c[3] = "d"
	fmt.Println("apd: ", c)

	d := make([]string, len(c))
	copy(d, c)
	fmt.Println("copy: ", d)

	l1 := b[2:5]
	fmt.Println("l1: ", l1)

	l2 := b[2:]
	fmt.Println("l2: ", l2)

	l3 := b[:5]
	fmt.Println("l3: ", l3)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inputLen := i + 1
		twoD[i] = make([]int, inputLen)
		for j := 0; j < inputLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}
