package model

import "time"

// NBABoxScoreStats - data model for NBA box score stats
type NBABoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// EventID is a unique ID that maps to the matchup e.g. 86833
	EventID int64 `json:"event_id"`
	// TeamID is the player's team's ID e.g. 21
	TeamID int64 `json:"team_id"`
	// Team is the player's team name e.g. Atlanta Braves
	Team string `json:"team"`
	// OpponentID is the opposing team's team id
	OpponentID int64 `json:"opponent_id"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent"`
	// PlayerID is the player's id
	PlayerID int64 `json:"player_id"`
	// Player is the player's name
	Player string `json:"player"`
	// Position is the player's position
	Position string `json:"position"`
	// Starter is whether the player was apart of the team's starting five during the event or not
	Starter bool `json:"starter"`
	// MinutesPlayed is minutes played during the event
	MinutesPlayed int `json:"minutes_played"`
	// FieldGoalsMade is the number of field goals made
	FieldGoalsMade int `json:"field_goals_made"`
	// FieldGoalAttempts is the number of field goals attempted
	FieldGoalAttempts int `json:"field_goal_attempts"`
	// ThreePointsMade is the number of three pointers made
	ThreePointsMade int `json:"three_points_made"`
	// ThreePointAttempts is the number of three point attempts
	ThreePointAttempts int `json:"three_point_attempts"`
	// FreeThrowsMade is the number of free throws made
	FreeThrowsMade int `json:"free_throws_made"`
	// FreeThrowAttempts is the number of free throw attempts
	FreeThrowAttempts int `json:"free_throw_attempts"`
	// OffensiveRebounds is the number of offensive rebounds
	OffensiveRebounds int `json:"offensive_rebounds"`
	// DefensiveRebounds is the number of defensive rebounds
	DefensiveRebounds int `json:"defensive_rebounds"`
	// TotalRebounds is the number of total rebounds
	TotalRebounds int `json:"total_rebounds"`
	// Assists is the number of assists
	Assists int `json:"assists"`
	// Steals is the number of steals
	Steals int `json:"steals"`
	// Blocks is the number of blocks
	Blocks int `json:"blocks"`
	// Turnovers is the number of turnovers
	Turnovers int `json:"turnovers"`
	// PersonalFouls is the number of personal fouls
	PersonalFouls int `json:"personal_fouls"`
	// Points is the number of points
	Points int `json:"points"`
}
