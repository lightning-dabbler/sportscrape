package jsonresponse

// BoxScoreMiscJSON - box score player misc json response by period (e.g. https://www.nba.com/game/chi-vs-uta-0022500240/box-score?type=misc&period=All)
type BoxScoreMiscJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period   int32             `json:"period"` //"period": 4
				EventID  string            `json:"gameId"` // "gameId": "0022500240"
				HomeTeam BoxScoreMiscStats `json:"homeTeam"`
				AwayTeam BoxScoreMiscStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreMiscStats struct {
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
			Minutes               string `json:"minutes"`               // "minutes": "28:39"
			PointsOffTurnovers    int32  `json:"pointsOffTurnovers"`    // "pointsOffTurnovers": 0
			PointsSecondChance    int32  `json:"pointsSecondChance"`    // "pointsSecondChance": 3
			PointsFastBreak       int32  `json:"pointsFastBreak"`       // "pointsFastBreak": 3
			PointsPaint           int32  `json:"pointsPaint"`           // "pointsPaint": 2
			OppPointsOffTurnovers int32  `json:"oppPointsOffTurnovers"` // "oppPointsOffTurnovers": 7
			OppPointsSecondChance int32  `json:"oppPointsSecondChance"` // "oppPointsSecondChance": 10
			OppPointsFastBreak    int32  `json:"oppPointsFastBreak"`    // "oppPointsFastBreak": 9
			OppPointsPaint        int32  `json:"oppPointsPaint"`        // "oppPointsPaint": 30
			Blocks                int32  `json:"blocks"`                // "blocks": 1
			BlocksAgainst         int32  `json:"blocksAgainst"`         // "blocksAgainst": 0
			FoulsPersonal         int32  `json:"foulsPersonal"`         // "foulsPersonal": 2
			FoulsDrawn            int32  `json:"foulsDrawn"`            // "foulsDrawn": 0
		} `json:"statistics"`
	} `json:"players"`
}
