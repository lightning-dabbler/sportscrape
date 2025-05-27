//go:build integration

package eventdata

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
	"github.com/stretchr/testify/assert"
)

func TestMLBProbableStartingPitcher(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.MLB),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-25"}),
	)
	matchups := matchupRunner.GetMatchups()
	scraper := MLBProbableStartingPitcherScraper{}
	scraper.League = foxsports.MLB
	runner := NewRunner(
		RunnerName("MLB probable starting pitcher"),
		RunnerConcurrency(1),
		RunnerScraper(
			&scraper,
		),
	)

	probablePitchers := runner.RunEventsDataScraper(matchups...)
	n_stats := len(probablePitchers)
	n_expected := 1
	assert.Equal(t, n_expected, n_stats, "1 set of opposing pitchers")
	result := probablePitchers[0].(model.MLBProbableStartingPitcher)
	assert.Equal(t, "Jack Flaherty", result.HomeStartingPitcher)
	assert.Equal(t, "1-2", result.HomeStartingPitcherRecord)
	assert.Equal(t, float32(7.36), result.HomeStartingPitcherERA)
	assert.Equal(t, int64(8249), result.HomeStartingPitcherID)
	assert.Equal(t, "Gerrit Cole", result.AwayStartingPitcher)
	assert.Equal(t, "1-0", result.AwayStartingPitcherRecord)
	assert.Equal(t, float32(2.17), result.AwayStartingPitcherERA)
	assert.Equal(t, int64(5539), result.AwayStartingPitcherID)
}
