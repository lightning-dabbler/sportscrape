package jsonresponse

// PlayByPlayJSON - Full play by play json response (e.g. https://www.wnba.com/game/dal-vs-ind-1022600254/play-by-play?period=All)
type PlayByPlayJSON struct {
	Props struct {
		PageProps struct {
			PlayByPlay struct {
				EventID string `json:"gameId"`
				Actions []struct {
					ActionNumber   int32   `json:"actionNumber"`
					Clock          string  `json:"clock"`
					Period         int32   `json:"period"`
					TeamID         int64   `json:"teamId"`
					TeamTricode    string  `json:"teamTricode"`
					PersonID       int64   `json:"personId"`
					PlayerName     string  `json:"playerName"`
					PlayerNameI    string  `json:"playerNameI"`
					XLegacy        int32   `json:"xLegacy"`
					YLegacy        int32   `json:"yLegacy"`
					ShotDistance   float32 `json:"shotDistance"`
					ShotResult     string  `json:"shotResult"`
					IsFieldGoal    int32   `json:"isFieldGoal"`
					ScoreHome      string  `json:"scoreHome"`
					ScoreAway      string  `json:"scoreAway"`
					PointsTotal    int32   `json:"pointsTotal"`
					Location       string  `json:"location"`
					Description    string  `json:"description"`
					ActionType     string  `json:"actionType"`
					SubType        string  `json:"subType"`
					VideoAvailable int32   `json:"videoAvailable"`
					ShotValue      int32   `json:"shotValue"`
					ActionID       int32   `json:"actionId"`
				} `json:"actions"`
			} `json:"playByPlay"`
		} `json:"pageProps"`
	} `json:"props"`
}
