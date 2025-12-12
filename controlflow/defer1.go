package controlflow

import "fmt"

func Defer1() {
	for i := 0; i < 10 ; i++ {
		defer fmt.Println("Defer", i)
	}
	fmt.Println("hello")
}