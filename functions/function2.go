// Named return values
package functions

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// return
	return x, y
	// return 10, 20
}

func Function2(){
	a, b := split(17)
	println(a, b)
}