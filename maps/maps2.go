package maps

import "fmt"

func Maps2() {
	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m3)

	// Maps entries are non addressable
	fmt.Println(&m3)
	// fmt.Println(&m3["a"])
}