package controlflow

import (
	"fmt"
	"runtime"
	"time"
)

func Switch1(){
	// init statement -> Variables declared in the init statement are visible only in the scope of switch statement
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
	}

	// switch without condition == switch true.
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
	}
}