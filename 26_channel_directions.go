package gobyexample

import "fmt"

// When using channels as function parameters, you can specify
// if a channel is meant to only send or receive values.
// This specificity increases the type-safety of the program.

// the `ping` function only accepts a channel for sending values into,
// it would be a compile-time error to try to receive from this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// the `pong` function accepts one channel for receives(`pings`),
// and another for sends (`pongs`)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// ChannelDirectionsDemo - demonstrates channel directions
func ChannelDirectionsDemo() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
