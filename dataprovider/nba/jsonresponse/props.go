package jsonresponse

// Each Props(.*)JSON struct represents a json response for a data feed
// These data feeds include traditional, advanced, misc, scoring, usage, fourfactors, tracking, hustle, defense, and matchups box scores as well as play by play.
// example URL template: https://www.nba.com/game/chi-vs-uta-0022500240/{feed}?period={period}&type={type} (the base URL is derivable from ShareURL from MatchupJSON)
// feed options: [box-score, play-by-play]
// period options: [All, Q1, Q2, Q3, Q4, 1stHalf, 2ndHalf, AllOT] (purposely omitting \d{1}OT until necessary)
// type options: [traditional, advanced, misc, scoring, usage, fourfactors, tracking, hustle, defense, matchups]
// No period params necessary for the following types (defaults to the whole game): [tracking, hustle, defense, matchups]
// element selector when document is retrieved: script#__NEXT_DATA__

// PropsBoxScoreMatchupsJSON - Full box score player matchup json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=matchups)
type PropsBoxScoreMatchupsJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				GameID   string               `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreMatchupsTeam `json:"homeTeam"`
				AwayTeam BoxScoreMatchupsTeam `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreMatchupsTeam struct {
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

// PropsPlayByPlayJSON - Full play by play json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/play-by-play?period=All)
type PropsPlayByPlayJSON struct {
	Props struct {
		PageProps struct {
			PlayByPlay struct {
				EventID string `json:"gameId"`
				Actions []struct {
					ActionNumber   int32  `json:"actionNumber"`   // "actionNumber": 8,
					Clock          string `json:"clock"`          // "clock": "PT11M39.00S",
					Period         int32  `json:"period"`         // "period": 1,
					TeamID         int64  `json:"teamId"`         // "teamId": 1610612762,
					TeamTricode    string `json:"teamTricode"`    // "teamTricode": "UTA",
					PersonID       int64  `json:"personId"`       // "personId": 1628374,
					PlayerName     string `json:"playerName"`     // "playerName": "Markkanen",
					PlayerNameI    string `json:"playerNameI"`    // "playerNameI": "L. Markkanen",
					XLegacy        int32  `json:"xLegacy"`        // "xLegacy": 33,
					YLegacy        int32  `json:"yLegacy"`        // "yLegacy": 31,
					ShotDistance   int32  `json:"shotDistance"`   // "shotDistance": 5,
					ShotResult     string `json:"shotResult"`     // "shotResult": "Made",
					IsFieldGoal    int32  `json:"isFieldGoal"`    // "isFieldGoal": 1,
					ScoreHome      string `json:"scoreHome"`      // "scoreHome": "2",
					ScoreAway      string `json:"scoreAway"`      // "scoreAway": "0",
					PointsTotal    int32  `json:"pointsTotal"`    // "pointsTotal": 2,
					Location       string `json:"location"`       // "location": "h",
					Description    string `json:"description"`    // "description": "Markkanen 5' Cutting Layup Shot (2 PTS) (Nurkic 1 AST)",
					ActionType     string `json:"actionType"`     // "actionType": "Made Shot",
					SubType        string `json:"subType"`        // "subType": "Cutting Layup Shot",
					VideoAvailable int32  `json:"videoAvailable"` // "videoAvailable": 1,
					ShotValue      int32  `json:"shotValue"`      // "shotValue": 2,
					ActionID       int32  `json:"actionId"`       // "actionId": 4 (unique identifier for each play)
				} `json:"actions"`
			} `json:"playByPlay"`
		} `json:"pageProps"`
	} `json:"props"`
}
