// Package gigasecond calculate the moment when someone has lived for 10^9 seconds..
package gigasecond

import (
	"time"
)

// testVersion should match the targetTestVersion in the test file.
const testVersion = 4

// AddGigasecond adds 10^9 seconds to a time value.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1e9) * time.Second)
}
