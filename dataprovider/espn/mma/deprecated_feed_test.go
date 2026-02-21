//go:build unit

package mma

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/lightning-dabbler/sportscrape/scraper"

	"github.com/stretchr/testify/assert"
)

func TestDeprecatedESPNMMMAMatchupScraper_PFL(T *testing.T) {
	matchupscraper := ESPNMMAMatchupScraper{Year: "2024", League: "pfl", BaseScraper: scraper.BaseScraper{Timeout: 3 * time.Minute}}

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)

	r, err := matchuprunner.Run()
	assert.Error(T, err, "deprecated")
	assert.Nil(T, r)
}
