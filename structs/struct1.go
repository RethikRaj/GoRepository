package structs

import "fmt"

type user struct {
	name string
	age  int
}

func Struct1() {
	// Struct creation
	var u1 user // Empty values of fields inside structs
	u2 := user{"Rethik", 20}
	u3 := user{age: 30, name: "Raj"}
	u4 := new(user)
	u5 := &user{name: "Raja", age: 50}

	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println(u4)
	fmt.Println(u5, u5.name, u5.age) // Go does auto dereferncing

	// Anonymous structs
	dog := struct{
		name string
		isGood bool
	} {
		name: "Doberman",
		isGood: false,
	}
	fmt.Println(dog)
}