package variables

import "fmt"

// Package level variables must be declared
var Pacakge_level_varibale string = "package level variable"
// p := 10 => Error

func Variable1() {

    var a string = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    var g = "banana"
    fmt.Println(f, g)
}

