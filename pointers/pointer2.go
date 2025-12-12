// Pointers with new keyword
package pointers

import "fmt"

func Pointer2() {
	a := new(int) // a is not nil , *a = 0
	fmt.Printf("Is a==nil : %v , value of *a : %v\n", a==nil, *a)
}