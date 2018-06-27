package timer

import (
	"time"
)

// GetTimer returns a new timer that waits for duration s.
func GetTimer(s time.Duration) *time.Timer {

	return time.NewTimer(time.Second * s)
}
