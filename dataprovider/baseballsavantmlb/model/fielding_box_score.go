package model

import "time"

// FieldingBoxScore represents the data model for MLB fielding box score stats scraped from baseballsavant.mlb.com
type FieldingBoxScore struct {
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
	// TeamID
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// Team is the player's team name
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// PlayerID
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// Player is the pitcher's name
	Player   string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
	// CaughtStealing - https://www.mlb.com/glossary/standard-stats/caught-stealing
	CaughtStealing int32 `json:"caught_stealing" parquet:"name=caught_stealing, type=INT32"`
	// StolenBases - https://www.mlb.com/glossary/standard-stats/stolen-base
	StolenBases int32 `json:"stolen_bases" parquet:"name=stolen_bases, type=INT32"`
	// Assists - https://www.mlb.com/glossary/standard-stats/assist
	Assists int32 `json:"assists" parquet:"name=assists, type=INT32"`
	// Putouts - https://www.mlb.com/glossary/standard-stats/putout
	Putouts int32 `json:"putouts" parquet:"name=putouts, type=INT32"`
	// Errors - https://www.mlb.com/glossary/standard-stats/error
	Errors int32 `json:"errors" parquet:"name=errors, type=INT32"`
	// Chances - https://www.mlb.com/glossary/standard-stats/total-chances
	Chances int32 `json:"chances" parquet:"name=chances, type=INT32"`
	// PassedBall - https://www.mlb.com/glossary/standard-stats/passed-ball
	PassedBall int32 `json:"passed_ball" parquet:"name=passed_ball, type=INT32"`
	// Pickoffs - https://www.mlb.com/glossary/standard-stats/pickoff
	Pickoffs int32 `json:"pickoffs" parquet:"name=pickoffs, type=INT32"`
}
