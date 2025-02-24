package sportsreferenceutil

import (
	"fmt"
	"time"
)

const (
	// time parse date template
	dateLayout = "2006-01-02"
)

// DateStrToTime takes in a date string in the form 2024-01-25
// Returns a timestamp and error
func DateStrToTime(date string) (time.Time, error) {
	dateParse, err := time.Parse(dateLayout, date)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not parse '%s': %w", date, err)
	}
	return dateParse, nil
}

// EventDate takes in a date string in the form 2024-01-25
// Returns a timezone (America/New_York) aware timestamp error
func EventDate(date string) (time.Time, error) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}, fmt.Errorf("An issue arose loading timezone: %w", err)
	}

	dateParse, err := time.ParseInLocation(dateLayout, date, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("Could not parse '%s': %w", date, err)
	}
	return dateParse, nil
}

type loserValues map[string]struct{}

// LoserValueExists is a helper to determine losers and winners in a matchup
var LoserValueExists loserValues = loserValues{"loser": struct{}{}, "winner": struct{}{}}

// ReturnUnemptyField validates that the extracted field from a selector is not an empty string
// Returns the string if not empty else raises a logs an issue and fails the process
func ReturnUnemptyField(str string, location string, field string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("No value @ %s for %s", location, field)
	}
	return str, nil
}
