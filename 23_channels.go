package gobyexample

import "fmt"

// ChannelsDemo - demonstrates sending into and receiving from channels
func ChannelsDemo() {
	// create a new channel with `make(chan val-type)`.
	// channels are typed by the values they convey
	messages := make(chan string)

	// send a value into a channel using the `channel <-` syntax
	go func() { messages <- "ping" }()

	// receive a value from the channel using the `<-channel` syntax
	msg := <-messages
	fmt.Println(msg)

	// By default, sends and receives block until both the sender and
	// receiver are ready. This property allowed us to wait at the end
	// of our program for the "ping" message without having to use any
	// other synchronization.
}
