//go:build integration

package nhl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkaterBoxScoreScraper(t *testing.T) {
	// https://api-web.nhle.com/v1/gamecenter/2024020250/boxscore
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2024-11-12", 2024020250)
	scraper := NewSkaterBoxScoreScraper()
	scraper.Fetcher = testFetcher
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.SkaterBoxScore]{
			Scraper: scraper,
		},
	)
	statlines, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	assert.Equal(t, 36, len(statlines), "36 statlines")

	forwards, defense := 0, 0
	huberdeauTested := false
	anderssonTested := false
	for _, s := range statlines {
		switch s.Position {
		case "C", "L", "R":
			forwards++
		case "D":
			defense++
		}
		switch s.PlayerID {
		case int64(8476456):
			huberdeauTested = true
			assert.Equal(t, int64(2024020250), s.EventID)
			assert.Equal(t, int64(20), s.TeamID)
			assert.Equal(t, "Flames", s.Team)
			assert.Equal(t, int64(23), s.OpponentID)
			assert.Equal(t, "Canucks", s.Opponent)
			assert.Equal(t, "Jonathan Huberdeau", s.Player)
			assert.Equal(t, int32(10), s.SweaterNumber)
			assert.Equal(t, "L", s.Position)
			assert.Equal(t, int32(0), s.Goals)
			assert.Equal(t, int32(0), s.Assists)
			assert.Equal(t, int32(0), s.Points)
			assert.Equal(t, int32(-1), s.PlusMinus)
			assert.Equal(t, int32(0), s.PIM)
			assert.Equal(t, int32(1), s.Hits)
			assert.Equal(t, int32(0), s.PowerPlayGoals)
			assert.Equal(t, int32(0), s.SOG)
			require.NotNil(t, s.FaceoffWinningPctg)
			assert.Equal(t, float32(0), *s.FaceoffWinningPctg)
			require.NotNil(t, s.TOI)
			assert.Equal(t, float32(18.75), *s.TOI)
			assert.Equal(t, int32(0), s.BlockedShots)
			assert.Equal(t, int32(20), s.Shifts)
			require.NotNil(t, s.Giveaways)
			assert.Equal(t, int32(2), *s.Giveaways)
			require.NotNil(t, s.Takeaways)
			assert.Equal(t, int32(0), *s.Takeaways)
		case int64(8478397):
			anderssonTested = true
			assert.Equal(t, int64(20), s.TeamID)
			assert.Equal(t, "Flames", s.Team)
			assert.Equal(t, "Rasmus Andersson", s.Player)
			assert.Equal(t, int32(4), s.SweaterNumber)
			assert.Equal(t, "D", s.Position)
			assert.Equal(t, int32(0), s.PlusMinus)
			assert.Equal(t, int32(0), s.Hits)
			assert.Equal(t, int32(2), s.SOG)
			require.NotNil(t, s.TOI)
			assert.Equal(t, float32(24.9), *s.TOI)
			assert.Equal(t, int32(33), s.Shifts)
			require.NotNil(t, s.Giveaways)
			assert.Equal(t, int32(1), *s.Giveaways)
			require.NotNil(t, s.Takeaways)
			assert.Equal(t, int32(0), *s.Takeaways)
		}
	}
	assert.Equal(t, 24, forwards, "24 forwards")
	assert.Equal(t, 12, defense, "12 defense")
	assert.True(t, huberdeauTested, "Jonathan Huberdeau statline should exist")
	assert.True(t, anderssonTested, "Rasmus Andersson statline should exist")
}
