package model

import (
	"time"
)

// NBABasicBoxScoreStats represents the data model for NBA basic box score stats scraped from basketball-reference.com
type NBABasicBoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// EventID is the event id of the matchup associated with the box score
	EventID string `json:"event_id"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// Team is the player's team name
	Team string `json:"team"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent"`
	// PlayerID is the player's id extracted from the player's basketball-reference profile url
	PlayerID string `json:"player_id"`
	// Player is the player's name
	Player string `json:"player"`
	// PlayerLink is the link to the player's basketball-reference profile
	PlayerLink string `json:"player_link"`
	// Starter is whether the player was apart of the team's starting five during the event or not
	Starter bool `json:"starter"`
	// MinutesPlayed is minutes played during the event
	MinutesPlayed float32 `json:"minutes_played"`
	// FieldGoalsMade is the number of field goals made
	FieldGoalsMade int `json:"field_goals_made"`
	// FieldGoalAttempts is the number of field goals attempted
	FieldGoalAttempts int `json:"field_goal_attempts"`
	// FieldGoalPercentage is the field goal percentage
	FieldGoalPercentage float32 `json:"field_goal_percentage"`
	// ThreePointsMade is the number of three pointers made
	ThreePointsMade int `json:"three_points_made"`
	// ThreePointAttempts is the number of three point attempts
	ThreePointAttempts int `json:"three_point_attempts"`
	// ThreePointPercentage is the three point percentage
	ThreePointPercentage float32 `json:"three_point_percentage"`
	// FreeThrowsMade is the number of free throws made
	FreeThrowsMade int `json:"free_throws_made"`
	// FreeThrowAttempts is the number of free throw attempts
	FreeThrowAttempts int `json:"free_throw_attempts"`
	// FreeThrowPercentage is the free throw percentage
	FreeThrowPercentage float32 `json:"free_throw_percentage"`
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
	// GameScore is a metric used to evaluate how well a player performs in a single game
	GameScore float32 `json:"game_score"`
	// PlusMinus is the plus-minus
	PlusMinus int `json:"plus_minus"`
}
