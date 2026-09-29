//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestFumblesBoxScoreScraper(t *testing.T) {
	s := NewFumblesBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 3, len(records), "3 players with a fumble stat")
	r := findPlayer(t, records, func(r model.FumblesBoxScore) string { return r.PlayerID }, "00-0035243")
	assert.Equal(t, dalTeamID, r.TeamID)
	assert.Equal(t, "Miles Sanders", r.Player)
	assert.Equal(t, "M.Sanders", r.PlayerShortName)
	assert.Equal(t, int32(1), r.Fumbles)
	assert.Equal(t, int32(1), r.Lost)
}
