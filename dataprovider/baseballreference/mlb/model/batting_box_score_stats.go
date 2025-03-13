package model

import "time"

// MLBBattingBoxScoreStats represents the data model for MLB batting box score stats scraped from baseball-reference.com
// https://www.mlb.com/glossary/standard-stats
type MLBBattingBoxScoreStats struct {
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
	// Position is the player's position
	Position string `json:"position"`
	// AtBat (AB) - https://www.mlb.com/glossary/standard-stats/at-bat
	AtBat int `json:"at_bat"`
	// Runs (R) - https://www.mlb.com/glossary/standard-stats/run
	Runs int `json:"runs"`
	// Hits (H) - https://www.mlb.com/glossary/standard-stats/hit
	Hits int `json:"hits"`
	// RunsBattedIn (RBI) - https://www.mlb.com/glossary/standard-stats/runs-batted-in
	RunsBattedIn int `json:"runs_batted_in"`
	// Walks (BB) - https://www.mlb.com/glossary/standard-stats/walk
	Walks int `json:"walks"`
	// Strikeouts (SO) - https://www.mlb.com/glossary/standard-stats/strikeout
	Strikeouts int `json:"strikeouts"`
	// PlateAppearances (PA) - https://www.mlb.com/glossary/standard-stats/plate-appearance
	PlateAppearances int `json:"plate_appearances"`
	// BattingAverage (BA) - https://www.mlb.com/glossary/standard-stats/batting-average
	BattingAverage *float32 `json:"batting_average"`
	// OnBasePercentage (OBP) - https://www.mlb.com/glossary/standard-stats/on-base-percentage
	OnBasePercentage *float32 `json:"on_base_percentage"`
	// SluggingPercentage (SLG) - https://www.mlb.com/glossary/standard-stats/slugging-percentage
	SluggingPercentage *float32 `json:"slugging_percentage"`
	// OnBasePlusSlugging (OPS) - https://www.mlb.com/glossary/standard-stats/on-base-plus-slugging
	OnBasePlusSlugging *float32 `json:"on_base_plus_slugging"`
	// Pitches Per Plate Appearance (P/PA) - https://www.mlb.com/glossary/advanced-stats/pitches-per-plate-appearance
	PitchesPerPlateAppearance *int `json:"pitches_per_plate_appearance"`
	// Strikes - includes both pitches in the zone and those swung at out of the zone.
	Strikes *int `json:"strikes"`
	// WinProbabilityAdded (WPA) - https://www.mlb.com/glossary/advanced-stats/win-probability-added
	WinProbabilityAdded *float32 `json:"win_probability_added"`
	// AverageLeverageIndex - the average pressure the pitcher or batter saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageLeverageIndex *float32 `json:"average_leverage_index"`
	// SumPositiveWinProbabilityAdded (WPA+) - Sum of positive events for batter
	SumPositiveWinProbabilityAdded *float32 `json:"sum_positive_win_probability_added"`
	// SumNegativeWinProbabilityAdded (WPA-) - Sum of negative events for batter
	SumNegativeWinProbabilityAdded *float32 `json:"sum_negative_win_probability_added"`
	// ChampionshipWinProbabilityAdded (cWPA) in percentage notation- https://www.reddit.com/r/baseball/comments/1agut0c/comment/kojk7c4
	ChampionshipWinProbabilityAdded *float32 `json:"championship_win_probability_added"`
	// AverageChampionshipLeverageIndex - the average pressure the pitcher or batter saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageChampionshipLeverageIndex *float32 `json:"average_championship_leverage_index"`
	// BaseOutRunsAdded (RE24) - Given the bases occupied/out situation, how many runs did the batter or baserunner add in the resulting play. Compared to average, so 0 is average, and above 0 is better than average
	BaseOutRunsAdded *float32 `json:"base_out_runs_added"`
	// Putout (PO) [Defense]- https://www.mlb.com/glossary/standard-stats/putout
	Putout int `json:"putout"`
	// Assist (A) [Defense] - https://www.mlb.com/glossary/standard-stats/assist
	Assist int `json:"assist"`
}
