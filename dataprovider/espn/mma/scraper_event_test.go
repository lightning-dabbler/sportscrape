package mma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventDataScraper(T *testing.T) {
	scraper := espnEventDataScraper{}
	model, err := scraper.Scrape("600040033")
	assert.NoError(T, err)
	assert.NotNil(T, model)
}
