package nfl

/*
All api.nfl.com endpoints require a bearer token (401 without one).
An anonymous token is minted with nfl.com's public web client credentials (see auth.go):

	POST https://api.nfl.com/identity/v3/token
	body: {"clientKey": ..., "clientSecret": ..., "deviceId": ..., "deviceInfo": "", "networkType": "other", "nflClaims": null}
	response: {"accessToken": ..., "refreshToken": ..., "expiresIn": <epoch seconds>}
	expiresIn is the mint time + 3600s (matches the JWT exp claim); tokens are re-minted 1 minute before expiry or after a 401
	The credentials can be overridden (always as a pair, otherwise the defaults are used for both) with the
	SPORTSCRAPE_NFL_DOT_COM_CLIENT_KEY and SPORTSCRAPE_NFL_DOT_COM_CLIENT_SECRET environment variables, or
	per Fetcher with TokenProvider = NewTokenProvider(clientKey, clientSecret), which takes precedence.

Rate limiting: no 429 has been observed and responses carry no Retry-After or rate limit headers
(they're served through Varnish with cache-control: max-age=60).

Matchup data:

	The date is resolved to its NFL week, the week's games are fetched and filtered to the ones
	whose kickoff falls on the date in America/New_York (the api's "date" field is the UTC date).

	time (event_time) is the scheduled kickoff from the 2013 season on; for the 2012 season and earlier every game's
	time is a 09:00 UTC placeholder on the game date (the date filter still works, 09:00 UTC is the same date in America/New_York).
	summary.startTime (start_time) is the actual start (null before kickoff and for older seasons, e.g. 2001), and is
	set for the 2012 season despite the placeholder time.

	URL template: https://api.nfl.com/football/v2/weeks/date/{date}
	date = YYYY-MM-DD
	e.g. https://api.nfl.com/football/v2/weeks/date/2025-09-04
	-> {"season": 2025, "seasonType": "REG", "week": 1, "weekType": "REG", "dateBegin": "2025-08-27", "dateEnd": "2025-09-10", ...}

	seasonType: PRE, REG, POST
	weekType: HOF (hall of fame game, PRE week 0), PRE, REG, WC, DIV, CONF, SB
	The Super Bowl week (POST week 4) spans the off-season e.g. dateEnd 2026-07-29
	Dates that aren't part of any week (off-season gaps e.g. 2022-07-27) respond with a 404; MatchupScraper returns no matchups for them

	URL template: https://api.nfl.com/football/v2/experience/weekly-game-details?season={season}&type={seasonType}&week={week}&...
	e.g. https://api.nfl.com/football/v2/experience/weekly-game-details?includeDriveChart=false&includeReplays=false&includeStandings=false&includeTaggedVideos=false&season=2025&type=REG&week=1

	gameType: UNSPECIFIED (pre-season and regular season), AFC_WC, NFC_WC, AFC_DIV, NFC_DIV, AFC_CONF, NFC_CONF, NFC_AFC_SB
	summary.phase e.g. PREGAME, FINAL, FINAL_OVERTIME; summary is null for games that are further out
	status is SCHEDULED even for completed games, so summary.phase is used instead

	Team abbreviations are not part of the game details and come from:
	URL template: https://api.nfl.com/experience/v1/teams?season={season}
	e.g. https://api.nfl.com/experience/v1/teams?season=2025
	Team names are always the franchise's current name, even for past seasons (e.g. Los Angeles Chargers for the
	2012 San Diego Chargers; the api has no historical names), while abbreviations are as of the season (e.g. SD in 2012).

Matchup periods (points per period) and play by play (all drives):

	URL template: https://api.nfl.com/experience/v2/gamedetails/{game_id}?includeDriveChart=true
	e.g. https://api.nfl.com/experience/v2/gamedetails/f5908b6d-311e-11f0-b670-ae1250fadad1?includeDriveChart=true

	driveChart.plays[].quarter: 1-4, 5 = OT1, 6 = OT2 (post season), ..., 0 for deleted plays
	summary.{away,home}Team.score has q1-q4 and a single ot bucket aggregating every overtime period, so the
	per period points come from driveChart.scoringSummaries[] (quarter + running awayScore/homeScore) instead.
	e.g. https://api.nfl.com/experience/v2/gamedetails/10012013-0112-00b3-6e81-cc137a8c006a?includeDriveChart=true
	(BAL @ DEN 2013-01-12, double overtime)

	Play by play is a superset of the "ALL DRIVES" drive chart on the nfl.com game page
	(e.g. https://www.nfl.com/games/falcons-at-packers-2026-reg-3?tab=recap): same drives, plays and order, plus the
	plays the page hides: plays that aren't part of a drive (driveSequence 0), deleted plays, END_QUARTER/END_GAME/UNSPECIFIED
	plays and official timeouts/two-minute warnings (TIMEOUT plays without stats).
	driveChart.plays[].driveSequence: 0 when the play isn't part of a drive (e.g. timeouts, end of quarter)

	Play by play stats: driveChart.plays[].stats[] credits each play's stats to players (gsisPlayerId, personId) or only a
	team (null player fields, e.g. first down rushing). statType is undocumented; the verified codes are decoded in stat_type.go.

Box score (players with stats only):

	URL template: https://api.nfl.com/football/v2/stats/live/player-statistics/{game_id}
	e.g. https://api.nfl.com/football/v2/stats/live/player-statistics/f5908b6d-311e-11f0-b670-ae1250fadad1

	404 for games that are further out; no players before kickoff.
	Player names are abbreviated (e.g. D.Prescott); players are identified by their GSIS ID (e.g. 00-0033077).

	Player full names (no endpoint scoped to the game provides them), looked up by the player's personId:
	URL template: https://api.nfl.com/football/v2/persons/{person_id}
	e.g. https://api.nfl.com/football/v2/persons/32005052-4528-5723-d1b2-96e92ebc1241 -> "displayName": "Dak Prescott"

Injuries (every weekly injury report of a season, the source of the nfl.com game page's injury report):

	URL template: https://api.nfl.com/football/v2/injuries?season={season}&seasonType={seasonType}&limit={limit}&pageToken={pageToken}
	e.g. https://api.nfl.com/football/v2/injuries?limit=500&season=2025&seasonType=REG
	-> {"injuries": [...], "pagination": {"limit": 500, "token": "..."}}

	Paging: the response's pagination.token is passed as the next request's pageToken (passing it as token returns the first page again);
	the last page has no token. week and teamId are ignored; season and seasonType are honored. PRE has no injuries (e.g. 2025).
	There's one entry per player per week (season, seasonType, week, person.id); a player's team can change between weeks.

	injuryStatus (the game status): null, OUT, DOUBTFUL, QUESTIONABLE; injuries[] are its reasons (e.g. Knee, "Not injury related - personal matter")
	practices[] are the reasons on the week's practice report (not tied to a practice day); practiceDays[] the participation per practice day:
	FULL, LIMITED, DIDNOT; practiceStatus is the latest practice day's participation.
	The current week's entries change during the week: practiceDays grow with each practice and injuryStatus/injuries are null until the final report.
*/
