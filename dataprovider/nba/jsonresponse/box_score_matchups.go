package jsonresponse

// BoxScoreMatchupsJSON - Full box score player matchup json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=matchups)
type BoxScoreMatchupsJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				EventID  string               `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreMatchupStats `json:"homeTeam"`
				AwayTeam BoxScoreMatchupStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreMatchupStats struct {
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
		Matchups   []struct {
			PersonID   int64  `json:"personId"`   // "personId": 1641824
			FirstName  string `json:"firstName"`  // "firstName": "Matas"
			FamilyName string `json:"familyName"` // "familyName": "Buzelis"
			JerseyNum  string `json:"jerseyNum"`  // "jerseyNum": "14"
			Statistics struct {
				MatchupMinutes                 string  `json:"matchupMinutes"`                 // "matchupMinutes": "0:41"
				MatchupMinutesSort             float32 `json:"matchupMinutesSort"`             // "matchupMinutesSort": 40.8
				PartialPossessions             float32 `json:"partialPossessions"`             // "partialPossessions": 4.4
				PercentageDefenderTotalTime    float32 `json:"percentageDefenderTotalTime"`    // "percentageDefenderTotalTime": 0.047
				PercentageOffensiveTotalTime   float32 `json:"percentageOffensiveTotalTime"`   // "percentageOffensiveTotalTime": 0.059
				PercentageTotalTimeBothOn      float32 `json:"percentageTotalTimeBothOn"`      // "percentageTotalTimeBothOn": 0.08
				SwitchesOn                     int32   `json:"switchesOn"`                     // "switchesOn": 0
				PlayerPoints                   int32   `json:"playerPoints"`                   // "playerPoints": 3
				TeamPoints                     int32   `json:"teamPoints"`                     // "teamPoints": 11
				MatchupAssists                 int32   `json:"matchupAssists"`                 // "matchupAssists": 0
				MatchupPotentialAssists        int32   `json:"matchupPotentialAssists"`        // "matchupPotentialAssists": 0
				MatchupTurnovers               int32   `json:"matchupTurnovers"`               // "matchupTurnovers": 0
				MatchupBlocks                  int32   `json:"matchupBlocks"`                  // "matchupBlocks": 0
				MatchupFieldGoalsMade          int32   `json:"matchupFieldGoalsMade"`          // "matchupFieldGoalsMade": 1
				MatchupFieldGoalsAttempted     int32   `json:"matchupFieldGoalsAttempted"`     // "matchupFieldGoalsAttempted": 1
				MatchupFieldGoalsPercentage    int32   `json:"matchupFieldGoalsPercentage"`    // "matchupFieldGoalsPercentage": 1
				MatchupThreePointersMade       int32   `json:"matchupThreePointersMade"`       // "matchupThreePointersMade": 1
				MatchupThreePointersAttempted  int32   `json:"matchupThreePointersAttempted"`  // "matchupThreePointersAttempted": 1
				MatchupThreePointersPercentage int32   `json:"matchupThreePointersPercentage"` // "matchupThreePointersPercentage": 1
				HelpBlocks                     int32   `json:"helpBlocks"`                     // "helpBlocks": 0
				HelpFieldGoalsMade             int32   `json:"helpFieldGoalsMade"`             // "helpFieldGoalsMade": 0
				HelpFieldGoalsAttempted        int32   `json:"helpFieldGoalsAttempted"`        // "helpFieldGoalsAttempted": 0
				HelpFieldGoalsPercentage       int32   `json:"helpFieldGoalsPercentage"`       // "helpFieldGoalsPercentage": 0
				MatchupFreeThrowsMade          int32   `json:"matchupFreeThrowsMade"`          // "matchupFreeThrowsMade": 0
				MatchupFreeThrowsAttempted     int32   `json:"matchupFreeThrowsAttempted"`     // "matchupFreeThrowsAttempted": 0
				ShootingFouls                  int32   `json:"shootingFouls"`                  // "shootingFouls": 0
			} `json:"statistics"`
		} `json:"matchups"`
	} `json:"players"`
}
