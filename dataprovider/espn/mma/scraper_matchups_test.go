package mma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Example for nba.MatchupRunner
func TestEventsDataScraper(T *testing.T) {
	scraper := ESPNMMAMatchupScraper{Year: "2024"}

	r := scraper.Scrape()
	assert.NoError(T, r.Error)
	assert.NotEmpty(T, r.Output)

	//// Test Unmarshalling the json.RawMessage into the map
	//var data map[string]interface{}
	//err = json.Unmarshal(model.Raw, &data)
	//assert.NoError(T, err)
	//
	//events := model.FilterScrapeableEvents()
	//assert.NotEmpty(T, events)
}
