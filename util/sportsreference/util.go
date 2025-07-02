package sportsreference

import (
	"fmt"
	"strings"
	"time"
)

const (
	// time parse date template
	dateLayout = "2006-01-02"
)

// DateStrToTime takes in a date string in the form 2024-01-25
//
// Parameter:
//   - date: the date string to parse
//
// Returns a time.Time{} and nilable error
func DateStrToTime(date string) (time.Time, error) {
	dateParse, err := time.Parse(dateLayout, date)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not parse '%s': %w", date, err)
	}
	return dateParse, nil
}

// EventDate takes in a date string in the form 2024-01-25
//
// Parameter:
//   - date: The date associated with the event
//
// Returns the event date as time.Time{} with timezone (America/New_York) awareness and nilable error
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

// extractID extracts an arbitrary id from a sportsreference related link
//
// Parameter:
//   - link: the url
//
// Returns the id parsed from the url
func extractID(link string) string {
	linkSplit := strings.Split(link, "/")
	return strings.Split(linkSplit[len(linkSplit)-1], ".")[0]
}

// EventID extracts the event id from a boxscore link
//
// Parameter:
//   - boxscoreLink: the boxscore link
//
// Returns the event id as a string and nilable error
func EventID(boxscoreLink string) (string, error) {
	id := extractID(boxscoreLink)
	if id == "" {
		return "", fmt.Errorf("error: Event ID is an empty string when parsing %s", boxscoreLink)
	}
	return id, nil
}

// PlayerID extracts the player id from a player link
//
// Parameter:
//   - playerLink: the player link
//
// Returns the player id as a string and nilable error
func PlayerID(playerLink string) (string, error) {
	id := extractID(playerLink)
	if id == "" {
		return "", fmt.Errorf("error: Player ID is an empty string when parsing %s", playerLink)
	}
	return id, nil
}

// TeamID
func TeamID(teamLink string) (string, error) {
	linkSplit := strings.Split(teamLink, "/")
	if len(linkSplit) != 4 {
		return "", fmt.Errorf("error: unexpected team link format %s", teamLink)
	}
	id := linkSplit[2]
	if id == "" {
		return "", fmt.Errorf("error: Team ID is an empty string when parsing %s", teamLink)
	}
	return id, nil
}

type loserValues map[string]struct{}

// LoserValueExists is a helper to determine losers and winners in a matchup
var LoserValueExists loserValues = loserValues{"loser": struct{}{}, "winner": struct{}{}}

type Headers []string

// ReturnUnemptyField validates that the extracted field from a selector is not an empty string
//
// Paramater:
//   - str: the value to validate
//   - location: selector
//   - field: the model field
//
// Returns the string if not empty and nilable error
func ReturnUnemptyField(str string, location string, field string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("No value @ %s for %s", location, field)
	}
	return str, nil
}
