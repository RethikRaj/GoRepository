// Struct with pointers
package pointers

import "fmt"

type vertex struct {
	x int
	y int
}

func Pointer4() {
	v := vertex{x : 1, y : 2}
	v.x = 10
	v.y = 20

	p := &v
	p.x = 100 // Same as doing (*p).x => Go does auto-dereferencing
	p.y = 200

	fmt.Println(*p)
}