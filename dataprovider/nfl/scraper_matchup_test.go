//go:build integration

package nfl

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/lightning-dabbler/sportscrape/scraper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testFetcher retries failed fetches after a 3s backoff (a 429's Retry-After takes precedence)
var testFetcher = Fetcher{FetchRetryBackoff: 3 * time.Second}

// retrieveMatchups runs the matchup scraper for date
func retrieveMatchups(t *testing.T, date string) []model.Matchup {
	t.Helper()
	scraper := NewMatchupScraper(WithMatchupDate(date))
	scraper.Fetcher = testFetcher
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: scraper,
		},
	)
	matchups, err := matchuprunner.Run()
	require.NoError(t, err)
	return matchups
}

// retrieveMatchup runs the matchup scraper for date and returns the matchup for eventID
func retrieveMatchup(t *testing.T, date string, eventID string) model.Matchup {
	t.Helper()
	for _, matchup := range retrieveMatchups(t, date) {
		if matchup.EventID == eventID {
			return matchup
		}
	}
	t.Fatalf("event %s not found for %s", eventID, date)
	return model.Matchup{}
}

// https://api.nfl.com/football/v2/stats/live/player-statistics/f5908b6d-311e-11f0-b670-ae1250fadad1
// DAL @ PHI
const (
	dalPhiDate    = "2025-09-04"
	dalPhiEventID = "f5908b6d-311e-11f0-b670-ae1250fadad1"
	dalTeamID     = "10401200-a308-98ca-ad5f-95df2fefea68"
	phiTeamID     = "10403700-b939-3cbd-3d16-24d4d6742fa2"
)

// https://api.nfl.com/experience/v2/gamedetails/10012015-0125-0020-8b38-7e57a77e8e95
// NFC Pro Bowl Team @ AFC Pro Bowl Team: the matchup uses the AFC and NFC Pro Bowl Team IDs,
// while the player statistics and drive chart use Pro Bowl team IDs (10408712-617b-8f0f-f061-488f2a55b7b5, 10408713-191c-6da2-c8ff-84762c48a270)
const (
	proBowlDate    = "2015-01-25"
	proBowlEventID = "10012015-0125-0020-8b38-7e57a77e8e95"
	afcProBowlID   = "10408600-77a1-8b0f-8f54-99b7e0a7d1b2"
	nfcProBowlID   = "10408700-5035-ce5c-b3dd-20413011bc6a"
)

// runBoxScore runs s against the matchup for eventID on date
func runBoxScore[E any](t *testing.T, s scraper.EventDataScraper[model.Matchup, E], date string, eventID string) []E {
	t.Helper()
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, date, eventID)
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, E]{
			Scraper: s,
		},
	)
	records, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	return records
}

// findPlayer returns the record whose player ID is playerID
func findPlayer[E any](t *testing.T, records []E, playerID func(E) string, id string) E {
	t.Helper()
	for _, record := range records {
		if playerID(record) == id {
			return record
		}
	}
	t.Fatalf("player %s not found", id)
	var zero E
	return zero
}

