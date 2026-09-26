package jsonresponse

// Score is the json response from https://api-web.nhle.com/v1/score/{date}
type Score struct {
	Games []Game `json:"games"`
}

type Game struct {
	ID     int64 `json:"id"`
	Season int64 `json:"season"`
	// GameType 1 = pre-season, 2 = regular season, 3 = post season (other values occur e.g. 19 for event 2024190001)
	GameType          int32           `json:"gameType"`
	GameDate          string          `json:"gameDate"`
	Venue             LocalizedString `json:"venue"`
	StartTimeUTC      string          `json:"startTimeUTC"`
	GameState         string          `json:"gameState"`
	GameScheduleState string          `json:"gameScheduleState"`
	AwayTeam          Team            `json:"awayTeam"`
	HomeTeam          Team            `json:"homeTeam"`
	// Period is absent before the game starts
	Period *int32 `json:"period"`
	// GameOutcome is absent before the game ends
	GameOutcome *GameOutcome `json:"gameOutcome"`
}

type Team struct {
	ID     int64           `json:"id"`
	Name   LocalizedString `json:"name"`
	Abbrev string          `json:"abbrev"`
	// Score is absent before the game starts
	Score *int32 `json:"score"`
	// SOG is absent before the game starts
	SOG *int32 `json:"sog"`
}

type GameOutcome struct {
	// LastPeriodType e.g. REG, OT, SO
	LastPeriodType string `json:"lastPeriodType"`
}
