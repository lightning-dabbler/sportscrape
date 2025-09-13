package mma

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchupScraper(T *testing.T) {
	scraper := NewScraperMatchups(2, []string{"2024"})
	out := scraper.Scrape()
	assert.NotEmpty(T, out)
}
