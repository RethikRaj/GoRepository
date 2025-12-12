// For range loops
package arraysandslices

import "fmt"

func Slice6() {
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		// v is copy => changes do not affect
		println(i, v)
		v = 5
	}

	fmt.Println(s)

}
