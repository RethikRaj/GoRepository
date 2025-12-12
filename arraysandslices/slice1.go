// Slice creation
package arraysandslices

import (
	"fmt"
	"slices"
)

func Slice1() {
	// 1. A nil slice. It has no underlying array, and its length and capacity are 0.
	var s1 []int
	fmt.Printf("s1 (nil slice): len=%d cap=%d %v\n", len(s1), cap(s1), s1)

	// 2. A slice literal. Go creates an underlying array and returns a slice that refers to it.
	s2 := []int{1, 2, 3}
	fmt.Printf("s2 (literal):     len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	// 3. Using make() with length. Capacity will be equal to the length.
	s3 := make([]int, 3)
	fmt.Printf("s3 (make len):    len=%d cap=%d %v\n", len(s3), cap(s3), s3)

	// 4. Using make() with length and capacity.
	s4 := make([]int, 3, 5)
	fmt.Printf("s4 (make cap):    len=%d cap=%d %v\n", len(s4), cap(s4), s4)

	// 5. Slicing from an array. The slice shares the underlying array.
	arr := [5]int{1, 2, 3, 4, 5}
	s5 := arr[1:4] // Elements from index 1 up to (but not including) 4.
	fmt.Printf("s5 (from array):  len=%d cap=%d %v\n", len(s5), cap(s5), s5)

	// 6. Slicing from another slice.
	// The high bound can go up to the capacity of the source slice.
	// s4 has len=3, cap=5. s4[1:4] is valid.
	// The new slice will have elements from index 1 up to 4 of the underlying array.
	s6 := s4[1:4]
	fmt.Printf("s6 (from slice):  len=%d cap=%d %v\n", len(s6), cap(s6), s6)

	// 7. Cloning a slice
	// This creates a new slice with its own separate underlying array,
	// containing the same elements as the original.
	s7 := slices.Clone(s2)
	fmt.Printf("s7 (clone):       len=%d cap=%d %v\n", len(s7), cap(s7), s7)

	// Modifying the clone does not affect the original.
	s7[0] = 99
	fmt.Printf("s2 (after s7 mod):len=%d cap=%d %v\n", len(s2), cap(s2), s2)
	fmt.Printf("s7 (modified):    len=%d cap=%d %v\n", len(s7), cap(s7), s7)
}

