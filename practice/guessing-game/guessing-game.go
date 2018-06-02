package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Range of system generated guess [0,n)
var guessRange *int64

// Player vs Computer mode
func byUser(secret int64, scanner *bufio.Scanner, random *rand.Rand) {
	var attempts int64

	start := time.Now()

	for {
		fmt.Print("Enter valid guess: ")

		// Get user input
		scanner.Scan()

		// Try converting guess to number
		guess, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 0, 64)

		if err != nil {
			continue
		}

		if guess > secret {
			attempts++

			fmt.Println("Too large")
		} else if guess < secret {
			attempts++

			fmt.Println("Too small")
		} else {
			break
		}
	}

	end := time.Now()

	fmt.Println()
	fmt.Printf("Secret: %d, Time elapsed: %f secs, Number of attempts: %d \n", secret, end.Sub(start).Seconds(), attempts)
	fmt.Println("Well done!")
	fmt.Println()
}

// Computer vs Computer mode
func byAi(secret int64, random *rand.Rand) {
	var attempts int64

	start := time.Now()

	for {
		guess := random.Int63n(*guessRange)

		if guess > secret || guess < secret {
			attempts++

			continue
		} else {
			break
		}
	}

	end := time.Now()

	fmt.Println()
	fmt.Printf("Secret: %d, Time elapsed: %f secs, Number of attempts: %d \n", secret, end.Sub(start).Seconds(), attempts)
	fmt.Println("Well done!")
	fmt.Println()
}

func main() {
	// Get guessing range
	guessRange = flag.Int64("guess-range", 100, "Range of guessing string from [0,n)")

	flag.Parse()

	fmt.Println("<================Guessing Game================>")

	// Create scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize rand
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate secret
	secret := random.Int63n(*guessRange)

outer:
	for {
		fmt.Println("Game Options: 1) Player Vs Computer 2) Computer Vs Computer 3) Quit")

		// Get user choice
		scanner.Scan()

		choice := strings.TrimSpace(scanner.Text())

		// Process choice
		switch choice {
		case "1":
			byUser(secret, scanner, random)
		case "2":
			byAi(secret, random)
		case "3":
			break outer
		default:
			continue
		}
	}
}
