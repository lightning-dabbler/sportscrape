package jsonresponse

// BoxScoreMiscJSON - box score player misc json response by period (e.g. https://www.wnba.com/game/phx-vs-chi-1022600224/box-score?type=misc&period=All)
type BoxScoreMiscJSON struct {
	Props struct {
		PageProps struct {
			Game struct {
				Period     int32             `json:"period"`
				GameStatus int32             `json:"gameStatus"`
				EventID    string            `json:"gameId"`
				HomeTeam   BoxScoreMiscStats `json:"homeTeam"`
				AwayTeam   BoxScoreMiscStats `json:"awayTeam"`
			} `json:"game"`
		} `json:"pageProps"`
	} `json:"props"`
}

type BoxScoreMiscStats struct {
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
			Minutes               string `json:"minutes"`
			PointsOffTurnovers    int32  `json:"pointsOffTurnovers"`
			PointsSecondChance    int32  `json:"pointsSecondChance"`
			PointsFastBreak       int32  `json:"pointsFastBreak"`
			PointsPaint           int32  `json:"pointsPaint"`
			OppPointsOffTurnovers int32  `json:"oppPointsOffTurnovers"`
			OppPointsSecondChance int32  `json:"oppPointsSecondChance"`
			OppPointsFastBreak    int32  `json:"oppPointsFastBreak"`
			OppPointsPaint        int32  `json:"oppPointsPaint"`
			Blocks                int32  `json:"blocks"`
			BlocksAgainst         int32  `json:"blocksAgainst"`
			FoulsPersonal         int32  `json:"foulsPersonal"`
			FoulsDrawn            int32  `json:"foulsDrawn"`
		} `json:"statistics"`
	} `json:"players"`
}
