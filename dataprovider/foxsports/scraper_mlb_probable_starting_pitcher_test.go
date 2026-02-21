//go:build integration

package foxsports

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMLBProbableStartingPitcherScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupScraper := NewMatchupScraper(
		MatchupScraperLeague(MLB),
		MatchupScraperSegmenter(&GeneralSegmenter{Date: "2024-10-25"}),
	)

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	boxscoreScraper := NewMLBProbableStartingPitcherScraper()
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MLBProbableStartingPitcher]{
			Scraper:     boxscoreScraper,
			Concurrency: 1,
		},
	)
	probablePitchers, err := boxscorerunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(probablePitchers)
	n_expected := 2
	assert.Equal(t, n_expected, n_records, "2 starting pitchers")
	homeStartingPitcher := probablePitchers[0]
	awayStartingPitcher := probablePitchers[1]

	assert.Equal(t, "Jack Flaherty", homeStartingPitcher.StartingPitcher)
	assert.Equal(t, "1-2", *homeStartingPitcher.StartingPitcherRecord)
	assert.Equal(t, float32(7.36), *homeStartingPitcher.StartingPitcherERA)
	assert.Equal(t, int64(8249), homeStartingPitcher.StartingPitcherID)

	assert.Equal(t, "Gerrit Cole", awayStartingPitcher.StartingPitcher)
	assert.Equal(t, "1-0", *awayStartingPitcher.StartingPitcherRecord)
	assert.Equal(t, float32(2.17), *awayStartingPitcher.StartingPitcherERA)
	assert.Equal(t, int64(5539), awayStartingPitcher.StartingPitcherID)
}
