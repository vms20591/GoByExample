package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()

	// Format time object to RFC3339 -> 2018-05-22T00:30:11+05:30
	p(now.Format(time.RFC3339))

	// Parse to time from RFC3339 <- 2018-05-22T00:30:11+05:30
	t1, err := time.Parse(time.RFC3339, "2018-05-22T00:30:11+05:30")
	if err != nil {
		panic(err)
	}
	p(t1)

}
