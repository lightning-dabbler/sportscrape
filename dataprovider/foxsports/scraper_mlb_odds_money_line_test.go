//go:build integration

package foxsports

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMLBOddsMoneyLineScraper(t *testing.T) {
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

	oddsScraper := NewMLBOddsMoneyLineScraper()
	oddsrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MLBOddsMoneyLine]{
			Scraper:     oddsScraper,
			Concurrency: 1,
		},
	)
	odds, err := oddsrunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(odds)
	n_expected := 1
	assert.Equal(t, n_expected, n_records, "1 odds record")
	record := odds[0]

	assert.Equal(t, int32(104), record.AwayTeamOdds)
	assert.Equal(t, int32(-123), record.HomeTeamOdds)
}
