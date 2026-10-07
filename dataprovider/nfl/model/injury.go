package model

import "time"

// Injury is a player's entry on a team's weekly injury report scraped from https://api.nfl.com/football/v2/injuries - composite key: season, season_type, week, team_id, player_id
//
// A player traded during a week can be on both teams' reports for that week (e.g. Christian McCaffrey, CAR and SF, 2022 REG week 7).
// The current week's entry changes during the week: PracticeDays grows with each practice and InjuryStatus/Injuries are set by the final report before the game.
// Past weeks are final. Entries without a GSIS ID (a duplicate person record repeating the player's entry e.g. Brock Wright, DET, 2021 REG week 18) are skipped.
type Injury struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// Season e.g. 2025 (a season's post-season weeks are played in the following year)
	Season int32 `json:"season" parquet:"name=season, type=INT32"`
	// SeasonType e.g. PRE, REG, POST
	SeasonType string `json:"season_type" parquet:"name=season_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Week of the season type e.g. 1 (POST weeks are 1-4)
	Week int32 `json:"week" parquet:"name=week, type=INT32"`
	// TeamID is the player's team for the week e.g. 10400610-c40e-a673-1743-2ce2a5d5d731
	TeamID string `json:"team_id" parquet:"name=team_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Team e.g. Buffalo Bills (the franchise's current name, even for past seasons e.g. Los Angeles Chargers for the 2012 San Diego Chargers)
	Team string `json:"team" parquet:"name=team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// TeamAbbreviation e.g. BUF, as of the season (e.g. SD for the 2012 Chargers)
	TeamAbbreviation string `json:"team_abbreviation" parquet:"name=team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PlayerID is the player's GSIS ID e.g. 00-0036162
	PlayerID string `json:"player_id" parquet:"name=player_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PersonID e.g. 32004241-5355-5802-e27b-153ecc1a567e
	PersonID string `json:"person_id" parquet:"name=person_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Player is the player's full name e.g. Tyler Bass
	Player string `json:"player" parquet:"name=player, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Position e.g. K; empty when the api has none (e.g. 2024 PRE week 3)
	Position string `json:"position" parquet:"name=position, type=BYTE_ARRAY, convertedtype=UTF8"`
	// InjuryStatus is the game status: OUT, DOUBTFUL, QUESTIONABLE; nil when the player has no game status (or before the final report)
	InjuryStatus *string `json:"injury_status" parquet:"name=injury_status, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Injuries is the comma separated reason(s) for InjuryStatus as reported e.g. "Hip, Groin", "Not injury related - personal matter"; nil when there are none
	Injuries *string `json:"injuries" parquet:"name=injuries, type=BYTE_ARRAY, convertedtype=UTF8"`
	// Practices is the comma separated reason(s) on the week's practice report as reported (not tied to a practice day)
	// e.g. "Not injury related - resting player, Illness"; nil when there are none
	Practices *string `json:"practices" parquet:"name=practices, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PracticeStatus is the participation on the latest practice day: FULL, LIMITED, DIDNOT; usually nil when there are no practice days (rarely set without them)
	PracticeStatus *string `json:"practice_status" parquet:"name=practice_status, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PracticeDays is the comma separated date:participation of each practice day as reported (oldest first)
	// e.g. "2025-09-03:LIMITED,2025-09-04:DIDNOT,2025-09-05:DIDNOT"; a day without a participation status is rendered as date:
	// e.g. "2024-12-25:FULL,2024-12-26:,2024-12-27:DIDNOT"; nil when there are no practice days
	PracticeDays *string `json:"practice_days" parquet:"name=practice_days, type=BYTE_ARRAY, convertedtype=UTF8"`
}
