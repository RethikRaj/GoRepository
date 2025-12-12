package packages

import (
	"fmt"
	"math/rand"
	"math"
)

func Package1(){
	fmt.Println("Random number generated is : ", rand.Intn(10));

	fmt.Println("The value of PI is : ", math.Pi);
	// fmt.Println("The value of PI is : ", math.pi);
}