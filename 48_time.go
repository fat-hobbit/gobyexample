package gobyexample

import (
	"fmt"
	"time"
)

// TimeDemo - Demonstrates usage of time in Go.
func TimeDemo() {
	p := fmt.Println

	// get current time
	now := time.Now()
	p(now)

	// we can build a `time` struct by providing year, month, day, etc.
	// times are always associated with a `Location`, i.e. timezone
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// extract various components of time value
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Monday to Sunday `Weekday` is also available
	p(then.Weekday())

	// compare two times
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// the `Sub` methods return a `Duration` representing the interval
	// between two times
	diff := now.Sub(then)
	p(diff)

	// compute length of duration in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Minutes())
	p(diff.Nanoseconds())

	// use `Add` to advance a time by a given duration, or with a `-`
	// to move backwards by a duration
	p(then.Add(diff))
	p(then.Add(-diff))
}
