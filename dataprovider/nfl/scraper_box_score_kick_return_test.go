//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestKickReturnBoxScoreScraper(t *testing.T) {
	s := NewKickReturnBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 4, len(records), "4 kick returners")
	r := findPlayer(t, records, func(r model.KickReturnBoxScore) string { return r.PlayerID }, "00-0037801")
	assert.Equal(t, dalTeamID, r.TeamID)
	assert.Equal(t, "KaVontae Turpin", r.Player)
	assert.Equal(t, "K.Turpin", r.PlayerShortName)
	assert.Equal(t, int32(4), r.Returns)
	assert.Equal(t, int32(81), r.Yards)
	assert.Equal(t, float32(20.3), r.Average)
	assert.Equal(t, int32(27), r.Longest)
}
