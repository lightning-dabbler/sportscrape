//go:build integration

package scraper

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/stretchr/testify/assert"
)

func TestMLBProbableStartingPitcher(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupScraper := NewMatchupScraper(
		MatchupScraperLeague(foxsports.MLB),
		MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-25"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
	if err != nil {
		t.Error(err)
	}

	boxscoreScraper := MLBProbableStartingPitcherScraper{}
	boxscoreScraper.League = foxsports.MLB
	runner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerConcurrency(1),
		sportscrape.EventDataRunnerScraper(
			&boxscoreScraper,
		),
	)
	probablePitchers, err := runner.RunEventsDataScraper(matchups...)
	if err != nil {
		t.Error(err)
	}
	n_records := len(probablePitchers)
	n_expected := 2
	assert.Equal(t, n_expected, n_records, "2 starting pitchers")
	homeStartingPitcher := probablePitchers[0].(model.MLBProbableStartingPitcher)
	awayStartingPitcher := probablePitchers[1].(model.MLBProbableStartingPitcher)

	assert.Equal(t, "Jack Flaherty", homeStartingPitcher.StartingPitcher)
	assert.Equal(t, "1-2", homeStartingPitcher.StartingPitcherRecord)
	assert.Equal(t, float32(7.36), homeStartingPitcher.StartingPitcherERA)
	assert.Equal(t, int64(8249), homeStartingPitcher.StartingPitcherID)

	assert.Equal(t, "Gerrit Cole", awayStartingPitcher.StartingPitcher)
	assert.Equal(t, "1-0", awayStartingPitcher.StartingPitcherRecord)
	assert.Equal(t, float32(2.17), awayStartingPitcher.StartingPitcherERA)
	assert.Equal(t, int64(5539), awayStartingPitcher.StartingPitcherID)
}
