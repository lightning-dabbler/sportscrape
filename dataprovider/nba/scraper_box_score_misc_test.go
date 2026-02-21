package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreMiscScraper(t *testing.T) {
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
	boxscorescraper := NewBoxScoreMiscScraper(
		WithBoxScoreMiscTimeout(2*time.Minute),
		WithBoxScoreMiscPeriod(Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreMisc]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 29, n_records, "29 stat lines")
	testRecord := records[15]

	assert.Equal(t, "0042400401", testRecord.EventID)
	assert.Equal(t, int32(3), testRecord.EventStatus)
	assert.Equal(t, "Final", testRecord.EventStatusText)
	assert.Equal(t, int64(1610612754), testRecord.TeamID)
	assert.Equal(t, "Pacers", testRecord.TeamName)
	assert.Equal(t, "Indiana Pacers", testRecord.TeamNameFull)
	assert.Equal(t, int64(1610612760), testRecord.OpponentID)
	assert.Equal(t, "Thunder", testRecord.OpponentName)
	assert.Equal(t, "Oklahoma City Thunder", testRecord.OpponentNameFull)
	assert.Equal(t, int64(1627783), testRecord.PlayerID)
	assert.Equal(t, "Pascal Siakam", testRecord.PlayerName)
	assert.Equal(t, "F", testRecord.Position)
	assert.Equal(t, true, testRecord.Starter)
	assert.Equal(t, float32(34.98), testRecord.Minutes)
	assert.Equal(t, int32(2), testRecord.PointsOffTurnovers)
	assert.Equal(t, int32(7), testRecord.PointsSecondChance)
	assert.Equal(t, int32(2), testRecord.PointsFastBreak)
	assert.Equal(t, int32(12), testRecord.PointsPaint)
	assert.Equal(t, int32(9), testRecord.OppPointsOffTurnovers)
	assert.Equal(t, int32(7), testRecord.OppPointsSecondChance)
	assert.Equal(t, int32(11), testRecord.OppPointsFastBreak)
	assert.Equal(t, int32(34), testRecord.OppPointsPaint)
	assert.Equal(t, int32(1), testRecord.Blocks)
	assert.Equal(t, int32(1), testRecord.BlocksAgainst)
	assert.Equal(t, int32(1), testRecord.FoulsPersonal)
	assert.Equal(t, int32(6), testRecord.FoulsDrawn)
}
