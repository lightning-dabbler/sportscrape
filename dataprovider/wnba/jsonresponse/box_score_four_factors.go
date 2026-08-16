package jsonresponse

// BoxScoreFourFactorsJSON - box score player four factors json response by period (e.g. https://www.wnba.com/game/phx-vs-chi-1022600224/box-score?type=fourfactors&period=All)
type BoxScoreFourFactorsJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period     int32                    `json:"period"`
				GameStatus int32                    `json:"gameStatus"`
				EventID    string                   `json:"gameId"`
				HomeTeam   BoxScoreFourFactorsStats `json:"homeTeam"`
				AwayTeam   BoxScoreFourFactorsStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreFourFactorsStats struct {
	TeamID      int64  `json:"teamId"`
	TeamName    string `json:"teamName"`
	TeamCity    string `json:"teamCity"`
	TeamTricode string `json:"teamTricode"`
	Players     []struct {
		PersonID   int64  `json:"personId"`
		FirstName  string `json:"firstName"`
		FamilyName string `json:"familyName"`
		Position   string `json:"position"`
		JerseyNum  string `json:"jerseyNum"`
		Statistics struct {
			Minutes                         string  `json:"minutes"`
			EffectiveFieldGoalPercentage    float32 `json:"effectiveFieldGoalPercentage"`
			FreeThrowAttemptRate            float32 `json:"freeThrowAttemptRate"`
			TeamTurnoverPercentage          float32 `json:"teamTurnoverPercentage"`
			OffensiveReboundPercentage      float32 `json:"offensiveReboundPercentage"`
			OppEffectiveFieldGoalPercentage float32 `json:"oppEffectiveFieldGoalPercentage"`
			OppFreeThrowAttemptRate         float32 `json:"oppFreeThrowAttemptRate"`
			OppTeamTurnoverPercentage       float32 `json:"oppTeamTurnoverPercentage"`
			OppOffensiveReboundPercentage   float32 `json:"oppOffensiveReboundPercentage"`
		} `json:"statistics"`
	} `json:"players"`
}
