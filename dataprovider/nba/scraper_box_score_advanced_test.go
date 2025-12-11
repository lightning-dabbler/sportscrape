package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreAdvancedScraper(t *testing.T) {
	// https://www.nba.com/games?date=2025-06-11
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2025-06-11"),
		WithMatchupTimeout(3*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	boxscorescraper := NewBoxScoreAdvancedScraper(
		WithBoxScoreAdvancedTimeout(2*time.Minute),
		WithBoxScoreAdvancedPeriod(Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 29, n_records, "29 stat lines")
	testRecord := records[20].(model.BoxScoreAdvanced)

	assert.Equal(t, "0042400403", testRecord.EventID)
	assert.Equal(t, int32(3), testRecord.EventStatus)
	assert.Equal(t, "Final", testRecord.EventStatusText)
	assert.Equal(t, int64(1610612760), testRecord.TeamID)
	assert.Equal(t, "Thunder", testRecord.TeamName)
	assert.Equal(t, "Oklahoma City Thunder", testRecord.TeamNameFull)
	assert.Equal(t, int64(1610612754), testRecord.OpponentID)
	assert.Equal(t, "Pacers", testRecord.OpponentName)
	assert.Equal(t, "Indiana Pacers", testRecord.OpponentNameFull)
	assert.Equal(t, int64(1627936), testRecord.PlayerID)
	assert.Equal(t, "Alex Caruso", testRecord.PlayerName)
	assert.Equal(t, "", testRecord.Position)
	assert.Equal(t, false, testRecord.Starter)
	assert.Equal(t, float32(32.23), testRecord.Minutes)
	assert.Equal(t, float32(91.4), testRecord.EstimatedOffensiveRating)
	assert.Equal(t, float32(91.4), testRecord.OffensiveRating)
	assert.Equal(t, float32(114.5), testRecord.EstimatedDefensiveRating)
	assert.Equal(t, float32(114.5), testRecord.DefensiveRating)
	assert.Equal(t, float32(-23.1), testRecord.EstimatedNetRating)
	assert.Equal(t, float32(-23.1), testRecord.NetRating)
	assert.Equal(t, float32(0.211), testRecord.AssistPercentage)
	assert.Equal(t, float32(4), testRecord.AssistToTurnover)
	assert.Equal(t, float32(36.4), testRecord.AssistRatio)
	assert.Equal(t, float32(0.03), testRecord.OffensiveReboundPercentage)
	assert.Equal(t, float32(0.138), testRecord.DefensiveReboundPercentage)
	assert.Equal(t, float32(0.081), testRecord.ReboundPercentage)
	assert.Equal(t, float32(9.1), testRecord.TurnoverRatio)
	assert.Equal(t, float32(0.6), testRecord.EffectiveFieldGoalPercentage)
	assert.Equal(t, float32(0.68), testRecord.TrueShootingPercentage)
	assert.Equal(t, float32(0.088), testRecord.UsagePercentage)
	assert.Equal(t, float32(0.088), testRecord.EstimatedUsagePercentage)
	assert.Equal(t, float32(103.48), testRecord.EstimatedPace)
	assert.Equal(t, float32(103.48), testRecord.Pace)
	assert.Equal(t, float32(86.24), testRecord.PacePer40)
	assert.Equal(t, int32(70), testRecord.Possessions)
	assert.Equal(t, float32(0.097), testRecord.PIE)
}
