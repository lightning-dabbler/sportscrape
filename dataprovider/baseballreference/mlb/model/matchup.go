package model

import (
	"time"
)

// MLBMatchup represents the data model for MLB matchups scraped from baseball-reference.com
type MLBMatchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// AwayTeam is the away team's name
	AwayTeam string `json:"away_team"`
	// HomeTeam is the home team's name
	HomeTeam string `json:"home_team"`
	// HomeScore is the home team's final score
	HomeScore int `json:"home_score"`
	// AwayScore is the away team's final score
	AwayScore int `json:"away_score"`
	// BoxScoreLink is the link to box score for related to the event
	BoxScoreLink string `json:"box_score_link"`
	// HomeTeamLink is the link to the home team's baseball-reference profile page
	HomeTeamLink string `json:"home_team_link"`
	// AwayTeamLink is the link to the away team's baseball-reference profile page
	AwayTeamLink string `json:"away_team_link"`
	// Loser is the losing team's name
	Loser string `json:"loser"`
	// PlayoffMatch is whether or not the matchup was a playoff game
	PlayoffMatch bool `json:"playoff_match"`
}
