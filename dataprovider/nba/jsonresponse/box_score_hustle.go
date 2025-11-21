package jsonresponse

// BoxScoreHustleJSON - Full box score player hustle json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=hustle)
type BoxScoreHustleJSON struct {
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
	} `json:"players"`
}
