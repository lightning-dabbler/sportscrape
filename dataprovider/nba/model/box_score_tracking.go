package model

import "time"

// BoxScoreTracking - composite key: event_id, player_id
type BoxScoreTracking struct {
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
	// Speed
	Speed float32 `json:"speed" parquet:"name=speed, type=FLOAT"`
	// Distance
	Distance float32 `json:"distance" parquet:"name=distance, type=FLOAT"`
	// ReboundChancesOffensive
	ReboundChancesOffensive int32 `json:"rebound_chances_offensive" parquet:"name=rebound_chances_offensive, type=INT32"`
	// ReboundChancesDefensive
	ReboundChancesDefensive int32 `json:"rebound_chances_defensive" parquet:"name=rebound_chances_defensive, type=INT32"`
	// ReboundChancesTotal
	ReboundChancesTotal int32 `json:"rebound_chances_total" parquet:"name=rebound_chances_total, type=INT32"`
	// Touches
	Touches int32 `json:"touches" parquet:"name=touches, type=INT32"`
	// SecondaryAssists
	SecondaryAssists int32 `json:"secondary_assists" parquet:"name=secondary_assists, type=INT32"`
	// FreeThrowAssists
	FreeThrowAssists int32 `json:"free_throw_assists" parquet:"name=free_throw_assists, type=INT32"`
	// Passes
	Passes int32 `json:"passes" parquet:"name=passes, type=INT32"`
	// Assists
	Assists int32 `json:"assists" parquet:"name=assists, type=INT32"`
	// ContestedFieldGoalsMade
	ContestedFieldGoalsMade int32 `json:"contested_field_goals_made" parquet:"name=contested_field_goals_made, type=INT32"`
	// ContestedFieldGoalsAttempted
	ContestedFieldGoalsAttempted int32 `json:"contested_field_goals_attempted" parquet:"name=contested_field_goals_attempted, type=INT32"`
	// ContestedFieldGoalPercentage
	ContestedFieldGoalPercentage float32 `json:"contested_field_goal_percentage" parquet:"name=contested_field_goal_percentage, type=FLOAT"`
	// UncontestedFieldGoalsMade
	UncontestedFieldGoalsMade int32 `json:"uncontested_field_goals_made" parquet:"name=uncontested_field_goals_made, type=INT32"`
	// UncontestedFieldGoalsAttempted
	UncontestedFieldGoalsAttempted int32 `json:"uncontested_field_goals_attempted" parquet:"name=uncontested_field_goals_attempted, type=INT32"`
	// UncontestedFieldGoalsPercentage
	UncontestedFieldGoalsPercentage float32 `json:"uncontested_field_goals_percentage" parquet:"name=uncontested_field_goals_percentage, type=FLOAT"`
	// FieldGoalPercentage
	FieldGoalPercentage float32 `json:"field_goal_percentage" parquet:"name=field_goal_percentage, type=FLOAT"`
	// DefendedAtRimFieldGoalsMade
	DefendedAtRimFieldGoalsMade int32 `json:"defended_at_rim_field_goals_made" parquet:"name=defended_at_rim_field_goals_made, type=INT32"`
	// DefendedAtRimFieldGoalsAttempted
	DefendedAtRimFieldGoalsAttempted int32 `json:"defended_at_rim_field_goals_attempted" parquet:"name=defended_at_rim_field_goals_attempted, type=INT32"`
	// DefendedAtRimFieldGoalPercentage
	DefendedAtRimFieldGoalPercentage float32 `json:"defended_at_rim_field_goalp_ercentage" parquet:"name=defended_at_rim_field_goalp_ercentage, type=FLOAT"`
}
