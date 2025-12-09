package util

import (
	"fmt"
	"strings"
	"time"
)

func TimeToDays(t time.Time) int32 {
	// Get seconds since epoch
	seconds := t.Unix()

	// Convert to days (86400 seconds in a day)
	return int32(seconds / 86400)
}

// DateStrToTime takes in a date string in the form 2024-01-25
//
// Parameter:
//   - date: the date string to parse
//
// Returns a time.Time{} and nilable error
func DateStrToTime(date string) (time.Time, error) {
	dateParse, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not parse '%s': %w", date, err)
	}
	return dateParse, nil
}

// RFC3339ToTime converts a RFC 3339 timestamp string to time.Time
func RFC3339ToTime(str string) (time.Time, error) {
	timestamp, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Time{}, fmt.Errorf("Could not convert RFC3339 string %s to time.Time: %w", str, err)
	}
	return timestamp, nil
}

// TransformMinutesPlayed
// minutesPlayed in the form \d+\:\d+ (e.g. 20:59)
func TransformMinutesPlayed(minutesPlayed string) (float32, error) {
	minutesPlayedSplit := strings.Split(minutesPlayed, ":")
	minutes, err := TextToInt(minutesPlayedSplit[0])
	if err != nil {
		return 0, fmt.Errorf("Could not convert minutes %s to integer: %w", minutesPlayedSplit[0], err)
	}

	seconds, err := TextToInt(minutesPlayedSplit[1])
	if err != nil {
		return 0, fmt.Errorf("Could not convert seconds %s to integer: %w", minutesPlayedSplit[1], err)
	}

	totalMinutes := float32(minutes) + float32(Round((float64(seconds)/float64(60)), 2))
	return totalMinutes, nil
}
