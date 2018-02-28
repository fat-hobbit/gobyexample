package gobyexample

import (
	"fmt"
	"os"
)

// ExitDemo - demonstrates using os.Exit to immediately exit with a given status.
func ExitDemo() {
	// defers will not be run when using os.Exit, so this line will never be run.
	defer fmt.Println("!")

	// Exit with status 3. Note that unlike e.g. C, Go does not use an integer return
	// value from `main` to indicate exit status. If you'd like to exit with a non-zero
	// status you should use os.Exit
	os.Exit(3)
}
