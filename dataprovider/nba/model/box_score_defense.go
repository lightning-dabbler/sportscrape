package model

import "time"

// BoxScoreDefense - composite key: event_id, player_id
type BoxScoreDefense struct {
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
	// MatchupMinutes
	MatchupMinutes float32 `json:"matchup_minutes" parquet:"name=matchup_minutes, type=FLOAT"`
	// PartialPossessions
	PartialPossessions float32 `json:"partial_possessions" parquet:"name=partial_possessions, type=FLOAT"`
	// SwitchesOn
	SwitchesOn int32 `json:"switches_on" parquet:"name=switches_on, type=INT32"`
	// PlayerPoints
	PlayerPoints int32 `json:"player_points" parquet:"name=player_points, type=INT32"`
	// DefensiveRebounds
	DefensiveRebounds int32 `json:"defensive_rebounds" parquet:"name=defensive_rebounds, type=INT32"`
	// MatchupAssists
	MatchupAssists int32 `json:"matchup_assists" parquet:"name=matchup_assists, type=INT32"`
	// MatchupTurnovers
	MatchupTurnovers int32 `json:"matchup_turnovers" parquet:"name=matchup_turnovers, type=INT32"`
	// Steals
	Steals int32 `json:"steals" parquet:"name=steals, type=INT32"`
	// Blocks
	Blocks int32 `json:"blocks" parquet:"name=blocks, type=INT32"`
	// MatchupFieldGoalsMade
	MatchupFieldGoalsMade int32 `json:"matchup_field_goals_made" parquet:"name=matchup_field_goals_made, type=INT32"`
	// MatchupFieldGoalsAttempted
	MatchupFieldGoalsAttempted int32 `json:"matchup_field_goals_attempted" parquet:"name=matchup_field_goals_attempted, type=INT32"`
	// MatchupFieldGoalPercentage
	MatchupFieldGoalPercentage float32 `json:"matchup_field_goal_percentage" parquet:"name=matchup_field_goal_percentage, type=FLOAT"`
	// MatchupThreePointersMade
	MatchupThreePointersMade int32 `json:"matchup_three_pointers_made" parquet:"name=matchup_three_pointers_made, type=INT32"`
	// MatchupThreePointersAttempted
	MatchupThreePointersAttempted int32 `json:"matchup_three_pointers_attempted" parquet:"name=matchup_three_pointers_attempted, type=INT32"`
	// MatchupThreePointerPercentage
	MatchupThreePointerPercentage float32 `json:"matchup_three_pointer_percentage" parquet:"name=matchup_three_pointer_percentage, type=FLOAT"`
}
