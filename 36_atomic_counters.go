package gobyexample

import (
	"fmt"
	"sync/atomic"
	"time"
)

// AtomicCountersDemo - demonstrates managing state of multiple
// goroutines using the sync/atomic package
func AtomicCountersDemo() {
	// Use an unsigned integer to represent our (always-positive) counter.
	var ops uint64

	// To simulate concurrent updates, we'll start 50 goroutines that each increment
	// the counter about once a millisecond.
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)

	// In order to safely use the counter while it's been updated by other goroutines,
	// we extract a copy of the current value into `opsFinal` via `atomic.LoadUint64`.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
