package gobyexample

import (
	"fmt"
	"math"
)

// interface: a named collection of method signatures
type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, we just need
// to implement all the methods in the interface.

// Here, we implement `geometry` on the `rectangle` struct
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

// And here, we implement `geometry` on the `circle` struct
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type, then we
// can call methods that are in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// InterfacesDemo - demonstrates interfaces in Go
func InterfacesDemo() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
