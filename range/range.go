package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println("num: ", num)
	}

	m := map[string]string{"k1": "v1", "k2": "v2"}
	for k, v := range m {
		fmt.Printf("%s->%s\n", k, v)
	}

	for k := range m {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
