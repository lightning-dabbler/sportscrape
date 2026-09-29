//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runMatchupPeriods(t *testing.T, matchup model.Matchup) []model.MatchupPeriods {
	t.Helper()
	scraper := NewMatchupPeriodsScraper()
	scraper.Fetcher = testFetcher
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
			Scraper: scraper,
		},
	)
	periods, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	return periods
}

type expectedPeriod struct {
	period     int32
	periodType string
	awayScore  int32
	homeScore  int32
}

func TestMatchupPeriodsScraper(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/f5908b6d-311e-11f0-b670-ae1250fadad1?includeDriveChart=true
	// DAL @ PHI, 20-24
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2025-09-04", "f5908b6d-311e-11f0-b670-ae1250fadad1")
	periods := runMatchupPeriods(t, matchup)
	require.Equal(t, 4, len(periods), "4 periods")

	expected := []expectedPeriod{
		{1, "REG", 7, 7},
		{2, "REG", 13, 14},
		{3, "REG", 0, 3},
		{4, "REG", 0, 0},
	}
	for i, e := range expected {
		p := periods[i]
		assert.Equal(t, "f5908b6d-311e-11f0-b670-ae1250fadad1", p.EventID)
		assert.Equal(t, "REG", p.SeasonType)
		assert.Equal(t, "FINAL", p.Phase)
		assert.Equal(t, "10401200-a308-98ca-ad5f-95df2fefea68", p.AwayTeamID)
		assert.Equal(t, "Dallas Cowboys", p.AwayTeam)
		assert.Equal(t, "DAL", p.AwayTeamAbbreviation)
		assert.Equal(t, "10403700-b939-3cbd-3d16-24d4d6742fa2", p.HomeTeamID)
		assert.Equal(t, "Philadelphia Eagles", p.HomeTeam)
		assert.Equal(t, "PHI", p.HomeTeamAbbreviation)
		assert.Equal(t, e.period, p.Period)
		assert.Equal(t, e.periodType, p.PeriodType)
		assert.Equal(t, e.awayScore, p.AwayTeamScore, "period %d away score", e.period)
		assert.Equal(t, e.homeScore, p.HomeTeamScore, "period %d home score", e.period)
	}
}

func TestMatchupPeriodsScraperOvertime(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/f6ced5de-311e-11f0-b670-ae1250fadad1?includeDriveChart=true
	// GB @ DAL, 40-40 (OT tie)
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2025-09-28", "f6ced5de-311e-11f0-b670-ae1250fadad1")
	periods := runMatchupPeriods(t, matchup)
	require.Equal(t, 5, len(periods), "5 periods")

	expected := []expectedPeriod{
		{1, "REG", 7, 0},
		{2, "REG", 6, 16},
		{3, "REG", 7, 7},
		{4, "REG", 17, 14},
		{5, "OT", 3, 3},
	}
	for i, e := range expected {
		p := periods[i]
		assert.Equal(t, "FINAL_OVERTIME", p.Phase)
		assert.Equal(t, "GB", p.AwayTeamAbbreviation)
		assert.Equal(t, "DAL", p.HomeTeamAbbreviation)
		assert.Equal(t, e.period, p.Period)
		assert.Equal(t, e.periodType, p.PeriodType)
		assert.Equal(t, e.awayScore, p.AwayTeamScore, "period %d away score", e.period)
		assert.Equal(t, e.homeScore, p.HomeTeamScore, "period %d home score", e.period)
	}
}

func TestMatchupPeriodsScraperPostSeasonOvertime(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/1ab6b0a3-f037-11f0-9442-5911216651e2?includeDriveChart=true
	// BUF @ DEN, 30-33 (OT), AFC divisional round
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2026-01-17", "1ab6b0a3-f037-11f0-9442-5911216651e2")
	periods := runMatchupPeriods(t, matchup)
	require.Equal(t, 5, len(periods), "5 periods")

	expected := []expectedPeriod{
		{1, "REG", 7, 3},
		{2, "REG", 3, 17},
		{3, "REG", 7, 3},
		{4, "REG", 13, 7},
		{5, "OT", 0, 3},
	}
	for i, e := range expected {
		p := periods[i]
		assert.Equal(t, "POST", p.SeasonType)
		assert.Equal(t, "FINAL_OVERTIME", p.Phase)
		assert.Equal(t, "BUF", p.AwayTeamAbbreviation)
		assert.Equal(t, "DEN", p.HomeTeamAbbreviation)
		assert.Equal(t, e.period, p.Period)
		assert.Equal(t, e.periodType, p.PeriodType)
		assert.Equal(t, e.awayScore, p.AwayTeamScore, "period %d away score", e.period)
		assert.Equal(t, e.homeScore, p.HomeTeamScore, "period %d home score", e.period)
	}
}

func TestMatchupPeriodsScraperMultipleOvertimes(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/10012013-0112-00b3-6e81-cc137a8c006a?includeDriveChart=true
	// BAL @ DEN, 38-35 (2OT), 2012 AFC divisional round
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2013-01-12", "10012013-0112-00b3-6e81-cc137a8c006a")
	periods := runMatchupPeriods(t, matchup)
	require.Equal(t, 6, len(periods), "6 periods")

	expected := []expectedPeriod{
		{1, "REG", 14, 14},
		{2, "REG", 7, 7},
		{3, "REG", 7, 7},
		{4, "REG", 7, 7},
		{5, "OT", 0, 0},
		{6, "OT", 3, 0},
	}
	for i, e := range expected {
		p := periods[i]
		assert.Equal(t, "FINAL_OVERTIME", p.Phase)
		assert.Equal(t, "BAL", p.AwayTeamAbbreviation)
		assert.Equal(t, "DEN", p.HomeTeamAbbreviation)
		assert.Equal(t, e.period, p.Period)
		assert.Equal(t, e.periodType, p.PeriodType)
		assert.Equal(t, e.awayScore, p.AwayTeamScore, "period %d away score", e.period)
		assert.Equal(t, e.homeScore, p.HomeTeamScore, "period %d home score", e.period)
	}
}