func TestMatchupScraper(t *testing.T) {
	// https://api.nfl.com/football/v2/weeks/date/2025-09-04
	// https://api.nfl.com/football/v2/experience/weekly-game-details?season=2025&type=REG&week=1&...
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	// DAL @ PHI kicks off 2025-09-04 8:20pm ET (2025-09-05T00:20:00Z); the rest of week 1 is filtered out
	matchups := retrieveMatchups(t, "2025-09-04")
	require.Equal(t, 1, len(matchups), "1 event")
	matchup := matchups[0]
	assert.Equal(t, "f5908b6d-311e-11f0-b670-ae1250fadad1", matchup.EventID)
	require.NotNil(t, matchup.GSISGameID)
	assert.Equal(t, "59843", *matchup.GSISGameID)
	require.NotNil(t, matchup.Slug)
	assert.Equal(t, "cowboys-at-eagles-2025-reg-1", *matchup.Slug)
	assert.Equal(t, time.Date(2025, 9, 5, 0, 20, 0, 0, time.UTC), matchup.EventTime)
	assert.Equal(t, "2025-09-04", matchup.GameDate)
	assert.Equal(t, int32(2025), matchup.Season)
	assert.Equal(t, "REG", matchup.SeasonType)
	assert.Equal(t, int32(1), matchup.Week)
	assert.Equal(t, "REG", matchup.WeekType)
	assert.Equal(t, "UNSPECIFIED", matchup.GameType)
	assert.Nil(t, matchup.Category)
	assert.False(t, matchup.NeutralSite)
	assert.False(t, matchup.International)
	require.NotNil(t, matchup.Venue)
	assert.Equal(t, "Lincoln Financial Field", *matchup.Venue)
	require.NotNil(t, matchup.VenueCity)
	assert.Equal(t, "Philadelphia", *matchup.VenueCity)
	require.NotNil(t, matchup.Phase)
	assert.Equal(t, "FINAL", *matchup.Phase)
	require.NotNil(t, matchup.Quarter)
	assert.Equal(t, "END_OF_GAME", *matchup.Quarter)
	require.NotNil(t, matchup.StartTime)
	assert.Equal(t, time.Date(2025, 9, 5, 0, 23, 1, 813000000, time.UTC), *matchup.StartTime)
	require.NotNil(t, matchup.Attendance)
	assert.Equal(t, int32(69879), *matchup.Attendance)
	require.NotNil(t, matchup.Weather)
	assert.Equal(t, "Rain Temp: 75° F, Humidity: 66%, Wind: S 11 mph", *matchup.Weather)
	assert.Equal(t, "10403700-b939-3cbd-3d16-24d4d6742fa2", matchup.HomeTeamID)
	assert.Equal(t, "Philadelphia Eagles", matchup.HomeTeam)
	assert.Equal(t, "PHI", matchup.HomeTeamAbbreviation)
	require.NotNil(t, matchup.HomeTeamScore)
	assert.Equal(t, int32(24), *matchup.HomeTeamScore)
	assert.Equal(t, "10401200-a308-98ca-ad5f-95df2fefea68", matchup.AwayTeamID)
	assert.Equal(t, "Dallas Cowboys", matchup.AwayTeam)
	assert.Equal(t, "DAL", matchup.AwayTeamAbbreviation)
	require.NotNil(t, matchup.AwayTeamScore)
	assert.Equal(t, int32(20), *matchup.AwayTeamScore)
}

func TestMatchupScraperSuperBowl(t *testing.T) {
	// https://api.nfl.com/football/v2/weeks/date/2026-02-08 (2025 season, POST week 4)
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchups := retrieveMatchups(t, "2026-02-08")
	require.Equal(t, 1, len(matchups), "1 event")
	matchup := matchups[0]
	assert.Equal(t, int32(2025), matchup.Season)
	assert.Equal(t, "POST", matchup.SeasonType)
	assert.Equal(t, int32(4), matchup.Week)
	assert.Equal(t, "SB", matchup.WeekType)
	assert.Equal(t, "NFC_AFC_SB", matchup.GameType)
	assert.True(t, matchup.NeutralSite)
	assert.Equal(t, "SEA", matchup.AwayTeamAbbreviation)
	assert.Equal(t, "NE", matchup.HomeTeamAbbreviation)
	require.NotNil(t, matchup.AwayTeamScore)
	assert.Equal(t, int32(29), *matchup.AwayTeamScore)
	require.NotNil(t, matchup.HomeTeamScore)
	assert.Equal(t, int32(13), *matchup.HomeTeamScore)
}

func TestMatchupScraperHallOfFameGame(t *testing.T) {
	// https://api.nfl.com/football/v2/weeks/date/2025-07-31 (PRE week 0)
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchups := retrieveMatchups(t, "2025-07-31")
	require.Equal(t, 1, len(matchups), "1 event")
	matchup := matchups[0]
	assert.Equal(t, "PRE", matchup.SeasonType)
	assert.Equal(t, int32(0), matchup.Week)
	assert.Equal(t, "HOF", matchup.WeekType)
	assert.Equal(t, "LAC", matchup.AwayTeamAbbreviation)
	assert.Equal(t, "DET", matchup.HomeTeamAbbreviation)
}

