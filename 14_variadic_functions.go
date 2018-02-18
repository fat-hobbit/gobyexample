package gobyexample

import "fmt"

// A function that takes an arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// CallVariadicFunctions - calls functions that accept multiple arguments
func CallVariadicFunctions() {
	sum(1, 2)
	sum(1, 2, 3)

	// if you already have multiple args in a slice, apply them to a
	// variadic function using `func(slice...)`
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
