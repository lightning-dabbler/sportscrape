package model

import "time"

// BoxScoreAdvanced - composite key: event_id, player_id
type BoxScoreAdvanced struct {
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
	// EstimatedOffensiveRating
	EstimatedOffensiveRating float32 `json:"estimated_offensive_rating" parquet:"name=estimated_offensive_rating, type=FLOAT"`
	// OffensiveRating
	OffensiveRating float32 `json:"offensive_rating" parquet:"name=offensive_rating, type=FLOAT"`
	// EstimatedDefensiveRating
	EstimatedDefensiveRating float32 `json:"estimated_defensive_rating" parquet:"name=estimated_defensive_rating, type=FLOAT"`
	// DefensiveRating
	DefensiveRating float32 `json:"defensive_rating" parquet:"name=defensive_rating, type=FLOAT"`
	// EstimatedNetRating
	EstimatedNetRating float32 `json:"estimated_net_rating" parquet:"name=estimated_net_rating, type=FLOAT"`
	// NetRating
	NetRating float32 `json:"net_rating" parquet:"name=net_rating, type=FLOAT"`
	// AssistPercentage
	AssistPercentage float32 `json:"assist_percentage" parquet:"name=assist_percentage, type=FLOAT"`
	// AssistToTurnover
	AssistToTurnover float32 `json:"assist_to_turnover" parquet:"name=assist_to_turnover, type=FLOAT"`
	// AssistRatio
	AssistRatio float32 `json:"assist_ratio" parquet:"name=assist_ratio, type=FLOAT"`
	// OffensiveReboundPercentage
	OffensiveReboundPercentage float32 `json:"offensive_rebound_percentage" parquet:"name=offensive_rebound_percentage, type=FLOAT"`
	// DefensiveReboundPercentage
	DefensiveReboundPercentage float32 `json:"defensive_rebound_percentage" parquet:"name=defensive_rebound_percentage, type=FLOAT"`
	// ReboundPercentage
	ReboundPercentage float32 `json:"rebound_percentage" parquet:"name=rebound_percentage, type=FLOAT"`
	// TurnoverRatio
	TurnoverRatio float32 `json:"turnover_ratio" parquet:"name=turnover_ratio, type=FLOAT"`
	// EffectiveFieldGoalPercentage
	EffectiveFieldGoalPercentage float32 `json:"effective_field_goal_percentage" parquet:"name=effective_field_goal_percentage, type=FLOAT"`
	// TrueShootingPercentage
	TrueShootingPercentage float32 `json:"true_shooting_percentage" parquet:"name=true_shooting_percentage, type=FLOAT"`
	// UsagePercentage
	UsagePercentage float32 `json:"usage_percentage" parquet:"name=usage_percentage, type=FLOAT"`
	// EstimatedUsagePercentage
	EstimatedUsagePercentage float32 `json:"estimated_usage_percentage" parquet:"name=estimated_usage_percentage, type=FLOAT"`
	// EstimatedPace
	EstimatedPace float32 `json:"estimated_pace" parquet:"name=estimated_pace, type=FLOAT"`
	// Pace
	Pace float32 `json:"pace" parquet:"name=pace, type=FLOAT"`
	// PacePer40
	PacePer40 float32 `json:"pace_per_40" parquet:"name=pace_per_40, type=FLOAT"`
	// Possessions
	Possessions int32 `json:"possessions" parquet:"name=possessions, type=INT32"`
	// PIE
	PIE float32 `json:"player_impact_estimate" parquet:"name=player_impact_estimate, type=FLOAT"`
}
