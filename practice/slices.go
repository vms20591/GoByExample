package main

import (
	"fmt"
)

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func SubtractOneFromSlice(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]

	return slice
}

func PtrSubtractOneFromSlice(slicePtr *[]byte) {
	slice := *slicePtr

	*slicePtr = slice[0 : len(slice)-1]
}

func main() {
	var buffer [256]byte

	var slice = buffer[100:150]

	// var slice2 = slice[5:10] // slice header -> length: 5, 0th elem: &buff[105]

	slice = slice[5:10] // slice header -> length: 5, 0th elem: &buff[105]

	slice = slice[1 : len(slice)-1] // slice header -> length: 3, 0th elem: &buff[106]

	// slashPos := bytes.IndexRune(slice, '/')

	slice = buffer[10:20]

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println("AddOneToEachElement->")
	fmt.Println("Before:", slice)
	AddOneToEachElement(slice)
	fmt.Println("After:", slice)
	fmt.Println()

	fmt.Println("SubtractOneFromSlice->")
	fmt.Println("Before len(slice):", len(slice))
	newSlice := SubtractOneFromSlice(slice)
	fmt.Println("After len(slice):", len(slice))
	fmt.Println("After len(newSlice):", len(newSlice))
	fmt.Println()

	fmt.Println("PtrSubtractOneFromSlice->")
	fmt.Println("Before len(slice):", len(slice))
	PtrSubtractOneFromSlice(&slice)
	fmt.Println("After len(slice):", len(slice))
}
