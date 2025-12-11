package model

import "time"

// Matchup
type Matchup struct {
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
	// EventStatus the numerical representation of the event status e.g. 3 (1=pregame, 2=in progress, 3=final)
	EventStatus int32 `json:"event_status" parquet:"name=event_status, type=INT32"`
	// EventStatusText (e.g. Final, Final/OT2, etc.)
	EventStatusText string `json:"event_status_text" parquet:"name=event_status_text, type=BYTE_ARRAY, convertedtype=UTF8"`
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
	// AwayTeamScore - total away team game score
	AwayTeamScore int32 `json:"away_team_score" parquet:"name=away_team_score, type=INT32"`
	// HomeTeamScore - total home team game score
	HomeTeamScore int32 `json:"home_team_score" parquet:"name=home_team_score, type=INT32"`
	// AwayTeamWins
	AwayTeamWins int32 `json:"away_team_wins" parquet:"name=away_team_wins, type=INT32"`
	// HomeTeamWins
	HomeTeamWins int32 `json:"home_team_wins" parquet:"name=home_team_wins, type=INT32"`
	// AwayTeamLosses
	AwayTeamLosses int32 `json:"away_team_losses" parquet:"name=away_team_losses, type=INT32"`
	// HomeTeamLosses
	HomeTeamLosses int32 `json:"home_team_losses" parquet:"name=home_team_losses, type=INT32"`
	// ShareURL (e.g. https://www.nba.com/game/gsw-vs-orl-0022500249)
	ShareURL string `json:"share_url" parquet:"name=share_url, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SeasonType (e.g. "Regular Season")
	SeasonType string `json:"season_type" parquet:"name=season_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SeasonYear (e.g 2025-26)
	SeasonYear string `json:"season_year" parquet:"name=season_year, type=BYTE_ARRAY, convertedtype=UTF8"`
	// LeagueID
	LeagueID string `json:"league_id" parquet:"name=league_id, type=BYTE_ARRAY, convertedtype=UTF8"`
}
