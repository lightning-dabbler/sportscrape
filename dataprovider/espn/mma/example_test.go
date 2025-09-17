package mma

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMatchupRunner(T *testing.T) {

	matchupRunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(
			ESPNMMAMatchupScraper{Year: "2024"},
		),
	)

	result, err := matchupRunner.Run()
	assert.NotEmpty(T, result)
	assert.NoError(T, err)

	eventRunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(ESPNMMAFightDetailsScraper{}),
	)

	result, err = eventRunner.Run(result[0:2]...)

	assert.NotEmpty(T, result)
	assert.NoError(T, err)

}
