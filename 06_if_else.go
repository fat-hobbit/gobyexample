package gobyexample

import "fmt"

// DoSomeIfElses - run some conditional logic
func DoSomeIfElses() {

	// if else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// just if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// a statement can precede conditionals
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
