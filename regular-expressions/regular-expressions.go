package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// Test whether a pattern matches a string
	matched, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(matched) // -> true

	// Compile regexp struct for optimization
	r, err := regexp.Compile("p([a-z]+)ch")

	if err != nil {
		panic("Damn the regexp didn't compile")
	}

	fmt.Println(r.MatchString("peach")) // -> true

	// Find match for regexp
	// Returns first match
	fmt.Println(r.FindString("peach punch")) // -> peach

	// Find index of the matched string
	// Returns first match
	// start index - inclusive
	// end index - exclusive
	fmt.Println(r.FindStringIndex("peach punch")) // -> [0 5]

	// Find whole match and submatch of first matching string
	// whole match -> p([a-z]+)ch
	// submatch -> ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch")) // -> [peach ea]

	// Find index of whole match and submatch of first matching string
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // -> [0 5 1 3]

	// Find all matches of regexp
	// -ve -> no limit of matches
	// +ve -> limit no of matches
	fmt.Println(r.FindAllString("peach punch", -1)) // -> [peach punch]

	// Find index of all matches of regexp
	fmt.Println(r.FindAllStringIndex("peach punch", -1)) // -> [[0 5] [6 11]]

	// Find whole match and submatch of all matching string
	// whole match -> p([a-z]+)ch
	// submatch -> ([a-z]+)
	fmt.Println(r.FindAllStringSubmatch("peach punch", -1)) // -> [[peach ea] [punch un]]

	// Find index whole match and submatch of all matching string
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch", -1)) // -> [[0 5 1 3] [6 11 7 9]]

	// Match byte arugments
	fmt.Println(r.Match([]byte("peach"))) // -> true

	// For constants use `MustCompile` variant
	// `MustCompile` -> regex
	// `Compile` -> (regex, error)
	r = regexp.MustCompile("p([a-z]+)ch") // -> panic if regexp has error
	fmt.Println(r)                        // -> p([a-z]+)ch

	// Replace all matching regexp string with other string
	fmt.Println(r.ReplaceAllString("It's a peach punch", "<fruit>")) // -> It's a <fruit> <fruit>

	// Transform every matched regexp string using a function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // -> a PEACH
}
