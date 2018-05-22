package main

import (
	"flag"
	"fmt"
)

func main() {
	p := fmt.Println

	// Create a string arg - argname, default, help
	// Returns a pointer
	wordPtr := flag.String("word", "foo", "a string")

	// Create a number arg - argname, default, help
	// Returns a pointer
	numPtr := flag.Int("num", 42, "an int")

	// Create a number arg - argname, default, help
	// Returns a pointer
	boolPtr := flag.Bool("fork", false, "a bool")

	// Create a arg from an existing variable
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a svar")

	// Parse args
	flag.Parse()

	// flags should be specified before tailing args
	// or they are considered as part of tailing args

	p("word: ", *wordPtr)
	p("num: ", *numPtr)
	p("bool: ", *boolPtr)
	p("svar: ", svar)
	// Print any tailing args
	p("tail: ", flag.Args())
}
