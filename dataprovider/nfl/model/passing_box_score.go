package model

import "time"

// PassingBoxScore - composite key: event_id, player_id
type PassingBoxScore struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. f5908b6d-311e-11f0-b670-ae1250fadad1
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventTime is the scheduled start time of the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// TeamID e.g. 10401200-a308-98ca-ad5f-95df2fefea68
	TeamID string `json:"team_id" parquet:"name=team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Team e.g. Dallas Cowboys (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID e.g. 10403700-b939-3cbd-3d16-24d4d6742fa2
	OpponentID string `json:"opponent_id" parquet:"name=opponent_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Opponent e.g. Philadelphia Eagles (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID is the player's GSIS ID e.g. 00-0033077
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PersonID e.g. 32005052-4528-5723-d1b2-96e92ebc1241
	PersonID string `json:"person_id" parquet:"name=person_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Player is the player's full name e.g. Dak Prescott (falls back to PlayerShortName when the full name is unavailable)
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerShortName is the abbreviated player name e.g. D.Prescott
	PlayerShortName string `json:"player_short_name" parquet:"name=player_short_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// JerseyNumber e.g. 04
	JerseyNumber string `json:"jersey_number" parquet:"name=jersey_number, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Completions
	Completions int32 `json:"completions" parquet:"name=completions, type=INT32"`
	// Attempts
	Attempts int32 `json:"attempts" parquet:"name=attempts, type=INT32"`
	// Yards
	Yards int32 `json:"yards" parquet:"name=yards, type=INT32"`
	// CompletionPercent - e.g. 61.8
	CompletionPercent float32 `json:"completion_percent" parquet:"name=completion_percent, type=FLOAT"`
	// YardsAverage - passing yards per attempt e.g. 5.53
	YardsAverage float32 `json:"yards_average" parquet:"name=yards_average, type=FLOAT"`
	// Touchdowns
	Touchdowns int32 `json:"touchdowns" parquet:"name=touchdowns, type=INT32"`
	// Interceptions - interceptions thrown
	Interceptions int32 `json:"interceptions" parquet:"name=interceptions, type=INT32"`
	// Long - longest completion
	Long int32 `json:"long" parquet:"name=long, type=INT32"`
	// LongestTouchdown - longest touchdown pass
	LongestTouchdown int32 `json:"longest_touchdown" parquet:"name=longest_touchdown, type=INT32"`
	// TimesSacked
	TimesSacked int32 `json:"times_sacked" parquet:"name=times_sacked, type=INT32"`
	// SackYardsLost
	SackYardsLost float32 `json:"sack_yards_lost" parquet:"name=sack_yards_lost, type=FLOAT"`
	// Rating - passer rating e.g. 76.6
	Rating float32 `json:"rating" parquet:"name=rating, type=FLOAT"`
	// TwoPointAttempts
	TwoPointAttempts int32 `json:"two_point_attempts" parquet:"name=two_point_attempts, type=INT32"`
	// TwoPointSuccesses
	TwoPointSuccesses int32 `json:"two_point_successes" parquet:"name=two_point_successes, type=INT32"`
}
