package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreScoringScraper(t *testing.T) {
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
	boxscorescraper := NewBoxScoreScoringScraper(
		WithBoxScoreScoringTimeout(2*time.Minute),
		WithBoxScoreScoringPeriod(H1),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreScoring]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 21, n_records, "21 stat lines")
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
	assert.Equal(t, float32(19.83), testRecord.Minutes)
	assert.Equal(t, float32(0.722), testRecord.PercentageFieldGoalsAttempted2pt)
	assert.Equal(t, float32(0.278), testRecord.PercentageFieldGoalsAttempted3pt)
	assert.Equal(t, float32(0.632), testRecord.PercentagePoints2pt)
	assert.Equal(t, float32(0.211), testRecord.PercentagePointsMidrange2pt)
	assert.Equal(t, float32(0.316), testRecord.PercentagePoints3pt)
	assert.Equal(t, float32(0), testRecord.PercentagePointsFastBreak)
	assert.Equal(t, float32(0.053), testRecord.PercentagePointsFreeThrow)
	assert.Equal(t, float32(0.105), testRecord.PercentagePointsOffTurnovers)
	assert.Equal(t, float32(0.421), testRecord.PercentagePointsPaint)
	assert.Equal(t, float32(0), testRecord.PercentageAssisted2pt)
	assert.Equal(t, float32(1), testRecord.PercentageUnassisted2pt)
	assert.Equal(t, float32(0.5), testRecord.PercentageAssisted3pt)
	assert.Equal(t, float32(0.5), testRecord.PercentageUnassisted3pt)
	assert.Equal(t, float32(0.125), testRecord.PercentageAssistedFGM)
	assert.Equal(t, float32(0.875), testRecord.PercentageUnassistedFGM)
}
