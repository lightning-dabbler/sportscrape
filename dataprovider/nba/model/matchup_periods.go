package model

import "time"

// MatchupPeriods - scores per period per game. Composite key: event_id, period
type MatchupPeriods struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 0022500249
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// HomeTeamID
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// HomeTeam
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamAbbreviation
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamID
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeam
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamAbbreviation
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Period is the quarter and/or overtime number, 1-4 represents quarter number and >4 represents Overtime
	Period int32 `json:"period" parquet:"name=period, type=INT32"`
	// AwayScore with respect to period and away team
	AwayScore int32 `json:"away_score" parquet:"name=away_score, type=INT32"`
	// HomeScore with respect to period and home team
	HomeScore int32 `json:"home_score" parquet:"name=home_score, type=INT32"`
	// SeasonType (e.g. "Regular Season")
	SeasonType string `json:"season_type" parquet:"name=season_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SeasonYear (e.g 2025-26)
	SeasonYear string `json:"season_year" parquet:"name=season_year, type=BYTE_ARRAY, convertedtype=UTF8"`
	// LeagueID
	LeagueID string `json:"league_id" parquet:"name=league_id, type=BYTE_ARRAY, convertedtype=UTF8"`
}
