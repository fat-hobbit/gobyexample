package gobyexample

import "fmt"

// declare a new type, `person`, which is a struct,
// with `name` and `age` fields
type person struct {
	name string
	age  int
}

// StructsDemo - demonstrate how structs work in Go
func StructsDemo() {
	// create a new `person` struct using an expression
	fmt.Println(person{"Bob", 20})

	// you can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// an `&` prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// you can also use dots with struct pointers - the pointers
	// are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// or you can also deference the pointer manually
	fmt.Println((*sp).name)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
