//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestKickingBoxScoreScraper(t *testing.T) {
	s := NewKickingBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 2, len(records), "2 kickers")
	r := findPlayer(t, records, func(r model.KickingBoxScore) string { return r.PlayerID }, "00-0037692")
	assert.Equal(t, dalTeamID, r.TeamID)
	assert.Equal(t, "Brandon Aubrey", r.Player)
	assert.Equal(t, "B.Aubrey", r.PlayerShortName)
	assert.Equal(t, int32(2), r.FieldGoalsMade)
	assert.Equal(t, int32(2), r.FieldGoalsAttempted)
	assert.Equal(t, int32(53), r.FieldGoalsLongestMade)
	assert.Equal(t, float32(47), r.FieldGoalsAverageLength)
	assert.Equal(t, int32(94), r.FieldGoalsTotalYards)
	assert.Equal(t, int32(2), r.ExtraPointsMade)
	assert.Equal(t, int32(2), r.ExtraPointsAttempted)
}
