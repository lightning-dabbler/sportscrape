package util

import "time"

func TimeToDays(t time.Time) int32 {
	// Get seconds since epoch
	seconds := t.Unix()

	// Convert to days (86400 seconds in a day)
	return int32(seconds / 86400)
}
