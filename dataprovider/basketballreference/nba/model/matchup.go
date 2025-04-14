package model

import "time"

// NBAMatchup represents the data model for NBA matchups scraped from basketball-reference.com
type NBAMatchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP_MILLIS, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// EventDateParquet is the timestamp associated with a given event (in days)
	EventDateParquet int32 `parquet:"name=event_date, type=INT32, convertedtype=DATE, logicaltype=DATE"`
	// AwayTeam is the away team's name
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY"`
	// HomeTeam is the home team's name
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY"`
	// HomeScore is the home team's final score
	HomeScore int32 `json:"home_score" parquet:"name=home_score, type=INT32"`
	// AwayScore is the away team's final score
	AwayScore int32 `json:"away_score" parquet:"name=away_score, type=INT32"`
	// BoxScoreLink is the link to box score for related to the event
	BoxScoreLink string `json:"box_score_link" parquet:"name=box_score_link, type=BYTE_ARRAY"`
	// AwayQ1Total is the away team's Q1 points scored
	AwayQ1Total int32 `json:"away_q1_total" parquet:"name=away_q1_total, type=INT32"`
	// AwayQ2Total is the away team's Q2 points scored
	AwayQ2Total int32 `json:"away_q2_total" parquet:"name=away_q2_total, type=INT32"`
	// AwayQ3Total is the away team's Q3 points scored
	AwayQ3Total int32 `json:"away_q3_total" parquet:"name=away_q3_total, type=INT32"`
	// AwayQ4Total is the away team's Q4 points scored
	AwayQ4Total int32 `json:"away_q4_total" parquet:"name=away_q4_total, type=INT32"`
	// HomeQ1Total is the home team's Q1 points scored
	HomeQ1Total int32 `json:"home_q1_total" parquet:"name=home_q1_total, type=INT32"`
	// HomeQ2Total is the home team's Q2 points scored
	HomeQ2Total int32 `json:"home_q2_total" parquet:"name=home_q2_total, type=INT32"`
	// HomeQ3Total is the home team's Q3 points scored
	HomeQ3Total int32 `json:"home_q3_total" parquet:"name=home_q3_total, type=INT32"`
	// HomeQ4Total is the home team's Q4 points scored
	HomeQ4Total int32 `json:"home_q4_total" parquet:"name=home_q4_total, type=INT32"`
	// HomeTeamLink is the link to the home team's basketball-reference profile page
	HomeTeamLink string `json:"home_team_link" parquet:"name=home_team_link, type=BYTE_ARRAY"`
	// AwayTeamLink is the link to the away team's basketball-reference profile page
	AwayTeamLink string `json:"away_team_link" parquet:"name=away_team_link, type=BYTE_ARRAY"`
	// Loser is the losing team's name
	Loser string `json:"loser" parquet:"name=loser, type=BOOLEAN"`
}
