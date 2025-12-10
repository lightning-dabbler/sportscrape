package model

import "time"

// BoxScoreMisc - composite key: event_id, player_id
type BoxScoreMisc struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a string ID that maps to the matchup e.g. 0022500249
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventStatus the numerical representation of the event status e.g. 3 (1=pregame, 2=in progress, 3=final)
	EventStatus int32 `json:"event_status" parquet:"name=event_status, type=INT32"`
	// EventStatusText (e.g. Final, Final/OT2, etc.)
	EventStatusText string `json:"event_status_text" parquet:"name=event_status_text, type=BYTE_ARRAY, convertedtype=UTF8"`
	// TeamID
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// TeamName
	TeamName string `json:"team_name" parquet:"name=team_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// TeamNameFull
	TeamNameFull string `json:"team_name_full" parquet:"name=team_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// OpponentName
	OpponentName string `json:"opponent_name" parquet:"name=opponent_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentNameFull
	OpponentNameFull string `json:"opponent_name_full" parquet:"name=opponent_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// PlayerName
	PlayerName string `json:"player_name" parquet:"name=player_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Position
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Starter
	Starter bool `json:"starter" parquet:"name=starter, type=BOOLEAN"`
	// Minutes
	Minutes float32 `json:"minutes" parquet:"name=minutes, type=FLOAT"`
	// PointsOffTurnovers
	PointsOffTurnovers int32 `json:"points_off_turnovers" parquet:"name=points_off_turnovers, type=INT32"`
	// PointsSecondChance
	PointsSecondChance int32 `json:"points_second_chance" parquet:"name=points_second_chance, type=INT32"`
	// PointsFastBreak
	PointsFastBreak int32 `json:"points_fast_break" parquet:"name=points_fast_break, type=INT32"`
	// PointsPaint
	PointsPaint int32 `json:"points_paint" parquet:"name=points_paint, type=INT32"`
	// OppPointsOffTurnovers
	OppPointsOffTurnovers int32 `json:"opp_points_off_turnovers" parquet:"name=opp_points_off_turnovers, type=INT32"`
	// OppPointsSecondChance
	OppPointsSecondChance int32 `json:"opp_points_second_chance" parquet:"name=opp_points_second_chance, type=INT32"`
	// OppPointsFastBreak
	OppPointsFastBreak int32 `json:"opp_points_fast_break" parquet:"name=opp_points_fast_break, type=INT32"`
	// OppPointsPaint
	OppPointsPaint int32 `json:"opp_points_paint" parquet:"name=opp_points_paint, type=INT32"`
	// Blocks
	Blocks int32 `json:"blocks" parquet:"name=blocks, type=INT32"`
	// BlocksAgainst
	BlocksAgainst int32 `json:"blocks_against" parquet:"name=blocks_against, type=INT32"`
	// FoulsPersonal
	FoulsPersonal int32 `json:"fouls_personal" parquet:"name=fouls_personal, type=INT32"`
	// FoulsDrawn
	FoulsDrawn int32 `json:"fouls_drawn" parquet:"name=fouls_drawn, type=INT32"`
}
