// Any interface

package interfaces

import "fmt"

// typeSwitch demonstrates the use of a type switch on an empty interface value.
// It checks the underlying concrete type of the value held by the interface `i`.
// The variable `v` in the switch statement holds the value of the concrete type.
func typeSwitch(i interface{}) {
	// The syntax 'v := i.(type)' is special and only works within a switch statement.
	// It extracts the underlying value and assigns it to 'v', and then checks its type.
	switch v := i.(type) {
	case string:
		// If the underlying type is string, 'v' is of type string.
		fmt.Println("String : ", v)
	case []int:
		// If the underlying type is a slice of integers, 'v' is of type []int.
		fmt.Println("[]int : ", v)
	case float64:
		// If the underlying type is float64, 'v' is of type float64.
		fmt.Println("float64", v)
	case int:
		// If the underlying type is int, 'v' is of type int.
		fmt.Println("Int", v)
	default:
		// The default case handles any other types not explicitly listed.
		// %T prints the type of the value.
		fmt.Printf("type %T\n", v)
	}
}

// Interface4 demonstrates the empty interface, type assertion, and type switches.
func Interface4() {
	// An empty interface variable can hold values of ANY type,
	// because all types implicitly satisfy the empty interface (`interface{}`).
	var i interface{}

	// i can hold any types => Bcoz all types implement interface{}
	i = 10              // i now holds an int
	i = make([]int, 3)  // i now holds a []int slice
	i = "Hello"         // i now holds a string

	// **Important Concept**: We cannot perform operations on `i` directly (like arithmetic or string concatenation)
	// because the compiler only knows `i` is an interface{}, not its concrete type.
	// fmt.Println(i + "World") // This line would cause a compile-time error.

	// **Type Assertion (Single return value)**:
	// Asserts that the concrete type inside `i` is `string`.
	// If the assertion fails (i.e., `i` doesn't hold a string), it will cause a **panic** (runtime error).
	s1 := i.(string)
	fmt.Println(s1)

	// **Type Assertion (Two return values - the "comma-ok" idiom)**:
	// This is the SAFE way to perform type assertion.
	// `s2` is the asserted value, and `ok` is a boolean indicating if the assertion succeeded.
	s2, ok := i.(string)
	fmt.Println(s2, ok) // Output: Hello true (Assertion succeeded)

	f1, ok := i.(float64)
	fmt.Println(f1, ok) // Output: 0 false (Assertion failed, ok is false, f1 is the zero value for float64)

	// The following causes a **panic** (runtime error) at runtime:
	// The assertion fails and there is no 'ok' check to handle the failure gracefully.
	// The error message would be: "panic: interface {} is string, not float64"
	// f2 := i.(float64)
	// fmt.Println(f2)

	// **Type Switch demonstration**:
	fmt.Println("\n--- Type Switch Results ---")
	typeSwitch(i)                 // Current value of i is "Hello" (string case)
	typeSwitch("Hwllo")           // String case
	typeSwitch(34)                // Int case
	typeSwitch(make([]int, 3))    // []int case
	typeSwitch(make(map[string]int)) // Default case (map[string]int)
	typeSwitch(3.14159)           // float64 case
}