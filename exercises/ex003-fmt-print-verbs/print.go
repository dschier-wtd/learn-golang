package main

import "fmt"

func main() {
	var name string = "Daniel"
	var age int = 37
	var height float64 = 1.80

	fmt.Println(name, age)
	fmt.Println("My Name is", name, "and I am", age, "years old")
	fmt.Printf("v: My Name is %v and I am %v years old.\n", name, age)
	fmt.Printf("s: My Name is %s and I am %s years old.\n", name, age)
	fmt.Printf("sd: My Name is %s and I am %d years old.\n", name, age)
	fmt.Printf("f: I am %.2f meters tall.\n", height)
	fmt.Printf("d: I am %04d years old.\n", age)
}
