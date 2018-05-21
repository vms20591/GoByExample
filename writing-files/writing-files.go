package main

import (
	"bufio"
	"fmt"
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
	// dump string or bytes into file
	d1 := []byte("hello golang\n")
	// creates file if it doesn't exist
	// truncates if it does
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	checkErr(err)

	// creates file if it doesn't exist
	// truncates if it does
	f, err := os.Create("/tmp/dat2")
	checkErr(err)

	// defer close immediately for writing
	defer f.Close()

	// write byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// robust method to write string
	n3, err := f.WriteString("hello golang\n")
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// flush to disk
	f.Sync()

	// bufferd writing
	w := bufio.NewWriter(f)
	d3 := []byte{115, 111, 109, 101, 10}
	n4, err := w.Write(d3)
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
