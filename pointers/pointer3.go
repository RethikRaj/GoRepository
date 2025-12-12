// new keyword => * It allocates zeroed storage for a value of type T and returns a *T
package pointers

import "fmt"

func Pointer3() {
	p := new([]int)
	fmt.Println(*p == nil) // true
	fmt.Println(p == nil) // false

	*p = make([]int, 5)
	fmt.Println(*p)
}