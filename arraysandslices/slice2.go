// Slicing operation

package arraysandslices

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Slice2() {
	s := []int{0, 1, 2, 3, 4}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:2]
	printSlice(s)

	// Extend to full capacity
	s = s[:cap(s)]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = s[1: 2]
	printSlice(s)
}
