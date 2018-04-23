package main

import "fmt"

const s = "string"

func needInt(x int) {

}

func needFloat(x float64) {

}

func main() {
	var s1 = "s1"
	var s2 string = "s2"
	s3 := "s3"

	fmt.Println(s1, s2, s3)

	var i = 10
	fmt.Println(i)

	var f = 10.01
	fmt.Println(f)

	var b = true
	fmt.Println(b)

	var c, d int = 1, 2
	fmt.Println(c, d)

	const n = 10000
	const e = 3e20 / n

	needInt(n)
	needFloat(f)
}
