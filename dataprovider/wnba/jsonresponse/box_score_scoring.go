package jsonresponse

// BoxScoreScoringJSON - box score player scoring json response by period (e.g. https://www.wnba.com/game/phx-vs-chi-1022600224/box-score?type=scoring&period=All)
type BoxScoreScoringJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period     int32                `json:"period"`
				GameStatus int32                `json:"gameStatus"`
				EventID    string               `json:"gameId"`
				HomeTeam   BoxScoreScoringStats `json:"homeTeam"`
				AwayTeam   BoxScoreScoringStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreScoringStats struct {
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
			Minutes                          string  `json:"minutes"`
			PercentageFieldGoalsAttempted2pt float32 `json:"percentageFieldGoalsAttempted2pt"`
			PercentageFieldGoalsAttempted3pt float32 `json:"percentageFieldGoalsAttempted3pt"`
			PercentagePoints2pt              float32 `json:"percentagePoints2pt"`
			PercentagePointsMidrange2pt      float32 `json:"percentagePointsMidrange2pt"`
			PercentagePoints3pt              float32 `json:"percentagePoints3pt"`
			PercentagePointsFastBreak        float32 `json:"percentagePointsFastBreak"`
			PercentagePointsFreeThrow        float32 `json:"percentagePointsFreeThrow"`
			PercentagePointsOffTurnovers     float32 `json:"percentagePointsOffTurnovers"`
			PercentagePointsPaint            float32 `json:"percentagePointsPaint"`
			PercentageAssisted2pt            float32 `json:"percentageAssisted2pt"`
			PercentageUnassisted2pt          float32 `json:"percentageUnassisted2pt"`
			PercentageAssisted3pt            float32 `json:"percentageAssisted3pt"`
			PercentageUnassisted3pt          float32 `json:"percentageUnassisted3pt"`
			PercentageAssistedFGM            float32 `json:"percentageAssistedFGM"`
			PercentageUnassistedFGM          float32 `json:"percentageUnassistedFGM"`
		} `json:"statistics"`
	} `json:"players"`
}
