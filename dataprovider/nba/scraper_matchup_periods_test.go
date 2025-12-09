//go:build integration

package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMatchupPeriodsScraper(t *testing.T) {
	// https://www.nba.com/games?date=2025-06-05
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	scraper := nba.NewMatchupPeriodsScraper(
		nba.WithMatchupPeriodsDate("2025-06-05"),
		nba.WithMatchupPeriodsTimeout(3*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(scraper),
	)

	records, err := matchuprunner.Run()
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 4, n_records, "4 records")
	testPeriod := matchup[3]
	assert.Equal(t, "0042400401", testPeriod.EventID)
	assert.Equal(t, int32(3), testPeriod.EventStatus)
	assert.Equal(t, "Final", testPeriod.EventStatusText)
	assert.Equal(t, int64(1610612760), testPeriod.HomeTeamID)
	assert.Equal(t, "Thunder", testPeriod.HomeTeam)
	assert.Equal(t, "OKC", testPeriod.HomeTeamAbbreviation)
	assert.Equal(t, int64(1610612754), testPeriod.AwayTeamID)
	assert.Equal(t, "Pacers", testPeriod.AwayTeam)
	assert.Equal(t, "IND", testPeriod.AwayTeamAbbreviation)
	assert.Equal(t, int32(4), testPeriod.Period)
	assert.Equal(t, int32(35), testPeriod.AwayTeamScore)
	assert.Equal(t, int32(25), testPeriod.HomeTeamScore)
	assert.Equal(t, "Playoffs", matchup.SeasonType)
	assert.Equal(t, "2024-25", matchup.SeasonYear)
	assert.Equal(t, "00", matchup.LeagueID)
}
