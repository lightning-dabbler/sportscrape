package jsonresponse

// BoxScoreTraditionalJSON - Full box score player traditional json response by period (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=traditional&period=All)
type BoxScoreTraditionalJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				EventID  string                   `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreTraditionalStats `json:"homeTeam"`
				AwayTeam BoxScoreTraditionalStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreTraditionalStats struct {
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
			Minutes                 string  `json:"minutes"`                 // "minutes": "28:39"
			FieldGoalsMade          int32   `json:"fieldGoalsMade"`          // "fieldGoalsMade": 4
			FieldGoalsAttempted     int32   `json:"fieldGoalsAttempted"`     // "fieldGoalsAttempted": 6
			FieldGoalsPercentage    float32 `json:"fieldGoalsPercentage"`    // "fieldGoalsPercentage": 0.667
			ThreePointersMade       int32   `json:"threePointersMade"`       // "threePointersMade": 2
			ThreePointersAttempted  int32   `json:"threePointersAttempted"`  // "threePointersAttempted": 3
			ThreePointersPercentage float32 `json:"threePointersPercentage"` // "threePointersPercentage": 0.667
			FreeThrowsMade          int32   `json:"freeThrowsMade"`          // "freeThrowsMade": 0
			FreeThrowsAttempted     int32   `json:"freeThrowsAttempted"`     // "freeThrowsAttempted": 0
			FreeThrowsPercentage    float32 `json:"freeThrowsPercentage"`    // "freeThrowsPercentage": 0
			ReboundsOffensive       int32   `json:"reboundsOffensive"`       // "reboundsOffensive": 0
			ReboundsDefensive       int32   `json:"reboundsDefensive"`       // "reboundsDefensive": 0
			ReboundsTotal           int32   `json:"reboundsTotal"`           // "reboundsTotal": 0
			Assists                 int32   `json:"assists"`                 // "assists": 3
			Steals                  int32   `json:"steals"`                  // "steals": 0
			Blocks                  int32   `json:"blocks"`                  // "blocks": 1
			Turnovers               int32   `json:"turnovers"`               // "turnovers": 0
			FoulsPersonal           int32   `json:"foulsPersonal"`           // "foulsPersonal": 2
			Points                  int32   `json:"points"`                  // "points": 10
			PlusMinusPoints         int32   `json:"plusMinusPoints"`         // "plusMinusPoints": -11
		} `json:"statistics"`
	} `json:"players"`
}
