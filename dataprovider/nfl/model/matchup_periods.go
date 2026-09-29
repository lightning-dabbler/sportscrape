package model

import "time"

// MatchupPeriods - composite key: event_id, period
type MatchupPeriods struct {
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
	// SeasonType e.g. PRE, REG, POST
	SeasonType string `json:"season_type" parquet:"name=season_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Phase e.g. FINAL, FINAL_OVERTIME (as of when the periods were scraped)
	Phase string `json:"phase" parquet:"name=phase, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamID
	HomeTeamID string `json:"home_team_id" parquet:"name=home_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeam e.g. Philadelphia Eagles (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamAbbreviation e.g. PHI, as of the matchup's season (e.g. SD for the 2012 Chargers)
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamID
	AwayTeamID string `json:"away_team_id" parquet:"name=away_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeam e.g. Dallas Cowboys (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamAbbreviation e.g. DAL, as of the matchup's season (e.g. SD for the 2012 Chargers)
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Period number 1-4, 5 = OT1, 6 = OT2, ...
	Period int32 `json:"period" parquet:"name=period, type=INT32"`
	// PeriodType REG (quarters) or OT (overtime periods)
	PeriodType string `json:"period_type" parquet:"name=period_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamScore - away team points in the period
	AwayTeamScore int32 `json:"away_team_score" parquet:"name=away_team_score, type=INT32"`
	// HomeTeamScore - home team points in the period
	HomeTeamScore int32 `json:"home_team_score" parquet:"name=home_team_score, type=INT32"`
}
