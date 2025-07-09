package model

import "time"

// PlayByPlay represents the data model for MLB event play by play scraped from baseballsavant.mlb.com
type PlayByPlay struct {
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
	// PlayID
	PlayID string `json:"play_id" parquet:"name=play_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Inning
	Inning int32 `json:"inning" parquet:"name=inning, type=INT32"`
	// AtBatNum
	AtBatNum int32 `json:"at_bat_number" parquet:"name=at_bat_number, type=INT32"`
	// Strikes
	Strikes int32 `json:"strikes" parquet:"name=strikes, type=INT32"`
	// Balls
	Balls int32 `json:"balls" parquet:"name=balls, type=INT32"`
	// Outs
	Outs int32 `json:"outs" parquet:"name=outs, type=INT32"`
	// PreStrikes - the number of strikes before the current play
	PreStrikes int32 `json:"pre_strikes" parquet:"name=pre_strikes, type=INT32"`
	// PreBalls - the number of balls before the current play
	PreBalls int32 `json:"pre_balls" parquet:"name=pre_balls, type=INT32"`
	// BatterID
	BatterID int64 `json:"batter_id" parquet:"name=batter_id, type=INT64"`
	// BatterStand - the abbreviation of the side of the home plate the batter is standing
	// R = the first-base side of the home plate, L = the third-base side of home plate
	BatterStand string `json:"batter_stand" parquet:"name=batter_stand, type=BYTE_ARRAY, convertedtype=UTF8"`
	// BatterName
	BatterName string `json:"batter_name" parquet:"name=batter_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PitcherID
	PitcherID int64 `json:"pitcher_id" parquet:"name=pitcher_id, type=INT32"`
	// PitcherThrow - https://www.mlb.com/glossary/pitch-types
	PitcherThrow string `json:"pitcher_throw" parquet:"name=pitcher_throw, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PitcherName
	PitcherName string `json:"pitcher_name" parquet:"name=pitcher_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Result - the terminal outcome of the batter
	Result *string `json:"result" parquet:"name=result, type=BYTE_ARRAY, convertedtype=UTF8"`
	// ResultDescription extrapolates on Result
	ResultDescription *string `json:"result_description" parquet:"name=result_description, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PitchType
	PitchType string `json:"pitch_type" parquet:"name=pitch_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PitchName
	PitchName string `json:"pitch_name" parquet:"name=pitch_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// CallName is the call on the play
	CallName string `json:"call_name" parquet:"name=call_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// CallDescription
	CallDescription string `json:"call_description" parquet:"name=call_description, type=BYTE_ARRAY, convertedtype=UTF8"`
	// IsStrikeSwinging
	IsStrikeSwinging bool `json:"is_strike_swinging" parquet:"name=is_strike_swinging, type=BOOLEAN"`
	// PitchStartSpeed
	PitchStartSpeed float32 `json:"pitch_start_speed" parquet:"name=pitch_start_speed, type=FLOAT"`
	// PitchEndSpeed
	PitchEndSpeed float32 `json:"pitch_end_speed" parquet:"name=pitch_end_speed, type=FLOAT"`
	// PitchNumber
	PitchNumber int32 `json:"pitch_number" parquet:"name=pitch_number, type=INT32"`
	// PitcherTotalPitches
	PitcherTotalPitches int32 `json:"pitcher_total_pitches" parquet:"name=pitcher_total_pitches, type=INT32"`
	// PitcherTotalPitchesByPitchType
	PitcherTotalPitchesByPitchType int32 `json:"pitcher_total_pitches_by_pitch_type" parquet:"name=pitcher_total_pitches_by_pitch_type, type=INT32"`
	// GameTotalPitches
	GameTotalPitches int32 `json:"game_total_pitches" parquet:"name=game_total_pitches, type=INT32"`
	// HitSpeed
	HitSpeed *string `json:"hit_speed" parquet:"name=hit_speed, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HitDistance
	HitDistance *string `json:"hit_distance" parquet:"name=hit_distance, type=BYTE_ARRAY, convertedtype=UTF8"`
}
