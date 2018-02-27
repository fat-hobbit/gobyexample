package gobyexample

// Go provides a `flag` package supporting basic command line flag parsing.
import (
	"flag"
	"fmt"
)

// CommandLineFlagsDemo - command line flags are a common way to
// specify options for command line programs. E.g. in `wc -l`, the
// `-l` is a command line flag.
func CommandLineFlagsDemo() {
	// Basic flag declarations are available for string, integer, and
	// boolean options. Here we declare a string flag `word` with a default
	// value "foo" and a short description.
	wordPtr := flag.String("word", "foo", "a string")

	// Declare more flags for integer and boolean options.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// It's also possible to declare an option that uses an existing var
	// declared elsewhere in the program. Note that we need to pass in a
	// pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call `flag.Parse()` to execute
	// command line parsing.
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
