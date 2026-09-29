//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestPuntingBoxScoreScraper(t *testing.T) {
	s := NewPuntingBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 2, len(records), "2 punters")
	r := findPlayer(t, records, func(r model.PuntingBoxScore) string { return r.PlayerID }, "00-0029692")
	assert.Equal(t, dalTeamID, r.TeamID)
	assert.Equal(t, "Bryan Anger", r.Player)
	assert.Equal(t, "B.Anger", r.PlayerShortName)
	assert.Equal(t, int32(2), r.Punts)
	assert.Equal(t, int32(87), r.Yards)
	assert.Equal(t, float32(43.5), r.AverageGross)
	assert.Equal(t, float32(43.5), r.AverageNet)
	assert.Equal(t, int32(49), r.Longest)
	assert.Equal(t, int32(1), r.Inside20)
}
