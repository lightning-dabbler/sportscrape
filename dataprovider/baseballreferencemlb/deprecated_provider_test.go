//go:build unit

package baseballreferencemlb

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestDeprecatedProvider(t *testing.T) {

	date := "2024-10-13"

	matchupscraper := NewMatchupScraper(
		WithMatchupDate(date),
		WithMatchupTimeout(5*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.MLBMatchup]{
			Scraper: matchupscraper,
		},
	)
	// Retrieve MLB matchups associated with date
	matchups, err := matchuprunner.Run()
	assert.Error(t, err, "deprecated")
	assert.Nil(t, matchups)
}
