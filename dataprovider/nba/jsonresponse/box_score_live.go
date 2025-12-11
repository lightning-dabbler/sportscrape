package jsonresponse

// BoxScoreLiveJSON - Full box score player live json response (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score)
type BoxScoreLiveJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				GameStatus int32             `json:"gameStatus"` // "gameStatus": 2
				EventID    string            `json:"gameId"`     // "gameId": "0022501204"
				HomeTeam   BoxScoreLiveStats `json:"homeTeam"`
				AwayTeam   BoxScoreLiveStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreLiveStats struct {
	TeamID      int64  `json:"teamId"`      // "teamId": 1610612762
	TeamName    string `json:"teamName"`    // "teamName": "Jazz"
	TeamCity    string `json:"teamCity"`    // "teamCity": "Utah"
	TeamTricode string `json:"teamTricode"` // "teamTricode": "UTA"
	Players     []struct {
		Status     string `json:"status"`     // "status": "ACTIVE"
		PersonID   int64  `json:"personId"`   // "personId": 1629004
		FirstName  string `json:"firstName"`  // "firstName": "Svi"
		FamilyName string `json:"familyName"` // "familyName": "Mykhailiuk"
		Position   string `json:"position"`   // "position": "F"
		JerseyNum  string `json:"jerseyNum"`  // "jerseyNum": "10"
		Statistics struct {
			Assists                 int32   `json:"assists"`                 // "assists": 1
			Blocks                  int32   `json:"blocks"`                  // "blocks": 1
			BlocksReceived          int32   `json:"blocksReceived"`          // "blocksReceived": 0
			FieldGoalsAttempted     int32   `json:"fieldGoalsAttempted"`     // "fieldGoalsAttempted": 8
			FieldGoalsMade          int32   `json:"fieldGoalsMade"`          // "fieldGoalsMade": 4
			FieldGoalsPercentage    float32 `json:"fieldGoalsPercentage"`    // "fieldGoalsPercentage": 0.5
			FoulsOffensive          int32   `json:"foulsOffensive"`          // "foulsOffensive": 0
			FoulsDrawn              int32   `json:"foulsDrawn"`              // "foulsDrawn": 1
			FoulsPersonal           int32   `json:"foulsPersonal"`           // "foulsPersonal": 1
			FoulsTechnical          int32   `json:"foulsTechnical"`          // "foulsTechnical": 0
			FreeThrowsAttempted     int32   `json:"freeThrowsAttempted"`     // "freeThrowsAttempted": 0
			FreeThrowsMade          int32   `json:"freeThrowsMade"`          // "freeThrowsMade": 0
			FreeThrowsPercentage    float32 `json:"freeThrowsPercentage"`    // "freeThrowsPercentage": 0
			Minus                   int32   `json:"minus"`                   // "minus": 36
			Minutes                 string  `json:"minutes"`                 // "minutes": "PT11M41.00S"
			MinutesCalculated       string  `json:"minutesCalculated"`       // "minutesCalculated": "PT12M"
			Plus                    int32   `json:"plus"`                    // "plus": 23
			PlusMinusPoints         int32   `json:"plusMinusPoints"`         // "plusMinusPoints": -13
			Points                  int32   `json:"points"`                  // "points": 8
			PointsFastBreak         int32   `json:"pointsFastBreak"`         // "pointsFastBreak": 2
			PointsInThePaint        int32   `json:"pointsInThePaint"`        // "pointsInThePaint": 6
			PointsSecondChance      int32   `json:"pointsSecondChance"`      // "pointsSecondChance": 0
			ReboundsDefensive       int32   `json:"reboundsDefensive"`       // "reboundsDefensive": 6
			ReboundsOffensive       int32   `json:"reboundsOffensive"`       // "reboundsOffensive": 1
			ReboundsTotal           int32   `json:"reboundsTotal"`           // "reboundsTotal": 7
			Steals                  int32   `json:"steals"`                  // "steals": 0
			ThreePointersAttempted  int32   `json:"threePointersAttempted"`  // "threePointersAttempted": 1
			ThreePointersMade       int32   `json:"threePointersMade"`       // "threePointersMade": 0
			ThreePointersPercentage float32 `json:"threePointersPercentage"` // "threePointersPercentage": 0
			Turnovers               int32   `json:"turnovers"`               // "turnovers": 2
			TwoPointersAttempted    int32   `json:"twoPointersAttempted"`    // "twoPointersAttempted": 7
			TwoPointersMade         int32   `json:"twoPointersMade"`         // "twoPointersMade": 4
			TwoPointersPercentage   float32 `json:"twoPointersPercentage"`   // "twoPointersPercentage": 0.5714285714285711
		} `json:"statistics"`
	} `json:"players"`
}
