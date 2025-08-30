//go:build integration

package foxsports

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMLBOddsTotalScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupScraper := NewMatchupScraper(
		MatchupScraperLeague(MLB),
		MatchupScraperSegmenter(&GeneralSegmenter{Date: "2024-10-25"}),
	)

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	oddsScraper := NewMLBOddsTotalScraper()
	oddsrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(1),
		runner.EventDataRunnerScraper(
			oddsScraper,
		),
	)
	odds, err := oddsrunner.Run(matchups...)
	assert.NoError(t, err)
	n_records := len(odds)
	n_expected := 1
	assert.Equal(t, n_expected, n_records, "1 odds record")
	record := odds[0].(model.MLBOddsTotal)

	assert.Equal(t, int32(-101), *record.OverOdds)
	assert.Equal(t, int32(-119), *record.UnderOdds)
	assert.Equal(t, float32(9), record.TotalLine)
}
