// We often need our programs to perform operations on collections of data,
// like selecting all items that satisfy a given predicate or mapping all
// items to a new collection with a custom function.
//
// Go does not provide any such functions out-of-the-box, it's common to
// provide collection functions if and when they are specifically needed
// for your program and data types.
//
// Here are some example collection functions for slices of strings that we
// can use to build our own. Note that in some cases it may be clearest to just
// inline the collection-manipulating code directly, instead of creating and
// calling a helper function.

package gobyexample

import (
	"fmt"
	"strings"
)

// `index` returns the first index of the target string t, or -1 if no match is found.
func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// `include` returns true if the target string t is in the slice.
func include(vs []string, t string) bool {
	return index(vs, t) >= 0
}

// `any` returns true if one of the strings in the slice satisfies
// the predicate `f`.
func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// `all` returns true if all of the strings in the slice satisfies
// the predicate `f`.
func all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// `filter` returns a new slice containing all strings in the slice that
// satisfy the predicate `f`.
func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// `map` returns a new slice containing the results of applying the function
// `f` to each string in the original slice.
func myMap(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// CollectionFunctionsDemo - demonstrates custom collection functions on slices of strings
func CollectionFunctionsDemo() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(index(strs, "pear"))

	fmt.Println(include(strs, "grape"))

	fmt.Println(any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(all(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(myMap(strs, strings.ToUpper))
}
