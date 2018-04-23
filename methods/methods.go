package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{10, 20}

	fmt.Println("area:", r.area())

	rp := &r

	fmt.Println("perimeter:", rp.perim())
}
