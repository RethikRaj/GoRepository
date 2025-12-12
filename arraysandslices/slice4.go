// Assigning operation
package arraysandslices

import "fmt"

func printSlicev2(s string, a []int){
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(a), cap(a), a)
}

func Slice4() {
	a := make([]int, 3, 5)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	printSlicev2("a : ", a)

	c := a // All header values of a are copied
	printSlicev2("c : ", c)

	b := a[1: 4]
	printSlicev2("b : ", b)

	b[0] = 99
	printSlicev2("a : ", a)
	printSlicev2("b : ", b)
	
	
	b = append(b, 19)
	printSlicev2("a : ", a)
	printSlicev2("b : ", b)
	
	a = append(a, 10)
	printSlicev2("a : ", a)
	printSlicev2("b : ", b)

	b = append(b , 30)
	printSlicev2("a : ", a)
	printSlicev2("b : ", b)

	b[0] = 999
	a = a[:cap(a)]
	printSlicev2("a : ", a)
	printSlicev2("b : ", b)
}

