package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreHustleScraper(t *testing.T) {
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
	boxscorescraper := NewBoxScoreHustleScraper(
		WithBoxScoreHustleTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreHustle]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 19, n_records, "19 stat lines")
	testRecord := records[10]

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
	assert.Equal(t, int32(19), testRecord.Points)
	assert.Equal(t, int32(12), testRecord.ContestedShots)
	assert.Equal(t, int32(8), testRecord.ContestedShots2pt)
	assert.Equal(t, int32(4), testRecord.ContestedShots3pt)
	assert.Equal(t, int32(2), testRecord.Deflections)
	assert.Equal(t, int32(0), testRecord.ChargesDrawn)
	assert.Equal(t, int32(0), testRecord.ScreenAssists)
	assert.Equal(t, int32(0), testRecord.ScreenAssistPoints)
	assert.Equal(t, int32(1), testRecord.LooseBallsRecoveredOffensive)
	assert.Equal(t, int32(0), testRecord.LooseBallsRecoveredDefensive)
	assert.Equal(t, int32(1), testRecord.LooseBallsRecoveredTotal)
	assert.Equal(t, int32(0), testRecord.OffensiveBoxOuts)
	assert.Equal(t, int32(1), testRecord.DefensiveBoxOuts)
	assert.Equal(t, int32(1), testRecord.BoxOutPlayerTeamRebounds)
	assert.Equal(t, int32(1), testRecord.BoxOutPlayerRebounds)
	assert.Equal(t, int32(1), testRecord.BoxOuts)
}
