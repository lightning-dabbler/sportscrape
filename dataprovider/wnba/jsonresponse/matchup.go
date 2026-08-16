package jsonresponse

// MatchupJSON - WNBA schedule JSON response (e.g. https://www.wnba.com/api/schedule?season=2026&regionId=1)
// Unlike NBA's __NEXT_DATA__/gameCardFeed shape, this is a plain JSON REST
// response. It always contains the entire season's gameDates, regardless of
// any `date` query param (which the API silently ignores).
type MatchupJSON struct {
	LeagueSchedule struct {
		SeasonYear int32  `json:"seasonYear,string"` // "2026"
		LeagueID   string `json:"leagueId"`          // "10"
		GameDates  []struct {
			GameDate string         `json:"gameDate"` // "08/03/2026 00:00:00"
			Games    []ScheduleGame `json:"games"`
		} `json:"gameDates"`
	} `json:"leagueSchedule"`
}

type ScheduleGame struct {
	GameID          string       `json:"gameId"`          // "1022600224"
	GameCode        string       `json:"gameCode"`        // "20260803/PHXCHI"
	SeasonType      string       `json:"seasonType"`      // "Regular Season"
	GameStatus      int32        `json:"gameStatus"`      // 3
	GameStatusText  string       `json:"gameStatusText"`  // "Final"
	GameDateTimeUTC string       `json:"gameDateTimeUTC"` // "2026-08-04T01:00:00Z"
	HomeTeam        ScheduleTeam `json:"homeTeam"`
	AwayTeam        ScheduleTeam `json:"awayTeam"`
}

type ScheduleTeam struct {
	TeamID      int64  `json:"teamId"`
	TeamName    string `json:"teamName"`    // "Sky"
	TeamCity    string `json:"teamCity"`    // "Chicago"
	TeamTricode string `json:"teamTricode"` // "CHI"
	Wins        int32  `json:"wins"`
	Losses      int32  `json:"losses"`
	Score       int32  `json:"score"`
}
