package nhl

/*
Matchup data:

	URL template: https://api-web.nhle.com/v1/score/{date}
	date = YYYY-MM-DD
	e.g. https://api-web.nhle.com/v1/score/2024-11-12

	gameType:
	- 1 = pre-season
	- 2 = regular season
	- 3 = post season
	other values occur e.g. 19 for event 2024190001 (https://api-web.nhle.com/v1/score/2025-02-12)

Matchup periods (linescore and shots on goal per period):

	URL template: https://api-web.nhle.com/v1/gamecenter/{game_id}/right-rail
	e.g. https://api-web.nhle.com/v1/gamecenter/2026010045/right-rail

Box score (skaters i.e. forwards + defense, goalies):

	URL template: https://api-web.nhle.com/v1/gamecenter/{game_id}/boxscore
	e.g. https://api-web.nhle.com/v1/gamecenter/2024020250/boxscore

	Box score player names are abbreviated (e.g. J. Huberdeau), so each player's
	first and last name is taken from the game's play-by-play rosterSpots
	(one request per game, see Play by play below).

Play by play:

	URL template: https://api-web.nhle.com/v1/gamecenter/{game_id}/play-by-play
	e.g. https://api-web.nhle.com/v1/gamecenter/2024020250/play-by-play
*/
