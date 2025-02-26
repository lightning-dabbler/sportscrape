package model

import (
	"time"
)

// NBAMatchup represents the data model for NBA matchups scraped from basketball-reference.com
type NBAMatchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// EventID is the parsed event id from the response payload of the matchup
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
	// AwayQ1Total is the away team's Q1 points scored
	AwayQ1Total int `json:"away_q1_total"`
	// AwayQ2Total is the away team's Q2 points scored
	AwayQ2Total int `json:"away_q2_total"`
	// AwayQ3Total is the away team's Q3 points scored
	AwayQ3Total int `json:"away_q3_total"`
	// AwayQ4Total is the away team's Q4 points scored
	AwayQ4Total int `json:"away_q4_total"`
	// HomeQ1Total is the home team's Q1 points scored
	HomeQ1Total int `json:"home_q1_total"`
	// HomeQ2Total is the home team's Q2 points scored
	HomeQ2Total int `json:"home_q2_total"`
	// HomeQ3Total is the home team's Q3 points scored
	HomeQ3Total int `json:"home_q3_total"`
	// HomeQ4Total is the home team's Q4 points scored
	HomeQ4Total int `json:"home_q4_total"`
	// HomeTeamLink is the link to the home team's basketball-reference profile page
	HomeTeamLink string `json:"home_team_link"`
	// AwayTeamLink is the link to the away team's basketball-reference profile page
	AwayTeamLink string `json:"away_team_link"`
	// Loser is the losing team's name
	Loser string `json:"loser"`
}
