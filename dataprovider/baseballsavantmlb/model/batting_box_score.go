package model

import "time"

// BattingBoxScore represents the data model for MLB batting box score stats scraped from baseballsavant.mlb.com
type BattingBoxScore struct {
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
	// TeamID
	TeamID string `json:"team_id" parquet:"name=team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Team is the player's team name
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID string `json:"opponent_id" parquet:"name=opponent_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Player is the pitcher's name
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// FlyOuts - https://www.mlb.com/glossary/standard-stats/flyout
	FlyOuts int32 `json:"fly_outs" parquet:"name=fly_outs, type=INT32"`
	// GroundOuts - https://www.mlb.com/glossary/standard-stats/groundout
	GroundOuts int32 `json:"ground_outs" parquet:"name=ground_outs, type=INT32"`
	// AirOuts - refers to a batted ball that is hit in the air, either a fly ball or a line drive, and is caught by a fielder for an out
	AirOuts int32 `json:"air_outs" parquet:"name=air_outs, type=INT32"`
	// Runs - https://www.mlb.com/glossary/standard-stats/run
	Runs int32 `json:"runs" parquet:"name=runs, type=INT32"`
	// Doubles - https://www.mlb.com/glossary/standard-stats/double
	Doubles int32 `json:"doubles" parquet:"name=doubles, type=INT32"`
	// Triples - https://www.mlb.com/glossary/standard-stats/triple
	Triples int32 `json:"triples" parquet:"name=triples, type=INT32"`
	// HomeRuns - https://www.mlb.com/glossary/standard-stats/home-run
	HomeRuns int32 `json:"home_runs" parquet:"name=home_runs, type=INT32"`
	// Strikeouts - https://www.mlb.com/glossary/standard-stats/strikeout
	Strikeouts int32 `json:"strikeouts" parquet:"name=strikeouts, type=INT32"`
	// Walks - https://www.mlb.com/glossary/standard-stats/walk
	Walks int32 `json:"walks" parquet:"name=walks, type=INT32"`
	// IntentionalWalks - https://www.mlb.com/glossary/standard-stats/intentional-walk
	IntentionalWalks int32 `json:"intentional_walks" parquet:"name=intentional_walks, type=INT32"`
	// Hits - https://www.mlb.com/glossary/standard-stats/hit
	Hits int32 `json:"hits" parquet:"name=hits, type=INT32"`
	// HitByPitch - https://www.mlb.com/glossary/standard-stats/hit-by-pitch
	HitByPitch int32 `json:"hit_by_pitch" parquet:"name=hit_by_pitch, type=INT32"`
	// AtBats - https://www.mlb.com/glossary/standard-stats/at-bat
	AtBats int32 `json:"at_bats" parquet:"name=at_bats, type=INT32"`
	// CaughtStealing - https://www.mlb.com/glossary/standard-stats/caught-stealing
	CaughtStealing int32 `json:"caught_stealing" parquet:"name=caught_stealing, type=INT32"`
	// StolenBases - https://www.mlb.com/glossary/standard-stats/stolen-base
	StolenBases int32 `json:"stolen_bases" parquet:"name=stolen_bases, type=INT32"`
	// GroundIntoDoublePlay - https://www.mlb.com/glossary/standard-stats/ground-into-double-play
	GroundIntoDoublePlay int32 `json:"ground_into_double_play" parquet:"name=ground_into_double_play, type=INT32"`
	// GroundIntoTriplePlay - https://www.mlb.com/glossary/standard-stats/triple-play
	GroundIntoTriplePlay int32 `json:"ground_into_triple_play" parquet:"name=ground_into_triple_play, type=INT32"`
	// PlateAppearances - https://www.mlb.com/glossary/standard-stats/plate-appearance
	PlateAppearances int32 `json:"plate_appearances" parquet:"name=plate_appearances, type=INT32"`
	// TotalBases - https://www.mlb.com/glossary/standard-stats/total-bases
	TotalBases int32 `json:"total_bases" parquet:"name=total_bases, type=INT32"`
	// RBI - https://www.mlb.com/glossary/standard-stats/runs-batted-in
	RBI int32 `json:"rbi" parquet:"name=rbi, type=INT32"`
	// LeftOnBase - https://www.mlb.com/glossary/standard-stats/left-on-base
	LeftOnBase int32 `json:"left_on_base" parquet:"name=left_on_base, type=INT32"`
	// SacBunts - https://www.mlb.com/glossary/standard-stats/sacrifice-bunt
	SacBunts int32 `json:"sac_bunts" parquet:"name=sac_bunts, type=INT32"`
	// SacFlies - https://www.mlb.com/glossary/standard-stats/sacrifice-fly
	SacFlies int32 `json:"sac_flies" parquet:"name=sac_flies, type=INT32"`
	// CatchersInterference - https://www.mlb.com/glossary/rules/catcher-interference
	CatchersInterference int32 `json:"catchers_interference" parquet:"name=catchers_interference, type=INT32"`
	// Pickoffs - https://www.mlb.com/glossary/standard-stats/pickoff
	Pickoffs int32 `json:"pickoffs" parquet:"name=pickoffs, type=INT32"`
	// AtBatsPerHomeRun - https://en.wikipedia.org/wiki/At_bats_per_home_run
	AtBatsPerHomeRun float32 `json:"at_bats_per_home_run" parquet:"name=at_bats_per_home_run, type=FLOAT"`
	// PopOuts - a pop fly that is caught for an out
	PopOuts int32 `json:"pop_outs" parquet:"name=pop_outs, type=INT32"`
	// LineOuts - a batter hits a line drive and a fielder catches the ball before it hits the ground
	LineOuts int32 `json:"line_outs" parquet:"name=line_outs, type=INT32"`
}
