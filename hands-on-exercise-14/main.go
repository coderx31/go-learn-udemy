package main

import "fmt"

var x = 40

const y = 41

func main() {
	var a int
	a = 10
	fmt.Printf("the value of a is %v and type of a is %T \n", a, a)
	var w = 30
	fmt.Printf("the value of w is %v and type of w is %T \n", w, w)
	fmt.Printf("the value of x is %v and type of x is %T \n", x, x)
	fmt.Printf("the value of y is %v and type of y is %T \n", y, y)
	z := 42
	fmt.Printf("the value of z is %v and type of z is %T \n", z, z)
}
