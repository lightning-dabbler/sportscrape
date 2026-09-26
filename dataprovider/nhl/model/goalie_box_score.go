package model

import "time"

// GoalieBoxScore - composite key: event_id, player_id
type GoalieBoxScore struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 2024020250
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the scheduled start time of the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// TeamID
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// Team e.g. Flames
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// Opponent e.g. Canucks
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// Player - full name from the player header e.g. Dan Vladar (falls back to the box score name e.g. D. Vladar only when the player header has no first/last name)
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SweaterNumber
	SweaterNumber int32 `json:"sweater_number" parquet:"name=sweater_number, type=INT32"`
	// Position e.g. G
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EvenStrengthSaves
	EvenStrengthSaves int32 `json:"even_strength_saves" parquet:"name=even_strength_saves, type=INT32"`
	// EvenStrengthShotsAgainst
	EvenStrengthShotsAgainst int32 `json:"even_strength_shots_against" parquet:"name=even_strength_shots_against, type=INT32"`
	// PowerPlaySaves
	PowerPlaySaves int32 `json:"power_play_saves" parquet:"name=power_play_saves, type=INT32"`
	// PowerPlayShotsAgainst
	PowerPlayShotsAgainst int32 `json:"power_play_shots_against" parquet:"name=power_play_shots_against, type=INT32"`
	// ShorthandedSaves
	ShorthandedSaves int32 `json:"shorthanded_saves" parquet:"name=shorthanded_saves, type=INT32"`
	// ShorthandedShotsAgainst
	ShorthandedShotsAgainst int32 `json:"shorthanded_shots_against" parquet:"name=shorthanded_shots_against, type=INT32"`
	// Saves - nil in limited scoring games
	Saves *int32 `json:"saves" parquet:"name=saves, type=INT32"`
	// ShotsAgainst
	ShotsAgainst int32 `json:"shots_against" parquet:"name=shots_against, type=INT32"`
	// SavePctg e.g. 0.90625; nil when no shots were faced
	SavePctg *float32 `json:"save_pctg" parquet:"name=save_pctg, type=FLOAT"`
	// EvenStrengthGoalsAgainst
	EvenStrengthGoalsAgainst int32 `json:"even_strength_goals_against" parquet:"name=even_strength_goals_against, type=INT32"`
	// PowerPlayGoalsAgainst
	PowerPlayGoalsAgainst int32 `json:"power_play_goals_against" parquet:"name=power_play_goals_against, type=INT32"`
	// ShorthandedGoalsAgainst
	ShorthandedGoalsAgainst int32 `json:"shorthanded_goals_against" parquet:"name=shorthanded_goals_against, type=INT32"`
	// GoalsAgainst
	GoalsAgainst int32 `json:"goals_against" parquet:"name=goals_against, type=INT32"`
	// PIM - penalty minutes; nil when absent from the box score (e.g. limited scoring and some preseason games)
	PIM *int32 `json:"pim" parquet:"name=pim, type=INT32"`
	// TOI - time on ice in minutes e.g. 57.6
	TOI float32 `json:"toi" parquet:"name=toi, type=FLOAT"`
	// Starter - nil when absent from the box score (e.g. limited scoring and some preseason games)
	Starter *bool `json:"starter" parquet:"name=starter, type=BOOLEAN"`
	// Decision e.g. W, L, O; nil when the goalie did not receive a decision
	Decision *string `json:"decision" parquet:"name=decision, type=BYTE_ARRAY, convertedtype=UTF8"`
}
