package gobyexample

import (
	"fmt"
	"time"
)

// SelectDemo - demonstrates `select` waiting on multiple
// channel operations
func SelectDemo() {
	// select across two channels.
	// each channel will receive a value after some amount of time,
	// to simulate e.g. blocking RPC operations executing in
	// concurrent goroutines
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// use `select` to await both of these values simultaneously,
	// printing each one as it arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}
