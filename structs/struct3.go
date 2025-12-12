// Nested structs => Composition ( has-a )
package structs

import "fmt"

type Address struct {
    city  string
    state string
}

type Employee struct {
    name    string
    age     int
    address Address
}

func Struct3() {
	addr := Address{city: "Los Angeles", state: "California"}
	employee := Employee{
		name: "Rethik",
		age : 20,
		address: addr,
	}
	fmt.Println(employee)
	fmt.Println(employee.name, employee.age, employee.address, employee.address.city, employee.address.state)
}