package model

import "time"

// DefenseBoxScore - composite key: event_id, player_id
type DefenseBoxScore struct {
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
	// Tackles - solo tackles
	Tackles float32 `json:"tackles" parquet:"name=tackles, type=FLOAT"`
	// TacklesAssists - assisted tackles
	TacklesAssists int32 `json:"tackles_assists" parquet:"name=tackles_assists, type=INT32"`
	// TacklesCombined - solo + assisted tackles
	TacklesCombined float32 `json:"tackles_combined" parquet:"name=tackles_combined, type=FLOAT"`
	// TacklesForLoss
	TacklesForLoss float32 `json:"tackles_for_loss" parquet:"name=tackles_for_loss, type=FLOAT"`
	// TacklesForLossYards
	TacklesForLossYards float32 `json:"tackles_for_loss_yards" parquet:"name=tackles_for_loss_yards, type=FLOAT"`
	// Sacks
	Sacks float32 `json:"sacks" parquet:"name=sacks, type=FLOAT"`
	// SackYards
	SackYards float32 `json:"sack_yards" parquet:"name=sack_yards, type=FLOAT"`
	// QuarterbackHits
	QuarterbackHits int32 `json:"quarterback_hits" parquet:"name=quarterback_hits, type=INT32"`
	// PassesDefended
	PassesDefended int32 `json:"passes_defended" parquet:"name=passes_defended, type=INT32"`
	// Interceptions
	Interceptions int32 `json:"interceptions" parquet:"name=interceptions, type=INT32"`
	// FumblesForced
	FumblesForced int32 `json:"fumbles_forced" parquet:"name=fumbles_forced, type=INT32"`
	// FumblesRecovered
	FumblesRecovered int32 `json:"fumbles_recovered" parquet:"name=fumbles_recovered, type=INT32"`
	// Safeties
	Safeties int32 `json:"safeties" parquet:"name=safeties, type=INT32"`
	// SpecialTeamsTackles
	SpecialTeamsTackles float32 `json:"special_teams_tackles" parquet:"name=special_teams_tackles, type=FLOAT"`
	// SpecialTeamsTacklesAssists
	SpecialTeamsTacklesAssists int32 `json:"special_teams_tackles_assists" parquet:"name=special_teams_tackles_assists, type=INT32"`
	// SpecialTeamsBlocks
	SpecialTeamsBlocks int32 `json:"special_teams_blocks" parquet:"name=special_teams_blocks, type=INT32"`
	// SpecialTeamsFumblesForced
	SpecialTeamsFumblesForced int32 `json:"special_teams_fumbles_forced" parquet:"name=special_teams_fumbles_forced, type=INT32"`
	// SpecialTeamsFumblesRecovered
	SpecialTeamsFumblesRecovered int32 `json:"special_teams_fumbles_recovered" parquet:"name=special_teams_fumbles_recovered, type=INT32"`
	// MiscellaneousTackles
	MiscellaneousTackles float32 `json:"miscellaneous_tackles" parquet:"name=miscellaneous_tackles, type=FLOAT"`
	// MiscellaneousTacklesAssists
	MiscellaneousTacklesAssists int32 `json:"miscellaneous_tackles_assists" parquet:"name=miscellaneous_tackles_assists, type=INT32"`
	// MiscellaneousFumblesForced
	MiscellaneousFumblesForced int32 `json:"miscellaneous_fumbles_forced" parquet:"name=miscellaneous_fumbles_forced, type=INT32"`
	// MiscellaneousFumblesRecovered
	MiscellaneousFumblesRecovered int32 `json:"miscellaneous_fumbles_recovered" parquet:"name=miscellaneous_fumbles_recovered, type=INT32"`
	// TwoPointAttempts - defensive two-point conversion return attempts
	TwoPointAttempts int32 `json:"two_point_attempts" parquet:"name=two_point_attempts, type=INT32"`
	// TwoPointSuccesses
	TwoPointSuccesses int32 `json:"two_point_successes" parquet:"name=two_point_successes, type=INT32"`
}
