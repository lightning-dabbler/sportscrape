package sportsreference

const (
	// Base url
	BasketballRefURL string = "https://www.basketball-reference.com"
	// templated matchup url
	BasketballRefMatchupURL string = BasketballRefURL + "/boxscores?month={month}&day={day}&year={year}"
)
