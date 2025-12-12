package functions

func swap(x, y string) (string, string){
	return y, x
}

func Function1(){
	a , b := swap("hello", "world")
	println(a, b)
}