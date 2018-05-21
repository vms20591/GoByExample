package main

import (
	"fmt"
	"strconv"
)

func main() {
	// "0" - infer base from string
	// "64" - result fits in 64 bits
	// Return type still Int64, but can be converted
	i, _ := strconv.ParseInt("10", 0, 64)
	fmt.Println(i)

	// "64" - parse to 64 bits
	// Return type still Float64, but can be converted
	f, _ := strconv.ParseFloat("12.23", 64)
	fmt.Println(f)

	// parse hexadecimal values
	d, _ := strconv.ParseInt("0x64", 0, 64)
	fmt.Println(d)

	// parse uint values
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// quick convert string to base 10
	k, _ := strconv.Atoi("189")
	fmt.Println(k)

	// incorrect parsing returns err
	_, err := strconv.Atoi("wat")
	fmt.Println(err)
}
