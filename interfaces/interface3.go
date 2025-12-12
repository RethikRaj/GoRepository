// Understanding method sets

package interfaces

import "fmt"
	
type geometryTwo interface {
    area() float64
}

type rectTwo struct {
	length float64
	breadth float64
}

func (r rectTwo) area() float64 {
	return r.length * r.breadth
}

func Interface3() {
	var g geometryTwo

	r1 := rectTwo{5,4}
	r2 := &rectTwo{10,11}

	g = r1
	fmt.Println(g.area()) 
	g = &r1
	fmt.Println(g.area()) 

	g = r2
	fmt.Println(g.area()) 
}