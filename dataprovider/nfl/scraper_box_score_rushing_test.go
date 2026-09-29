//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestRushingBoxScoreScraper(t *testing.T) {
	s := NewRushingBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 8, len(records), "8 rushers")
	r := findPlayer(t, records, func(r model.RushingBoxScore) string { return r.PlayerID }, "00-0036389")
	assert.Equal(t, phiTeamID, r.TeamID)
	assert.Equal(t, "Jalen Hurts", r.Player)
	assert.Equal(t, "J.Hurts", r.PlayerShortName)
	assert.Equal(t, int32(14), r.Attempts)
	assert.Equal(t, int32(62), r.Yards)
	assert.Equal(t, float32(4.4), r.Average)
	assert.Equal(t, int32(2), r.Touchdowns)
	assert.Equal(t, int32(15), r.Long)
	assert.Equal(t, int32(8), r.LongestTouchdown)
}
