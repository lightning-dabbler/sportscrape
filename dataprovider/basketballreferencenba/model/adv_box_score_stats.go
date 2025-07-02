package model

import "time"

// NBAAdvBoxScoreStats represents the data model for NBA advanced box score stats scraped from basketball-reference.com
type NBAAdvBoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// EventDateParquet is the timestamp associated with a given event (in days)
	EventDateParquet int32 `json:"-" parquet:"name=event_date, type=INT32, convertedtype=DATE, logicaltype=DATE"`
	// TeamID
	TeamID string `json:"team_id" parquet:"name=team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Team is the player's team name
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID string `json:"opponent_id" parquet:"name=opponent_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID is the player's id extracted from the player's basketball-reference profile url
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Player is the player's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerLink is the link to the player's basketball-reference profile
	PlayerLink string `json:"player_link" parquet:"name=player_link, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Starter is whether the player was apart of the team's starting five during the event or not
	Starter bool `json:"starter" parquet:"name=starter, type=BOOLEAN"`
	// MinutesPlayed is minutes played during the event
	MinutesPlayed float32 `json:"minutes_played" parquet:"name=minutes_played, type=FLOAT"`
	// TrueShootingPercentage - A measure of shooting efficiency that takes into account 2-point field goals, 3-point field goals, and free throws (in decimal notation).
	TrueShootingPercentage float32 `json:"true_shooting_percentage" parquet:"name=true_shooting_percentage, type=FLOAT"`
	// EffectiveFieldGoalPercentage - This statistic adjusts for the fact that a 3-point field goal is worth one more point than a 2-point field goal (in decimal notation).
	EffectiveFieldGoalPercentage float32 `json:"effective_field_goal_percentage" parquet:"name=effective_field_goal_percentage, type=FLOAT"`
	// ThreePointAttemptRate - Percentage of FG Attempts from 3-Point Range (in decimal notation).
	ThreePointAttemptRate float32 `json:"three_point_attempt_rate" parquet:"name=three_point_attempt_rate, type=FLOAT"`
	// FreeThrowAttemptRate - Number of FT Attempts Per FG Attempt (in decimal notation).
	FreeThrowAttemptRate float32 `json:"free_throw_attempt_rate" parquet:"name=free_throw_attempt_rate, type=FLOAT"`
	// OffensiveReboundPercentage - An estimate of the percentage of available offensive rebounds a player grabbed while they were on the floor (in percent notation).
	OffensiveReboundPercentage float32 `json:"offensive_rebound_percentage" parquet:"name=offensive_rebound_percentage, type=FLOAT"`
	// DefensiveReboundPercentage - An estimate of the percentage of available defensive rebounds a player grabbed while they were on the floor (in percent notation).
	DefensiveReboundPercentage float32 `json:"defensive_rebound_percentage" parquet:"name=defensive_rebound_percentage, type=FLOAT"`
	// TotalReboundPercentage - An estimate of the percentage of available rebounds a player grabbed while they were on the floor (in percent notation).
	TotalReboundPercentage float32 `json:"total_rebound_percentage" parquet:"name=total_rebound_percentage, type=FLOAT"`
	// AssistPercentage - An estimate of the percentage of teammate field goals a player assisted while they were on the floor (in percent notation).
	AssistPercentage float32 `json:"assist_percentage" parquet:"name=assist_percentage, type=FLOAT"`
	// StealPercentage - An estimate of the percentage of opponent possessions that end with a steal by the player while they were on the floor (in percent notation).
	StealPercentage float32 `json:"steal_percentage" parquet:"name=steal_percentage, type=FLOAT"`
	// BlockPercentage - An estimate of the percentage of opponent two-point field goal attempts blocked by the player while they were on the floor (in percent notation).
	BlockPercentage float32 `json:"block_percentage" parquet:"name=block_percentage, type=FLOAT"`
	// TurnoverPercentage - An estimate of turnovers committed per 100 plays (in percent notation).
	TurnoverPercentage float32 `json:"turnover_percentage" parquet:"name=turnover_percentage, type=FLOAT"`
	// UsagePercentage - An estimate of the percentage of team plays used by a player while they were on the floor (in percent notation).
	UsagePercentage float32 `json:"usage_percentage" parquet:"name=usage_percentage, type=FLOAT"`
	// OffensiveRating - An estimate of points produced (players) or scored (teams) per 100 possessions
	OffensiveRating int32 `json:"offensive_rating" parquet:"name=offensive_rating, type=INT32"`
	// DefensiveRating - An estimate of points allowed per 100 possessions
	DefensiveRating int32 `json:"defensive_rating" parquet:"name=defensive_rating, type=INT32"`
	// BoxPlusMinus - A box score estimate of the points per 100 possessions a player contributed above a league-average player, translated to an average team.
	BoxPlusMinus float32 `json:"box_plus_minus" parquet:"name=box_plus_minus, type=FLOAT"`
}
