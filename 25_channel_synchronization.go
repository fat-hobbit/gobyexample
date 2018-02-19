package gobyexample

import (
	"fmt"
	"time"
)

// the `done` channel will be used to notify another
// goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value into the channel to indicate we're done
	done <- true
}

// ChannelSyncDemo - demonstrates channel synchronization
func ChannelSyncDemo() {
	// start a worker goroutine, giving it the channel
	// to notify on
	done := make(chan bool, 1)
	go worker(done)

	// block until we receive a notification from the worker
	// on the channel
	<-done
}
