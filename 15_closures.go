package gobyexample

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// ClosureDemo - demonstrates closures in Go
func ClosureDemo() {
	// closure that captures its own i value
	nextInt := intSeq()

	// see effect of the closure by calling function a few times
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// create a new function to confirm that each function has a unique state
	newInts := intSeq()
	fmt.Println(newInts())
}
