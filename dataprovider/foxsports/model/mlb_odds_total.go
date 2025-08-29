package model

import "time"

// MLBOddsTotal - data model for MLB total bettings odds
type MLBOddsTotal struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 86833
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// HomeTeamID is the home team's ID e.g. 21
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// HomeTeamNameFull is the home team's full name e.g. Atlanta Braves
	HomeTeamNameFull string `json:"home_team_name_full" parquet:"name=home_team_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamID is the away team's ID e.g. 8
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeamNameFull is the away team's full name e.g. Los Angeles Angels
	AwayTeamNameFull string `json:"away_team_name_full" parquet:"name=away_team_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OverOdds
	OverOdds int32 `json:"over_odds" parquet:"name=over_odds, type=INT32"`
	// UnderOdds
	UnderOdds int32 `json:"under_odds" parquet:"name=under_odds, type=INT32"`
	// TotalLine
	TotalLine float32 `json:"total_line" parquet:"name=total_line, type=FLOAT"`
}
