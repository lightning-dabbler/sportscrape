//go:build unit

package runner

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/internal/mocks/scraper"
	"github.com/stretchr/testify/assert"
)

func TestMatchupRunner(t *testing.T) {
	mockscraper := scraper.NewMockMatchupScraper[any](t)
	dummyoutput := sportscrape.MatchupOutput[any]{
		Context: sportscrape.MatchupContext{},
	}
	mockscraper.EXPECT().Init().Return()
	mockscraper.EXPECT().Scrape().Return(dummyoutput).Once()
	mockscraper.EXPECT().Feed().Return(sportscrape.DummyFeed)
	mockscraper.EXPECT().Provider().Return(sportscrape.DummyProvider).Once()
	matchuprunner := NewMatchupRunner(
		MatchupRunnerConfig[any]{Scraper: mockscraper},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	assert.Nil(t, matchups)
}

func TestEventDataRunner(t *testing.T) {
	type fakeMatchup struct{}
	type fakeEvent struct{}
	mockscraper := scraper.NewMockEventDataScraper[fakeMatchup, fakeEvent](t)
	dummyoutput := sportscrape.EventDataOutput[fakeEvent]{
		Context: sportscrape.EventDataContext{},
	}
	mockscraper.EXPECT().Init().Return()
	mockscraper.EXPECT().Scrape(fakeMatchup{}).Return(dummyoutput).Once()
	mockscraper.EXPECT().Feed().Return(sportscrape.DummyFeed).Times(3)
	mockscraper.EXPECT().Provider().Return(sportscrape.DummyProvider).Once()
	eventDataRunner := NewEventDataRunner(
		EventDataRunnerConfig[fakeMatchup, fakeEvent]{
			Concurrency: 1,
			Scraper:     mockscraper,
		},
	)
	data, err := eventDataRunner.Run([]fakeMatchup{{}})
	assert.NoError(t, err)
	assert.Empty(t, data)
}
