// Package clock provides a clock type, date independent and with minutes precision.
package clock

import "fmt"

const (
	testVersion = 4
	day         = 24 * 60
)

// A Clock represents an instant in day with minutes precision.
type Clock int

// New creates a new clock pointing the specified hour and minute.
func New(hour, minute int) Clock {
	return Clock(normalize(60*hour + minute))
}

// String returns a string represetantion of the clocl.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", int(c)/60, int(c)%60)
}

// Add adds (or removes when 'm' is negative) 'm' minutes from the current time on the clock
func (c Clock) Add(m int) Clock {
	c = Clock(normalize(int(c) + m))
	return c
}

func normalize(v int) int {
	return ((v % day) + day) % day
}
