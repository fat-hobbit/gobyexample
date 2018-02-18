package gobyexample

import "fmt"

// pass by value - the function `zeroval` gets a copie of `ival`
// distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// pass by reference
func zeroptr(iptr *int) {
	// *iptr deferences the pointer from its memory address
	// to the current value at that address
	*iptr = 0
}

// PointersDemo - demonstrate pointers in Go
func PointersDemo() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// the `&i` syntax passes the memory address of i to `zeroptr`
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// pointers can be printed too
	fmt.Println("pointer:", &i)
}
