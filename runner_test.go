//go:build unit

package sportscrape

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestMatchupRunner(t *testing.T) {
	mockscraper := NewMockMatchupScraper(t)
	dummyoutput := MatchupOutput{
		Context: MatchupContext{},
	}
	mockscraper.EXPECT().Init().Return()
	mockscraper.EXPECT().Scrape().Return(dummyoutput).Once()
	mockscraper.EXPECT().Feed().Return(DummyFeed)
	mockscraper.EXPECT().Provider().Return(DummyProvider).Once()
	runner := NewMatchupRunner(
		MatchupRunnerScraper(mockscraper),
	)
	matchups, err := runner.Run()
	assert.NoError(t, err)
	assert.Nil(t, matchups)
}

func TestEventDataRunner(t *testing.T) {
	mockscraper := NewMockEventDataScraper(t)
	dummyoutput := EventDataOutput{
		Context: EventDataContext{},
	}
	mockscraper.EXPECT().Init().Return()
	mockscraper.EXPECT().Scrape(mock.Anything).Return(dummyoutput).Once()
	mockscraper.EXPECT().Feed().Return(DummyFeed)
	mockscraper.EXPECT().Provider().Return(DummyProvider).Once()
	runner := NewEventDataRunner(
		EventDataRunnerScraper(mockscraper),
	)
	data, err := runner.Run(2)
	assert.NoError(t, err)
	assert.Nil(t, data)
}
