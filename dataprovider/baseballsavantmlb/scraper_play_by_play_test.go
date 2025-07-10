package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/stretchr/testify/assert"
)

func TestPlayByPlayScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchupscraper := NewMatchupScraper(
		MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	playbyplayscraper := NewPlayByPlayScraper()
	playbyplayrunner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerScraper(playbyplayscraper),
		sportscrape.EventDataRunnerConcurrency(1),
	)
	plays, err := playbyplayrunner.Run(matchups...)
	assert.NoError(t, err)
	assert.Equal(t, 342, len(plays), "342 plays")
}
