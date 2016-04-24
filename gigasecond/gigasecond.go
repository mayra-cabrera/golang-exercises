// Package clause.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4
const gigasecond = 1000000000

// API function.  It uses a type from the Go standard library.
func AddGigasecond(start time.Time) time.Time {
	return start.Add(gigasecond * time.Second)
}
