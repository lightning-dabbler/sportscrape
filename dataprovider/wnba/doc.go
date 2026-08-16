package wnba

/*
Matchup data:

	URL template: https://www.wnba.com/api/schedule?season={year}&regionId=1
	season = YYYY (year component of the requested Date), REQUIRED by the API
	regionId = 1 (broadcaster localization only; kept fixed)
	e.g. https://www.wnba.com/api/schedule?season=2026&regionId=1

	IMPORTANT: this endpoint always returns the full season's gameDates[],
	regardless of any `date` query param (which the API silently ignores).
	The scraper fetches the whole season once per requested year and filters
	gameDates[] client-side to the requested Date/EndDate range. This is a
	plain JSON REST call (net/http, no headless browser) via
	scraper.BaseJsonScraper, unlike NBA's __NEXT_DATA__/chromedp matchup
	scraper.

	ShareURL is not present in the schedule JSON; it is derived as:
	https://www.wnba.com/game/{away_tricode_lower}-vs-{home_tricode_lower}-{gameId}

Event data (box scores):

	Each BoxScore(.*)JSON struct represents a json response for a box score data feed.
	Structurally identical to NBA's mechanism: __NEXT_DATA__ SSR JSON scraped via chromedp.
	example URL template: https://www.wnba.com/game/phx-vs-chi-1022600224/box-score?period={period}&type={type} (the base URL is derivable from ShareURL from MatchupJSON)
	period options: [All, Q1, Q2, Q3, Q4, 1stHalf, 2ndHalf, AllOT]
	box score type options currently supported: [traditional, advanced, misc, scoring, usage, fourfactors]

Play by play:

	Sourced the same way as box scores: __NEXT_DATA__ on the game's
	play-by-play page, e.g. https://www.wnba.com/game/dal-vs-ind-1022600254/play-by-play?period=All
	Confirmed field-for-field identical to NBA's actions[] shape (actionNumber,
	clock, period, teamId, teamTricode, personId, playerName, scoreHome,
	scoreAway, description, actionType, subType, etc.), including internal
	consistency (final action's scoreHome/scoreAway match the game's final
	score). Always requests the full game (period=All); no period selection.

Matchup periods:

	Unlike NBA, where MatchupPeriodsScraper is a discovery-phase feed reading
	Periods straight off the same schedule/gameCardFeed page used for
	matchup discovery, WNBA's schedule API (jsonresponse.MatchupJSON) carries
	no period/quarter data at all - confirmed by scanning an entire season's
	response for any period/quarter/linescore key. Instead, per-quarter team
	scores live on the individual game's box-score page, as
	game.homeTeam.periods[]/game.awayTeam.periods[] - a sibling of
	players/statistics on the same __NEXT_DATA__ envelope, present regardless
	of the box-score `type` requested.

	Consequently MatchupPeriodsScraper here is an EventDataScraper (one
	chromedp fetch per game, requires the matchup step first), not a
	BaseMatchupScraper-family feed like NBA's. It internally requests the
	Traditional box-score type purely as a lightweight carrier for the
	periods data and never reads player stats. No period selection is
	exposed - it always returns whatever period breakdown the box-score page
	currently has (partial mid-game, full once Final).

Element selector when document is retrieved (event data only): script#__NEXT_DATA__
The selected element contains the relevant JSON
*/
