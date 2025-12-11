//go:build unit

package baseballreferencemlb

import (
	"testing"
	"time"

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
		runner.MatchupRunnerScraper(matchupscraper),
	)
	// Retrieve MLB matchups associated with date
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err, "deprecated")
	assert.Nil(t, matchups)
}
