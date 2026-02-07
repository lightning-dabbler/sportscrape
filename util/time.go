package util

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	colonFormatRe   = regexp.MustCompile(`^(\d+):(\d+)$`)                   // 03:50
	iso8601FormatRe = regexp.MustCompile(`^PT(\d+)M(?:(\d+(?:\.\d+)?)S)?$`) // PT10M43.00S | PT11M
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
		return time.Time{}, fmt.Errorf("could not convert RFC3339 string %s to time.Time: %w", str, err)
	}
	return timestamp, nil
}

// TransformMinutesPlayed parses "10:43", "PT10M43.00S", or "PT10M" formats
// Returns duration in minutes as float32 rounded to 2 decimal places
func TransformMinutesPlayed(minutesPlayed string) (float32, error) {

	// Check for MM:SS format
	if matches := colonFormatRe.FindStringSubmatch(minutesPlayed); matches != nil {
		return parseColonFormat(matches)
	}

	// Check for ISO 8601 format PT[M]M[S]S or PT[M]M
	if matches := iso8601FormatRe.FindStringSubmatch(minutesPlayed); matches != nil {
		return parseISO8601Format(matches)
	}

	return 0, fmt.Errorf("invalid duration format: %s", minutesPlayed)
}

func parseColonFormat(matches []string) (float32, error) {
	// matches[0] is the full match, matches[1] and matches[2] are capture groups
	if len(matches) < 3 {
		return 0, fmt.Errorf("insufficient match groups for colon format")
	}

	minutes, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %w", err)
	}

	seconds, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %w", err)
	}

	dur := time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	return float32(Round(dur.Minutes(), 2)), nil
}

func parseISO8601Format(matches []string) (float32, error) {
	// matches[0] is full match, matches[1] is minutes, matches[2] is optional seconds
	if len(matches) < 2 {
		return 0, fmt.Errorf("insufficient match groups for ISO 8601 format")
	}

	minutes, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %w", err)
	}

	var seconds float64

	// Seconds are optional (matches[2])
	if len(matches) > 2 && matches[2] != "" {
		seconds, err = strconv.ParseFloat(matches[2], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid seconds: %w", err)
		}
	}

	dur := time.Duration(minutes)*time.Minute +
		time.Duration(seconds*float64(time.Second))
	return float32(Round(dur.Minutes(), 2)), nil
}
