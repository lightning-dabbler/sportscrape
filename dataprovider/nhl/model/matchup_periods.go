package model

import "time"

// MatchupPeriods - composite key: event_id, period
type MatchupPeriods struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 2026010045
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the scheduled start time of the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// GameType 1 = pre-season, 2 = regular season, 3 = post season (other values occur e.g. 19 for event 2024190001)
	GameType int32 `json:"game_type" parquet:"name=game_type, type=INT32"`
	// GameState e.g. FUT, FINAL, OFF (as of when the matchup was scraped)
	GameState string `json:"game_state" parquet:"name=game_state, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamID
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// HomeTeam e.g. Rangers
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamAbbreviation e.g. NYR
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamID
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeam e.g. Devils
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamAbbreviation e.g. NJD
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Period number e.g. 1, 2, 3, 4 (OT), 5 (SO)
	Period int32 `json:"period" parquet:"name=period, type=INT32"`
	// PeriodType e.g. REG, OT, SO
	PeriodType string `json:"period_type" parquet:"name=period_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamScore - away team goals in the period (linescore); for the SO period, 1 for the shootout winner and 0 otherwise
	AwayTeamScore *int32 `json:"away_team_score" parquet:"name=away_team_score, type=INT32"`
	// HomeTeamScore - home team goals in the period (linescore); for the SO period, 1 for the shootout winner and 0 otherwise
	HomeTeamScore *int32 `json:"home_team_score" parquet:"name=home_team_score, type=INT32"`
	// AwayTeamSOG - away team shots on goal in the period; 0 for the SO period
	AwayTeamSOG *int32 `json:"away_team_sog" parquet:"name=away_team_sog, type=INT32"`
	// HomeTeamSOG - home team shots on goal in the period; 0 for the SO period
	HomeTeamSOG *int32 `json:"home_team_sog" parquet:"name=home_team_sog, type=INT32"`
}
