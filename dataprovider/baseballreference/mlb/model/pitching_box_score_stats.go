package model

import "time"

// MLBBattingBoxScoreStats represents the data model for MLB batting box score stats scraped from baseball-reference.com
// https://www.mlb.com/glossary/standard-stats
// https://www.mlb.com/glossary/advanced-stats
type MLBPitchingBoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// EventDateParquet is the timestamp associated with a given event (in days)
	EventDateParquet int32 `parquet:"name=event_date, type=INT32, convertedtype=DATE, logicaltype=DATE"`
	// Team is the player's team name
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY"`
	// PlayerID is the player's id extracted from the player's baseball-reference profile url
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY"`
	// Player is the player's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY"`
	// PlayerLink is the link to the player's baseball-reference profile
	PlayerLink string `json:"player_link" parquet:"name=player_link, type=BYTE_ARRAY"`
	// PitchingOrder - The sequence of pitchers who played during the event per team
	PitchingOrder int32 `json:"pitching_order" parquet:"name=pitching_order, type=INT32"`
	// InningsPitched (IP) - https://www.mlb.com/glossary/standard-stats/innings-pitched
	InningsPitched float32 `json:"innings_pitched" parquet:"name=innings_pitched, type=FLOAT"`
	// HitsAllowed (H) - https://www.mlb.com/glossary/standard-stats/hit
	HitsAllowed int32 `json:"hits_allowed" parquet:"name=hits_allowed, type=INT32"`
	// RunsAllowed (R) - https://www.mlb.com/glossary/standard-stats/run
	RunsAllowed int32 `json:"runs_allowed" parquet:"name=runs_allowed, type=INT32"`
	// EarnedRunsAllowed (ER) - https://www.mlb.com/glossary/standard-stats/earned-run
	EarnedRunsAllowed int32 `json:"earned_runs_allowed" parquet:"name=earned_runs_allowed, type=INT32"`
	// Walks (BB) - https://www.mlb.com/glossary/standard-stats/walk
	Walks int32 `json:"walks" parquet:"name=walks, type=INT32"`
	// StrikeOuts (SO) - https://www.mlb.com/glossary/standard-stats/strikeout
	Strikeouts int32 `json:"strikeouts" parquet:"name=strikeouts, type=INT32"`
	// HomeRunsAllowed (HR) - https://www.mlb.com/glossary/standard-stats/home-run
	HomeRunsAllowed int32 `json:"home_runs_allowed" parquet:"name=home_runs_allowed, type=INT32"`
	// EarnedRunAverage (ERA) - https://www.mlb.com/glossary/standard-stats/earned-run-average
	EarnedRunAverage float32 `json:"earned_run_average" parquet:"name=earned_run_average, type=FLOAT"`
	// BattersFaced (BF) - https://www.mlb.com/glossary/standard-stats/batters-faced
	BattersFaced int32 `json:"batters_faced" parquet:"name=batters_faced, type=INT32"`
	// Pitches Per Plate Appearance (P/PA) - https://www.mlb.com/glossary/advanced-stats/pitches-per-plate-appearance
	PitchesPerPlateAppearance int32 `json:"pitches_per_plate_appearance" parquet:"name=pitches_per_plate_appearance, type=INT32"`
	// Strikes - includes both pitches in the zone and those swung at out of the zone.
	Strikes int32 `json:"strikes" parquet:"name=strikes, type=INT32"`
	// StrikesByContact - Strikes from foul balls or balls put into play
	StrikesByContact int32 `json:"strikes_by_contact" parquet:"name=strikes_by_contact, type=INT32"`
	// StrikesSwinging - Strikes due to a swing and a miss
	StrikesSwinging int32 `json:"strikes_swinging" parquet:"name=strikes_swinging, type=INT32"`
	// StrikesLooking - Strikes called by the umpire
	StrikesLooking int32 `json:"strikes_looking" parquet:"name=strikes_looking, type=INT32"`
	// GroundBalls - Includes bunts and all other ground balls.
	GroundBalls int32 `json:"ground_balls" parquet:"name=ground_balls, type=INT32"`
	// FlyBalls - Includes Fly Balls, Line Drives, and Pop-Ups.
	FlyBalls int32 `json:"fly_balls" parquet:"name=fly_balls, type=INT32"`
	// LineDrives - These are double-counted in Fly Balls as well.
	LineDrives int32 `json:"line_drives" parquet:"name=line_drives, type=INT32"`
	// UnknownBattedBallType - A ball in play for which we don’t know the type.
	UnknownBattedBallType int32 `json:"unknown_batted_ball_type" parquet:"name=unknown_batted_ball_type, type=INT32"`
	// GameScore - https://www.mlb.com/glossary/advanced-stats/game-score
	GameScore *int32 `json:"game_score" parquet:"name=game_score, type=INT32"`
	// InheritedRunners (IR) - https://www.mlb.com/glossary/standard-stats/inherited-runner
	InheritedRunners *int32 `json:"inherited_runners" parquet:"name=inherited_runners, type=INT32"`
	// InheritedScore refers to runs scored by inherited runners against a relief pitcher,
	// meaning runners on base when a relief pitcher enters the game. These runs are not charged
	// to the relief pitcher's ERA but are tracked in statistics like Inherited Runs Allowed (IR-A)
	// and Inherited Runs Allowed Percentage (IR-A%).
	InheritedScore *int32 `json:"inherited_score" parquet:"name=inherited_score, type=INT32"`
	// WinProbabilityAdded (WPA) - https://www.mlb.com/glossary/advanced-stats/win-probability-added
	WinProbabilityAdded float32 `json:"win_probability_added" parquet:"name=win_probability_added, type=FLOAT"`
	// AverageLeverageIndex - the average pressure the pitcher saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageLeverageIndex float32 `json:"average_leverage_index" parquet:"name=average_leverage_index, type=FLOAT"`
	// ChampionshipWinProbabilityAdded (cWPA) in percentage notation- https://www.reddit.com/r/baseball/comments/1agut0c/comment/kojk7c4
	ChampionshipWinProbabilityAdded float32 `json:"championship_win_probability_added" parquet:"name=championship_win_probability_added, type=FLOAT"`
	// AverageChampionshipLeverageIndex - the average pressure the pitcher saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageChampionshipLeverageIndex float32 `json:"average_championship_leverage_index" parquet:"name=average_championship_leverage_index, type=FLOAT"`
	// BaseOutRunsSaved (RE24) - Given the bases occupied/out situation, how many runs did the pitcher save in the resulting play. Compared to average, so 0 is average, and above 0 is better than average
	BaseOutRunsSaved float32 `json:"base_out_runs_added" parquet:"name=base_out_runs_added, type=FLOAT"`
}
