package model

import "time"

// PlayByPlay - composite key: event_id, play_id
type PlayByPlay struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. f5908b6d-311e-11f0-b670-ae1250fadad1
	EventID string `json:"event_id" parquet:"name=event_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventTime is the scheduled start time of the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// HomeTeamID
	HomeTeamID string `json:"home_team_id" parquet:"name=home_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamID
	AwayTeamID string `json:"away_team_id" parquet:"name=away_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayID is unique within the matchup e.g. 40
	PlayID int64 `json:"play_id" parquet:"name=play_id, type=INT64"`
	// PlaySequenceNumber orders the plays e.g. 40
	PlaySequenceNumber float64 `json:"play_sequence_number" parquet:"name=play_sequence_number, type=DOUBLE"`
	// Quarter 1-4, 5 = OT1, 6 = OT2, ...; 0 for deleted plays
	Quarter int32 `json:"quarter" parquet:"name=quarter, type=INT32"`
	// Clock - game clock at the start of the play e.g. 14:54; nil for some plays e.g. END_QUARTER
	Clock *string `json:"clock" parquet:"name=clock, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayType e.g. GAME_START, KICK_OFF, RUSH, PASS, SACK, PUNT, FIELD_GOAL, XP_KICK, PENALTY, TIMEOUT, END_QUARTER, END_GAME, COMMENT, UNSPECIFIED
	PlayType string `json:"play_type" parquet:"name=play_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Down 1-4; 0 when the play isn't from scrimmage e.g. kickoffs, extra points, timeouts
	Down int32 `json:"down" parquet:"name=down, type=INT32"`
	// Distance - yards to go for a first down; 0 when the play isn't from scrimmage
	Distance int32 `json:"distance" parquet:"name=distance, type=INT32"`
	// YardLine - line of scrimmage e.g. DAL 47, 50; nil for plays without one e.g. timeouts
	YardLine *string `json:"yard_line" parquet:"name=yard_line, type=BYTE_ARRAY, convertedtype=UTF8"`
	// YardsGained e.g. 7
	YardsGained int32 `json:"yards_gained" parquet:"name=yards_gained, type=INT32"`
	// PrePlayByPlay - pre-snap summary e.g. PHI  KO  PHI 35
	PrePlayByPlay *string `json:"pre_play_by_play" parquet:"name=pre_play_by_play, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Description e.g. J.Williams left tackle to PHI 46 for 7 yards (C.DeJean).
	Description *string `json:"description" parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8"`
	// DescriptionWithJerseyNumbers e.g. 4-J.Elliott kicks 60 yards from PHI 35 to DAL 5.
	DescriptionWithJerseyNumbers *string `json:"description_with_jersey_numbers" parquet:"name=description_with_jersey_numbers, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayStartTime is the wall clock time the play started
	PlayStartTime *time.Time `json:"play_start_time"`
	// PlayStartTimeParquet is the wall clock time the play started (in milliseconds)
	PlayStartTimeParquet *int64 `json:"-" parquet:"name=play_start_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// PlayEndTime is the wall clock time the play ended
	PlayEndTime *time.Time `json:"play_end_time"`
	// PlayEndTimeParquet is the wall clock time the play ended (in milliseconds)
	PlayEndTimeParquet *int64 `json:"-" parquet:"name=play_end_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// IsGoalToGo
	IsGoalToGo bool `json:"is_goal_to_go" parquet:"name=is_goal_to_go, type=BOOLEAN"`
	// IsEndOfQuarter
	IsEndOfQuarter bool `json:"is_end_of_quarter" parquet:"name=is_end_of_quarter, type=BOOLEAN"`
	// Scored
	Scored bool `json:"scored" parquet:"name=scored, type=BOOLEAN"`
	// ScoringPlayType e.g. TOUCHDOWN, PAT, PAT2, FIELD_GOAL, UNSPECIFIED
	ScoringPlayType string `json:"scoring_play_type" parquet:"name=scoring_play_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// ScoringTeamID - the matchup's team ID (also for the Pro Bowl, whose drive chart uses Pro Bowl team IDs); nil when the play didn't score
	ScoringTeamID *string `json:"scoring_team_id" parquet:"name=scoring_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SpecialTeamsPlayType e.g. PENALTY, UNSPECIFIED
	SpecialTeamsPlayType string `json:"special_teams_play_type" parquet:"name=special_teams_play_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// NextPlayType e.g. PLAY_FROM_SCRIMMAGE, FREE_KICK, XP_KICK, UNSPECIFIED
	NextPlayType string `json:"next_play_type" parquet:"name=next_play_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// NextPlayIsGoalToGo
	NextPlayIsGoalToGo bool `json:"next_play_is_goal_to_go" parquet:"name=next_play_is_goal_to_go, type=BOOLEAN"`
	// Deleted - the play was nullified/removed
	Deleted bool `json:"deleted" parquet:"name=deleted, type=BOOLEAN"`
	// DriveSequence - the drive the play is part of e.g. 1; nil when the play isn't part of a drive (e.g. timeouts)
	DriveSequence *int32 `json:"drive_sequence" parquet:"name=drive_sequence, type=INT32"`
	// DriveTeamID - matchup's team ID of the team in possession for the drive (also for the Pro Bowl); nil when the play isn't part of a drive
	DriveTeamID *string `json:"drive_team_id" parquet:"name=drive_team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// DriveStart - how the drive started e.g. Kickoff, Punt, Fumble, Downs; nil when the play isn't part of a drive
	DriveStart *string `json:"drive_start" parquet:"name=drive_start, type=BYTE_ARRAY, convertedtype=UTF8"`
	// DriveResult - how the drive ended e.g. Touchdown, Field Goal, Punt, Fumble, Downs, End of Half, End of Game; nil when the play isn't part of a drive
	DriveResult *string `json:"drive_result" parquet:"name=drive_result, type=BYTE_ARRAY, convertedtype=UTF8"`
}
