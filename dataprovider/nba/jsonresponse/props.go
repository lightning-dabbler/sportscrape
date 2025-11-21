package jsonresponse

// Each Props(.*)JSON struct represents a json response for a data feed
// These data feeds include traditional, advanced, misc, scoring, usage, fourfactors, tracking, hustle, defense, and matchups box scores as well as play by play.
// example URL template: https://www.nba.com/game/chi-vs-uta-0022500240/{feed}?period={period}&type={type} (the base URL is derivable from ShareURL from MatchupJSON)
// feed options: [box-score, play-by-play]
// period options: [All, Q1, Q2, Q3, Q4, 1stHalf, 2ndHalf, AllOT] (purposely omitting \d{1}OT until necessary)
// type options: [traditional, advanced, misc, scoring, usage, fourfactors, tracking, hustle, defense, matchups]
// No period params necessary for the following types (defaults to the whole game): [tracking, hustle, defense, matchups]
// element selector when document is retrieved: script#__NEXT_DATA__

// PropsBoxScoreTrackingJSON - Full box score player tracing json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=tracking)
type PropsBoxScoreTrackingJSON struct {
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

// PropsBoxScoreHustleJSON - Full box score player hustle json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=hustle)
type PropsBoxScoreHustleJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				EventID  string              `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreHustleStats `json:"homeTeam"`
				AwayTeam BoxScoreHustleStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreHustleStats struct {
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
			Minutes                      string `json:"minutes"`                      // "minutes": "38:02",
			Points                       int32  `json:"points"`                       // "points": 5,
			ContestedShots               int32  `json:"contestedShots"`               // "contestedShots": 11,
			ContestedShots2pt            int32  `json:"contestedShots2pt"`            // "contestedShots2pt": 4,
			ContestedShots3pt            int32  `json:"contestedShots3pt"`            // "contestedShots3pt": 7,
			Deflections                  int32  `json:"deflections"`                  // "deflections": 5,
			ChargesDrawn                 int32  `json:"chargesDrawn"`                 // "chargesDrawn": 0,
			ScreenAssists                int32  `json:"screenAssists"`                // "screenAssists": 5,
			ScreenAssistPoints           int32  `json:"screenAssistPoints"`           // "screenAssistPoints": 14,
			LooseBallsRecoveredOffensive int32  `json:"looseBallsRecoveredOffensive"` // "looseBallsRecoveredOffensive": 0,
			LooseBallsRecoveredDefensive int32  `json:"looseBallsRecoveredDefensive"` // "looseBallsRecoveredDefensive": 0,
			LooseBallsRecoveredTotal     int32  `json:"looseBallsRecoveredTotal"`     // "looseBallsRecoveredTotal": 0,
			OffensiveBoxOuts             int32  `json:"offensiveBoxOuts"`             // "offensiveBoxOuts": 0,
			DefensiveBoxOuts             int32  `json:"defensiveBoxOuts"`             // "defensiveBoxOuts": 0,
			BoxOutPlayerTeamRebounds     int32  `json:"boxOutPlayerTeamRebounds"`     // "boxOutPlayerTeamRebounds": 0,
			BoxOutPlayerRebounds         int32  `json:"boxOutPlayerRebounds"`         // "boxOutPlayerRebounds": 0,
			BoxOuts                      int32  `json:"boxOuts"`                      // "boxOuts": 0
		} `json:"statistics"`
	}
}

// PropsBoxScoreDefenseJSON - Full box score player defense json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=defense)
type PropsBoxScoreDefenseJSON struct {
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
	}
}

// PropsBoxScoreMatchupsJSON - Full box score player matchup json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=matchups)
type PropsBoxScoreMatchupsJSON struct {
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
