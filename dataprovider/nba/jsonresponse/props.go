package jsonresponse

// e.g. https://www.nba.com/game/chi-vs-uta-0022500240?period=All (the uri is derivable from ShareURL in MatchupJSON)
// Primarily using this to retrieve play by play data for all periods
// selector: script#__NEXT_DATA__
type PropsJSON struct {
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
