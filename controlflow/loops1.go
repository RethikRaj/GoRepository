package controlflow

func Loop1() {
	// Basic for loop : variables declared in init statement are visible only in the scope of for statement
	for i := 0; i < 10; i++ {
		println(i)
	}

	// Infinite loop
	// for {
	// 	println("Infinite loop")
	// }

	// The init and post statements are optional.
	j := 0
	for ; j < 10 ; {
		println(j)
		j++
	}

	// While loop using for
	k := 0
	for k < 10 {
		println(k)
		k++
	}

	// Break and continue
	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}