package gobyexample

import (
	"fmt"
	"time"
)

// EpochDemo - demonstrates getting number of seconds etc since Unix epoch
func EpochDemo() {
	// use `time.Now` with `Unix` or `UnixNano` to get elapsed time since
	// Unix epoch in seconds or nanoseconds respectively
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// there is no `UnixMilli`, so we'll need to manually divide from
	// nanoseconds
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// we can also convert integer seconds or nanoseconds since the
	// epoch into the corresponding `time`
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
