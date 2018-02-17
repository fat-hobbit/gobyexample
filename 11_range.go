package gobyexample

import "fmt"

// WorkWithRanges - create some ranges to iterate over elements in a
// variety of data structures
func WorkWithRanges() {
	// use `range` to sum numbers in a slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// using index of element provided by `range`
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` can also iterate over just the keys of map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range`` on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune,
	// and the second value is the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
