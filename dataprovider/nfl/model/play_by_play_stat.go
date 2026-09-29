package model

import "time"

// PlayByPlayStat is a stat credited to a player (or only a team) on a play - composite key: event_id, play_stat_id
type PlayByPlayStat struct {
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
	// PlayID is the play the stat belongs to (joins with PlayByPlay.PlayID) e.g. 71
	PlayID int64 `json:"play_id" parquet:"name=play_id, type=INT64"`
	// PlayStatID is unique per stat e.g. 10330059-8430-0400-0410-00142cd91712
	PlayStatID string `json:"play_stat_id" parquet:"name=play_stat_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Quarter of the play 1-4, 5 = OT1, 6 = OT2, ...
	Quarter int32 `json:"quarter" parquet:"name=quarter, type=INT32"`
	// StatType is the stat's code e.g. 10
	StatType int32 `json:"stat_type" parquet:"name=stat_type, type=INT32"`
	// StatTypeDescription e.g. rushing yards; nil for codes that haven't been verified (see dataprovider/nfl/stat_type.go)
	StatTypeDescription *string `json:"stat_type_description" parquet:"name=stat_type_description, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Yards e.g. 7; nil for stats without yards (e.g. solo tackle, target)
	Yards *int32 `json:"yards" parquet:"name=yards, type=INT32"`
	// TeamID of the team credited with the stat
	TeamID string `json:"team_id" parquet:"name=team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Team of the team credited with the stat e.g. Dallas Cowboys (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID is the player's GSIS ID e.g. 00-0033077; nil for team stats (e.g. first down rushing, penalty on a team)
	PlayerID *string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PersonID e.g. 32005052-4528-5723-d1b2-96e92ebc1241; nil for team stats
	PersonID *string `json:"person_id" parquet:"name=person_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Player is the player's full name e.g. Dak Prescott (falls back to PlayerShortName when the full name is unavailable); nil for team stats
	Player *string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerShortName is the abbreviated player name e.g. D.Prescott; nil for team stats
	PlayerShortName *string `json:"player_short_name" parquet:"name=player_short_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// JerseyNumber e.g. 04; nil for team stats
	JerseyNumber *string `json:"jersey_number" parquet:"name=jersey_number, type=BYTE_ARRAY, convertedtype=UTF8"`
}
