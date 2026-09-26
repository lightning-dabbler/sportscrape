package model

import "time"

// Matchup represents an NHL matchup scraped from https://api-web.nhle.com/v1/score/{date}
type Matchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 2024020250
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the scheduled start time of the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// GameDate is the local game date e.g. 2024-11-12
	GameDate string `json:"game_date" parquet:"name=game_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Season e.g. 20242025
	Season int64 `json:"season" parquet:"name=season, type=INT64"`
	// GameType 1 = pre-season, 2 = regular season, 3 = post season (other values occur e.g. 19 for event 2024190001)
	GameType int32 `json:"game_type" parquet:"name=game_type, type=INT32"`
	// GameState e.g. FUT, FINAL, OFF
	GameState string `json:"game_state" parquet:"name=game_state, type=BYTE_ARRAY, convertedtype=UTF8"`
	// GameScheduleState e.g. OK
	GameScheduleState string `json:"game_schedule_state" parquet:"name=game_schedule_state, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Venue e.g. Rogers Arena
	Venue string `json:"venue" parquet:"name=venue, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamID
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// HomeTeam e.g. Canucks
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamAbbreviation e.g. VAN
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamScore - nil before the game starts
	HomeTeamScore *int32 `json:"home_team_score" parquet:"name=home_team_score, type=INT32"`
	// HomeTeamSOG - home team shots on goal; nil before the game starts
	HomeTeamSOG *int32 `json:"home_team_sog" parquet:"name=home_team_sog, type=INT32"`
	// AwayTeamID
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeam e.g. Flames
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamAbbreviation e.g. CGY
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamScore - nil before the game starts
	AwayTeamScore *int32 `json:"away_team_score" parquet:"name=away_team_score, type=INT32"`
	// AwayTeamSOG - away team shots on goal; nil before the game starts
	AwayTeamSOG *int32 `json:"away_team_sog" parquet:"name=away_team_sog, type=INT32"`
	// Period - current (or last) period number; nil before the game starts
	Period *int32 `json:"period" parquet:"name=period, type=INT32"`
	// LastPeriodType e.g. REG, OT, SO; nil before the game ends
	LastPeriodType *string `json:"last_period_type" parquet:"name=last_period_type, type=BYTE_ARRAY, convertedtype=UTF8"`
}
