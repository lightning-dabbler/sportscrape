package model

import "time"

// Matchup represents the data model for NBA, MLB, and NCAAB matchups scraped from foxsports.com
type Matchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// EventID is a unique ID that maps to the matchup e.g. 86833
	EventID int64 `json:"event_id"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventStatus the numerical representation of the event status e.g. 3
	EventStatus int32 `json:"event_status"`
	// StatusLine the string representation of the event status e.g. FINAL
	StatusLine string `json:"status_line"`
	// HomeTeamId is the home team's ID e.g. 21
	HomeTeamId int32 `json:"home_team_id"`
	// HomeTeamAbbreviation is the abbreviation of the home team's name e.g. ATL
	HomeTeamAbbreviation string `json:"home_team_abbreviation"`
	// HomeTeamNameLong is the home team's longer name but not the full name e.g. Braves
	HomeTeamNameLong string `json:"home_team_name_long"`
	// HomeTeamNameFull is the home team's full name e.g. Atlanta Braves
	HomeTeamNameFull string `json:"home_team_name_full"`
	// HomeRecord is the home team's record at time of request e.g. 69-37
	HomeRecord string `json:"home_record"`
	// HomeScore is the home team's score at time of request e.g. 12
	HomeScore int32 `json:"home_score"`
	// HomeRank is an optional rank. The values are expected for some NCAAB teams
	HomeRank *int32 `json:"home_rank"`
	// AwayTeamId is the away team's ID e.g. 8
	AwayTeamId int32 `json:"away_team_id"`
	// AwayTeambbreviation is the abbreviation of the away team's name e.g. LAA
	AwayTeambbreviation string `json:"away_team_abbreviation"`
	// AwayTeamNameLong is the away team's longer name but not the full name e.g. Angels
	AwayTeamNameLong string `json:"away_team_name_long"`
	// AwayTeamNameFull is the away team's full name e.g. Los Angeles Angels
	AwayTeamNameFull string `json:"away_team_name_full"`
	// AwayRecord is the away team's record at time of request e.g. 56-53
	AwayRecord string `json:"away_record"`
	// AwayScore is the away team's score at time of request e.g. 5
	AwayScore int32 `json:"away_score"`
	// AwayRank is an optional rank. The values are expected for some NCAAB teams
	AwayRank *int32 `json:"away_rank"`
	// Loser is an optional team id associated with the loser. Optional because some games could be live a loser will not be determined until the matchup is complete.
	Loser *int32 `json:"loser"`
	// IsPlayoff indicates whether the matchup is a playoff game (optional)
	IsPlayoff *bool `json:"is_playoff"`
}
