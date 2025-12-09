//go:build integration

package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreUsageScraper(t *testing.T) {
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
	boxscorescraper := NewBoxScoreUsageScraper(
		WithBoxScoreUsageTimeout(2*time.Minute),
		WithBoxScoreUsagePeriod(Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 29, n_records, "29 stat lines")
	testRecord := records[19].(model.BoxScoreUsage)

	assert.Equal(t, "0042400403", testRecord.EventID)
	assert.Equal(t, int32(3), testRecord.EventStatus)
	assert.Equal(t, "Final", testRecord.EventStatusText)
	assert.Equal(t, int64(1610612760), testRecord.TeamID)
	assert.Equal(t, "Thunder", testRecord.TeamName)
	assert.Equal(t, "Oklahoma City Thunder", testRecord.TeamNameFull)
	assert.Equal(t, int64(1610612754), testRecord.OpponentID)
	assert.Equal(t, "Pacers", testRecord.OpponentName)
	assert.Equal(t, "Indiana Pacers", testRecord.OpponentNameFull)
	assert.Equal(t, int64(1628983), testRecord.PlayerID)
	assert.Equal(t, "Shai Gilgeous-Alexander", testRecord.PlayerName)
	assert.Equal(t, "G", testRecord.Position)
	assert.Equal(t, true, testRecord.Starter)
	assert.Equal(t, float32(42.03), testRecord.Minutes)
	assert.Equal(t, float32(0.287), testRecord.UsagePercentage)
	assert.Equal(t, float32(0.281), testRecord.PercentageFieldGoalsMade)
	assert.Equal(t, float32(0.29), testRecord.PercentageFieldGoalsAttempted)
	assert.Equal(t, float32(0.111), testRecord.PercentageThreePointersMade)
	assert.Equal(t, float32(0.143), testRecord.PercentageThreePointersAttempted)
	assert.Equal(t, float32(0.25), testRecord.PercentageFreeThrowsMade)
	assert.Equal(t, float32(0.222), testRecord.PercentageFreeThrowsAttempted)
	assert.Equal(t, float32(0.222), testRecord.PercentageReboundsOffensive)
	assert.Equal(t, float32(0.207), testRecord.PercentageReboundsDefensive)
	assert.Equal(t, float32(0.211), testRecord.PercentageReboundsTotal)
	assert.Equal(t, float32(0.308), testRecord.PercentageAssists)
	assert.Equal(t, float32(0.375), testRecord.PercentageTurnovers)
	assert.Equal(t, float32(0), testRecord.PercentageSteals)
	assert.Equal(t, float32(0.75), testRecord.PercentageBlocks)
	assert.Equal(t, float32(0.1), testRecord.PercentageBlocksAllowed)
	assert.Equal(t, float32(0.125), testRecord.PercentagePersonalFouls)
	assert.Equal(t, float32(0.227), testRecord.PercentagePersonalFoulsDrawn)
	assert.Equal(t, float32(0.258), testRecord.PercentagePoints)
}
