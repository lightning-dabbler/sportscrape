//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestKickoffBoxScoreScraper(t *testing.T) {
	s := NewKickoffBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 2, len(records), "2 kickoff specialists")
	r := findPlayer(t, records, func(r model.KickoffBoxScore) string { return r.PlayerID }, "00-0033787")
	assert.Equal(t, phiTeamID, r.TeamID)
	assert.Equal(t, "Jake Elliott", r.Player)
	assert.Equal(t, "J.Elliott", r.PlayerShortName)
	assert.Equal(t, int32(5), r.Kickoffs)
	assert.Equal(t, int32(298), r.Yards)
	assert.Equal(t, int32(1), r.Inside20)
	assert.Equal(t, int32(101), r.ReturnYards)
}
