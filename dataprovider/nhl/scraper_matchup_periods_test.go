//go:build integration

package nhl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatchupPeriodsScraper(t *testing.T) {
	// https://api-web.nhle.com/v1/gamecenter/2026010045/right-rail
	// NJD @ NYR, 3-4 (SO)
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2026-09-24", 2026010045)
	scraper := NewMatchupPeriodsScraper()
	scraper.Fetcher = testFetcher
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
			Scraper: scraper,
		},
	)
	periods, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	require.Equal(t, 5, len(periods), "5 periods")

	expected := []struct {
		period     int32
		periodType string
		awayScore  int32
		homeScore  int32
		awaySOG    int32
		homeSOG    int32
	}{
		{1, "REG", 1, 0, 5, 4},
		{2, "REG", 1, 0, 7, 8},
		{3, "REG", 1, 3, 8, 8},
		{4, "OT", 0, 0, 3, 3},
		{5, "SO", 0, 1, 0, 0},
	}
	for i, e := range expected {
		p := periods[i]
		assert.Equal(t, int64(2026010045), p.EventID)
		assert.Equal(t, int32(1), p.GameType)
		assert.Equal(t, int64(1), p.AwayTeamID)
		assert.Equal(t, "Devils", p.AwayTeam)
		assert.Equal(t, "NJD", p.AwayTeamAbbreviation)
		assert.Equal(t, int64(3), p.HomeTeamID)
		assert.Equal(t, "Rangers", p.HomeTeam)
		assert.Equal(t, "NYR", p.HomeTeamAbbreviation)
		assert.Equal(t, e.period, p.Period)
		assert.Equal(t, e.periodType, p.PeriodType)
		require.NotNil(t, p.AwayTeamScore)
		require.NotNil(t, p.HomeTeamScore)
		require.NotNil(t, p.AwayTeamSOG)
		require.NotNil(t, p.HomeTeamSOG)
		assert.Equal(t, e.awayScore, *p.AwayTeamScore, "period %d away score", e.period)
		assert.Equal(t, e.homeScore, *p.HomeTeamScore, "period %d home score", e.period)
		assert.Equal(t, e.awaySOG, *p.AwayTeamSOG, "period %d away sog", e.period)
		assert.Equal(t, e.homeSOG, *p.HomeTeamSOG, "period %d home sog", e.period)
	}
}
