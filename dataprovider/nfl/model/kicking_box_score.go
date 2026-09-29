package model

import "time"

// KickingBoxScore - composite key: event_id, player_id
type KickingBoxScore struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. f5908b6d-311e-11f0-b670-ae1250fadad1
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventTime is the scheduled start time of the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// TeamID e.g. 10401200-a308-98ca-ad5f-95df2fefea68
	TeamID string `json:"team_id" parquet:"name=team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Team e.g. Dallas Cowboys (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID e.g. 10403700-b939-3cbd-3d16-24d4d6742fa2
	OpponentID string `json:"opponent_id" parquet:"name=opponent_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Opponent e.g. Philadelphia Eagles (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID is the player's GSIS ID e.g. 00-0033077
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PersonID e.g. 32005052-4528-5723-d1b2-96e92ebc1241
	PersonID string `json:"person_id" parquet:"name=person_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Player is the player's full name e.g. Dak Prescott (falls back to PlayerShortName when the full name is unavailable)
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerShortName is the abbreviated player name e.g. D.Prescott
	PlayerShortName string `json:"player_short_name" parquet:"name=player_short_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// JerseyNumber e.g. 04
	JerseyNumber string `json:"jersey_number" parquet:"name=jersey_number, type=BYTE_ARRAY, convertedtype=UTF8"`
	// FieldGoalsMade
	FieldGoalsMade int32 `json:"field_goals_made" parquet:"name=field_goals_made, type=INT32"`
	// FieldGoalsAttempted
	FieldGoalsAttempted int32 `json:"field_goals_attempted" parquet:"name=field_goals_attempted, type=INT32"`
	// FieldGoalsMissed
	FieldGoalsMissed int32 `json:"field_goals_missed" parquet:"name=field_goals_missed, type=INT32"`
	// FieldGoalsBlocked
	FieldGoalsBlocked int32 `json:"field_goals_blocked" parquet:"name=field_goals_blocked, type=INT32"`
	// FieldGoalsLongestMade
	FieldGoalsLongestMade int32 `json:"field_goals_longest_made" parquet:"name=field_goals_longest_made, type=INT32"`
	// FieldGoalsAverageLength
	FieldGoalsAverageLength float32 `json:"field_goals_average_length" parquet:"name=field_goals_average_length, type=FLOAT"`
	// FieldGoalsTotalYards
	FieldGoalsTotalYards int32 `json:"field_goals_total_yards" parquet:"name=field_goals_total_yards, type=INT32"`
	// ExtraPointsMade
	ExtraPointsMade int32 `json:"extra_points_made" parquet:"name=extra_points_made, type=INT32"`
	// ExtraPointsAttempted
	ExtraPointsAttempted int32 `json:"extra_points_attempted" parquet:"name=extra_points_attempted, type=INT32"`
	// ExtraPointsMissed
	ExtraPointsMissed int32 `json:"extra_points_missed" parquet:"name=extra_points_missed, type=INT32"`
	// ExtraPointsBlocked
	ExtraPointsBlocked int32 `json:"extra_points_blocked" parquet:"name=extra_points_blocked, type=INT32"`
}
