package jsonresponse

// PlayByPlay is the json response from https://api-web.nhle.com/v1/gamecenter/{game_id}/play-by-play
type PlayByPlay struct {
	ID        int64  `json:"id"`
	GameState string `json:"gameState"`
	Plays     []Play `json:"plays"`
}

type Play struct {
	EventID          int64            `json:"eventId"`
	PeriodDescriptor PeriodDescriptor `json:"periodDescriptor"`
	// TimeInPeriod e.g. 19:00
	TimeInPeriod string `json:"timeInPeriod"`
	// TimeRemaining e.g. 01:00
	TimeRemaining         string `json:"timeRemaining"`
	SituationCode         string `json:"situationCode"`
	HomeTeamDefendingSide string `json:"homeTeamDefendingSide"`
	TypeCode              int32  `json:"typeCode"`
	// TypeDescKey e.g. goal, faceoff, penalty
	TypeDescKey string `json:"typeDescKey"`
	SortOrder   int32  `json:"sortOrder"`
	// Details is absent for some play types e.g. period-start
	Details *PlayDetails `json:"details"`
}

type PlayDetails struct {
	EventOwnerTeamID    *int64  `json:"eventOwnerTeamId"`
	XCoord              *int32  `json:"xCoord"`
	YCoord              *int32  `json:"yCoord"`
	ZoneCode            *string `json:"zoneCode"`
	ShotType            *string `json:"shotType"`
	Reason              *string `json:"reason"`
	SecondaryReason     *string `json:"secondaryReason"`
	TypeCode            *string `json:"typeCode"`
	DescKey             *string `json:"descKey"`
	Duration            *int32  `json:"duration"`
	AwayScore           *int32  `json:"awayScore"`
	HomeScore           *int32  `json:"homeScore"`
	AwaySOG             *int32  `json:"awaySOG"`
	HomeSOG             *int32  `json:"homeSOG"`
	GoalInGame          *int32  `json:"goalInGame"`
	ScoringPlayerID     *int64  `json:"scoringPlayerId"`
	ScoringPlayerTotal  *int32  `json:"scoringPlayerTotal"`
	Assist1PlayerID     *int64  `json:"assist1PlayerId"`
	Assist1PlayerTotal  *int32  `json:"assist1PlayerTotal"`
	Assist2PlayerID     *int64  `json:"assist2PlayerId"`
	Assist2PlayerTotal  *int32  `json:"assist2PlayerTotal"`
	GoalieInNetID       *int64  `json:"goalieInNetId"`
	ShootingPlayerID    *int64  `json:"shootingPlayerId"`
	BlockingPlayerID    *int64  `json:"blockingPlayerId"`
	HittingPlayerID     *int64  `json:"hittingPlayerId"`
	HitteePlayerID      *int64  `json:"hitteePlayerId"`
	WinningPlayerID     *int64  `json:"winningPlayerId"`
	LosingPlayerID      *int64  `json:"losingPlayerId"`
	CommittedByPlayerID *int64  `json:"committedByPlayerId"`
	DrawnByPlayerID     *int64  `json:"drawnByPlayerId"`
	ServedByPlayerID    *int64  `json:"servedByPlayerId"`
	PlayerID            *int64  `json:"playerId"`
}
