package model

import "time"

// PitchingBoxScore represents the data model for MLB pitching box score stats scraped from baseballsavant.mlb.com
type PitchingBoxScore struct {
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
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// Team is the player's team name
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Opponent is the opposing team name
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// PlayerID
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// Player is the pitcher's name
	Player   string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
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
	// NumberOfPitches https://www.mlb.com/glossary/standard-stats/number-of-pitches
	NumberOfPitches int32 `json:"number_of_pitches" parquet:"name=number_of_pitches, type=INT32"`
	// InningsPitched - https://www.mlb.com/glossary/standard-stats/innings-pitched
	InningsPitched float32 `json:"innings_pitched" parquet:"name=innings_pitched, type=FLOAT"`
	// Saves - https://www.mlb.com/glossary/standard-stats/save
	Saves int32 `json:"saves" parquet:"name=saves, type=INT32"`
	// BlownSaves - https://www.mlb.com/glossary/standard-stats/blown-save
	BlownSaves int32 `json:"blown_saves" parquet:"name=blown_saves, type=INT32"`
	// EarnedRuns - https://www.mlb.com/glossary/standard-stats/earned-run
	EarnedRuns int32 `json:"earned_runs" parquet:"name=earned_runs, type=INT32"`
	// BattersFaced - https://www.mlb.com/glossary/standard-stats/batters-faced
	BattersFaced int32 `json:"batters_faced" parquet:"name=batters_faced, type=INT32"`
	// Outs - https://www.mlb.com/glossary/standard-stats/out
	Outs int32 `json:"outs" parquet:"name=outs, type=INT32"`
	// Shutouts - https://www.mlb.com/glossary/standard-stats/shutout
	Shutouts int32 `json:"shutouts" parquet:"name=shutouts, type=INT32"`
	// Balls - pitches out of the strike zone
	Balls int32 `json:"balls" parquet:"name=balls, type=INT32"`
	// Strikes - https://en.wikipedia.org/wiki/Glossary_of_baseball_terms#strike
	Strikes int32 `json:"strikes" parquet:"name=strikes, type=INT32"`
	// Balk - https://www.mlb.com/glossary/standard-stats/balk
	Balks int32 `json:"balks" parquet:"name=balks, type=INT32"`
	// WildPitches - https://www.mlb.com/glossary/standard-stats/wild-pitch
	WildPitches int32 `json:"wild_pitches" parquet:"name=wild_pitches, type=INT32"`
	// Pickoffs - https://www.mlb.com/glossary/standard-stats/pickoff
	Pickoffs int32 `json:"pickoffs" parquet:"name=pickoffs, type=INT32"`
	// RBI - https://www.mlb.com/glossary/standard-stats/runs-batted-in
	RBI int32 `json:"rbi" parquet:"name=rbi, type=INT32"`
	// EarnedRunAverage - https://www.mlb.com/glossary/standard-stats/earned-run-average
	EarnedRunAverage float32 `json:"earned_run_average" parquet:"name=earned_run_average, type=FLOAT"`
	// InheritedRunners - https://www.mlb.com/glossary/standard-stats/inherited-runner
	InheritedRunners       int32 `json:"inherited_runners" parquet:"name=inherited_runners, type=INT32"`
	InheritedRunnersScored int32 `json:"inherited_runners_scored" parquet:"name=inherited_runners_scored, type=INT32"`
	// CatchersInterference - https://www.mlb.com/glossary/rules/catcher-interference
	CatchersInterference int32 `json:"catchers_interference" parquet:"name=catchers_interference, type=INT32"`
	// SacBunts - https://www.mlb.com/glossary/standard-stats/sacrifice-bunt
	SacBunts int32 `json:"sac_bunts" parquet:"name=sac_bunts, type=INT32"`
	// SacFlies - https://www.mlb.com/glossary/standard-stats/sacrifice-fly
	SacFlies int32 `json:"sac_flies" parquet:"name=sac_flies, type=INT32"`
	// PassedBall -https://www.mlb.com/glossary/standard-stats/passed-ball
	PassedBall int32 `json:"passed_ball" parquet:"name=passed_ball, type=INT32"`
	// PopOuts - a pop fly that is caught for an out
	PopOuts int32 `json:"pop_outs" parquet:"name=pop_outs, type=INT32"`
	// LineOuts - a batter hits a line drive and a fielder catches the ball before it hits the ground
	LineOuts int32 `json:"line_outs" parquet:"name=line_outs, type=INT32"`
}
