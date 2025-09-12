package mma

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/stretchr/testify/assert"
)

func TestScrapeAll(T *testing.T) {

	eventsFeedScraper := ESPNMMAEventsFeedScraper{
		Year: "2024",
	}
	eventScraper := ESPNMMAEventDataScraper{}

	// Events Feed
	feed, err := eventsFeedScraper.Scrape("2024")
	assert.NoError(T, err)

	allData := make([]*model.ESPNEventData, 0)

	for _, event := range feed.FilterScrapeableEvents() {
		id := event.ID
		data, err := eventScraper.Scrape(id)
		assert.NoError(T, err)
		allData = append(allData, data)
	}

	assert.NotEmpty(T, allData)
}
