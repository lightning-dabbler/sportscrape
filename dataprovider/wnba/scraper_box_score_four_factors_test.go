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

func TestBoxScoreFourFactorsScraper(t *testing.T) {
	// https://www.wnba.com/game/dal-vs-ind-1022600254/box-score?type=fourfactors&period=All
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

	boxscorescraper := NewBoxScoreFourFactorsScraper(
		WithBoxScoreFourFactorsTimeout(3*time.Minute),
		WithBoxScoreFourFactorsPeriod(Full),
	)
	boxscorescraper.NetworkHeaders = NetworkHeaders
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreFourFactors]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run([]model.Matchup{matchup})
	assert.NoError(t, err)
	require.NotEmpty(t, records, "expected at least one player stat line")

	for _, r := range records {
		assert.Equal(t, "1022600254", r.EventID)
		assert.Contains(t, []string{"Fever", "Wings"}, r.TeamName)
		assert.Contains(t, []string{"Fever", "Wings"}, r.OpponentName)
		assert.NotEqual(t, r.TeamName, r.OpponentName)
	}

	// One record: validate every field against the real, known response
	// captured 2026-08-15 (Lexie Hull, Indiana Fever).
	var hull *model.BoxScoreFourFactors
	for i := range records {
		if records[i].PlayerID == 1631086 {
			hull = &records[i]
			break
		}
	}
	require.NotNil(t, hull, "expected to find player 1631086 (Lexie Hull)")

	assert.WithinDuration(t, time.Now().UTC(), hull.PullTimestamp, time.Minute)
	assert.Equal(t, "1022600254", hull.EventID)
	assert.Equal(t, "2026-08-14T23:30:00Z", hull.EventTime.Format(time.RFC3339))
	assert.Equal(t, int32(3), hull.EventStatus)
	assert.Equal(t, "Final", hull.EventStatusText)
	assert.Equal(t, int64(1611661325), hull.TeamID)
	assert.Equal(t, "Fever", hull.TeamName)
	assert.Equal(t, "Indiana Fever", hull.TeamNameFull)
	assert.Equal(t, int64(1611661321), hull.OpponentID)
	assert.Equal(t, "Wings", hull.OpponentName)
	assert.Equal(t, "Dallas Wings", hull.OpponentNameFull)
	assert.Equal(t, int64(1631086), hull.PlayerID)
	assert.Equal(t, "Lexie Hull", hull.PlayerName)
	assert.Equal(t, "F", hull.Position)
	assert.True(t, hull.Starter)
	assert.Equal(t, float32(12.4), hull.Minutes)
	assert.Equal(t, float32(0.5), hull.EffectiveFieldGoalPercentage)
	assert.Equal(t, float32(0), hull.FreeThrowAttemptRate)
	assert.Equal(t, float32(0.208), hull.TeamTurnoverPercentage)
	assert.Equal(t, float32(0), hull.OffensiveReboundPercentage)
	assert.Equal(t, float32(0.417), hull.OppEffectiveFieldGoalPercentage)
	assert.Equal(t, float32(0.556), hull.OppFreeThrowAttemptRate)
	assert.Equal(t, float32(0.214), hull.OppTeamTurnoverPercentage)
	assert.Equal(t, float32(0.364), hull.OppOffensiveReboundPercentage)
}
