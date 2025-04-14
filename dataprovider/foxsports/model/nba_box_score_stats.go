package model

import "time"

// NBABoxScoreStats - data model for NBA box score stats
type NBABoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP_MILLIS, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 86833
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP_MILLIS, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// TeamID is the player's team's ID e.g. 21
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// Team is the player's team name e.g. Atlanta Braves
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY"`
	// OpponentID is the opposing team's team id
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY"`
	// PlayerID is the player's id
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// Player is the player's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY"`
	// Position is the player's position
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY"`
	// Starter is whether the player was apart of the team's starting five during the event or not
	Starter bool `json:"starter" parquet:"name=starter, type=BOOLEAN"`
	// MinutesPlayed is minutes played during the event
	MinutesPlayed int32 `json:"minutes_played" parquet:"name=minutes_played, type=INT32"`
	// FieldGoalsMade is the number of field goals made
	FieldGoalsMade int32 `json:"field_goals_made" parquet:"name=field_goals_made, type=INT32"`
	// FieldGoalAttempts is the number of field goals attempted
	FieldGoalAttempts int32 `json:"field_goal_attempts" parquet:"name=field_goal_attempts, type=INT32"`
	// ThreePointsMade is the number of three pointers made
	ThreePointsMade int32 `json:"three_points_made" parquet:"name=three_points_made, type=INT32"`
	// ThreePointAttempts is the number of three point attempts
	ThreePointAttempts int32 `json:"three_point_attempts" parquet:"name=three_point_attempts, type=INT32"`
	// FreeThrowsMade is the number of free throws made
	FreeThrowsMade int32 `json:"free_throws_made" parquet:"name=free_throws_made, type=INT32"`
	// FreeThrowAttempts is the number of free throw attempts
	FreeThrowAttempts int32 `json:"free_throw_attempts" parquet:"name=free_throw_attempts, type=INT32"`
	// OffensiveRebounds is the number of offensive rebounds
	OffensiveRebounds int32 `json:"offensive_rebounds" parquet:"name=offensive_rebounds, type=INT32"`
	// DefensiveRebounds is the number of defensive rebounds
	DefensiveRebounds int32 `json:"defensive_rebounds" parquet:"name=defensive_rebounds, type=INT32"`
	// TotalRebounds is the number of total rebounds
	TotalRebounds int32 `json:"total_rebounds" parquet:"name=total_rebounds, type=INT32"`
	// Assists is the number of assists
	Assists int32 `json:"assists" parquet:"name=assists, type=INT32"`
	// Steals is the number of steals
	Steals int32 `json:"steals" parquet:"name=steals, type=INT32"`
	// Blocks is the number of blocks
	Blocks int32 `json:"blocks" parquet:"name=blocks, type=INT32"`
	// Turnovers is the number of turnovers
	Turnovers int32 `json:"turnovers" parquet:"name=turnovers, type=INT32"`
	// PersonalFouls is the number of personal fouls
	PersonalFouls int32 `json:"personal_fouls" parquet:"name=personal_fouls, type=INT32"`
	// Points is the number of points
	Points int32 `json:"points" parquet:"name=points, type=INT32"`
}
