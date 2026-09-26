//go:build integration

package nhl

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testFetcher retries failed fetches after a 3s backoff (a 429's Retry-After takes precedence)
var testFetcher = Fetcher{FetchRetryBackoff: 3 * time.Second}

// retrieveMatchup runs the matchup scraper for date and returns the matchup for eventID
func retrieveMatchup(t *testing.T, date string, eventID int64) model.Matchup {
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
	for _, matchup := range matchups {
		if matchup.EventID == eventID {
			return matchup
		}
	}
	t.Fatalf("event %d not found for %s", eventID, date)
	return model.Matchup{}
}

func TestMatchupScraper(t *testing.T) {
	// https://api-web.nhle.com/v1/score/2024-11-12
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	scraper := NewMatchupScraper(WithMatchupDate("2024-11-12"))
	scraper.Fetcher = testFetcher
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: scraper,
		},
	)
	matchups, err := matchuprunner.Run()
	require.NoError(t, err)
	assert.Equal(t, 7, len(matchups), "7 events")

	var matchup *model.Matchup
	for i := range matchups {
		if matchups[i].EventID == int64(2024020250) {
			matchup = &matchups[i]
		}
	}
	require.NotNil(t, matchup, "event 2024020250 should exist")
	assert.Equal(t, time.Date(2024, 11, 13, 3, 0, 0, 0, time.UTC), matchup.EventTime)
	assert.Equal(t, "2024-11-12", matchup.GameDate)
	assert.Equal(t, int64(20242025), matchup.Season)
	assert.Equal(t, int32(2), matchup.GameType)
	assert.Equal(t, "OFF", matchup.GameState)
	assert.Equal(t, "OK", matchup.GameScheduleState)
	assert.Equal(t, "Rogers Arena", matchup.Venue)
	assert.Equal(t, int64(23), matchup.HomeTeamID)
	assert.Equal(t, "Canucks", matchup.HomeTeam)
	assert.Equal(t, "VAN", matchup.HomeTeamAbbreviation)
	assert.Equal(t, int32(3), *matchup.HomeTeamScore)
	assert.Equal(t, int32(32), *matchup.HomeTeamSOG)
	assert.Equal(t, int64(20), matchup.AwayTeamID)
	assert.Equal(t, "Flames", matchup.AwayTeam)
	assert.Equal(t, "CGY", matchup.AwayTeamAbbreviation)
	assert.Equal(t, int32(1), *matchup.AwayTeamScore)
	assert.Equal(t, int32(29), *matchup.AwayTeamSOG)
	assert.Equal(t, int32(3), *matchup.Period)
	assert.Equal(t, "REG", *matchup.LastPeriodType)
}

func TestMatchupScraperNoGames(t *testing.T) {
	// https://api-web.nhle.com/v1/score/2026-07-15
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	scraper := NewMatchupScraper(WithMatchupDate("2026-07-15"))
	scraper.Fetcher = testFetcher
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: scraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(matchups), "0 events")
}
