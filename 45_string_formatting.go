package gobyexample

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

// StringFormattingDemo - demonstrates Go's string formatting capabilities
func StringFormattingDemo() {
	p := point{1, 2}

	// %v - print instance of our struct
	fmt.Printf("%v\n", p)

	// %+v - print instance of our struct with field names
	fmt.Printf("%+v\n", p)

	// %#v - print Go representation of the value, i.e. the source
	// code snippet that would produce that value
	fmt.Printf("%#v\n", p)

	// %T - print type of value
	fmt.Printf("%T\n", p)

	// %t - format booleans
	fmt.Printf("%t\n", true)

	// %d - format integer in base 10
	fmt.Printf("%d\n", 123)

	// %b - print binary representation
	fmt.Printf("%b\n", 14)

	// %c - print character corresponding to given integer
	fmt.Printf("%c\n", 33)

	// %x - print hexadecimal representation
	fmt.Printf("%x\n", 456)

	// %f - basic decimal formatting
	fmt.Printf("%f\n", 78.9)

	// %e, %E - scientific notation
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// %s - basic string printing
	fmt.Printf("%s\n", "\"string\"")

	// %q - double-quote strings as in Go source
	fmt.Printf("%q\n", "\"string\"")

	// %x - render string in base-16, two output chars
	// per byte of input
	fmt.Printf("%x\n", "hex this")

	// %p - print representation of pointer
	fmt.Printf("%p\n", &p)

	// specify width of integers
	// by default result is right-justified and padded with spaces
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// specify width of floats
	// restrict decimal precision at the same time with
	// `width.precision` syntax
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// to left-justify, use the - flag
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// control width when formatting strings
	// right-justified by default
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// left-justify with the - flag
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// `Sprintf` formats and returns a string without
	// printing it anywhere
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// format and print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
