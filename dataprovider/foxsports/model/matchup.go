package model

import "time"

// Matchup represents the data model for NBA, MLB, NFL, and NCAAB matchups scraped from foxsports.com
type Matchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 86833
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventStatus the numerical representation of the event status e.g. 3
	EventStatus int32 `json:"event_status" parquet:"name=event_status, type=INT32"`
	// StatusLine the string representation of the event status e.g. FINAL
	StatusLine string `json:"status_line" parquet:"name=status_line, type=BYTE_ARRAY"`
	// HomeTeamID is the home team's ID e.g. 21
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// HomeTeamAbbreviation is the abbreviation of the home team's name e.g. ATL
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY"`
	// HomeTeamNameLong is the home team's longer name but not the full name e.g. Braves
	HomeTeamNameLong string `json:"home_team_name_long" parquet:"name=home_team_name_long, type=BYTE_ARRAY"`
	// HomeTeamNameFull is the home team's full name e.g. Atlanta Braves
	HomeTeamNameFull string `json:"home_team_name_full" parquet:"name=home_team_name_full, type=BYTE_ARRAY"`
	// HomeRecord is the home team's record at time of request e.g. 69-37
	HomeRecord string `json:"home_record" parquet:"name=home_record, type=BYTE_ARRAY"`
	// HomeScore is the home team's score at time of request e.g. 12
	HomeScore int32 `json:"home_score" parquet:"name=home_score, type=INT32"`
	// HomeRank is an optional rank. The values are expected for some NCAAB teams
	HomeRank *int32 `json:"home_rank" parquet:"name=home_rank, type=INT32"`
	// AwayTeamID is the away team's ID e.g. 8
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeamAbbreviation is the abbreviation of the away team's name e.g. LAA
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY"`
	// AwayTeamNameLong is the away team's longer name but not the full name e.g. Angels
	AwayTeamNameLong string `json:"away_team_name_long" parquet:"name=away_team_name_long, type=BYTE_ARRAY"`
	// AwayTeamNameFull is the away team's full name e.g. Los Angeles Angels
	AwayTeamNameFull string `json:"away_team_name_full" parquet:"name=away_team_name_full, type=BYTE_ARRAY"`
	// AwayRecord is the away team's record at time of request e.g. 56-53
	AwayRecord string `json:"away_record" parquet:"name=away_record, type=BYTE_ARRAY"`
	// AwayScore is the away team's score at time of request e.g. 5
	AwayScore int32 `json:"away_score" parquet:"name=away_score, type=INT32"`
	// AwayRank is an optional rank. The values are expected for some NCAAB teams
	AwayRank *int32 `json:"away_rank" parquet:"name=away_rank, type=INT32"`
	// Loser is an optional team id associated with the loser. Optional because some games could be live a loser will not be determined until the matchup is complete.
	Loser *int64 `json:"loser" parquet:"name=loser, type=INT64"`
	// IsPlayoff indicates whether the matchup is a playoff game
	IsPlayoff bool `json:"is_playoff" parquet:"name=is_playoff, type=BOOLEAN"`
}
