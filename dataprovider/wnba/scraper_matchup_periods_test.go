//go:build integration

package wnba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatchupPeriodsScraper(t *testing.T) {
	// https://www.wnba.com/game/dal-vs-ind-1022600254/box-score?period=All&type=traditional
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(WithMatchupDate("2026-08-14"))
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{Scraper: matchupScraper},
	)
	matchups, err := matchuprunner.Run()
	require.NoError(t, err)

	var matchup model.Matchup
	var found bool
	for _, m := range matchups {
		if m.EventID == "1022600254" {
			matchup = m
			found = true
			break
		}
	}
	require.True(t, found, "expected to find game 1022600254 (DAL @ IND) on 2026-08-14")
	assert.Equal(t, "IND", matchup.HomeTeamAbbreviation)
	assert.Equal(t, "DAL", matchup.AwayTeamAbbreviation)

	periodsscraper := NewMatchupPeriodsScraper(
		WithMatchupPeriodsTimeout(3 * time.Minute),
	)
	periodsscraper.NetworkHeaders = NetworkHeaders
	periodsrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
			Scraper:     periodsscraper,
			Concurrency: 1,
		},
	)

	records, err := periodsrunner.Run([]model.Matchup{matchup})
	assert.NoError(t, err)
	require.Len(t, records, 4, "4 regular quarters, no OT")

	var homeSum, awaySum int32
	for i, r := range records {
		assert.Equal(t, "1022600254", r.EventID)
		assert.Equal(t, int32(i+1), r.Period)
		assert.Equal(t, "REGULAR", r.PeriodType)
		homeSum += r.HomeTeamScore
		awaySum += r.AwayTeamScore
	}
	assert.Equal(t, int32(98), homeSum, "home (IND) period scores should sum to final score")
	assert.Equal(t, int32(87), awaySum, "away (DAL) period scores should sum to final score")

	// Record 0 (1st period): validate every field against the real, known
	// response captured 2026-08-15.
	q1 := records[0]
	assert.WithinDuration(t, time.Now().UTC(), q1.PullTimestamp, time.Minute)
	assert.Equal(t, "1022600254", q1.EventID)
	assert.Equal(t, "2026-08-14T23:30:00Z", q1.EventTime.Format(time.RFC3339))
	assert.Equal(t, int32(3), q1.EventStatus)
	assert.Equal(t, "Final", q1.EventStatusText)
	assert.Equal(t, int64(1611661325), q1.HomeTeamID)
	assert.Equal(t, "Fever", q1.HomeTeam)
	assert.Equal(t, "IND", q1.HomeTeamAbbreviation)
	assert.Equal(t, int64(1611661321), q1.AwayTeamID)
	assert.Equal(t, "Wings", q1.AwayTeam)
	assert.Equal(t, "DAL", q1.AwayTeamAbbreviation)
	assert.Equal(t, int32(1), q1.Period)
	assert.Equal(t, "REGULAR", q1.PeriodType)
	assert.Equal(t, int32(15), q1.AwayTeamScore)
	assert.Equal(t, int32(17), q1.HomeTeamScore)
	assert.Equal(t, "Regular Season", q1.SeasonType)
	assert.Equal(t, "2026", q1.SeasonYear)
	assert.Equal(t, "10", q1.LeagueID)
}
