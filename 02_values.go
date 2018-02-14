package gobyexample

import "fmt"

// PrintSomeValues - print some operations on basic types
func PrintSomeValues() {
	fmt.Println("Running function PrintSomeValues()")

	// String concatenation
	fmt.Println("Go" + "lang")

	// Integers and Floats
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// Booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)

	fmt.Println()
}
