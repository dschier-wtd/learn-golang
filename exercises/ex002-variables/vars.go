package main

import "fmt"

func main() {
	// assign values to variables
	var autotype = "autotype"
	direct := "direct"
	var t string = "typed"
	fmt.Println(autotype, direct, t)

	// strings
	var helloText string = "Hello, World!"
	fmt.Println(helloText)

	// integers
	var x int = 42
	var y int = 27
	var z int = x + y
	fmt.Println(z)

	// floats
	var pi float64 = 3.141592653589793
	var e float64 = 2.718281828459045
	var r float64 = pi * e
	fmt.Println(r)

	// booleans
	var b bool = true
	var c bool = false
	var d bool = b && c
	fmt.Println(d)
}
