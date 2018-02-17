package gobyexample

import "fmt"

// WorkWithSlices - Create slices and print them to stdout
func WorkWithSlices() {
	// create zero-valued empty slice with built-in make function
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// get and set a slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// get length of a slice
	fmt.Println("len:", len(s))

	// the built-in function `append` returns a slice
	// containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy a slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice operator - from index 2 up to 5, excluding 5
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slice up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// slice up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initalize a variable for a slice in one line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// multi-dimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
