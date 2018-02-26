package gobyexample

import (
	"fmt"
	"time"
)

// TimeFormattingAndParsingDemo - demonstrates time formatting and parsing
func TimeFormattingAndParsingDemo() {
	p := fmt.Println

	// basic example of formatting a time according to RFC3339,
	// using the corresponding layout constant
	t := time.Now()
	p(t.Format(time.RFC3339))

	// time parsing uses the same layout values as `Format`
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` and `Parse` use example-based layouts, usually a constant from `time`
	// we can also supply custom layouts, which must use the reference time
	// `Mon Jan 2 15:04:05 MST 2006` to show the pattern with which to format / parse
	// a given time/string
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// for purely numeric representations, we can also use standard string formatting with
	// the extracted components of the time value
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
