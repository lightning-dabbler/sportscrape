package model

import "time"

// BoxScoreUsage - composite key: event_id, player_id
type BoxScoreUsage struct {
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
	// UsagePercentage
	UsagePercentage float32 `json:"usage_percentage" parquet:"name=usage_percentage, type=FLOAT"`
	// PercentageFieldGoalsMade
	PercentageFieldGoalsMade float32 `json:"percentage_field_goals_made" parquet:"name=percentage_field_goals_made, type=FLOAT"`
	// PercentageFieldGoalsAttempted
	PercentageFieldGoalsAttempted float32 `json:"percentage_field_goals_attempted" parquet:"name=percentage_field_goals_attempted, type=FLOAT"`
	// PercentageThreePointersMade
	PercentageThreePointersMade float32 `json:"percentage_three_pointers_made" parquet:"name=percentage_three_pointers_made, type=FLOAT"`
	// PercentageThreePointersAttempted
	PercentageThreePointersAttempted float32 `json:"percentage_three_pointers_attempted" parquet:"name=percentage_three_pointers_attempted, type=FLOAT"`
	// PercentageFreeThrowsMade
	PercentageFreeThrowsMade float32 `json:"percentage_free_throws_made" parquet:"name=percentage_free_throws_made, type=FLOAT"`
	// PercentageFreeThrowsAttempted
	PercentageFreeThrowsAttempted float32 `json:"percentage_free_throws_attempted" parquet:"name=percentage_free_throws_attempted, type=FLOAT"`
	// PercentageReboundsOffensive
	PercentageReboundsOffensive float32 `json:"percentage_rebounds_offensive" parquet:"name=percentage_rebounds_offensive, type=FLOAT"`
	// PercentageReboundsDefensive
	PercentageReboundsDefensive float32 `json:"percentage_rebounds_defensive" parquet:"name=percentage_rebounds_defensive, type=FLOAT"`
	// PercentageReboundsTotal
	PercentageReboundsTotal float32 `json:"percentage_rebounds_total" parquet:"name=percentage_rebounds_total, type=FLOAT"`
	// PercentageAssists
	PercentageAssists float32 `json:"percentage_assists" parquet:"name=percentage_assists, type=FLOAT"`
	// PercentageTurnovers
	PercentageTurnovers float32 `json:"percentage_turnovers" parquet:"name=percentage_turnovers, type=FLOAT"`
	// PercentageSteals
	PercentageSteals float32 `json:"percentage_steals" parquet:"name=percentage_steals, type=FLOAT"`
	// PercentageBlocks
	PercentageBlocks float32 `json:"percentage_blocks" parquet:"name=percentage_blocks, type=FLOAT"`
	// PercentageBlocksAllowed
	PercentageBlocksAllowed float32 `json:"percentage_blocks_allowed" parquet:"name=percentage_blocks_allowed, type=FLOAT"`
	// PercentagePersonalFouls
	PercentagePersonalFouls float32 `json:"percentage_personal_fouls" parquet:"name=percentage_personal_fouls, type=FLOAT"`
	// PercentagePersonalFoulsDrawn
	PercentagePersonalFoulsDrawn float32 `json:"percentage_personal_fouls_drawn" parquet:"name=percentage_personal_fouls_drawn, type=FLOAT"`
	// PercentagePoints
	PercentagePoints float32 `json:"percentage_points" parquet:"name=percentage_points, type=FLOAT"`
}
