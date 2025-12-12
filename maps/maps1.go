package maps

import "fmt"

func Map1() {
	// 1.
	// var m1 map[string]int // m1 is nil map
	// m1["a"] = 1 // Panic  : can't assign to nil map

	// 2.
	m2 := make(map[string]int)
	// Set
	m2["a"] = 1
	m2["b"] = 2
	fmt.Println(m2)

	// get 
	fmt.Println(m2["a"])

	// delete
	delete(m2, "a")
	fmt.Println(m2)

	// checking existence of key => map by default return zero value if a key doesn't exist . Therefore to distinguish between missing keys and keys with zero value(intentionally) , we use the following syntax
	v, ok := m2["a"]
	fmt.Println(v, ok)

	// len
	fmt.Println(len(m2))

	// range 
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// clear
	clear(m2)
	fmt.Println(m2)


	// 3. Map literal
	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m3)

	// Maps entries are non addressable
	fmt.Println(&m3)
	// fmt.Println(&m3["a"])

}