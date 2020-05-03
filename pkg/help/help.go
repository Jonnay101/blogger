package help

import "time"

// GetCurrentUTCTime returns the current Universal Time to the nearest second
func GetCurrentUTCTime() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}