func TestMatchupScraperNoGames(t *testing.T) {
	// A Wednesday between weeks 1 and 2 of the 2025 season, a date in the off-season (part of the Super Bowl week),
	// and an off-season date that isn't part of any week (https://api.nfl.com/football/v2/weeks/date/2022-07-27 responds with a 404)
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	for _, date := range []string{"2025-09-10", "2026-06-15", "2022-07-27"} {
		assert.Equal(t, 0, len(retrieveMatchups(t, date)), "0 events on %s", date)
	}
}

func TestMatchupScraperPostSeason(t *testing.T) {
	// https://api.nfl.com/football/v2/weeks/date/2026-01-17 (2025 season, POST week 2, divisional round)
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	// SF @ SEA kicks off 2026-01-17 8pm ET (2026-01-18T01:00:00Z)
	matchups := retrieveMatchups(t, "2026-01-17")
	require.Equal(t, 2, len(matchups), "2 events")
	expected := []struct {
		eventID   string
		eventTime time.Time
		gameType  string
		phase     string
		away      string
		home      string
		awayScore int32
		homeScore int32
	}{
		{"1ab6b0a3-f037-11f0-9442-5911216651e2", time.Date(2026, 1, 17, 21, 30, 0, 0, time.UTC), "AFC_DIV", "FINAL_OVERTIME", "BUF", "DEN", 30, 33},
		{"5b3fe531-f037-11f0-9442-5911216651e2", time.Date(2026, 1, 18, 1, 0, 0, 0, time.UTC), "NFC_DIV", "FINAL", "SF", "SEA", 6, 41},
	}
	for i, e := range expected {
		matchup := matchups[i]
		assert.Equal(t, e.eventID, matchup.EventID)
		assert.Equal(t, e.eventTime, matchup.EventTime)
		assert.Equal(t, "2026-01-17", matchup.GameDate)
		assert.Equal(t, int32(2025), matchup.Season)
		assert.Equal(t, "POST", matchup.SeasonType)
		assert.Equal(t, int32(2), matchup.Week)
		assert.Equal(t, "DIV", matchup.WeekType)
		assert.Equal(t, e.gameType, matchup.GameType)
		require.NotNil(t, matchup.Phase)
		assert.Equal(t, e.phase, *matchup.Phase)
		assert.Equal(t, e.away, matchup.AwayTeamAbbreviation)
		assert.Equal(t, e.home, matchup.HomeTeamAbbreviation)
		require.NotNil(t, matchup.AwayTeamScore)
		assert.Equal(t, e.awayScore, *matchup.AwayTeamScore)
		require.NotNil(t, matchup.HomeTeamScore)
		assert.Equal(t, e.homeScore, *matchup.HomeTeamScore)
	}
}

func TestMatchupScraperPlaceholderEventTime(t *testing.T) {
	// https://api.nfl.com/football/v2/experience/weekly-game-details?season=2012&type=REG&week=1&...
	// the 2012 season and earlier have a 09:00 UTC placeholder time; startTime is the actual start
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	// DAL @ NYG started 2012-09-05 8:41pm ET
	matchups := retrieveMatchups(t, "2012-09-05")
	require.Equal(t, 1, len(matchups), "1 event")
	matchup := matchups[0]
	assert.Equal(t, "10012012-0905-0027-c93d-2a27a30b22e7", matchup.EventID)
	assert.Equal(t, time.Date(2012, 9, 5, 9, 0, 0, 0, time.UTC), matchup.EventTime)
	assert.Equal(t, "2012-09-05", matchup.GameDate)
	require.NotNil(t, matchup.StartTime)
	assert.Equal(t, time.Date(2012, 9, 6, 0, 41, 22, 0, time.UTC), *matchup.StartTime)

	// CIN @ BAL has no startTime, SD @ OAK has one
	startTimes := make(map[string]*time.Time)
	for _, matchup := range retrieveMatchups(t, "2012-09-10") {
		startTimes[matchup.EventID] = matchup.StartTime
	}
	require.Equal(t, 2, len(startTimes), "2 events")
	assert.Nil(t, startTimes["10012012-0910-00ef-37dc-71dd49beb289"])
	require.NotNil(t, startTimes["10012012-0910-0194-6297-c3b0a8a8ed29"])
	assert.Equal(t, time.Date(2012, 9, 11, 2, 25, 26, 0, time.UTC), *startTimes["10012012-0910-0194-6297-c3b0a8a8ed29"])
}
