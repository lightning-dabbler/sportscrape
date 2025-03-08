package baseballreference

const (
	URI         string = "https://www.baseball-reference.com"
	MatchupURL  string = URI + "/boxes?month={month}&day={day}&year={year}"
	BoxScoreURL string = URI + "{route}"
	Domain      string = "www.baseball-reference.com"
)
