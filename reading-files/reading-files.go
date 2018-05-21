package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Manage error checking
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Reading entire file to memeory
	dat, err := ioutil.ReadFile("/tmp/test.go")
	checkErr(err)
	fmt.Println(string(dat))

	// Get a file handle
	f, err := os.Open("/tmp/test.go")
	checkErr(err)

	// Buffer to hold data
	b1 := make([]byte, 5)
	// Reads upto 5 bytes
	n1, err := f.Read(b1)
	checkErr(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Seek to location from starting
	o2, err := f.Seek(6, 0)
	checkErr(err)
	// Read upto 3 bytes from seeked location
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	checkErr(err)
	b3 := make([]byte, 2)
	// io.ReadAtLeast is robust when compared
	// to above example
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// go back to beginning of file
	_, err = f.Seek(0, 0)
	checkErr(err)

	// bufio provides buffered reading
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	checkErr(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// use defer after opening file to automatically
	// close
	f.Close()
}
