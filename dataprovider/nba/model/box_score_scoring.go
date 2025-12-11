package model

import "time"

// BoxScoreScoring - composite key: event_id, player_id
type BoxScoreScoring struct {
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
	// PercentageFieldGoalsAttempted2pt
	PercentageFieldGoalsAttempted2pt float32 `json:"percentage_field_goals_attempted_2pt" parquet:"name=percentage_field_goals_attempted_2pt, type=FLOAT"`
	// PercentageFieldGoalsAttempted3pt
	PercentageFieldGoalsAttempted3pt float32 `json:"percentage_field_goals_attempted_3pt" parquet:"name=percentage_field_goals_attempted_3pt, type=FLOAT"`
	// PercentagePoints2pt
	PercentagePoints2pt float32 `json:"percentage_points_2pt" parquet:"name=percentage_points_2pt, type=FLOAT"`
	// PercentagePointsMidrange2pt
	PercentagePointsMidrange2pt float32 `json:"percentage_points_midrange_2pt" parquet:"name=percentage_points_midrange_2pt, type=FLOAT"`
	// PercentagePoints3pt
	PercentagePoints3pt float32 `json:"percentage_points_3pt" parquet:"name=percentage_points_3pt, type=FLOAT"`
	// PercentagePointsFastBreak
	PercentagePointsFastBreak float32 `json:"percentage_points_fast_break" parquet:"name=percentage_points_fast_break, type=FLOAT"`
	// PercentagePointsFreeThrow
	PercentagePointsFreeThrow float32 `json:"percentage_points_free_throw" parquet:"name=percentage_points_free_throw, type=FLOAT"`
	// PercentagePointsOffTurnovers
	PercentagePointsOffTurnovers float32 `json:"percentage_points_off_turnovers" parquet:"name=percentage_points_off_turnovers, type=FLOAT"`
	// PercentagePointsPaint
	PercentagePointsPaint float32 `json:"percentage_points_paint" parquet:"name=percentage_points_paint, type=FLOAT"`
	// PercentageAssisted2pt
	PercentageAssisted2pt float32 `json:"percentage_assisted_2pt" parquet:"name=percentage_assisted_2pt, type=FLOAT"`
	// PercentageUnassisted2pt
	PercentageUnassisted2pt float32 `json:"percentage_unassisted_2pt" parquet:"name=percentage_unassisted_2pt, type=FLOAT"`
	// PercentageAssisted3pt
	PercentageAssisted3pt float32 `json:"percentage_assisted_3pt" parquet:"name=percentage_assisted_3pt, type=FLOAT"`
	// PercentageUnassisted3pt
	PercentageUnassisted3pt float32 `json:"percentage_unassisted_3pt" parquet:"name=percentage_unassisted_3pt, type=FLOAT"`
	// PercentageAssistedFGM
	PercentageAssistedFGM float32 `json:"percentage_assisted_fgm" parquet:"name=percentage_assisted_fgm, type=FLOAT"`
	// PercentageUnassistedFGM
	PercentageUnassistedFGM float32 `json:"percentage_unassisted_fgm" parquet:"name=percentage_unassisted_fgm, type=FLOAT"`
}
