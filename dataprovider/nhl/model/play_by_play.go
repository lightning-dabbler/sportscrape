package model

import "time"

// PlayByPlay - composite key: event_id, play_event_id
type PlayByPlay struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the matchup e.g. 2024020250
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the scheduled start time of the matchup
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the scheduled start time of the matchup (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// LimitedScoring is the API's limitedScoring flag. "Scoring" means stat recording (scorekeeping), not goals:
	// true means the NHL recorded only a limited set of stats for the game, and the play-by-play contains
	// only goal and penalty events (no shots, hits, faceoffs, etc.).
	LimitedScoring bool `json:"limited_scoring" parquet:"name=limited_scoring, type=BOOLEAN"`
	// HomeTeamID
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// AwayTeamID
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// PlayEventID is the ID of the play within the game e.g. 464
	PlayEventID int64 `json:"play_event_id" parquet:"name=play_event_id, type=INT64"`
	// SortOrder orders plays chronologically
	SortOrder int32 `json:"sort_order" parquet:"name=sort_order, type=INT32"`
	// Period number e.g. 1, 2, 3, 4 (OT), 5 (SO)
	Period int32 `json:"period" parquet:"name=period, type=INT32"`
	// PeriodType e.g. REG, OT, SO
	PeriodType string `json:"period_type" parquet:"name=period_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// TimeInPeriod - elapsed time in the period in minutes e.g. 19.0
	TimeInPeriod float32 `json:"time_in_period" parquet:"name=time_in_period, type=FLOAT"`
	// TimeRemaining - time remaining in the period in minutes e.g. 1.0
	TimeRemaining float32 `json:"time_remaining" parquet:"name=time_remaining, type=FLOAT"`
	// SituationCode e.g. 1551
	SituationCode string `json:"situation_code" parquet:"name=situation_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeamDefendingSide e.g. left, right
	HomeTeamDefendingSide string `json:"home_team_defending_side" parquet:"name=home_team_defending_side, type=BYTE_ARRAY, convertedtype=UTF8"`
	// TypeCode e.g. 505
	TypeCode int32 `json:"type_code" parquet:"name=type_code, type=INT32"`
	// TypeDescKey e.g. goal, faceoff, penalty
	TypeDescKey string `json:"type_desc_key" parquet:"name=type_desc_key, type=BYTE_ARRAY, convertedtype=UTF8"`
	// EventOwnerTeamID
	EventOwnerTeamID *int64 `json:"event_owner_team_id" parquet:"name=event_owner_team_id, type=INT64"`
	// XCoord
	XCoord *int32 `json:"x_coord" parquet:"name=x_coord, type=INT32"`
	// YCoord
	YCoord *int32 `json:"y_coord" parquet:"name=y_coord, type=INT32"`
	// ZoneCode e.g. O, D, N
	ZoneCode *string `json:"zone_code" parquet:"name=zone_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	// ShotType e.g. wrist
	ShotType *string `json:"shot_type" parquet:"name=shot_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Reason
	Reason *string `json:"reason" parquet:"name=reason, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SecondaryReason
	SecondaryReason *string `json:"secondary_reason" parquet:"name=secondary_reason, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PenaltyTypeCode e.g. MIN
	PenaltyTypeCode *string `json:"penalty_type_code" parquet:"name=penalty_type_code, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PenaltyDescKey e.g. high-sticking
	PenaltyDescKey *string `json:"penalty_desc_key" parquet:"name=penalty_desc_key, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PenaltyDuration in minutes e.g. 2
	PenaltyDuration *int32 `json:"penalty_duration" parquet:"name=penalty_duration, type=INT32"`
	// AwayScore - score after a goal event; on a shootout (period_type SO) goal it is the pre-shootout score
	AwayScore *int32 `json:"away_score" parquet:"name=away_score, type=INT32"`
	// HomeScore - score after a goal event; on a shootout (period_type SO) goal it is the pre-shootout score
	HomeScore *int32 `json:"home_score" parquet:"name=home_score, type=INT32"`
	// AwaySOG - running shots on goal as reported on shot-on-goal events; excludes goals and, in a shootout,
	// includes shootout attempts, so it is not the official total (use the matchup's away_team_sog)
	AwaySOG *int32 `json:"away_sog" parquet:"name=away_sog, type=INT32"`
	// HomeSOG - running shots on goal as reported on shot-on-goal events; excludes goals and, in a shootout,
	// includes shootout attempts, so it is not the official total (use the matchup's home_team_sog)
	HomeSOG *int32 `json:"home_sog" parquet:"name=home_sog, type=INT32"`
	// GoalInGame
	GoalInGame *int32 `json:"goal_in_game" parquet:"name=goal_in_game, type=INT32"`
	// ScoringPlayerID
	ScoringPlayerID *int64 `json:"scoring_player_id" parquet:"name=scoring_player_id, type=INT64"`
	// ScoringPlayerTotal
	ScoringPlayerTotal *int32 `json:"scoring_player_total" parquet:"name=scoring_player_total, type=INT32"`
	// Assist1PlayerID
	Assist1PlayerID *int64 `json:"assist1_player_id" parquet:"name=assist1_player_id, type=INT64"`
	// Assist1PlayerTotal
	Assist1PlayerTotal *int32 `json:"assist1_player_total" parquet:"name=assist1_player_total, type=INT32"`
	// Assist2PlayerID
	Assist2PlayerID *int64 `json:"assist2_player_id" parquet:"name=assist2_player_id, type=INT64"`
	// Assist2PlayerTotal
	Assist2PlayerTotal *int32 `json:"assist2_player_total" parquet:"name=assist2_player_total, type=INT32"`
	// GoalieInNetID
	GoalieInNetID *int64 `json:"goalie_in_net_id" parquet:"name=goalie_in_net_id, type=INT64"`
	// ShootingPlayerID
	ShootingPlayerID *int64 `json:"shooting_player_id" parquet:"name=shooting_player_id, type=INT64"`
	// BlockingPlayerID
	BlockingPlayerID *int64 `json:"blocking_player_id" parquet:"name=blocking_player_id, type=INT64"`
	// HittingPlayerID
	HittingPlayerID *int64 `json:"hitting_player_id" parquet:"name=hitting_player_id, type=INT64"`
	// HitteePlayerID
	HitteePlayerID *int64 `json:"hittee_player_id" parquet:"name=hittee_player_id, type=INT64"`
	// WinningPlayerID - faceoff winner
	WinningPlayerID *int64 `json:"winning_player_id" parquet:"name=winning_player_id, type=INT64"`
	// LosingPlayerID - faceoff loser
	LosingPlayerID *int64 `json:"losing_player_id" parquet:"name=losing_player_id, type=INT64"`
	// CommittedByPlayerID - penalty committed by
	CommittedByPlayerID *int64 `json:"committed_by_player_id" parquet:"name=committed_by_player_id, type=INT64"`
	// DrawnByPlayerID - penalty drawn by
	DrawnByPlayerID *int64 `json:"drawn_by_player_id" parquet:"name=drawn_by_player_id, type=INT64"`
	// ServedByPlayerID - penalty served by
	ServedByPlayerID *int64 `json:"served_by_player_id" parquet:"name=served_by_player_id, type=INT64"`
	// PlayerID - player associated with the play e.g. giveaway, takeaway
	PlayerID *int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
}
