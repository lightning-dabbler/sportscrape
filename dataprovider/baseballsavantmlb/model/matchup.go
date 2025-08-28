package model

import "time"

// Matchup represents the data model for MLB matchups scraped from baseballsavant.mlb.com
type Matchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 86833
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// Status is the string representation of the event status e.g. Final
	Status string `json:"status" parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
	// VenueID
	VenueID int64 `json:"venue_id" parquet:"name=venue_id, type=INT64"`
	// VenueName
	VenueName string `json:"venue_name" parquet:"name=venue_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamLeagueID
	HomeTeamLeagueID int64 `json:"home_team_league_id" parquet:"name=home_team_league_id, type=INT64"`
	// HomeTeamLeagueName
	HomeTeamLeagueName string `json:"home_team_league_name" parquet:"name=home_team_league_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamDivisionID
	HomeTeamDivisionID int64 `json:"home_team_division_id" parquet:"name=home_team_division_id, type=INT64"`
	// HomeTeamDivisionName
	HomeTeamDivisionName string `json:"home_team_division_name" parquet:"name=home_team_division_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamID is the home team's ID
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// HomeTeamAbbreviation is the abbreviation of the home team's name e.g. DET
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamName is the home team's full name e.g. Detroit Tigers
	HomeTeamName string `json:"home_team_name" parquet:"name=home_team_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeRecord is the home team's number of wins
	HomeWins int32 `json:"home_wins" parquet:"name=home_wins, type=INT32"`
	// HomeLosses is the home team's number of losses
	HomeLosses int32 `json:"home_losses" parquet:"name=home_losses, type=INT32"`
	// HomeScore is the home team's score at time of request
	HomeScore *int32 `json:"home_score" parquet:"name=home_score, type=INT32"`
	// HomeStartingPitcherID is the player id for the home team's starting pitcher (nillable because it may not be yet announced when fetching a scheduled event)
	HomeStartingPitcherID *int64 `json:"home_starting_pitcher_id" parquet:"name=home_starting_pitcher_id, type=INT64"`
	// HomeStartingPitcher the name of the home team's starting pitcher
	HomeStartingPitcher *string `json:"home_starting_pitcher" parquet:"name=home_starting_pitcher, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeStartingPitcherPitchHand - R for right, L for left
	HomeStartingPitcherPitchHand *string `json:"home_starting_pitcher_pitch_hand" parquet:"name=home_starting_pitcher_pitch_hand, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamLeagueID
	AwayTeamLeagueID int64 `json:"away_team_league_id" parquet:"name=away_team_league_id, type=INT64"`
	// AwayTeamLeagueName
	AwayTeamLeagueName string `json:"away_team_league_name" parquet:"name=away_team_league_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamDivisionID
	AwayTeamDivisionID int64 `json:"away_team_division_id" parquet:"name=away_team_division_id, type=INT64"`
	// AwayTeamDivisionName
	AwayTeamDivisionName string `json:"away_team_division_name" parquet:"name=away_team_division_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamID is the away team's ID e.g. 8
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeamAbbreviation is the abbreviation of the away team's name e.g. LAA
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamName is the away team's full name
	AwayTeamName string `json:"away_team_name" parquet:"name=away_team_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayRecord is the away team's number of wins
	AwayWins int32 `json:"away_wins" parquet:"name=away_wins, type=INT32"`
	// AwayLosses is the away team's number of losses
	AwayLosses int32 `json:"away_losses" parquet:"name=away_losses, type=INT32"`
	// AwayScore is the away team's score at time of request
	AwayScore *int32 `json:"away_score" parquet:"name=away_score, type=INT32"`
	// AwayStartingPitcherID is the player id for the away team's starting pitcher (nillable because it may not be yet announced when fetching a scheduled event)
	AwayStartingPitcherID *int64 `json:"away_starting_pitcher_id" parquet:"name=away_starting_pitcher_id, type=INT64"`
	// AwayStartingPitcher the name of the away team's starting pitcher
	AwayStartingPitcher *string `json:"away_starting_pitcher" parquet:"name=away_starting_pitcher, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayStartingPitcherPitchHand - R for right, L for left
	AwayStartingPitcherPitchHand *string `json:"away_starting_pitcher_pitch_hand" parquet:"name=away_starting_pitcher_pitch_hand, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Loser is an optional team id associated with the loser. Optional because some games could be live a loser will not be determined until the matchup is complete.
	Loser *int64 `json:"loser" parquet:"name=loser, type=INT64"`
	// GameType is an abbreviation that indicates the type of event taking place (e.g. "R", "S", "W", etc.)
	GameType string `json:"game_type" parquet:"name=game_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SeriesDescription acts as supplemental information for GameType (e.g. "Regular Season", "Spring Training", "World Series", etc.)
	SeriesDescription string `json:"series_description" parquet:"name=series_description, type=BYTE_ARRAY, convertedtype=UTF8"`
	// GamesInSeries is the number of games being played in this series of matchup
	GamesInSeries int32 `json:"games_in_series" parquet:"name=games_in_series, type=INT32"`
	// SeriesGameNumber is the series game number of the event
	SeriesGameNumber int32 `json:"series_game_number" parquet:"name=series_game_number, type=INT32"`
	// Season - e.g. 2025
	Season int32 `json:"season" parquet:"name=season, type=INT32"`
}
