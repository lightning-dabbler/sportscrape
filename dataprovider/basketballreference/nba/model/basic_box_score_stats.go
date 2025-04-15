package model

import "time"

// NBABasicBoxScoreStats represents the data model for NBA basic box score stats scraped from basketball-reference.com
type NBABasicBoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// EventDateParquet is the timestamp associated with a given event (in days)
	EventDateParquet int32 `parquet:"name=event_date, type=INT32, convertedtype=DATE, logicaltype=DATE"`
	// Team is the player's team name
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY"`
	// PlayerID is the player's id extracted from the player's basketball-reference profile url
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY"`
	// Player is the player's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY"`
	// PlayerLink is the link to the player's basketball-reference profile
	PlayerLink string `json:"player_link" parquet:"name=player_link, type=BYTE_ARRAY"`
	// Starter is whether the player was apart of the team's starting five during the event or not
	Starter bool `json:"starter" parquet:"name=starter, type=BOOLEAN"`
	// MinutesPlayed is minutes played during the event
	MinutesPlayed float32 `json:"minutes_played" parquet:"name=minutes_played, type=FLOAT"`
	// FieldGoalsMade is the number of field goals made
	FieldGoalsMade int32 `json:"field_goals_made" parquet:"name=field_goals_made, type=INT32"`
	// FieldGoalAttempts is the number of field goals attempted
	FieldGoalAttempts int32 `json:"field_goal_attempts" parquet:"name=field_goal_attempts, type=INT32"`
	// FieldGoalPercentage is the field goal percentage
	FieldGoalPercentage float32 `json:"field_goal_percentage" parquet:"name=field_goal_percentage, type=FLOAT"`
	// ThreePointsMade is the number of three pointers made
	ThreePointsMade int32 `json:"three_points_made" parquet:"name=three_points_made, type=INT32"`
	// ThreePointAttempts is the number of three point attempts
	ThreePointAttempts int32 `json:"three_point_attempts" parquet:"name=three_point_attempts, type=INT32"`
	// ThreePointPercentage is the three point percentage
	ThreePointPercentage float32 `json:"three_point_percentage" parquet:"name=three_point_percentage, type=FLOAT"`
	// FreeThrowsMade is the number of free throws made
	FreeThrowsMade int32 `json:"free_throws_made" parquet:"name=free_throws_made, type=INT32"`
	// FreeThrowAttempts is the number of free throw attempts
	FreeThrowAttempts int32 `json:"free_throw_attempts" parquet:"name=free_throw_attempts, type=INT32"`
	// FreeThrowPercentage is the free throw percentage
	FreeThrowPercentage float32 `json:"free_throw_percentage" parquet:"name=free_throw_percentage, type=FLOAT"`
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
	// GameScore is a metric used to evaluate how well a player performs in a single game
	GameScore float32 `json:"game_score" parquet:"name=game_score, type=FLOAT"`
	// PlusMinus is the plus-minus
	PlusMinus int32 `json:"plus_minus" parquet:"name=plus_minus, type=INT32"`
}
