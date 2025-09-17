package mma

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/stretchr/testify/assert"
)

func TestMatchupScraper(T *testing.T) {
	scraper := ESPNMMAFightDetailsScraper{}

	mockTime := time.Now()
	matchup := model.Matchup{
		PullTimestamp: mockTime,
		EventID:       "600041054",
		EventTime:     mockTime,
	}
	result := scraper.Scrape(matchup)
	assert.NoError(T, result.Error)
	assert.NotEmpty(T, result.Output)

}
