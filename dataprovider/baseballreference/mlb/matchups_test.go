//go:build integration

package mlb

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb/model"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchups(t *testing.T) {

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
		{
			name:               "2024-09-02 MLB matches",
			date:               "2024-09-02",
			expectedNumMatches: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := NewMatchupRunner(
				WithMatchupTimeout(2 * time.Minute),
			)
			matchups := runner.GetMatchups(tt.date)
			for _, matchup := range matchups {
				structured_matchup := matchup.(model.MLBMatchup)
				assert.Equal(t, tt.playoff, structured_matchup.PlayoffMatch)
			}
			assert.Equal(t, tt.expectedNumMatches, len(matchups))
		})
	}
}
