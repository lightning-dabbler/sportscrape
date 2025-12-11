//go:build unit

package mma

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/runner"
	scraper2 "github.com/lightning-dabbler/sportscrape/scraper"

	"github.com/stretchr/testify/assert"
)

func TestDeprecatedESPNMMMAMatchupScraper_PFL(T *testing.T) {
	scraper := ESPNMMAMatchupScraper{Year: "2024", League: "pfl", BaseScraper: scraper2.BaseScraper{Timeout: 3 * time.Minute}}

	matchupRunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(
			scraper,
		),
	)

	r, err := matchupRunner.Run()
	assert.Error(T, err, "deprecated")
	assert.Nil(t, r)
}
