//go:build integration

package baseballreferencemlb

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb/model"
	"github.com/stretchr/testify/assert"
)

func TestMatchupScraper(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	tests := []struct {
		name               string
		date               string
		expectedNumMatches int
		playoff            bool
	}{
		{
			name:               "2024-10-02 MLB matches",
			date:               "2024-10-02",
			expectedNumMatches: 4,
			playoff:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scraper := NewMatchupScraper(
				WithMatchupDate(tt.date),
				WithMatchupTimeout(5*time.Minute),
			)
			runner := sportscrape.NewMatchupRunner(
				sportscrape.MatchupRunnerScraper(scraper),
			)
			matchups, err := runner.Run()
			if err != nil {
				t.Error(err)
			}
			for _, matchup := range matchups {
				structured_matchup := matchup.(model.MLBMatchup)
				assert.Equal(t, tt.playoff, structured_matchup.PlayoffMatch)
			}
			assert.Equal(t, tt.expectedNumMatches, len(matchups))
		})
	}
}
