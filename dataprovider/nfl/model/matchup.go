package model

import "time"

// Matchup represents an NFL matchup scraped from https://api.nfl.com/football/v2/experience/weekly-game-details
type Matchup struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. f5908b6d-311e-11f0-b670-ae1250fadad1
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// GSISGameID is the matchup's GSIS ID e.g. 59843
	GSISGameID *string `json:"gsis_game_id" parquet:"name=gsis_game_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Slug is the nfl.com game page slug e.g. cowboys-at-eagles-2025-reg-1 (https://www.nfl.com/games/cowboys-at-eagles-2025-reg-1)
	Slug *string `json:"slug" parquet:"name=slug, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventTime is the scheduled start time of the matchup. A 09:00 UTC placeholder on the game date for the 2012 season and earlier
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// StartTime is the actual start time of the game (e.g. a few minutes after EventTime, later after a delay); nil before kickoff and for older seasons (e.g. 2001)
	StartTime *time.Time `json:"start_time"`
	// StartTimeParquet is the actual start time of the game (in milliseconds)
	StartTimeParquet *int64 `json:"-" parquet:"name=start_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// GameDate is the scheduled start date in America/New_York e.g. 2025-09-04
	GameDate string `json:"game_date" parquet:"name=game_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Season e.g. 2025 (a season's post-season games are played in the following year)
	Season int32 `json:"season" parquet:"name=season, type=INT32"`
	// SeasonType e.g. PRE, REG, POST
	SeasonType string `json:"season_type" parquet:"name=season_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Week of the season type e.g. 1 (PRE week 0 is the Hall of Fame game)
	Week int32 `json:"week" parquet:"name=week, type=INT32"`
	// WeekType e.g. HOF, PRE, REG, WC, DIV, CONF, SB
	WeekType string `json:"week_type" parquet:"name=week_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// GameType e.g. UNSPECIFIED (pre-season and regular season), AFC_WC, NFC_DIV, AFC_CONF, NFC_AFC_SB
	GameType string `json:"game_type" parquet:"name=game_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Category e.g. TNF, SNF, MNF; nil otherwise
	Category *string `json:"category" parquet:"name=category, type=BYTE_ARRAY, convertedtype=UTF8"`
	// NeutralSite
	NeutralSite bool `json:"neutral_site" parquet:"name=neutral_site, type=BOOLEAN"`
	// International
	International bool `json:"international" parquet:"name=international, type=BOOLEAN"`
	// Venue e.g. Lincoln Financial Field
	Venue *string `json:"venue" parquet:"name=venue, type=BYTE_ARRAY, convertedtype=UTF8"`
	// VenueCity e.g. Philadelphia
	VenueCity *string `json:"venue_city" parquet:"name=venue_city, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Phase e.g. PREGAME, FINAL, FINAL_OVERTIME (as of when the matchup was scraped); nil when the game is further out
	Phase *string `json:"phase" parquet:"name=phase, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Quarter e.g. Q1, END_OF_GAME (as of when the matchup was scraped); nil when the game is further out
	Quarter *string `json:"quarter" parquet:"name=quarter, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Clock - game clock e.g. 15:00 (as of when the matchup was scraped); nil when the game is further out
	Clock *string `json:"clock" parquet:"name=clock, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Attendance e.g. 69879; nil before the game
	Attendance *int32 `json:"attendance" parquet:"name=attendance, type=INT32"`
	// Weather e.g. Rain Temp: 75° F, Humidity: 66%, Wind: S 11 mph; nil before the game
	Weather *string `json:"weather" parquet:"name=weather, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamID e.g. 10403700-b939-3cbd-3d16-24d4d6742fa2
	HomeTeamID string `json:"home_team_id" parquet:"name=home_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeam e.g. Philadelphia Eagles (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamAbbreviation e.g. PHI, as of the matchup's season (e.g. SD for the 2012 Chargers)
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamScore - nil when the game is further out
	HomeTeamScore *int32 `json:"home_team_score" parquet:"name=home_team_score, type=INT32"`
	// AwayTeamID e.g. 10401200-a308-98ca-ad5f-95df2fefea68
	AwayTeamID string `json:"away_team_id" parquet:"name=away_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeam e.g. Dallas Cowboys (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamAbbreviation e.g. DAL, as of the matchup's season (e.g. SD for the 2012 Chargers)
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamScore - nil when the game is further out
	AwayTeamScore *int32 `json:"away_team_score" parquet:"name=away_team_score, type=INT32"`
}
