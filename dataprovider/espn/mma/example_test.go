package mma

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMatchupRunner(T *testing.T) {

	matchupRunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(
			NewScraperMatchups(2, []string{"2024"}),
		),
	)

	result, err := matchupRunner.Run()

	assert.NotEmpty(T, result)
	assert.NoError(T, err)

}
