package model

import "time"

// FumblesBoxScore - composite key: event_id, player_id
type FumblesBoxScore struct {
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
	// Fumbles
	Fumbles int32 `json:"fumbles" parquet:"name=fumbles, type=INT32"`
	// Lost
	Lost int32 `json:"lost" parquet:"name=lost, type=INT32"`
	// Forced
	Forced int32 `json:"forced" parquet:"name=forced, type=INT32"`
	// OutOfBounds
	OutOfBounds int32 `json:"out_of_bounds" parquet:"name=out_of_bounds, type=INT32"`
	// OwnRecoveries
	OwnRecoveries int32 `json:"own_recoveries" parquet:"name=own_recoveries, type=INT32"`
	// OwnRecoveryYards
	OwnRecoveryYards int32 `json:"own_recovery_yards" parquet:"name=own_recovery_yards, type=INT32"`
	// OwnRecoveryTouchdowns
	OwnRecoveryTouchdowns int32 `json:"own_recovery_touchdowns" parquet:"name=own_recovery_touchdowns, type=INT32"`
	// OpponentRecoveries
	OpponentRecoveries int32 `json:"opponent_recoveries" parquet:"name=opponent_recoveries, type=INT32"`
	// OpponentRecoveryYards
	OpponentRecoveryYards int32 `json:"opponent_recovery_yards" parquet:"name=opponent_recovery_yards, type=INT32"`
	// OpponentRecoveryTouchdowns
	OpponentRecoveryTouchdowns int32 `json:"opponent_recovery_touchdowns" parquet:"name=opponent_recovery_touchdowns, type=INT32"`
	// RecoveredInEndZoneForTouchdown
	RecoveredInEndZoneForTouchdown int32 `json:"recovered_in_end_zone_for_touchdown" parquet:"name=recovered_in_end_zone_for_touchdown, type=INT32"`
}
