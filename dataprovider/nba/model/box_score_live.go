package model

import "time"

// BoxScoreLive - composite key: event_id, player_id
type BoxScoreLive struct {
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
	// Status
	Status string `json:"status" parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Assists
	Assists int32 `json:"assists" parquet:"name=assists, type=INT32"`
	// Blocks
	Blocks int32 `json:"blocks" parquet:"name=blocks, type=INT32"`
	// BlocksReceived
	BlocksReceived int32 `json:"blocks_received" parquet:"name=blocks_received, type=INT32"`
	// FieldGoalsAttempted
	FieldGoalsAttempted int32 `json:"field_goals_attempted" parquet:"name=field_goals_attempted, type=INT32"`
	// FieldGoalsMade
	FieldGoalsMade int32 `json:"field_goals_made" parquet:"name=field_goals_made, type=INT32"`
	// FieldGoalsPercentage
	FieldGoalsPercentage float32 `json:"field_goals_percentage" parquet:"name=field_goals_percentage, type=FLOAT"`
	// FoulsOffensive
	FoulsOffensive int32 `json:"fouls_offensive" parquet:"name=fouls_offensive, type=INT32"`
	// FoulsDrawn
	FoulsDrawn int32 `json:"fouls_drawn" parquet:"name=fouls_drawn, type=INT32"`
	// FoulsPersonal
	FoulsPersonal int32 `json:"fouls_personal" parquet:"name=fouls_personal, type=INT32"`
	// FoulsTechnical
	FoulsTechnical int32 `json:"fouls_technical" parquet:"name=fouls_technical, type=INT32"`
	// FreeThrowsAttempted
	FreeThrowsAttempted int32 `json:"free_throws_attempted" parquet:"name=free_throws_attempted, type=INT32"`
	// FreeThrowsMade
	FreeThrowsMade int32 `json:"free_throws_made" parquet:"name=free_throws_made, type=INT32"`
	// FreeThrowsPercentage
	FreeThrowsPercentage float32 `json:"free_throws_percentage" parquet:"name=free_throws_percentage, type=FLOAT"`
	// Minus
	Minus int32 `json:"minus" parquet:"name=minus, type=INT32"`
	// Minutes
	Minutes float32 `json:"minutes" parquet:"name=minutes, type=FLOAT"`
	// MinutesCalculated
	MinutesCalculated float32 `json:"minutes_calculated" parquet:"name=minutes_calculated, type=FLOAT"`
	// Plus
	Plus int32 `json:"plus" parquet:"name=plus, type=INT32"`
	// PlusMinusPoints
	PlusMinusPoints int32 `json:"plus_minus_points" parquet:"name=plus_minus_points, type=INT32"`
	// Points
	Points int32 `json:"points" parquet:"name=points, type=INT32"`
	// PointsFastBreak
	PointsFastBreak int32 `json:"points_fast_break" parquet:"name=points_fast_break, type=INT32"`
	// PointsInThePaint
	PointsInThePaint int32 `json:"points_in_the_paint" parquet:"name=points_in_the_paint, type=INT32"`
	// PointsSecondChance
	PointsSecondChance int32 `json:"points_second_chance" parquet:"name=points_second_chance, type=INT32"`
	// ReboundsDefensive
	ReboundsDefensive int32 `json:"rebounds_defensive" parquet:"name=rebounds_defensive, type=INT32"`
	// ReboundsOffensive
	ReboundsOffensive int32 `json:"rebounds_offensive" parquet:"name=rebounds_offensive, type=INT32"`
	// ReboundsTotal
	ReboundsTotal int32 `json:"rebounds_total" parquet:"name=rebounds_total, type=INT32"`
	// Steals
	Steals int32 `json:"steals" parquet:"name=steals, type=INT32"`
	// ThreePointersAttempted
	ThreePointersAttempted int32 `json:"three_pointers_attempted" parquet:"name=three_pointers_attempted, type=INT32"`
	// ThreePointersMade
	ThreePointersMade int32 `json:"three_pointers_made" parquet:"name=three_pointers_made, type=INT32"`
	// ThreePointersPercentage
	ThreePointersPercentage float32 `json:"three_pointers_percentage" parquet:"name=three_pointers_percentage, type=FLOAT"`
	// Turnovers
	Turnovers int32 `json:"turnovers" parquet:"name=turnovers, type=INT32"`
	// TwoPointersAttempted
	TwoPointersAttempted int32 `json:"two_pointers_attempted" parquet:"name=two_pointers_attempted, type=INT32"`
	// TwoPointersMade
	TwoPointersMade int32 `json:"two_pointers_made" parquet:"name=two_pointers_made, type=INT32"`
	// TwoPointersPercentage
	TwoPointersPercentage float32 `json:"two_pointers_percentage" parquet:"name=two_pointers_percentage, type=FLOAT"`
}
