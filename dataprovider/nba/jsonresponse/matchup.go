package jsonresponse

// MatchupJSON - Matchup details JSON response (e.g. - https://www.nba.com/games?date=2025-11-19)
type MatchupJSON struct {
	Props struct {
		PageProps struct {
			GameCardFeed struct {
				Modules []struct {
					Cards []struct {
						CardData struct {
							EventID        string          `json:"gameId"` // "0022500249"
							AwayTeam       MatchupCardTeam `json:"awayTeam"`
							HomeTeam       MatchupCardTeam `json:"homeTeam"`
							EventTime      string          `json:"gameTimeUtc"`    // "2025-11-19T00:00:00Z"
							SeasonType     string          `json:"seasonType"`     // "Regular Season"
							SeasonYear     string          `json:"seasonYear"`     // "2025-26"
							ShareUrl       string          `json:"shareUrl"`       // "https://www.nba.com/game/gsw-vs-orl-0022500249"
							LeagueID       string          `json:"leagueId"`       // "00"
							GameStatus     int32           `json:"gameStatus"`     // 3
							GameStatusText string          `json:"gameStatusText"` // "Final"
						} `json:"cardData"`
					} `json:"cards"`
				} `json:"modules"`
			} `json:"gameCardFeed"`
		} `json:"pageProps"`
	} `json:"props"`
}

type MatchupCardTeam struct {
	TeamID      int64  `json:"teamId"`
	TeamName    string `json:"teamName"`    // "Magic"
	TeamTricode string `json:"teamTricode"` // "ORL"
	Score       int32  `json:"score"`
	Losses      int32  `json:"losses"`
	Wins        int32  `json:"wins"`
	Periods     []struct {
		Period int32 `json:"period"`
		Score  int32 `json:"score"`
	} `json:"periods"`
}
