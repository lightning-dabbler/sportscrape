//go:build integration

package nba

import (
	"testing"
	"time"

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
	}{
		{
			name:               "2025-02-13 NBA matches",
			date:               "2025-02-13",
			expectedNumMatches: 5,
		},
		{
			name:               "2025-02-19 NBA matches",
			date:               "2025-02-19",
			expectedNumMatches: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := NewMatchupRunner(
				WithMatchupTimeout(2 * time.Minute),
			)
			matchups := runner.GetMatchups(tt.date)
			assert.Equal(t, tt.expectedNumMatches, len(matchups))
		})
	}
}
