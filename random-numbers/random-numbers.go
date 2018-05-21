package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println

	// Returns value from 0 to 99
	p(rand.Intn(100))

	// Returns value from 0.0 to 1.0
	p(rand.Float64())

	p((rand.Float64() * 5) + 5)

	// default random number generator is deterministic
	// produces same sequence
	// use seed
	// not safe in general, use `crypto/rand`
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// produces different numbers
	p(r1.Intn(100))
	p(r1.Intn(100))
	p(r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	p(r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)

	p(r3.Intn(100))
}
