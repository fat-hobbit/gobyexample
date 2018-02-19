package gobyexample

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	// By convention, errors are the last return value
	// and have type `error`, a built-in interface
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	// a `nil` in the error position indicates that
	// there was no error
	return arg + 3, nil
}

// We can use custom types as `errors` by implementing
// the `Error()` method on them.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// Here we use the &argError syntax to build a new
// struct, supplying values for the two fields
// `arg` and `prob`
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

// ErrorsDemo - demonstrate error handling in Go
func ErrorsDemo() {
	// call f1
	for _, i := range []int{7, 42} {
		// note that the use of an inline error check
		// on the `if` line is a common idiom in Go
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	// call f2
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// if you want to programatically use the data in a custom error,
	// you'll need to get the error as an instance of the custom error
	// type via type assertion.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
