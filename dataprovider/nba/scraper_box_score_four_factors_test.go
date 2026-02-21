package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreFourFactorsScraper(t *testing.T) {
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
	boxscorescraper := NewBoxScoreFourFactorsScraper(
		WithBoxScoreFourFactorsTimeout(2*time.Minute),
		WithBoxScoreFourFactorsPeriod(Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreFourFactors]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 29, n_records, "29 stat lines")
	testRecord := records[22]

	assert.Equal(t, "0042400401", testRecord.EventID)
	assert.Equal(t, int32(3), testRecord.EventStatus)
	assert.Equal(t, "Final", testRecord.EventStatusText)
	assert.Equal(t, int64(1610612754), testRecord.TeamID)
	assert.Equal(t, "Pacers", testRecord.TeamName)
	assert.Equal(t, "Indiana Pacers", testRecord.TeamNameFull)
	assert.Equal(t, int64(1610612760), testRecord.OpponentID)
	assert.Equal(t, "Thunder", testRecord.OpponentName)
	assert.Equal(t, "Oklahoma City Thunder", testRecord.OpponentNameFull)
	assert.Equal(t, int64(1631097), testRecord.PlayerID)
	assert.Equal(t, "Bennedict Mathurin", testRecord.PlayerName)
	assert.Equal(t, "", testRecord.Position)
	assert.Equal(t, false, testRecord.Starter)
	assert.Equal(t, float32(15.88), testRecord.Minutes)
	assert.Equal(t, float32(0.586), testRecord.EffectiveFieldGoalPercentage)
	assert.Equal(t, float32(0.207), testRecord.FreeThrowAttemptRate)
	assert.Equal(t, float32(0.178), testRecord.TeamTurnoverPercentage)
	assert.Equal(t, float32(0.2), testRecord.OffensiveReboundPercentage)
	assert.Equal(t, float32(0.567), testRecord.OppEffectiveFieldGoalPercentage)
	assert.Equal(t, float32(0.5), testRecord.OppFreeThrowAttemptRate)
	assert.Equal(t, float32(0.14), testRecord.OppTeamTurnoverPercentage)
	assert.Equal(t, float32(0.4), testRecord.OppOffensiveReboundPercentage)
}
