//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestInterceptionsBoxScoreScraper(t *testing.T) {
	s := NewInterceptionsBoxScoreScraper()
	s.Fetcher = testFetcher
	// DAL @ PHI had no interceptions
	assert.Equal(t, 0, len(runBoxScore(t, s, dalPhiDate, dalPhiEventID)), "0 interceptions")

	// https://api.nfl.com/football/v2/stats/live/player-statistics/f61df2b7-311e-11f0-b670-ae1250fadad1
	// MIN @ CHI
	records := runBoxScore(t, s, "2025-09-08", "f61df2b7-311e-11f0-b670-ae1250fadad1")
	r := findPlayer(t, records, func(r model.InterceptionsBoxScore) string { return r.PlayerID }, "00-0036986")
	assert.Equal(t, "Chicago Bears", r.Team)
	assert.Equal(t, "Minnesota Vikings", r.Opponent)
	assert.Equal(t, "Nahshon Wright", r.Player)
	assert.Equal(t, "N.Wright", r.PlayerShortName)
	assert.Equal(t, int32(1), r.Interceptions)
	assert.Equal(t, int32(74), r.Yards)
	assert.Equal(t, int32(74), r.Long)
	assert.Equal(t, int32(74), r.LongestTouchdown)
	assert.Equal(t, int32(1), r.Touchdowns)
}
