package jsonresponse

// BoxScoreTrackingJSON - Full box score player tracing json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=tracking)
type BoxScoreTrackingJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				EventID  string                `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreTrackingStats `json:"homeTeam"`
				AwayTeam BoxScoreTrackingStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreTrackingStats struct {
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
			Minutes                          string  `json:"minutes"`                          // "minutes": "28:39",
			Speed                            float32 `json:"speed"`                            // "speed": 4.7,
			Distance                         float32 `json:"distance"`                         // "distance": 2.11,
			ReboundChancesOffensive          int32   `json:"reboundChancesOffensive"`          // "reboundChancesOffensive": 0,
			ReboundChancesDefensive          int32   `json:"reboundChancesDefensive"`          // "reboundChancesDefensive": 6,
			ReboundChancesTotal              int32   `json:"reboundChancesTotal"`              // "reboundChancesTotal": 6,
			Touches                          int32   `json:"touches"`                          // "touches": 37,
			SecondaryAssists                 int32   `json:"secondaryAssists"`                 // "secondaryAssists": 0,
			FreeThrowAssists                 int32   `json:"freeThrowAssists"`                 // "freeThrowAssists": 1,
			Passes                           int32   `json:"passes"`                           // "passes": 31,
			Assists                          int32   `json:"assists"`                          // "assists": 3,
			ContestedFieldGoalsMade          int32   `json:"contestedFieldGoalsMade"`          // "contestedFieldGoalsMade": 1,
			ContestedFieldGoalsAttempted     int32   `json:"contestedFieldGoalsAttempted"`     // "contestedFieldGoalsAttempted": 1,
			ContestedFieldGoalPercentage     int32   `json:"contestedFieldGoalPercentage"`     // "contestedFieldGoalPercentage": 1,
			UncontestedFieldGoalsMade        int32   `json:"uncontestedFieldGoalsMade"`        // "uncontestedFieldGoalsMade": 3,
			UncontestedFieldGoalsAttempted   int32   `json:"uncontestedFieldGoalsAttempted"`   // "uncontestedFieldGoalsAttempted": 5,
			UncontestedFieldGoalsPercentage  float32 `json:"uncontestedFieldGoalsPercentage"`  // "uncontestedFieldGoalsPercentage": 0.6,
			FieldGoalPercentage              float32 `json:"fieldGoalPercentage"`              // "fieldGoalPercentage": 0.667,
			DefendedAtRimFieldGoalsMade      int32   `json:"defendedAtRimFieldGoalsMade"`      // "defendedAtRimFieldGoalsMade": 1,
			DefendedAtRimFieldGoalsAttempted int32   `json:"defendedAtRimFieldGoalsAttempted"` // "defendedAtRimFieldGoalsAttempted": 5,
			DefendedAtRimFieldGoalPercentage float32 `json:"defendedAtRimFieldGoalPercentage"` // "defendedAtRimFieldGoalPercentage": 0.2
		} `json:"statistics"`
	}
}
