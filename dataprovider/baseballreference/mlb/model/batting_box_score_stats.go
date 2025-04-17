package model

import "time"

// MLBBattingBoxScoreStats represents the data model for MLB batting box score stats scraped from baseball-reference.com
// https://www.mlb.com/glossary/standard-stats
// https://www.mlb.com/glossary/advanced-stats
type MLBBattingBoxScoreStats struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is the parsed event id from the box score link of the matchup
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY"`
	// EventDate is the timestamp associated with a given event
	EventDate time.Time `json:"event_date"`
	// EventDateParquet is the timestamp associated with a given event (in days)
	EventDateParquet int32 `json:"-" parquet:"name=event_date, type=INT32, convertedtype=DATE, logicaltype=DATE"`
	// Team is the player's team name e.g. Atlanta Braves
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY"`
	// PlayerID is the player's id
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY"`
	// Player is the player's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY"`
	// PlayerLink is the link to the player's baseball-reference profile
	PlayerLink string `json:"player_link" parquet:"name=player_link, type=BYTE_ARRAY"`
	// Position is the player's position
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY"`
	// AtBat (AB) - https://www.mlb.com/glossary/standard-stats/at-bat
	AtBat int32 `json:"at_bat" parquet:"name=at_bat, type=INT32"`
	// Runs (R) - https://www.mlb.com/glossary/standard-stats/run
	Runs int32 `json:"runs" parquet:"name=runs, type=INT32"`
	// Hits (H) - https://www.mlb.com/glossary/standard-stats/hit
	Hits int32 `json:"hits" parquet:"name=hits, type=INT32"`
	// RunsBattedIn (RBI) - https://www.mlb.com/glossary/standard-stats/runs-batted-in
	RunsBattedIn int32 `json:"runs_batted_in" parquet:"name=runs_batted_in, type=INT32"`
	// Walks (BB) - https://www.mlb.com/glossary/standard-stats/walk
	Walks int32 `json:"walks" parquet:"name=walks, type=INT32"`
	// Strikeouts (SO) - https://www.mlb.com/glossary/standard-stats/strikeout
	Strikeouts int32 `json:"strikeouts" parquet:"name=strikeouts, type=INT32"`
	// PlateAppearances (PA) - https://www.mlb.com/glossary/standard-stats/plate-appearance
	PlateAppearances int32 `json:"plate_appearances" parquet:"name=plate_appearances, type=INT32"`
	// BattingAverage (BA) - https://www.mlb.com/glossary/standard-stats/batting-average
	BattingAverage *float32 `json:"batting_average" parquet:"name=batting_average, type=FLOAT"`
	// OnBasePercentage (OBP) - https://www.mlb.com/glossary/standard-stats/on-base-percentage
	OnBasePercentage *float32 `json:"on_base_percentage" parquet:"name=on_base_percentage, type=FLOAT"`
	// SluggingPercentage (SLG) - https://www.mlb.com/glossary/standard-stats/slugging-percentage
	SluggingPercentage *float32 `json:"slugging_percentage" parquet:"name=slugging_percentage, type=FLOAT"`
	// OnBasePlusSlugging (OPS) - https://www.mlb.com/glossary/standard-stats/on-base-plus-slugging
	OnBasePlusSlugging *float32 `json:"on_base_plus_slugging" parquet:"name=on_base_plus_slugging, type=FLOAT"`
	// Pitches Per Plate Appearance (P/PA) - https://www.mlb.com/glossary/advanced-stats/pitches-per-plate-appearance
	PitchesPerPlateAppearance *int32 `json:"pitches_per_plate_appearance" parquet:"name=pitches_per_plate_appearance, type=INT32"`
	// Strikes - includes both pitches in the zone and those swung at out of the zone.
	Strikes *int32 `json:"strikes" parquet:"name=strikes, type=INT32"`
	// WinProbabilityAdded (WPA) - https://www.mlb.com/glossary/advanced-stats/win-probability-added
	WinProbabilityAdded *float32 `json:"win_probability_added" parquet:"name=win_probability_added, type=FLOAT"`
	// AverageLeverageIndex - the average pressure the batter saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageLeverageIndex *float32 `json:"average_leverage_index" parquet:"name=average_leverage_index, type=FLOAT"`
	// SumPositiveWinProbabilityAdded (WPA+) - Sum of positive events for batter
	SumPositiveWinProbabilityAdded *float32 `json:"sum_positive_win_probability_added" parquet:"name=sum_positive_win_probability_added, type=FLOAT"`
	// SumNegativeWinProbabilityAdded (WPA-) - Sum of negative events for batter
	SumNegativeWinProbabilityAdded *float32 `json:"sum_negative_win_probability_added" parquet:"name=sum_negative_win_probability_added, type=FLOAT"`
	// ChampionshipWinProbabilityAdded (cWPA) in percentage notation- https://www.reddit.com/r/baseball/comments/1agut0c/comment/kojk7c4
	ChampionshipWinProbabilityAdded *float32 `json:"championship_win_probability_added" parquet:"name=championship_win_probability_added, type=FLOAT"`
	// AverageChampionshipLeverageIndex - the average pressure the batter saw in this game or season. 1.0 is average pressure, below 1.0 is low pressure and above 1.0 is high pressure. https://www.mlb.com/glossary/advanced-stats/leverage-index
	AverageChampionshipLeverageIndex *float32 `json:"average_championship_leverage_index" parquet:"name=average_championship_leverage_index, type=FLOAT"`
	// BaseOutRunsAdded (RE24) - Given the bases occupied/out situation, how many runs did the batter or baserunner add in the resulting play. Compared to average, so 0 is average, and above 0 is better than average
	BaseOutRunsAdded *float32 `json:"base_out_runs_added" parquet:"name=base_out_runs_added, type=FLOAT"`
	// Putout (PO) [Defense]- https://www.mlb.com/glossary/standard-stats/putout
	Putout int32 `json:"putout" parquet:"name=putout, type=INT32"`
	// Assist (A) [Defense] - https://www.mlb.com/glossary/standard-stats/assist
	Assist int32 `json:"assist" parquet:"name=assist, type=INT32"`
}
