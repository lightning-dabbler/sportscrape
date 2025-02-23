package sportsreferenceutil

import (
	"fmt"
	"log"
	"time"
)

const (
	// time parse date template
	dateLayout = "2006-01-02"
)

// DateStrToTime takes in a date string in the form 2024-01-25
// Returns a timestamp
func DateStrToTime(date string) time.Time {
	dateParse, timeParseErr := time.Parse(dateLayout, date)
	if timeParseErr != nil {
		fmt.Println("Could not parse: '" + date + "'")
		log.Fatalln(timeParseErr)
	}
	return dateParse
}

// EventDate takes in a date string in the form 2024-01-25
// Returns a timezone (America/New_York) aware timestamp
func EventDate(date string) time.Time {
	loc, locationErr := time.LoadLocation("America/New_York")
	if locationErr != nil {
		log.Fatalln(locationErr)
	}

	dateParse, timeParseErr := time.ParseInLocation(dateLayout, date, loc)
	if timeParseErr != nil {
		fmt.Println("Could not parse: '" + date + "'")
		log.Fatalln(timeParseErr)
	}
	return dateParse
}

type loserValues map[string]struct{}

// LoserValueExists is a helper to determine losers and winners in a matchup
var LoserValueExists loserValues = loserValues{"loser": struct{}{}, "winner": struct{}{}}

// ReturnUnemptyField validates that the extracted field from a selector is not an empty string
// Returns the string if not empty else raises a logs an issue and fails the process
func ReturnUnemptyField(str string, location string, field string) string {
	if str == "" {
		log.Fatalf("No value @ %s for %s\n", location, field)
	}
	return str
}
