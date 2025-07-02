//go:build unit

package sportscrape

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestEventDataRunner(t *testing.T) {}
