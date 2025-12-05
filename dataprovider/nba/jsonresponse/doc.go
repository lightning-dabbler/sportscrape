package jsonresponse

/*
Matchup data:

	URL template: https://www.nba.com/games?date={date}
	date = YYYY-MM-DD
	e.g. https://www.nba.com/games?date=2025-11-19

Event data:

	Each BoxScore(.*)JSON struct represents a json response for a box score data feed
	These data feeds include traditional, advanced, misc, scoring, usage, fourfactors, tracking, hustle, defense, and matchups box scores as well as play by play.
	example URL template: https://www.nba.com/game/chi-vs-uta-0022500240/{feed}?period={period}&type={type} (the base URL is derivable from ShareURL from MatchupJSON)
	feed options: [box-score, play-by-play]
	period options: [All, Q1, Q2, Q3, Q4, 1stHalf, 2ndHalf, AllOT] (purposely omitting \d{1}OT until necessary)
	box score type options: [traditional, advanced, misc, scoring, usage, fourfactors, tracking, hustle, defense, matchups]
	No period params necessary for the following box score types (defaults to the whole game): [tracking, hustle, defense, matchups]
	play by play will only be exported for the whole game

Element selector when document is retrieved: script#__NEXT_DATA__
The selected element contains the relevant JSON
*/
