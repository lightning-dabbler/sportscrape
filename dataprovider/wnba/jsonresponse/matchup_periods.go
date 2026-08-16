package jsonresponse

// MatchupPeriodsJSON - period-by-period team score json response, sourced
// from the box-score page (e.g. https://www.wnba.com/game/dal-vs-ind-1022600254/box-score?period=All&type=traditional).
// Unlike NBA, WNBA's schedule API (jsonresponse.MatchupJSON) carries no
// period/quarter data at all - it lives on the box-score page instead, as a
// sibling of players/statistics on the same game.homeTeam/awayTeam envelope,
// present regardless of `type`.
type MatchupPeriodsJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period     int32              `json:"period"`
				GameStatus int32              `json:"gameStatus"`
				EventID    string             `json:"gameId"`
				HomeTeam   MatchupPeriodsTeam `json:"homeTeam"`
				AwayTeam   MatchupPeriodsTeam `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type MatchupPeriodsTeam struct {
	Score   int32 `json:"score"`
	Periods []struct {
		Period     int32  `json:"period"`
		PeriodType string `json:"periodType"` // e.g. "REGULAR"
		Score      int32  `json:"score"`
	} `json:"periods"`
}
