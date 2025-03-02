package model

import (
	"time"
)

// NBAAdvBoxScoreStats represents the data model for NBA advanced box score stats scraped from basketball-reference.com
type NBAAdvBoxScoreStats struct {
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
	// TrueShootingPercentage - A measure of shooting efficiency that takes into account 2-point field goals, 3-point field goals, and free throws (in decimal notation).
	TrueShootingPercentage float32 `json:"true_shooting_percentage"`
	// EffectiveFieldGoalPercentage - This statistic adjusts for the fact that a 3-point field goal is worth one more point than a 2-point field goal (in decimal notation).
	EffectiveFieldGoalPercentage float32 `json:"effective_field_goal_percentage"`
	// ThreePointAttemptRate - Percentage of FG Attempts from 3-Point Range (in decimal notation).
	ThreePointAttemptRate float32 `json:"three_point_attempt_rate"`
	// FreeThrowAttemptRate - Number of FT Attempts Per FG Attempt (in decimal notation).
	FreeThrowAttemptRate float32 `json:"free_throw_attempt_rate"`
	// OffensiveReboundPercentage - An estimate of the percentage of available offensive rebounds a player grabbed while they were on the floor (in percent notation).
	OffensiveReboundPercentage float32 `json:"offensive_rebound_percentage"`
	// DefensiveReboundPercentage - An estimate of the percentage of available defensive rebounds a player grabbed while they were on the floor (in percent notation).
	DefensiveReboundPercentage float32 `json:"defensive_rebound_percentage"`
	// TotalReboundPercentage - An estimate of the percentage of available rebounds a player grabbed while they were on the floor (in percent notation).
	TotalReboundPercentage float32 `json:"total_rebound_percentage"`
	// AssistPercentage - An estimate of the percentage of teammate field goals a player assisted while they were on the floor (in percent notation).
	AssistPercentage float32 `json:"assist_percentage"`
	// StealPercentage - An estimate of the percentage of opponent possessions that end with a steal by the player while they were on the floor (in percent notation).
	StealPercentage float32 `json:"steal_percentage"`
	// BlockPercentage - An estimate of the percentage of opponent two-point field goal attempts blocked by the player while they were on the floor (in percent notation).
	BlockPercentage float32 `json:"block_percentage"`
	// TurnoverPercentage - An estimate of turnovers committed per 100 plays (in percent notation).
	TurnoverPercentage float32 `json:"turnover_percentage"`
	// UsagePercentage - An estimate of the percentage of team plays used by a player while they were on the floor (in percent notation).
	UsagePercentage float32 `json:"usage_percentage"`
	// OffensiveRating - An estimate of points produced (players) or scored (teams) per 100 possessions
	OffensiveRating int `json:"offensive_rating"`
	// DefensiveRating - An estimate of points allowed per 100 possessions
	DefensiveRating int `json:"defensive_rating"`
	// BoxPlusMinus - A box score estimate of the points per 100 possessions a player contributed above a league-average player, translated to an average team.
	BoxPlusMinus float32 `json:"box_plus_minus"`
}
