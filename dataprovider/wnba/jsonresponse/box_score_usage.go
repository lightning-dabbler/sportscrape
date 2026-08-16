package jsonresponse

// BoxScoreUsageJSON - Full box score player usage json response by period (e.g. https://www.wnba.com/game/phx-vs-chi-1022600224/box-score?type=usage&period=All)
type BoxScoreUsageJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period     int32              `json:"period"`
				GameStatus int32              `json:"gameStatus"`
				EventID    string             `json:"gameId"`
				HomeTeam   BoxScoreUsageStats `json:"homeTeam"`
				AwayTeam   BoxScoreUsageStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreUsageStats struct {
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
			UsagePercentage                  float32 `json:"usagePercentage"`
			PercentageFieldGoalsMade         float32 `json:"percentageFieldGoalsMade"`
			PercentageFieldGoalsAttempted    float32 `json:"percentageFieldGoalsAttempted"`
			PercentageThreePointersMade      float32 `json:"percentageThreePointersMade"`
			PercentageThreePointersAttempted float32 `json:"percentageThreePointersAttempted"`
			PercentageFreeThrowsMade         float32 `json:"percentageFreeThrowsMade"`
			PercentageFreeThrowsAttempted    float32 `json:"percentageFreeThrowsAttempted"`
			PercentageReboundsOffensive      float32 `json:"percentageReboundsOffensive"`
			PercentageReboundsDefensive      float32 `json:"percentageReboundsDefensive"`
			PercentageReboundsTotal          float32 `json:"percentageReboundsTotal"`
			PercentageAssists                float32 `json:"percentageAssists"`
			PercentageTurnovers              float32 `json:"percentageTurnovers"`
			PercentageSteals                 float32 `json:"percentageSteals"`
			PercentageBlocks                 float32 `json:"percentageBlocks"`
			PercentageBlocksAllowed          float32 `json:"percentageBlocksAllowed"`
			PercentagePersonalFouls          float32 `json:"percentagePersonalFouls"`
			PercentagePersonalFoulsDrawn     float32 `json:"percentagePersonalFoulsDrawn"`
			PercentagePoints                 float32 `json:"percentagePoints"`
		} `json:"statistics"`
	} `json:"players"`
}
