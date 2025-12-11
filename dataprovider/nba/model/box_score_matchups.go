package model

import "time"

// BoxScoreMatchups - composite key: event_id, player_id, opponent_player_id
type BoxScoreMatchups struct {
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
	// OpponentPlayerID
	OpponentPlayerID int64 `json:"opponent_player_id" parquet:"name=opponent_player_id, type=INT64"`
	// OpponentPlayerName
	OpponentPlayerName string `json:"opponent_player_name" parquet:"name=opponent_player_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// MatchupMinutes
	MatchupMinutes float32 `json:"matchup_minutes" parquet:"name=matchup_minutes, type=FLOAT"`
	// MatchupMinutesSort
	MatchupMinutesSort float32 `json:"matchup_minutes_sort" parquet:"name=matchup_minutes_sort, type=FLOAT"`
	// PartialPossessions
	PartialPossessions float32 `json:"partial_possessions" parquet:"name=partial_possessions, type=FLOAT"`
	// PercentageDefenderTotalTime
	PercentageDefenderTotalTime float32 `json:"percentage_defender_total_time" parquet:"name=percentage_defender_total_time, type=FLOAT"`
	// PercentageOffensiveTotalTime
	PercentageOffensiveTotalTime float32 `json:"percentage_offensive_total_time" parquet:"name=percentage_offensive_total_time, type=FLOAT"`
	// PercentageTotalTimeBothOn
	PercentageTotalTimeBothOn float32 `json:"percentage_total_time_both_on" parquet:"name=percentage_total_time_both_on, type=FLOAT"`
	// SwitchesOn
	SwitchesOn int32 `json:"switches_on" parquet:"name=switches_on, type=INT32"`
	// PlayerPoints
	PlayerPoints int32 `json:"player_points" parquet:"name=player_points, type=INT32"`
	// TeamPoints
	TeamPoints int32 `json:"team_points" parquet:"name=team_points, type=INT32"`
	// MatchupAssists
	MatchupAssists int32 `json:"matchup_assists" parquet:"name=matchup_assists, type=INT32"`
	// MatchupPotentialAssists
	MatchupPotentialAssists int32 `json:"matchup_potential_assists" parquet:"name=matchup_potential_assists, type=INT32"`
	// MatchupTurnovers
	MatchupTurnovers int32 `json:"matchup_turnovers" parquet:"name=matchup_turnovers, type=INT32"`
	// MatchupBlocks
	MatchupBlocks int32 `json:"matchup_blocks" parquet:"name=matchup_blocks, type=INT32"`
	// MatchupFieldGoalsMade
	MatchupFieldGoalsMade int32 `json:"matchup_field_goals_made" parquet:"name=matchup_field_goals_made, type=INT32"`
	// MatchupFieldGoalsAttempted
	MatchupFieldGoalsAttempted int32 `json:"matchup_field_goals_attempted" parquet:"name=matchup_field_goals_attempted, type=INT32"`
	// MatchupFieldGoalsPercentage
	MatchupFieldGoalsPercentage float32 `json:"matchup_field_goals_percentage" parquet:"name=matchup_field_goals_percentage, type=FLOAT"`
	// MatchupThreePointersMade
	MatchupThreePointersMade int32 `json:"matchup_three_pointers_made" parquet:"name=matchup_three_pointers_made, type=INT32"`
	// MatchupThreePointersAttempted
	MatchupThreePointersAttempted int32 `json:"matchup_three_pointers_attempted" parquet:"name=matchup_three_pointers_attempted, type=INT32"`
	// MatchupThreePointersPercentage
	MatchupThreePointersPercentage float32 `json:"matchup_three_pointers_percentage" parquet:"name=matchup_three_pointers_percentage, type=FLOAT"`
	// HelpBlocks
	HelpBlocks int32 `json:"help_blocks" parquet:"name=help_blocks, type=INT32"`
	// HelpFieldGoalsMade
	HelpFieldGoalsMade int32 `json:"help_field_goals_made" parquet:"name=help_field_goals_made, type=INT32"`
	// HelpFieldGoalsAttempted
	HelpFieldGoalsAttempted int32 `json:"help_field_goals_attempted" parquet:"name=help_field_goals_attempted, type=INT32"`
	// HelpFieldGoalsPercentage
	HelpFieldGoalsPercentage float32 `json:"help_field_goals_percentage" parquet:"name=help_field_goals_percentage, type=FLOAT"`
	// MatchupFreeThrowsMade
	MatchupFreeThrowsMade int32 `json:"matchup_free_throws_made" parquet:"name=matchup_free_throws_made, type=INT32"`
	// MatchupFreeThrowsAttempted
	MatchupFreeThrowsAttempted int32 `json:"matchup_free_throws_attempted" parquet:"name=matchup_free_throws_attempted, type=INT32"`
	// ShootingFouls
	ShootingFouls int32 `json:"shooting_fouls" parquet:"name=shooting_fouls, type=INT32"`
}
