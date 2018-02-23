package gobyexample

import (
	"fmt"
	"sort"
)

// In order to sort by a custom function in Go, we need a corresponding type.
// Here we've created a `byLength` type that's just an alias for the builtin
// `[]string` type.
type byLength []string

// We implement `sort.Interface` - `Len`, `Less`, and `Swap` on our type, so that
// we can use the `sort` package's generic `Sort` function.
// `Len` and `Swap` will usually be similar across types, and `Less` will hold
// the actual sorting logic.
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// SortingByFunctionsDemo - demonstrates custom sorting using functions
func SortingByFunctionsDemo() {
	// With the above in place, we can implement our custom sort by
	// casting the original `fruits` slice to `byLength`, then use
	// `sort.Sort` on that typed slice.
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
