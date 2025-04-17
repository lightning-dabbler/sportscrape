package model

import "time"

// MLBPitchingBoxScoreStats - data model for MLB pitching box score stats
type MLBPitchingBoxScoreStats struct {
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
	// TeamID is the player's team's ID e.g. 21
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// Team is the player's team name e.g. Atlanta Braves
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY"`
	// OpponentID is the opposing team's team id
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY"`
	// PlayerID is the player's id
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// Player is the player's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY"`
	// Record is starting pitcher's record
	Record *string `json:"record" parquet:"name=record, type=BYTE_ARRAY"`
	// PitchingOrder - The sequence of pitchers who played during the event per team (starting from 1)
	PitchingOrder int32 `json:"pitching_order" parquet:"name=pitching_order, type=INT32"`
	// InningsPitched (IP) - https://www.mlb.com/glossary/standard-stats/innings-pitched
	InningsPitched float32 `json:"innings_pitched" parquet:"name=innings_pitched, type=FLOAT"`
	// HitsAllowed (H) - https://www.mlb.com/glossary/standard-stats/hit
	HitsAllowed int32 `json:"hits_allowed" parquet:"name=hits_allowed, type=INT32"`
	// RunsAllowed (R) - https://www.mlb.com/glossary/standard-stats/run
	RunsAllowed int32 `json:"runs_allowed" parquet:"name=runs_allowed, type=INT32"`
	// EarnedRunsAllowed (ER) - https://www.mlb.com/glossary/standard-stats/earned-run
	EarnedRunsAllowed int32 `json:"earned_runs_allowed" parquet:"name=earned_runs_allowed, type=INT32"`
	// Walks (BB) - https://www.mlb.com/glossary/standard-stats/walk
	Walks int32 `json:"walks" parquet:"name=walks, type=INT32"`
	// StrikeOuts (SO) - https://www.mlb.com/glossary/standard-stats/strikeout
	Strikeouts int32 `json:"strikeouts" parquet:"name=strikeouts, type=INT32"`
	// HomeRunsAllowed (HR) - https://www.mlb.com/glossary/standard-stats/home-run
	HomeRunsAllowed int32 `json:"home_runs_allowed" parquet:"name=home_runs_allowed, type=INT32"`
	// EarnedRunAverage (ERA) - https://www.mlb.com/glossary/standard-stats/earned-run-average
	EarnedRunAverage float32 `json:"earned_run_average" parquet:"name=earned_run_average, type=FLOAT"`
}
