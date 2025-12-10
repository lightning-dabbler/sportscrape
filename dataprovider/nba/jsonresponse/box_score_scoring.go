package jsonresponse

// BoxScoreScoringJSON - box score player scoring json response by period (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=scoring&period=All)
type BoxScoreScoringJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period   int32                `json:"period"` //"period": 4
				EventID  string               `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreScoringStats `json:"homeTeam"`
				AwayTeam BoxScoreScoringStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreScoringStats struct {
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
			Minutes                          string  `json:"minutes"`                          // "minutes": "28:39"
			PercentageFieldGoalsAttempted2pt float32 `json:"percentageFieldGoalsAttempted2pt"` // "percentageFieldGoalsAttempted2pt": 0.5
			PercentageFieldGoalsAttempted3pt float32 `json:"percentageFieldGoalsAttempted3pt"` // "percentageFieldGoalsAttempted3pt": 0.5
			PercentagePoints2pt              float32 `json:"percentagePoints2pt"`              // "percentagePoints2pt": 0.4
			PercentagePointsMidrange2pt      float32 `json:"percentagePointsMidrange2pt"`      // "percentagePointsMidrange2pt": 0.2
			PercentagePoints3pt              float32 `json:"percentagePoints3pt"`              // "percentagePoints3pt": 0.6
			PercentagePointsFastBreak        float32 `json:"percentagePointsFastBreak"`        // "percentagePointsFastBreak": 0.3
			PercentagePointsFreeThrow        float32 `json:"percentagePointsFreeThrow"`        // "percentagePointsFreeThrow": 0
			PercentagePointsOffTurnovers     float32 `json:"percentagePointsOffTurnovers"`     // "percentagePointsOffTurnovers": 0
			PercentagePointsPaint            float32 `json:"percentagePointsPaint"`            // "percentagePointsPaint": 0.2
			PercentageAssisted2pt            float32 `json:"percentageAssisted2pt"`            // "percentageAssisted2pt": 1
			PercentageUnassisted2pt          float32 `json:"percentageUnassisted2pt"`          // "percentageUnassisted2pt": 0
			PercentageAssisted3pt            float32 `json:"percentageAssisted3pt"`            // "percentageAssisted3pt": 1
			PercentageUnassisted3pt          float32 `json:"percentageUnassisted3pt"`          // "percentageUnassisted3pt": 0
			PercentageAssistedFGM            float32 `json:"percentageAssistedFGM"`            // "percentageAssistedFGM": 1
			PercentageUnassistedFGM          float32 `json:"percentageUnassistedFGM"`          // "percentageUnassistedFGM": 0
		} `json:"statistics"`
	} `json:"players"`
}
