package gobyexample

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomNumbersDemo - demonstrates random numbers
func RandomNumbersDemo() {
	// `rand.Intn` returns a random int n 0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` returns a float64 f, 0.0 <= f < 1.0,
	// which can be used to generate random floats in other ranges
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Default number generator is deterministic, so it'll produce the
	// same sequence of numbers each time by default. To produce varying
	// sequences, give it a seed that changes. Note that this is not
	// cryptographically secure, use `crypto/rand` for that.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// If you seed a source with the same number, it produces the
	// same sequence of random numbers.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Println(r3.Intn(100))
}
