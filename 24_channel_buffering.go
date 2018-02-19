package gobyexample

import "fmt"

// By default, channels are unbuffered, meaning that they
// will only accept sends if there is a corresponding receive
// ready to receive the sent value.
// Buffered channels accept a limited number of values without
// a corresponding receiver for those values.

// ChannelBufferingDemo - demonstrates use of channel buffering
func ChannelBufferingDemo() {
	// make channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// because the channel is buffered, we can send these
	// values into it without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// later, we can receive those two values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
