//go:build integration

package nhl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlayByPlayScraper(t *testing.T) {
	// https://api-web.nhle.com/v1/gamecenter/2024020250/play-by-play
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2024-11-12", 2024020250)
	scraper := NewPlayByPlayScraper()
	scraper.Fetcher = testFetcher
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper: scraper,
		},
	)
	plays, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	assert.Equal(t, 323, len(plays), "323 plays")

	goalTested := false
	penaltyTested := false
	for _, p := range plays {
		switch p.PlayEventID {
		case int64(464):
			goalTested = true
			assert.Equal(t, int64(2024020250), p.EventID)
			assert.Equal(t, int64(23), p.HomeTeamID)
			assert.Equal(t, int64(20), p.AwayTeamID)
			assert.Equal(t, int32(284), p.SortOrder)
			assert.Equal(t, int32(1), p.Period)
			assert.Equal(t, "REG", p.PeriodType)
			assert.Equal(t, float32(19), p.TimeInPeriod)
			assert.Equal(t, float32(1), p.TimeRemaining)
			assert.Equal(t, "1551", p.SituationCode)
			assert.Equal(t, "right", p.HomeTeamDefendingSide)
			assert.Equal(t, int32(505), p.TypeCode)
			assert.Equal(t, "goal", p.TypeDescKey)
			assert.Equal(t, int64(20), *p.EventOwnerTeamID)
			assert.Equal(t, int32(82), *p.XCoord)
			assert.Equal(t, int32(3), *p.YCoord)
			assert.Equal(t, "O", *p.ZoneCode)
			assert.Equal(t, "wrist", *p.ShotType)
			assert.Equal(t, int64(8477993), *p.ScoringPlayerID)
			assert.Equal(t, int32(2), *p.ScoringPlayerTotal)
			assert.Equal(t, int64(8480947), *p.GoalieInNetID)
			assert.Equal(t, int32(1), *p.AwayScore)
			assert.Equal(t, int32(0), *p.HomeScore)
		case int64(32):
			penaltyTested = true
			assert.Equal(t, int32(298), p.SortOrder)
			assert.Equal(t, float32(20), p.TimeInPeriod)
			assert.Equal(t, float32(0), p.TimeRemaining)
			assert.Equal(t, "1560", p.SituationCode)
			assert.Equal(t, int32(509), p.TypeCode)
			assert.Equal(t, "penalty", p.TypeDescKey)
			assert.Equal(t, "MIN", *p.PenaltyTypeCode)
			assert.Equal(t, "high-sticking", *p.PenaltyDescKey)
			assert.Equal(t, int32(2), *p.PenaltyDuration)
			assert.Equal(t, int64(8482624), *p.CommittedByPlayerID)
			assert.Equal(t, int64(8480012), *p.DrawnByPlayerID)
			assert.Equal(t, int64(20), *p.EventOwnerTeamID)
			assert.Equal(t, "D", *p.ZoneCode)
		}
	}
	assert.True(t, goalTested, "goal play 464 should exist")
	assert.True(t, penaltyTested, "penalty play 32 should exist")
}
