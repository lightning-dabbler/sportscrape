package model

import "time"

// MLBBattingBoxScoreStats represents the data model for MLB batting box score stats scraped from baseball-reference.com
// https://www.mlb.com/glossary/standard-stats
// https://www.mlb.com/glossary/advanced-stats
type MLBPitchingBoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// Team is the player's team name
	Team string `json:"team"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent"`
	// PlayerID is the player's id extracted from the player's baseball-reference profile url
	PlayerID string `json:"player_id"`
	// Player is the player's name
	Player string `json:"player"`
	// PlayerLink is the link to the player's baseball-reference profile
	PlayerLink string `json:"player_link"`
	// PitchingOrder - The sequence of pitchers who played during the event per team
	PitchingOrder int `json:"pitching_order"`
	// InningsPitched (IP) - https://www.mlb.com/glossary/standard-stats/innings-pitched
	InningsPitched float32 `json:"innings_pitched"`
	// HitsAllowed (H) - https://www.mlb.com/glossary/standard-stats/hit
	HitsAllowed int `json:"hits_allowed"`
	// RunsAllowed (R) - https://www.mlb.com/glossary/standard-stats/run
	RunsAllowed int `json:"runs_allowed"`
	// EarnedRunsAllowed (ER) - https://www.mlb.com/glossary/standard-stats/earned-run
	EarnedRunsAllowed int `json:"earned_runs_allowed"`
	// Walks (BB) - https://www.mlb.com/glossary/standard-stats/walk
	Walks int `json:"walks"`
	// StrikeOuts (SO) - https://www.mlb.com/glossary/standard-stats/strikeout
	Strikeouts int `json:"strikeouts"`
	// HomeRunsAllowed (HR) - https://www.mlb.com/glossary/standard-stats/home-run
	HomeRunsAllowed int `json:"home_runs_allowed"`
	// EarnedRunAverage (ERA) - https://www.mlb.com/glossary/standard-stats/earned-run-average
	EarnedRunAverage float32 `json:"earned_run_average"`
	// BattersFaced (BF) - https://www.mlb.com/glossary/standard-stats/batters-faced
	BattersFaced int `json:"batters_faced"`
	// Pitches Per Plate Appearance (P/PA) - https://www.mlb.com/glossary/advanced-stats/pitches-per-plate-appearance
	PitchesPerPlateAppearance int `json:"pitches_per_plate_appearance"`
	// Strikes - includes both pitches in the zone and those swung at out of the zone.
	Strikes int `json:"strikes"`
	// StrikesByContact - Strikes from foul balls or balls put into play
	StrikesByContact int `json:"strikes_by_contact"`
	// StrikesSwinging - Strikes due to a swing and a miss
	StrikesSwinging int `json:"strikes_swinging"`
	// StrikesLooking - Strikes called by the umpire
	StrikesLooking int `json:"strikes_looking"`
	// GroundBalls - Includes bunts and all other ground balls.
	GroundBalls int `json:"ground_balls"`
	// FlyBalls - Includes Fly Balls, Line Drives, and Pop-Ups.
	FlyBalls int `json:"fly_balls"`
	// LineDrives - These are double-counted in Fly Balls as well.
	LineDrives int `json:"line_drives"`
	// UnknownBattedBallType - A ball in play for which we don’t know the type.
	UnknownBattedBallType int `json:"unknown_batted_ball_type"`
	// GameScore - https://www.mlb.com/glossary/advanced-stats/game-score
	GameScore *int `json:"game_score"`
	// InheritedRunners (IR) - https://www.mlb.com/glossary/standard-stats/inherited-runner
	InheritedRunners *int `json:"inherited_runners"`
	// InheritedScore refers to runs scored by inherited runners against a relief pitcher,
	// meaning runners on base when a relief pitcher enters the game. These runs are not charged
	// to the relief pitcher's ERA but are tracked in statistics like Inherited Runs Allowed (IR-A)
	// and Inherited Runs Allowed Percentage (IR-A%).
	InheritedScore *int `json:"inherited_score"`
	// WinProbabilityAdded (WPA) - https://www.mlb.com/glossary/advanced-stats/win-probability-added
	WinProbabilityAdded float32 `json:"win_probability_added"`
	// AverageLeverageIndex - the average pressure the pitcher saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageLeverageIndex float32 `json:"average_leverage_index"`
	// ChampionshipWinProbabilityAdded (cWPA) in percentage notation- https://www.reddit.com/r/baseball/comments/1agut0c/comment/kojk7c4
	ChampionshipWinProbabilityAdded float32 `json:"championship_win_probability_added"`
	// AverageChampionshipLeverageIndex - the average pressure the pitcher saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageChampionshipLeverageIndex float32 `json:"average_championship_leverage_index"`
	// BaseOutRunsSaved (RE24) - Given the bases occupied/out situation, how many runs did the pitcher save in the resulting play. Compared to average, so 0 is average, and above 0 is better than average
	BaseOutRunsSaved float32 `json:"base_out_runs_added"`
}
