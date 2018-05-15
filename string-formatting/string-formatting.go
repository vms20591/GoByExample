package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// Format general Go values
	fmt.Printf("%v\n", p) // -> {1 2}

	// Include struct field names
	fmt.Printf("%+v\n", p) // -> {x:1, y:2}

	// Go syntax representation of the value
	fmt.Printf("%#v\n", p) // -> main.point{x:1, y:2}

	// Print type of value
	fmt.Printf("%T\n", p) // -> main.point

	// Format boolean
	fmt.Printf("%t\n", true) // -> true

	// Format integers
	fmt.Printf("%d\n", 123)

	// Binary representation
	fmt.Printf("%b\n", 10) // -> 1010

	// Character representation of integer
	fmt.Printf("%c\n", 65) // -> A

	// Hex representation
	fmt.Printf("%x\n", 10) // -> a

	// Format floats
	fmt.Printf("%f\n", 10.33) // -> 10.330000

	// Scientific notation
	fmt.Printf("%e\n", 123003400.0) // -> 1.230034e+08
	fmt.Printf("%E\n", 123003400.0) // -> 1.230034e+08

	// Format string
	fmt.Printf("%s\n", "\"string\"") // -> "string"

	// Double quote string
	fmt.Printf("%q\n", "\"string\"") // -> "\"string\""

	// Escape non-printable characters
	fmt.Printf("%q\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98") // -> "\xbd\xb2=\xbc ⌘"

	// Escape non-ascii characters
	fmt.Printf("%+q\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98") // -> "\xbd\xb2=\xbc \u2318"

	// Render string in base 16
	fmt.Printf("% x\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98") // -> bd b2 3d bc 20 e2 8c 98

	// Pointer representation
	fmt.Printf("%p\n", &p)

	// Integer width precision
	// right justified by defult and padded by spaces
	// %(width) => %6
	// %(filler)(width) => %06
	fmt.Printf("|%6d|%6d|\n", 12, 35)   // -> |    12|    35|
	fmt.Printf("|%06d|%06d|\n", 12, 35) // -> |000012|000035|

	// Float precision
	// %(width).(precision) => %6.2
	fmt.Printf("|%6.2f|%6.2f|\n", 12.352, 35.98) // -> | 12.35| 35.98|

	// Left justify
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12.352, 35.98) // -> |12.35 |35.98 |

	// String precision
	fmt.Printf("|%6s|%6s|\n", "go", "lang") // -> |    go|  lang|

	// Format string and return content
	s := fmt.Sprintf("a %s", "string") // -> a string
	fmt.Println(s)                     // -> a string

	// Format string and write to writer
	// os.Stderr -> writer
	fmt.Fprintf(os.Stderr, "an %s\n", "error") // -> an error
}
