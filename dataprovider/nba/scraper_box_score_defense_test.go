package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreDefenseScraper(t *testing.T) {
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
	boxscorescraper := NewBoxScoreDefenseScraper(
		WithBoxScoreDefenseTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreDefense]{
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
	assert.Equal(t, float32(15.82), testRecord.MatchupMinutes)
	assert.Equal(t, float32(81), testRecord.PartialPossessions)
	assert.Equal(t, int32(0), testRecord.SwitchesOn)
	assert.Equal(t, int32(17), testRecord.PlayerPoints)
	assert.Equal(t, int32(5), testRecord.DefensiveRebounds)
	assert.Equal(t, int32(7), testRecord.MatchupAssists)
	assert.Equal(t, int32(5), testRecord.MatchupTurnovers)
	assert.Equal(t, int32(3), testRecord.Steals)
	assert.Equal(t, int32(0), testRecord.Blocks)
	assert.Equal(t, int32(7), testRecord.MatchupFieldGoalsMade)
	assert.Equal(t, int32(12), testRecord.MatchupFieldGoalsAttempted)
	assert.Equal(t, float32(0.583), testRecord.MatchupFieldGoalPercentage)
	assert.Equal(t, int32(3), testRecord.MatchupThreePointersMade)
	assert.Equal(t, int32(3), testRecord.MatchupThreePointersAttempted)
	assert.Equal(t, float32(1), testRecord.MatchupThreePointerPercentage)
}
