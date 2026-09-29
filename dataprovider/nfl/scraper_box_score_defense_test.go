//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestDefenseBoxScoreScraper(t *testing.T) {
	s := NewDefenseBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 37, len(records), "37 defenders")
	r := findPlayer(t, records, func(r model.DefenseBoxScore) string { return r.PlayerID }, "00-0040708")
	assert.Equal(t, phiTeamID, r.TeamID)
	assert.Equal(t, "Jihaad Campbell", r.Player)
	assert.Equal(t, "J.Campbell", r.PlayerShortName)
	assert.Equal(t, float32(0), r.Tackles)
	assert.Equal(t, int32(3), r.TacklesAssists)
	assert.Equal(t, float32(3), r.TacklesCombined)
	assert.Equal(t, int32(1), r.PassesDefended)
	assert.Equal(t, int32(1), r.FumblesForced)
}
