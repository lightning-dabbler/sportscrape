package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreTrackingScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2025-06-05"),
		WithMatchupTimeout(3*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	boxscorescraper := NewBoxScoreTrackingScraper(
		WithBoxScoreTrackingTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 29, n_records, "29 stat lines")
	testRecord := records[18].(model.BoxScoreTracking)

	assert.Equal(t, "0042400401", testRecord.EventID)
	assert.Equal(t, int32(3), testRecord.EventStatus)
	assert.Equal(t, "Final", testRecord.EventStatusText)
	assert.Equal(t, int64(1610612754), testRecord.TeamID)
	assert.Equal(t, "Pacers", testRecord.TeamName)
	assert.Equal(t, "Indiana Pacers", testRecord.TeamNameFull)
	assert.Equal(t, int64(1610612760), testRecord.OpponentID)
	assert.Equal(t, "Thunder", testRecord.OpponentName)
	assert.Equal(t, "Oklahoma City Thunder", testRecord.OpponentNameFull)
	assert.Equal(t, int64(1630169), testRecord.PlayerID)
	assert.Equal(t, "Tyrese Haliburton", testRecord.PlayerName)
	assert.Equal(t, "G", testRecord.Position)
	assert.Equal(t, true, testRecord.Starter)
	assert.Equal(t, float32(38.92), testRecord.Minutes)
	assert.Equal(t, float32(4.45), testRecord.Speed)
	assert.Equal(t, float32(3.06), testRecord.Distance)
	assert.Equal(t, int32(0), testRecord.ReboundChancesOffensive)
	assert.Equal(t, int32(13), testRecord.ReboundChancesDefensive)
	assert.Equal(t, int32(13), testRecord.ReboundChancesTotal)
	assert.Equal(t, int32(108), testRecord.Touches)
	assert.Equal(t, int32(2), testRecord.SecondaryAssists)
	assert.Equal(t, int32(1), testRecord.FreeThrowAssists)
	assert.Equal(t, int32(89), testRecord.Passes)
	assert.Equal(t, int32(6), testRecord.Assists)
	assert.Equal(t, int32(1), testRecord.ContestedFieldGoalsMade)
	assert.Equal(t, int32(3), testRecord.ContestedFieldGoalsAttempted)
	assert.Equal(t, float32(0.333), testRecord.ContestedFieldGoalPercentage)
	assert.Equal(t, int32(5), testRecord.UncontestedFieldGoalsMade)
	assert.Equal(t, int32(10), testRecord.UncontestedFieldGoalsAttempted)
	assert.Equal(t, float32(0.5), testRecord.UncontestedFieldGoalsPercentage)
	assert.Equal(t, float32(0.462), testRecord.FieldGoalPercentage)
	assert.Equal(t, int32(3), testRecord.DefendedAtRimFieldGoalsMade)
	assert.Equal(t, int32(5), testRecord.DefendedAtRimFieldGoalsAttempted)
	assert.Equal(t, float32(0.6), testRecord.DefendedAtRimFieldGoalPercentage)
}
