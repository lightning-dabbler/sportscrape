//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestReceivingBoxScoreScraper(t *testing.T) {
	s := NewReceivingBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 15, len(records), "15 receivers")
	r := findPlayer(t, records, func(r model.ReceivingBoxScore) string { return r.PlayerID }, "00-0036358")
	assert.Equal(t, dalTeamID, r.TeamID)
	assert.Equal(t, "CeeDee Lamb", r.Player)
	assert.Equal(t, "C.Lamb", r.PlayerShortName)
	assert.Equal(t, int32(13), r.Targets)
	assert.Equal(t, int32(7), r.Receptions)
	assert.Equal(t, int32(110), r.Yards)
	assert.Equal(t, float32(15.7), r.Average)
	assert.Equal(t, int32(32), r.Long)
	assert.Equal(t, int32(44), r.YardsAfterCatch)
}
