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
	// TeamID is the team's ID e.g. 21
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// TeamNameFull is the team's full name e.g. Atlanta Braves
	TeamNameFull string `json:"team_name_full" parquet:"name=team_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// StartingPitcherID is the player id for the team's starting pitcher
	StartingPitcherID int64 `json:"starting_pitcher_id" parquet:"name=starting_pitcher_id, type=INT64"`
	// StartingPitcher the name of the team's starting pitcher
	StartingPitcher string `json:"starting_pitcher" parquet:"name=starting_pitcher, type=BYTE_ARRAY, convertedtype=UTF8"`
	// StartingPitcherRecord is the record of the team's starting pitcher
	// Please note: the API does not provide PIT starting pitcher record! This field is not reliable for PIT calculations.
	StartingPitcherRecord *string `json:"starting_pitcher_record" parquet:"name=starting_pitcher_record, type=BYTE_ARRAY, convertedtype=UTF8"`
	// StartingPitcherERA is the team's starting pitcher's earned run average - https://www.mlb.com/glossary/standard-stats/earned-run-average
	// Please note: the API does not provide PIT starting pitcher ERA! This field is not reliable for PIT calculations.
	StartingPitcherERA *float32 `json:"starting_pitcher_era" parquet:"name=starting_pitcher_era, type=FLOAT"`
}
