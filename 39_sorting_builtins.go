package gobyexample

import (
	"fmt"
	"sort"
)

// SortBuiltinsDemo - demonstrates sorting of builtin types
func SortBuiltinsDemo() {
	// sort methods are specific to the builtin type
	// in-place sort, modifies the given slice
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// similar example for ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// check if a slice is already sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
