package arraysandslices

import "fmt"

func Array1() {
	// 1. Array Declaration and Initialization
	// Declare an integer array 'a' of fixed size 5.
	// Elements are automatically initialized to their zero value (0 for int).
	var a [5]int
    fmt.Println("emp:", a) // Output: emp: [0 0 0 0 0]

	// 2. Setting and Getting Values
	// Assign a value to the element at index 4 (the fifth element).
    a[4] = 100
    fmt.Println("set:", a)  // Output: set: [0 0 0 0 100]
	// Retrieve the value at index 4.
    fmt.Println("get:", a[4]) // Output: get: 100

	// 3. Array Length
	// Get the fixed length of the array 'a'.
    fmt.Println("len:", len(a)) // Output: len: 5

	// 4. Declaration with Initialization (Shorthand)
	// Declare and initialize an array 'b' of size 5 with given values.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println(b) // Output: dcl: [1 2 3 4 5]

	// 5. Array Length Ellipsis (...)
	// The size of the array is determined by the number of initializer values.
	// This is equivalent to `[5]int{...}`.
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println(b) // Output: dcl: [1 2 3 4 5]

	// 6. Declaration with Keyed Elements
	// Initialize an array using specific indices (keys).
	// [100, 0, 0, 400, 500]. Unspecified elements get the zero value (0).
	// The array size is implicitly 5, as 400 is at index 3, and 500 is at index 4.
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b) // Output: idx: [100 0 0 400 500]

	// 7. Multi-dimensional Arrays (2D)
	// Declare a 2x3 two-dimensional array of integers.
    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD) // Output: 2d: [[0 1 2] [1 2 3]]

	// 8. Multi-dimensional Array Initialization Shorthand
	// Initialize the 2D array directly.
    twoD = [2][3]int{
        {1, 2, 3}, // First row
        {1, 2, 3}, // Second row
    }
    fmt.Println("2d: ", twoD) // Output: 2d: [[1 2 3] [1 2 3]]
}