// Understanding method sets

package interfaces

import "fmt"
	
type geometry interface {
    area() float64
}

type rect struct {
	length float64
	breadth float64
}

func (r *rect) area() float64 {
	return r.length * r.breadth
}

func Interface2() {
	var g geometry

	r1 := rect{5,4}
	r2 := &rect{10,11}

	g = r2
	fmt.Println(g.area()) 
	g = &r1
	fmt.Println(g.area()) 

	// g = r1 => Error
}