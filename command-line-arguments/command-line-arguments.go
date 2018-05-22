package main

import (
	"fmt"
	"os"
)

func main() {
	p := fmt.Println

	// args[0] - path to program
	// args[1:] - passed arguments
	p(os.Args)
}
