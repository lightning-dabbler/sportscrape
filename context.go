package sportscrape

import "time"

type MatchupContext struct {
	Errors int
	Skips  int
}

type EventDataContext struct {
	PullTimestamp time.Time
	EventTime     time.Time
	EventID       interface{}
	URL           string
	AwayID        interface{}
	AwayTeam      string
	HomeID        interface{}
	HomeTeam      string
}
