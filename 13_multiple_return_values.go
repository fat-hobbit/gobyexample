package gobyexample

import "fmt"

// (int, int) in the function signature shows that it returns two ints
func vals() (int, int) {
	return 3, 7
}

// MultipleReturnValues - demonstrate returning multiple return values
func MultipleReturnValues() {
	// use multiple assignment with the return values from the function
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Use the blank identifier if you only want a subset of the return values
	_, c := vals()
	fmt.Println(c)
}
