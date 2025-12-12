package methods

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) print() {
	fmt.Println("Name : ", u.name, "Age : ", u.age)
}

func (u *user) changeAge(age int) {
	u.age = age
}

func Method1() {
	u1 := user{name : "Rethik", age : 20}
	u1.print()
	u1.changeAge(21) // Go does auto referencing => (&u1).changeAge(21)
	u1.print()

	u2 := &user{name : "Raj", age: 30}
	u2.print()	// Go does auto dereferencing => (*u2).print()
	u2.changeAge(31)
	u2.print() // Go does auto dereferencing => (*u2).print()

}