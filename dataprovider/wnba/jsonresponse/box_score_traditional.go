package jsonresponse

// BoxScoreTraditionalJSON - Full box score player traditional json response by period (e.g. https://www.wnba.com/game/phx-vs-chi-1022600224/box-score?type=traditional&period=All)
type BoxScoreTraditionalJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period     int32                    `json:"period"`     //"period": 4
				GameStatus int32                    `json:"gameStatus"` // "gameStatus": 3
				EventID    string                   `json:"gameId"`     // "gameId": "1022600224"
				HomeTeam   BoxScoreTraditionalStats `json:"homeTeam"`
				AwayTeam   BoxScoreTraditionalStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreTraditionalStats struct {
	TeamID      int64  `json:"teamId"`
	TeamName    string `json:"teamName"`
	TeamCity    string `json:"teamCity"`
	TeamTricode string `json:"teamTricode"`
	Players     []struct {
		PersonID   int64  `json:"personId"`
		FirstName  string `json:"firstName"`
		FamilyName string `json:"familyName"`
		Position   string `json:"position"`
		JerseyNum  string `json:"jerseyNum"`
		Statistics struct {
			Minutes                 string  `json:"minutes"`
			FieldGoalsMade          int32   `json:"fieldGoalsMade"`
			FieldGoalsAttempted     int32   `json:"fieldGoalsAttempted"`
			FieldGoalsPercentage    float32 `json:"fieldGoalsPercentage"`
			ThreePointersMade       int32   `json:"threePointersMade"`
			ThreePointersAttempted  int32   `json:"threePointersAttempted"`
			ThreePointersPercentage float32 `json:"threePointersPercentage"`
			FreeThrowsMade          int32   `json:"freeThrowsMade"`
			FreeThrowsAttempted     int32   `json:"freeThrowsAttempted"`
			FreeThrowsPercentage    float32 `json:"freeThrowsPercentage"`
			ReboundsOffensive       int32   `json:"reboundsOffensive"`
			ReboundsDefensive       int32   `json:"reboundsDefensive"`
			ReboundsTotal           int32   `json:"reboundsTotal"`
			Assists                 int32   `json:"assists"`
			Steals                  int32   `json:"steals"`
			Blocks                  int32   `json:"blocks"`
			Turnovers               int32   `json:"turnovers"`
			FoulsPersonal           int32   `json:"foulsPersonal"`
			Points                  int32   `json:"points"`
			PlusMinusPoints         int32   `json:"plusMinusPoints"`
		} `json:"statistics"`
	} `json:"players"`
}
