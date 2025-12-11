package model

import "time"

// BoxScoreHustle - composite key: event_id, player_id
type BoxScoreHustle struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a string ID that maps to the matchup e.g. 0022500249
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventStatus the numerical representation of the event status e.g. 3 (1=pregame, 2=in progress, 3=final)
	EventStatus int32 `json:"event_status" parquet:"name=event_status, type=INT32"`
	// EventStatusText (e.g. Final, Final/OT2, etc.)
	EventStatusText string `json:"event_status_text" parquet:"name=event_status_text, type=BYTE_ARRAY, convertedtype=UTF8"`
	// TeamID
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// TeamName
	TeamName string `json:"team_name" parquet:"name=team_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// TeamNameFull
	TeamNameFull string `json:"team_name_full" parquet:"name=team_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// OpponentName
	OpponentName string `json:"opponent_name" parquet:"name=opponent_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentNameFull
	OpponentNameFull string `json:"opponent_name_full" parquet:"name=opponent_name_full, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// PlayerName
	PlayerName string `json:"player_name" parquet:"name=player_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Position
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Starter
	Starter bool `json:"starter" parquet:"name=starter, type=BOOLEAN"`
	// Minutes
	Minutes float32 `json:"minutes" parquet:"name=minutes, type=FLOAT"`
	// Points
	Points int32 `json:"points" parquet:"name=points, type=INT32"`
	// ContestedShots
	ContestedShots int32 `json:"contested_shots" parquet:"name=contested_shots, type=INT32"`
	// ContestedShots2pt
	ContestedShots2pt int32 `json:"contested_shots_2pt" parquet:"name=contested_shots_2pt, type=INT32"`
	// ContestedShots3pt
	ContestedShots3pt int32 `json:"contested_shots_3pt" parquet:"name=contested_shots_3pt, type=INT32"`
	// Deflections
	Deflections int32 `json:"deflections" parquet:"name=deflections, type=INT32"`
	// ChargesDrawn
	ChargesDrawn int32 `json:"charges_drawn" parquet:"name=charges_drawn, type=INT32"`
	// ScreenAssists
	ScreenAssists int32 `json:"screen_assists" parquet:"name=screen_assists, type=INT32"`
	// ScreenAssistPoints
	ScreenAssistPoints int32 `json:"screen_assist_points" parquet:"name=screen_assist_points, type=INT32"`
	// LooseBallsRecoveredOffensive
	LooseBallsRecoveredOffensive int32 `json:"loose_balls_recovered_offensive" parquet:"name=loose_balls_recovered_offensive, type=INT32"`
	// LooseBallsRecoveredDefensive
	LooseBallsRecoveredDefensive int32 `json:"loose_balls_recovered_defensive" parquet:"name=loose_balls_recovered_defensive, type=INT32"`
	// LooseBallsRecoveredTotal
	LooseBallsRecoveredTotal int32 `json:"loose_balls_recovered_total" parquet:"name=loose_balls_recovered_total, type=INT32"`
	// OffensiveBoxOuts
	OffensiveBoxOuts int32 `json:"offensive_box_outs" parquet:"name=offensive_box_outs, type=INT32"`
	// DefensiveBoxOuts
	DefensiveBoxOuts int32 `json:"defensive_box_outs" parquet:"name=defensive_box_outs, type=INT32"`
	// BoxOutPlayerTeamRebounds
	BoxOutPlayerTeamRebounds int32 `json:"box_out_player_team_rebounds" parquet:"name=box_out_player_team_rebounds, type=INT32"`
	// BoxOutPlayerRebounds
	BoxOutPlayerRebounds int32 `json:"box_out_player_rebounds" parquet:"name=box_out_player_rebounds, type=INT32"`
	// BoxOuts
	BoxOuts int32 `json:"box_outs" parquet:"name=box_outs, type=INT32"`
}
