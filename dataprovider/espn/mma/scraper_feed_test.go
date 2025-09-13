package mma

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Example for nba.MatchupRunner
func TestEventsDataScraper(T *testing.T) {
	scraper := espnScraperFeed{}

	// Test Retrieving the model for a single year
	model, err := scraper.Scrape("2024")
	assert.NoError(T, err)
	assert.NotNil(T, model)

	// Test Unmarshalling the json.RawMessage into the map
	var data map[string]interface{}
	err = json.Unmarshal(model.Raw, &data)
	assert.NoError(T, err)

	events := model.FilterScrapeableEvents()
	assert.NotEmpty(T, events)
}
