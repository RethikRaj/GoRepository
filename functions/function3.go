// First class citizens

package functions

import "fmt"

// Functions can be passed as arguments to other function
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	firstResult := arithmetic(a, b)
	secondResult := arithmetic(firstResult, c)
	return secondResult
}

func Function3() {
	// Functions can be assigned to variables
	sum := func (a int, b int) int {
		return a + b
	}

	summation := aggregate(2, 3, 4, sum)
	fmt.Println(summation)

	mul := func (a int , b int) int {
		return a * b
	}

	product := aggregate(2, 3, 4, mul)
	fmt.Println(product)
}