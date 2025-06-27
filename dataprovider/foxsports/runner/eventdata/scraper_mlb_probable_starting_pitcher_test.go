//go:build integration

package eventdata

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
	"github.com/lightning-dabbler/sportscrape/util/runner/eventdata"
	matchuputil "github.com/lightning-dabbler/sportscrape/util/runner/matchup"
	"github.com/stretchr/testify/assert"
)

func TestMLBProbableStartingPitcher(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.MLB),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-25"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("MLB Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()

	scraper := MLBProbableStartingPitcherScraper{}
	scraper.League = foxsports.MLB
	runner := eventdata.NewRunner(
		eventdata.RunnerName("MLB probable starting pitcher"),
		eventdata.RunnerConcurrency(1),
		eventdata.RunnerScraper(
			&scraper,
		),
	)

	probablePitchers := runner.RunEventsDataScraper(matchups...)
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
