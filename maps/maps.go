package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("m: ", m)

	n := map[string]int{"k1": 1, "k2": 2}
	fmt.Println("dcl: ", n)

	fmt.Println("len: ", len(n))

	delete(m, "k2")
	fmt.Println("del: ", m)

	var _, prs = m["k2"]
	fmt.Println("prs: ", prs)
}
