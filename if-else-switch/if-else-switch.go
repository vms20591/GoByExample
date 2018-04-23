package main

import "fmt"
import "time"

func main() {
	if 10%2 == 0 {
		fmt.Println("even")
	}

	var a, b int = 1, 2

	if a > b {
		fmt.Println("a is greater")
	} else if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("b is greater")
	}

	if c := 0; c > -1 {
		fmt.Println("c is positive")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday")
	}

	switch {
	case a > b:
		fmt.Println("a > b")
	case a > 2:
		fmt.Println("a > c")
	default:
		fmt.Println("a is smallest")
	}
}
