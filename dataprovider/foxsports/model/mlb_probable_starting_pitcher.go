package model

import "time"

// MLBProbableStartingPitcher - data model for MLB probable starting pitcher
type MLBProbableStartingPitcher struct {
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
	// HomeStartingPitcherID is the player id for the home team's starting pitcher
	HomeStartingPitcherID int64 `json:"home_starting_pitcher_id" parquet:"name=home_starting_pitcher_id, type=INT64"`
	// HomeStartingPitcher the name of the Home team's starting pitcher
	HomeStartingPitcher string `json:"home_starting_pitcher" parquet:"name=home_starting_pitcher, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeStartingPitcherRecord is the record of the home team's starting pitcher
	HomeStartingPitcherRecord string `json:"home_starting_pitcher_record" parquet:"name=home_starting_pitcher_record, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeStartingPitcherERA is the home team's starting pitcher's earned run average - https://www.mlb.com/glossary/standard-stats/earned-run-average
	HomeStartingPitcherERA float32 `json:"home_starting_pitcher_era" parquet:"name=home_starting_pitcher_era, type=FLOAT"`
	// AwayTeamID is the away team's ID e.g. 8
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeamNameFull is the away team's full name e.g. Los Angeles Angels
	AwayTeamNameFull string `json:"away_team_name_full" parquet:"name=away_team_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayStartingPitcherID is the player id for the home team's starting pitcher
	AwayStartingPitcherID int64 `json:"away_starting_pitcher_id" parquet:"name=away_starting_pitcher_id, type=INT64"`
	// AwayStartingPitcher the name of the Away team's starting pitcher
	AwayStartingPitcher string `json:"away_starting_pitcher" parquet:"name=away_starting_pitcher, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayStartingPitcherRecord is the record of the away team's starting pitcher
	AwayStartingPitcherRecord string `json:"away_starting_pitcher_record" parquet:"name=away_starting_pitcher_record, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeStartingPitcherERA is the away team's starting pitcher's earned run average - https://www.mlb.com/glossary/standard-stats/earned-run-average
	AwayStartingPitcherERA float32 `json:"away_starting_pitcher_era" parquet:"name=away_starting_pitcher_era, type=FLOAT"`
}
