//go:build integration

package nhl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoalieBoxScoreScraper(t *testing.T) {
	// https://api-web.nhle.com/v1/gamecenter/2024020250/boxscore
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2024-11-12", 2024020250)
	scraper := NewGoalieBoxScoreScraper()
	scraper.Fetcher = testFetcher
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.GoalieBoxScore]{
			Scraper: scraper,
		},
	)
	statlines, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	assert.Equal(t, 4, len(statlines), "4 statlines")

	vladarTested := false
	wolfTested := false
	for _, s := range statlines {
		switch s.PlayerID {
		case int64(8478435):
			vladarTested = true
			assert.Equal(t, int64(2024020250), s.EventID)
			assert.Equal(t, int64(20), s.TeamID)
			assert.Equal(t, "Flames", s.Team)
			assert.Equal(t, int64(23), s.OpponentID)
			assert.Equal(t, "Canucks", s.Opponent)
			assert.Equal(t, "Dan Vladar", s.Player)
			assert.Equal(t, int32(80), s.SweaterNumber)
			assert.Equal(t, "G", s.Position)
			assert.Equal(t, int32(26), s.EvenStrengthSaves)
			assert.Equal(t, int32(28), s.EvenStrengthShotsAgainst)
			assert.Equal(t, int32(2), s.PowerPlaySaves)
			assert.Equal(t, int32(3), s.PowerPlayShotsAgainst)
			assert.Equal(t, int32(1), s.ShorthandedSaves)
			assert.Equal(t, int32(1), s.ShorthandedShotsAgainst)
			assert.Equal(t, int32(29), s.Saves)
			assert.Equal(t, int32(32), s.ShotsAgainst)
			require.NotNil(t, s.SavePctg)
			assert.Equal(t, float32(0.90625), *s.SavePctg)
			assert.Equal(t, int32(2), s.EvenStrengthGoalsAgainst)
			assert.Equal(t, int32(1), s.PowerPlayGoalsAgainst)
			assert.Equal(t, int32(0), s.ShorthandedGoalsAgainst)
			assert.Equal(t, int32(3), s.GoalsAgainst)
			assert.Equal(t, int32(0), s.PIM)
			assert.Equal(t, float32(57.6), s.TOI)
			assert.True(t, s.Starter)
			require.NotNil(t, s.Decision)
			assert.Equal(t, "L", *s.Decision)
		case int64(8481692):
			wolfTested = true
			assert.Equal(t, int32(32), s.SweaterNumber)
			assert.Equal(t, int32(0), s.ShotsAgainst)
			assert.Equal(t, float32(0), s.TOI)
			assert.False(t, s.Starter)
			assert.Nil(t, s.SavePctg)
			assert.Nil(t, s.Decision)
		}
	}
	assert.True(t, vladarTested, "Dan Vladar statline should exist")
	assert.True(t, wolfTested, "player 8481692 statline should exist")
}
