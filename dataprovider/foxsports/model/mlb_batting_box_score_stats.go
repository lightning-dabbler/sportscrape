package model

import "time"

// MLBBattingBoxScoreStats - data model for MLB batting box score stats
type MLBBattingBoxScoreStats struct {
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
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID is the opposing team's team id
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID is the player's id
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// Player is the player's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Position is the player's position
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AtBat (AB) - https://www.mlb.com/glossary/standard-stats/at-bat
	AtBat int32 `json:"at_bat" parquet:"name=at_bat, type=INT32"`
	// Runs (R) - https://www.mlb.com/glossary/standard-stats/run
	Runs int32 `json:"runs" parquet:"name=runs, type=INT32"`
	// Hits (H) - https://www.mlb.com/glossary/standard-stats/hit
	Hits int32 `json:"hits" parquet:"name=hits, type=INT32"`
	// RunsBattedIn (RBI) - https://www.mlb.com/glossary/standard-stats/runs-batted-in
	RunsBattedIn int32 `json:"runs_batted_in" parquet:"name=runs_batted_in, type=INT32"`
	// Walks (BB) - https://www.mlb.com/glossary/standard-stats/walk
	Walks int32 `json:"walks" parquet:"name=walks, type=INT32"`
	// Strikeouts (SO) - https://www.mlb.com/glossary/standard-stats/strikeout
	Strikeouts int32 `json:"strikeouts" parquet:"name=strikeouts, type=INT32"`
	// LeftOnBase (LOB) - https://www.mlb.com/glossary/standard-stats/left-on-base
	LeftOnBase int32 `json:"left_on_base" parquet:"name=left_on_base, type=INT32"`
	// BattingAverage (AVG) - https://www.mlb.com/glossary/standard-stats/batting-average
	BattingAverage float32 `json:"batting_average" parquet:"name=batting_average, type=FLOAT"`
}
