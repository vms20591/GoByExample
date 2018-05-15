package main

import (
	"fmt"
	"strings"
)

func Index(str []string, t string) int {
	for i, s := range str {
		if s == t {
			return i
		}
	}

	return -1
}

func Any(str []string, f func(string) bool) bool {
	for _, s := range str {
		if f(s) {
			return true
		}
	}

	return false
}

func main() {
	str := []string{"avocado", "orange", "peach"}

	fmt.Println("Index of peach:", Index(str, "peach"))

	fmt.Println("Has prefix a?", Any(str, func(s string) bool {
		return strings.HasPrefix(s, "a")
	}))
}
