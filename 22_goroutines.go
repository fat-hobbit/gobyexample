package gobyexample

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// GoroutinesDemo - demonstrate goroutines
func GoroutinesDemo() {
	// synchronous, blocking function call
	f("direct")

	// invoke function in a goroutine.
	// this new goroutine will execute concurrently
	// with the calling one.
	go f("goroutine")

	// invoke goroutine for an anonymous function.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through to here
	// the `Scanln` call requires that we press a key before the
	// program exits.
	fmt.Scanln()
	fmt.Println("done")
}
