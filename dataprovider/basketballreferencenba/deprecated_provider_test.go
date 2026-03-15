//go:build unit

package basketballreferencenba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestDeprecatedProvider(t *testing.T) {
	date := "2025-02-19"

	matchupscraper := NewMatchupScraper(
		WithMatchupDate(date),
		WithMatchupTimeout(5*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.NBAMatchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.Error(t, err, "deprecated")
	assert.Nil(t, matchups)
}
