package pointers

import "fmt"

func Pointer1() {
	var b *int // b is nil
	// fmt.Println(*b) // nil pointer dereferencing => error

	a := 10
	b = &a
	fmt.Println(*b)

	c := &a
	*c = 20
	fmt.Println(a)
}