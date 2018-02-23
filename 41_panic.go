package gobyexample

import "os"

// PanicDemo - demonstrates panics
func PanicDemo() {
	// A panic typically means something went unexpectedly wrong.
	// Mostly we use it to fail fast on errors that shouldn't occur
	// during normal operation, or that we aren't prepared to handle
	// gracefully.
	panic("We've got a BIG PROBLEM")

	// In Go, it's idiomatic to use error-handling return values
	// wherever possible.
	// A common use of panic is to abort if a function returns an error
	// value that we don't know how to (or want to) handle.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
