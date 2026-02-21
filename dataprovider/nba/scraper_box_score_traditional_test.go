package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreTraditionalScraper(t *testing.T) {
	// https://www.nba.com/games?date=2025-06-05
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2025-06-05"),
		WithMatchupTimeout(3*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	boxscorescraper := NewBoxScoreTraditionalScraper(
		WithBoxScoreTraditionalTimeout(2*time.Minute),
		WithBoxScoreTraditionalPeriod(Q1),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreTraditional]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 20, n_records, "20 stat lines")
	testRecord := records[4]

	assert.Equal(t, "0042400401", testRecord.EventID)
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
	assert.Equal(t, float32(12), testRecord.Minutes)
	assert.Equal(t, int32(5), testRecord.FieldGoalsMade)
	assert.Equal(t, int32(12), testRecord.FieldGoalsAttempted)
	assert.Equal(t, float32(0.417), testRecord.FieldGoalsPercentage)
	assert.Equal(t, int32(1), testRecord.ThreePointersMade)
	assert.Equal(t, int32(1), testRecord.ThreePointersAttempted)
	assert.Equal(t, float32(1), testRecord.ThreePointersPercentage)
	assert.Equal(t, int32(1), testRecord.FreeThrowsMade)
	assert.Equal(t, int32(2), testRecord.FreeThrowsAttempted)
	assert.Equal(t, float32(0.5), testRecord.FreeThrowsPercentage)
	assert.Equal(t, int32(0), testRecord.ReboundsOffensive)
	assert.Equal(t, int32(2), testRecord.ReboundsDefensive)
	assert.Equal(t, int32(2), testRecord.ReboundsTotal)
	assert.Equal(t, int32(1), testRecord.Assists)
	assert.Equal(t, int32(1), testRecord.Steals)
	assert.Equal(t, int32(0), testRecord.Blocks)
	assert.Equal(t, int32(1), testRecord.Turnovers)
	assert.Equal(t, int32(0), testRecord.FoulsPersonal)
	assert.Equal(t, int32(12), testRecord.Points)
	assert.Equal(t, int32(9), testRecord.PlusMinusPoints)
}
