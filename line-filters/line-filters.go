package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Enables reading line by line
	scanner := bufio.NewScanner(os.Stdin)

	// Scan for line
	for scanner.Scan() {
		// Convert each line to upper case
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	// EOF is expect and not reported as an error
	// Explicity check for error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}
