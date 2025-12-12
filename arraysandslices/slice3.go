// Append operation

package arraysandslices

import "fmt"

func example1(){
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func example2() {
	s := make([]int, 3, 5)
	printSlice(s)

	s = append(s, 4)
	printSlice(s)

	s = append(s, 5)
	printSlice(s)

	// Slice grows
	s = append(s, 6)
	printSlice(s)
}

func example3() {
	s := make([]int, 3, 5)
	printSlice(s)

	s = s[2 : cap(s)]
	printSlice(s)

	// Slice grows
	s = append(s, 6)
	printSlice(s)
}

func Slice3() {
	example1()
	fmt.Println("=========")
	example2()
	fmt.Println("=========")
	example3()
}