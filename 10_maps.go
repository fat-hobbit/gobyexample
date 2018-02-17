package gobyexample

import "fmt"

// WorkWithMaps - Create maps and print them to stdout
func WorkWithMaps() {
	// create empty map with builtin `make`` function:
	// `make(map[key-type]val-type`
	m := make(map[string]int)

	// set key/value pairs
	m["k1"] = 7
	m["k2"] = 13

	// print all key/value pairs
	fmt.Println("map:", m)

	// get value associated with key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// get number of key/value pairs in a map
	fmt.Println("len:", len(m))

	// remove key/value pairs from a map with builtin `delete` function
	delete(m, "k2")
	fmt.Println("map:", m)

	// optional second return value when getting a value from a map
	// indicates if the key was present in the map
	// use the blank identifier if you don't need the value
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
