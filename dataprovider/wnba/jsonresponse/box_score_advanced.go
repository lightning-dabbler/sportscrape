package jsonresponse

// BoxScoreAdvancedJSON - Full box score player advanced json response by period (e.g. https://www.wnba.com/game/phx-vs-chi-1022600224/box-score?type=advanced&period=All)
type BoxScoreAdvancedJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period     int32                 `json:"period"`
				GameStatus int32                 `json:"gameStatus"`
				EventID    string                `json:"gameId"`
				HomeTeam   BoxScoreAdvancedStats `json:"homeTeam"`
				AwayTeam   BoxScoreAdvancedStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreAdvancedStats struct {
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
			Minutes                      string  `json:"minutes"`
			EstimatedOffensiveRating     float32 `json:"estimatedOffensiveRating"`
			OffensiveRating              float32 `json:"offensiveRating"`
			EstimatedDefensiveRating     float32 `json:"estimatedDefensiveRating"`
			DefensiveRating              float32 `json:"defensiveRating"`
			EstimatedNetRating           float32 `json:"estimatedNetRating"`
			NetRating                    float32 `json:"netRating"`
			AssistPercentage             float32 `json:"assistPercentage"`
			AssistToTurnover             float32 `json:"assistToTurnover"`
			AssistRatio                  float32 `json:"assistRatio"`
			OffensiveReboundPercentage   float32 `json:"offensiveReboundPercentage"`
			DefensiveReboundPercentage   float32 `json:"defensiveReboundPercentage"`
			ReboundPercentage            float32 `json:"reboundPercentage"`
			TurnoverRatio                float32 `json:"turnoverRatio"`
			EffectiveFieldGoalPercentage float32 `json:"effectiveFieldGoalPercentage"`
			TrueShootingPercentage       float32 `json:"trueShootingPercentage"`
			UsagePercentage              float32 `json:"usagePercentage"`
			EstimatedUsagePercentage     float32 `json:"estimatedUsagePercentage"`
			EstimatedPace                float32 `json:"estimatedPace"`
			Pace                         float32 `json:"pace"`
			PacePer40                    float32 `json:"pacePer40"`
			Possessions                  int32   `json:"possessions"`
			PIE                          float32 `json:"PIE"`
		} `json:"statistics"`
	} `json:"players"`
}
