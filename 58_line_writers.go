package gobyexample

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LineFiltersDemo - A line filter is a common type of program that
// reads input on stdin, processes it, and then prints some derived
// result on stdout. grep and sed are examples of line filters.
func LineFiltersDemo() {
	// Wrapping the unbuffered os.Stdin with a buffered scanner gives us
	// a convenient Scan method that advances the scanner to the next token,
	// which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text returns the current token (the next line here)
		// from the input
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Check for errors during Scan. EOF is expected and not
	// reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
