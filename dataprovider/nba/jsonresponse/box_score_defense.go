package jsonresponse

// BoxScoreDefenseJSON - Full box score player defense json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=defense)
type BoxScoreDefenseJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				EventID  string               `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreDefenseStats `json:"homeTeam"`
				AwayTeam BoxScoreDefenseStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreDefenseStats struct {
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
			MatchupMinutes                string  `json:"matchupMinutes"`                // "matchupMinutes": "10:57",
			PartialPossessions            float32 `json:"partialPossessions"`            // "partialPossessions": 65.7,
			SwitchesOn                    int32   `json:"switchesOn"`                    // "switchesOn": 0,
			PlayerPoints                  int32   `json:"playerPoints"`                  // "playerPoints": 11,
			DefensiveRebounds             int32   `json:"defensiveRebounds"`             // "defensiveRebounds": 0,
			MatchupAssists                int32   `json:"matchupAssists"`                // "matchupAssists": 6,
			MatchupTurnovers              int32   `json:"matchupTurnovers"`              // "matchupTurnovers": 2,
			Steals                        int32   `json:"steals"`                        // "steals": 0,
			Blocks                        int32   `json:"blocks"`                        // "blocks": 1,
			MatchupFieldGoalsMade         int32   `json:"matchupFieldGoalsMade"`         // "matchupFieldGoalsMade": 3,
			MatchupFieldGoalsAttempted    int32   `json:"matchupFieldGoalsAttempted"`    // "matchupFieldGoalsAttempted": 13,
			MatchupFieldGoalPercentage    float32 `json:"matchupFieldGoalPercentage"`    // "matchupFieldGoalPercentage": 0.231,
			MatchupThreePointersMade      int32   `json:"matchupThreePointersMade"`      // "matchupThreePointersMade": 1,
			MatchupThreePointersAttempted int32   `json:"matchupThreePointersAttempted"` // "matchupThreePointersAttempted": 5,
			MatchupThreePointerPercentage float32 `json:"matchupThreePointerPercentage"` // "matchupThreePointerPercentage": 0.2
		} `json:"statistics"`
	} `json:"players"`
}
