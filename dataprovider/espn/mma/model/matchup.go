package model

import (
	"time"
)

type Matchup struct {

	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 86833
	EventID string `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the timestamp associated with the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`

	LeagueID               string `json:"league_id" parquet:"name=league_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	LeagueName             string `json:"league_name" parquet:"name=league_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Date                   string `json:"date" parquet:"name=date, type=BYTE_ARRAY, convertedtype=UTF8"`
	Completed              bool   `json:"completed" parquet:"name=completed, type=BOOLEAN"`
	Link                   string `json:"link" parquet:"name=link, type=BYTE_ARRAY, convertedtype=UTF8"`
	Name                   string `json:"name" parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8"`
	IsPostponedOrCancelled bool   `json:"is_postponed_or_cancelled" parquet:"name=is_postponed_or_cancelled, type=BOOLEAN"`
	StatusID               string `json:"id" parquet:"name=id, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusState            string `json:"state" parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusDetail           string `json:"detail" parquet:"name=detail, type=BYTE_ARRAY, convertedtype=UTF8"`
}
