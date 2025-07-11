//go:build integration

package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestFieldingBoxScoreScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	date := "2024-10-07"
	matchupscraper := NewMatchupScraper(
		MatchupScraperDate(date),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	boxscorescraper := NewFieldingBoxScoreScraper()
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)
	boxScoreStats, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	assert.Equal(t, 64, len(boxScoreStats), "64 statlines")

	playerToTest := map[string]bool{
		"Austin Wells": false,
	}
	for _, statline := range boxScoreStats {
		stats := statline.(model.FieldingBoxScore)
		if stats.Player == "Austin Wells" {
			playerToTest[stats.Player] = true
			assert.Equal(t, int64(775332), stats.EventID)
			assert.Equal(t, int64(147), stats.TeamID)
			assert.Equal(t, "New York Yankees", stats.Team)
			assert.Equal(t, "Kansas City Royals", stats.Opponent)
			assert.Equal(t, int64(118), stats.OpponentID)
			assert.Equal(t, int64(669224), stats.PlayerID)
			assert.Equal(t, "Austin Wells", stats.Player)
			assert.Equal(t, "Catcher", stats.Position)
			assert.Equal(t, int32(1), stats.CaughtStealing)
			assert.Equal(t, int32(2), stats.StolenBases)
			assert.Equal(t, int32(1), stats.Assists)
			assert.Equal(t, int32(15), stats.Putouts)
			assert.Equal(t, int32(0), stats.Errors)
			assert.Equal(t, int32(16), stats.Chances)
			assert.Equal(t, int32(0), stats.PassedBall)
			assert.Equal(t, int32(0), stats.Pickoffs)
		}
	}
	assert.True(t, playerToTest["Austin Wells"], "Austin Wells's fielding stats were collected")
}
