package baseballreference

const (
	URL         string = "https://www.baseball-reference.com"
	MatchupURL  string = URL + "/boxes?month={month}&day={day}&year={year}"
	BoxScoreURL string = URL + "{route}"
	Domain      string = "www.baseball-reference.com"
)
