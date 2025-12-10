package jsonresponse

// BoxScoreAdvancedJSON - Full box score player traditional json response by period (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=advanced&period=All)
type BoxScoreAdvancedJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				EventID  string                `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreAdvancedStats `json:"homeTeam"`
				AwayTeam BoxScoreAdvancedStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreAdvancedStats struct {
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
			Minutes                      string  `json:"minutes"`                      // "minutes": "28:39"
			EstimatedOffensiveRating     float32 `json:"estimatedOffensiveRating"`     // "estimatedOffensiveRating": 104.7
			OffensiveRating              float32 `json:"offensiveRating"`              // "offensiveRating": 104.7
			EstimatedDefensiveRating     float32 `json:"estimatedDefensiveRating"`     // "estimatedDefensiveRating": 114.7
			DefensiveRating              float32 `json:"defensiveRating"`              // "defensiveRating": 114.7
			EstimatedNetRating           float32 `json:"estimatedNetRating"`           // "estimatedNetRating": -10
			NetRating                    float32 `json:"netRating"`                    // "netRating": -10
			AssistPercentage             float32 `json:"assistPercentage"`             // "assistPercentage": 0.167
			AssistToTurnover             float32 `json:"assistToTurnover"`             // "assistToTurnover": 0
			AssistRatio                  float32 `json:"assistRatio"`                  // "assistRatio": 33.3
			OffensiveReboundPercentage   float32 `json:"offensiveReboundPercentage"`   // "offensiveReboundPercentage": 0
			DefensiveReboundPercentage   float32 `json:"defensiveReboundPercentage"`   // "defensiveReboundPercentage": 0
			ReboundPercentage            float32 `json:"reboundPercentage"`            // "reboundPercentage": 0
			TurnoverRatio                float32 `json:"turnoverRatio"`                // "turnoverRatio": 0
			EffectiveFieldGoalPercentage float32 `json:"effectiveFieldGoalPercentage"` // "effectiveFieldGoalPercentage": 0.833
			TrueShootingPercentage       float32 `json:"trueShootingPercentage"`       // "trueShootingPercentage": 0.833
			UsagePercentage              float32 `json:"usagePercentage"`              // "usagePercentage": 0.082
			EstimatedUsagePercentage     float32 `json:"estimatedUsagePercentage"`     // "estimatedUsagePercentage": 0.082
			EstimatedPace                float32 `json:"estimatedPace"`                // "estimatedPace": 110.6
			Pace                         float32 `json:"pace"`                         // "pace": 110.6
			PacePer40                    float32 `json:"pacePer40"`                    // "pacePer40": 92.17
			Possessions                  int32   `json:"possessions"`                  // "possessions": 64
			PIE                          float32 `json:"PIE"`                          // "PIE": 0.069
		} `json:"statistics"`
	} `json:"players"`
}
