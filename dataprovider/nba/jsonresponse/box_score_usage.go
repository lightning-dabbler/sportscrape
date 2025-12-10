package jsonresponse

// BoxScoreUsageJSON - Full box score player traditional json response by period (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=usage&period=All)
type BoxScoreUsageJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period   int32              `json:"period"` //"period": 4
				EventID  string             `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreUsageStats `json:"homeTeam"`
				AwayTeam BoxScoreUsageStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreUsageStats struct {
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
			UsagePercentage                  float32 `json:"usagePercentage"`                  // "usagePercentage": 0.082
			PercentageFieldGoalsMade         float32 `json:"percentageFieldGoalsMade"`         // "percentageFieldGoalsMade": 0.182
			PercentageFieldGoalsAttempted    float32 `json:"percentageFieldGoalsAttempted"`    // "percentageFieldGoalsAttempted": 0.102
			PercentageThreePointersMade      float32 `json:"percentageThreePointersMade"`      // "percentageThreePointersMade": 0.2
			PercentageThreePointersAttempted float32 `json:"percentageThreePointersAttempted"` // "percentageThreePointersAttempted": 0.111
			PercentageFreeThrowsMade         float32 `json:"percentageFreeThrowsMade"`         // "percentageFreeThrowsMade": 0
			PercentageFreeThrowsAttempted    float32 `json:"percentageFreeThrowsAttempted"`    // "percentageFreeThrowsAttempted": 0
			PercentageReboundsOffensive      float32 `json:"percentageReboundsOffensive"`      // "percentageReboundsOffensive": 0
			PercentageReboundsDefensive      float32 `json:"percentageReboundsDefensive"`      // "percentageReboundsDefensive": 0
			PercentageReboundsTotal          float32 `json:"percentageReboundsTotal"`          // "percentageReboundsTotal": 0
			PercentageAssists                float32 `json:"percentageAssists"`                // "percentageAssists": 0.214
			PercentageTurnovers              float32 `json:"percentageTurnovers"`              // "percentageTurnovers": 0
			PercentageSteals                 float32 `json:"percentageSteals"`                 // "percentageSteals": 0
			PercentageBlocks                 float32 `json:"percentageBlocks"`                 // "percentageBlocks": 0.5
			PercentageBlocksAllowed          float32 `json:"percentageBlocksAllowed"`          // "percentageBlocksAllowed": 0
			PercentagePersonalFouls          float32 `json:"percentagePersonalFouls"`          // "percentagePersonalFouls": 0.143
			PercentagePersonalFoulsDrawn     float32 `json:"percentagePersonalFoulsDrawn"`     // "percentagePersonalFoulsDrawn": 0
			PercentagePoints                 float32 `json:"percentagePoints"`                 // "percentagePoints": 0.149
		} `json:"statistics"`
	} `json:"players"`
}
