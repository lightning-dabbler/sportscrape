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

	LeagueID               string `json:"league_id"`
	LeagueName             string `json:"league_name"`
	Date                   string `json:"date"`
	Completed              bool   `json:"completed"`
	Link                   string `json:"link"`
	Name                   string `json:"name"`
	IsPostponedOrCancelled bool   `json:"isPostponedOrCanceled"`
	StatusID               string `json:"id"`
	StatusState            string `json:"state"`
	StatusDetail           string `json:"detail"`
}
