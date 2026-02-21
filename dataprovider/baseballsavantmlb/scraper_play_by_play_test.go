package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestPlayByPlayScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchupscraper := NewMatchupScraper(
		MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	playbyplayscraper := NewPlayByPlayScraper()
	playbyplayrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper:     playbyplayscraper,
			Concurrency: 1,
		},
	)
	plays, err := playbyplayrunner.Run(matchups)
	assert.NoError(t, err)
	assert.Equal(t, 342, len(plays), "342 plays")
}
