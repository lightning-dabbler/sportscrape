package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreMatchupsScraper(t *testing.T) {
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
	boxscorescraper := NewBoxScoreMatchupsScraper(
		WithBoxScoreMatchupsTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 185, n_records, "185 stat lines")
	testRecord := records[34].(model.BoxScoreMatchups)

	assert.Equal(t, "0042400401", testRecord.EventID)
	assert.Equal(t, int32(3), testRecord.EventStatus)
	assert.Equal(t, "Final", testRecord.EventStatusText)
	assert.Equal(t, int64(1610612754), testRecord.TeamID)
	assert.Equal(t, "Pacers", testRecord.TeamName)
	assert.Equal(t, "Indiana Pacers", testRecord.TeamNameFull)
	assert.Equal(t, int64(1610612760), testRecord.OpponentID)
	assert.Equal(t, "Thunder", testRecord.OpponentName)
	assert.Equal(t, "Oklahoma City Thunder", testRecord.OpponentNameFull)
	assert.Equal(t, int64(1628983), testRecord.PlayerID)
	assert.Equal(t, "Shai Gilgeous-Alexander", testRecord.PlayerName)
	assert.Equal(t, "G", testRecord.Position)
	assert.Equal(t, true, testRecord.Starter)

	assert.Equal(t, int64(1630169), testRecord.OpponentPlayerID)
	assert.Equal(t, "Tyrese Haliburton", testRecord.OpponentPlayerName)
	assert.Equal(t, float32(0.68), testRecord.MatchupMinutes)
	assert.Equal(t, float32(41), testRecord.MatchupMinutesSort)
	assert.Equal(t, int32(5), testRecord.PartialPossessions)
	assert.Equal(t, float32(0.047), testRecord.PercentageDefenderTotalTime)
	assert.Equal(t, float32(0.047), testRecord.PercentageOffensiveTotalTime)
	assert.Equal(t, float32(0.052), testRecord.PercentageTotalTimeBothOn)
	assert.Equal(t, int32(0), testRecord.SwitchesOn)
	assert.Equal(t, int32(11), testRecord.PlayerPoints)
	assert.Equal(t, int32(13), testRecord.TeamPoints)
	assert.Equal(t, int32(1), testRecord.MatchupAssists)
	assert.Equal(t, int32(0), testRecord.MatchupPotentialAssists)
	assert.Equal(t, int32(0), testRecord.MatchupTurnovers)
	assert.Equal(t, int32(0), testRecord.MatchupBlocks)
	assert.Equal(t, int32(3), testRecord.MatchupFieldGoalsMade)
	assert.Equal(t, int32(3), testRecord.MatchupFieldGoalsAttempted)
	assert.Equal(t, float32(1), testRecord.MatchupFieldGoalsPercentage)
	assert.Equal(t, int32(1), testRecord.MatchupThreePointersMade)
	assert.Equal(t, int32(1), testRecord.MatchupThreePointersAttempted)
	assert.Equal(t, float32(1), testRecord.MatchupThreePointersPercentage)
	assert.Equal(t, int32(0), testRecord.HelpBlocks)
	assert.Equal(t, int32(0), testRecord.HelpFieldGoalsMade)
	assert.Equal(t, int32(0), testRecord.HelpFieldGoalsAttempted)
	assert.Equal(t, float32(0), testRecord.HelpFieldGoalsPercentage)
	assert.Equal(t, int32(4), testRecord.MatchupFreeThrowsMade)
	assert.Equal(t, int32(4), testRecord.MatchupFreeThrowsAttempted)
	assert.Equal(t, int32(1), testRecord.ShootingFouls)

}
