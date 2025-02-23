package basketballreference

const (
	// Base url
	URL string = "https://www.basketball-reference.com"
	// templated matchup url
	MatchupURL string = URL + "/boxscores?month={month}&day={day}&year={year}"
	// templated box scores url
	BoxScoreURL string = URL + "/boxscores/{event_id}.html"
)
