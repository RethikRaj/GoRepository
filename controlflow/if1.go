package controlflow

import "fmt"

func If1() {
	// Basic if-else statement -> Variables declared in the init statement are visible only in the scope of if-elseif-...-else block
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}