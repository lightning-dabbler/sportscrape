package jsonresponse

// BoxScoreFourFactorsJSON - box score player four factors json response by period (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=fourfactors&period=All)
type BoxScoreFourFactorsJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				EventID  string                   `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreFourFactorsStats `json:"homeTeam"`
				AwayTeam BoxScoreFourFactorsStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreFourFactorsStats struct {
	TeamID      int64  `json:"teamId"`      // "teamId": 1610612762
	TeamName    string `json:"teamName"`    // "teamName": "Jazz"
	TeamCity    string `json:"teamCity"`    // "teamCity": "Utah"
	TeamTricode string `json:"teamTricode"` // "teamTricode": "UTA"
	Players     []struct {
		PersonID   int64  `json:"personId"`   // "personId": 1629004
		FirstName  string `json:"firstName"`  // "firstName": "Svi"
		FamilyName string `json:"familyName"` // "familyName": "Mykhailiuk"
		Position   string `json:"position"`   // "position": "F"
		JerseyNum  string `json:"jerseyNum"`  // "jerseyNum": "10"
		Statistics struct {
			Minutes                         string  `json:"minutes"`                         // "minutes": "28:39",
			EffectiveFieldGoalPercentage    float32 `json:"effectiveFieldGoalPercentage"`    // "effectiveFieldGoalPercentage": 0.458
			FreeThrowAttemptRate            float32 `json:"freeThrowAttemptRate"`            // "freeThrowAttemptRate": 0.254
			TeamTurnoverPercentage          float32 `json:"teamTurnoverPercentage"`          // "teamTurnoverPercentage": 0.063
			OffensiveReboundPercentage      float32 `json:"offensiveReboundPercentage"`      // "offensiveReboundPercentage": 0.036
			OppEffectiveFieldGoalPercentage float32 `json:"oppEffectiveFieldGoalPercentage"` // "oppEffectiveFieldGoalPercentage": 0.468
			OppFreeThrowAttemptRate         float32 `json:"oppFreeThrowAttemptRate"`         // "oppFreeThrowAttemptRate": 0.323
			OppTeamTurnoverPercentage       float32 `json:"oppTeamTurnoverPercentage"`       // "oppTeamTurnoverPercentage": 0.075
			OppOffensiveReboundPercentage   float32 `json:"oppOffensiveReboundPercentage"`   // "oppOffensiveReboundPercentage": 0.273
		} `json:"statistics"`
	} `json:"players"`
}
