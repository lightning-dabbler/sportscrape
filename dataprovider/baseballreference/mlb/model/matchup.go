package model

import "time"

// MLBMatchup represents the data model for MLB matchups scraped from baseball-reference.com
type MLBMatchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// EventDateParquet is the timestamp associated with a given event (in days)
	EventDateParquet int32 `json:"-" parquet:"name=event_date, type=INT32, convertedtype=DATE, logicaltype=DATE"`
	// AwayTeam is the away team's name
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeam is the home team's name
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeScore is the home team's final score
	HomeScore int32 `json:"home_score" parquet:"name=home_score, type=INT32"`
	// AwayScore is the away team's final score
	AwayScore int32 `json:"away_score" parquet:"name=away_score, type=INT32"`
	// BoxScoreLink is the link to box score for related to the event
	BoxScoreLink string `json:"box_score_link" parquet:"name=box_score_link, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamLink is the link to the home team's baseball-reference profile page
	HomeTeamLink string `json:"home_team_link" parquet:"name=home_team_link, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamLink is the link to the away team's baseball-reference profile page
	AwayTeamLink string `json:"away_team_link" parquet:"name=away_team_link, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Loser is the losing team's name
	Loser string `json:"loser" parquet:"name=loser, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayoffMatch is whether or not the matchup was a playoff game
	PlayoffMatch bool `json:"playoff_match" parquet:"name=playoff_match, type=BOOLEAN"`
}
