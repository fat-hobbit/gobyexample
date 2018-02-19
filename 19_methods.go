package gobyexample

import "fmt"

// Go supports methods defined on struct types
type rect struct {
	width, height int
}

// this method has a receiver type of *rect (pointer receiver type)
func (r *rect) area() int {
	return r.width * r.height
}

// this method is declared for a value receiver type
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// MethodsDemo - demonstrate method calls in Go
func MethodsDemo() {
	// Go automatically handles conversion between values and pointers
	// for method calls. You may want to use a pointer receiver type to
	// avoid copying on method calls or to allow the method to mutate
	// the receiver.

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
