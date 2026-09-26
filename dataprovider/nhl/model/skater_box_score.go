package model

import "time"

// SkaterBoxScore is a forward or defense box score statline - composite key: event_id, player_id
type SkaterBoxScore struct {
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
	// true means the NHL recorded only a limited set of stats for the game. Goals, assists, points, plus/minus,
	// penalty minutes, power play goals and shots on goal are recorded; TOI, giveaways, takeaways and faceoff
	// winning pctg are nil; hits, blocked shots and shifts are reported as 0 because they were not tracked.
	LimitedScoring bool `json:"limited_scoring" parquet:"name=limited_scoring, type=BOOLEAN"`
	// TeamID
	TeamID int64 `json:"team_id" parquet:"name=team_id, type=INT64"`
	// Team e.g. Flames
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// OpponentID
	OpponentID int64 `json:"opponent_id" parquet:"name=opponent_id, type=INT64"`
	// Opponent e.g. Canucks
	Opponent string `json:"opponent" parquet:"name=opponent, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID
	PlayerID int64 `json:"player_id" parquet:"name=player_id, type=INT64"`
	// Player - full name from the play-by-play rosterSpots e.g. Jonathan Huberdeau (falls back to the box score name e.g. J. Huberdeau when the player has no first/last name there)
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// SweaterNumber
	SweaterNumber int32 `json:"sweater_number" parquet:"name=sweater_number, type=INT32"`
	// Position e.g. C, L, R, D
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Goals
	Goals int32 `json:"goals" parquet:"name=goals, type=INT32"`
	// Assists
	Assists int32 `json:"assists" parquet:"name=assists, type=INT32"`
	// Points
	Points int32 `json:"points" parquet:"name=points, type=INT32"`
	// PlusMinus
	PlusMinus int32 `json:"plus_minus" parquet:"name=plus_minus, type=INT32"`
	// PIM - penalty minutes
	PIM int32 `json:"pim" parquet:"name=pim, type=INT32"`
	// Hits
	Hits int32 `json:"hits" parquet:"name=hits, type=INT32"`
	// PowerPlayGoals
	PowerPlayGoals int32 `json:"power_play_goals" parquet:"name=power_play_goals, type=INT32"`
	// SOG - shots on goal
	SOG int32 `json:"sog" parquet:"name=sog, type=INT32"`
	// FaceoffWinningPctg e.g. 0.5; nil when absent from the box score
	FaceoffWinningPctg *float32 `json:"faceoff_winning_pctg" parquet:"name=faceoff_winning_pctg, type=FLOAT"`
	// TOI - time on ice in minutes e.g. 18.75; nil in limited scoring games
	TOI *float32 `json:"toi" parquet:"name=toi, type=FLOAT"`
	// BlockedShots
	BlockedShots int32 `json:"blocked_shots" parquet:"name=blocked_shots, type=INT32"`
	// Shifts
	Shifts int32 `json:"shifts" parquet:"name=shifts, type=INT32"`
	// Giveaways - nil in limited scoring games
	Giveaways *int32 `json:"giveaways" parquet:"name=giveaways, type=INT32"`
	// Takeaways - nil in limited scoring games
	Takeaways *int32 `json:"takeaways" parquet:"name=takeaways, type=INT32"`
}
