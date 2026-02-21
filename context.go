package sportscrape

import "time"

type MatchupContext struct {
	Errors int
	Skips  int
}

type EventDataContext struct {
	PullTimestamp time.Time
	EventTime     time.Time
	EventID       any
	URL           string
	AwayID        any
	AwayTeam      string
	HomeID        any
	HomeTeam      string
}
