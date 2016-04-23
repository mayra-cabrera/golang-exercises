// Clock stub file

package clock

import "time"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (c Clock) String() string {
	d := time.Date(2016, time.April, 10, c.hour, c.minute, 0, 0, time.UTC)
	return d.Format("15:04")
}

func (c Clock) Add(minutes int) Clock {
	c.minute = (c.minute + minutes) % 1440
	return c
}
