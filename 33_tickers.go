package gobyexample

import (
	"fmt"
	"time"
)

// Timers are for when you want to something once in the future.
// Tickers are for when you want to do something repeatedly at
// regular intervals.

// TickersDemo - demonstrates using Go's built-in tickers.
func TickersDemo() {
	// Tickers use a similar mchanism to timers: a channel that is sent values.
	// Here we'll use the `range` built-in on the channel to iterate over the
	// values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped it won't receive
	// any more values on its channel.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
