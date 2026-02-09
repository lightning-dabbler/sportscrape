package model

import "time"

// PlayByPlay - composite key: event_id, action_id
type PlayByPlay struct {
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
	// ActionNumber
	ActionNumber int32 `json:"action_number" parquet:"name=action_number, type=INT32"`
	// Clock
	Clock float32 `json:"clock" parquet:"name=clock, type=FLOAT"`
	// Period
	Period int32 `json:"period" parquet:"name=period, type=INT32"`
	// TeamID
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// TeamAbbreviation
	TeamAbbreviation string `json:"team_abbreviation" parquet:"name=team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PersonID
	PersonID int64 `json:"person_id" parquet:"name=person_id, type=INT64"`
	// PlayerName
	PlayerName string `json:"player_name" parquet:"name=player_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerNameInitial
	PlayerNameInitial string `json:"player_name_initial" parquet:"name=player_name_initial, type=BYTE_ARRAY, convertedtype=UTF8"`
	// ShotDistance
	ShotDistance float32 `json:"shot_distance" parquet:"name=shot_distance, type=FLOAT"`
	// ShotResult
	ShotResult string `json:"shot_result" parquet:"name=shot_result, type=BYTE_ARRAY, convertedtype=UTF8"`
	// IsFieldGoal
	IsFieldGoal int32 `json:"is_field_goal" parquet:"name=is_field_goal, type=INT32"`
	// ScoreHome
	ScoreHome string `json:"score_home" parquet:"name=score_home, type=BYTE_ARRAY, convertedtype=UTF8"`
	// ScoreAway
	ScoreAway string `json:"score_away" parquet:"name=score_away, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PointsTotal
	PointsTotal int32 `json:"points_total" parquet:"name=points_total, type=INT32"`
	// Location
	Location string `json:"location" parquet:"name=location, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Description
	Description string `json:"description" parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8"`
	// ActionType
	ActionType string `json:"action_type" parquet:"name=action_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SubType
	SubType string `json:"sub_type" parquet:"name=sub_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// ShotValue
	ShotValue int32 `json:"shot_value" parquet:"name=shot_value, type=INT32"`
	// ActionID
	ActionID int32 `json:"action_id" parquet:"name=action_id, type=INT32"`
}
