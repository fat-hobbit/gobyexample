package gobyexample

import "fmt"

// PrintValues - print some operations on basic types
func PrintValues() {
	// String concatenation
	fmt.Println("Go" + "lang")

	// Integers and Floats
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// Booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}
