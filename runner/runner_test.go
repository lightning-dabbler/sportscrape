//go:build unit

package runner

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/internal/mocks/scraper"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestMatchupRunner(t *testing.T) {
	mockscraper := scraper.NewMockMatchupScraper(t)
	dummyoutput := sportscrape.MatchupOutput{
		Context: sportscrape.MatchupContext{},
	}
	mockscraper.EXPECT().Init().Return()
	mockscraper.EXPECT().Scrape().Return(dummyoutput).Once()
	mockscraper.EXPECT().Feed().Return(sportscrape.DummyFeed)
	mockscraper.EXPECT().Provider().Return(sportscrape.DummyProvider).Once()
	runner := NewMatchupRunner(
		MatchupRunnerScraper(mockscraper),
	)
	matchups, err := runner.Run()
	assert.NoError(t, err)
	assert.Nil(t, matchups)
}

func TestEventDataRunner(t *testing.T) {
	mockscraper := scraper.NewMockEventDataScraper(t)
	dummyoutput := sportscrape.EventDataOutput{
		Context: sportscrape.EventDataContext{},
	}
	mockscraper.EXPECT().Init().Return()
	mockscraper.EXPECT().Scrape(mock.Anything).Return(dummyoutput).Once()
	mockscraper.EXPECT().Feed().Return(sportscrape.DummyFeed)
	mockscraper.EXPECT().Provider().Return(sportscrape.DummyProvider).Once()
	runner := NewEventDataRunner(
		EventDataRunnerScraper(mockscraper),
	)
	data, err := runner.Run(2)
	assert.NoError(t, err)
	assert.Nil(t, data)
}
