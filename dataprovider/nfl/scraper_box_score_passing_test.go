//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/stretchr/testify/assert"
)

func TestPassingBoxScoreScraper(t *testing.T) {
	s := NewPassingBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, dalPhiDate, dalPhiEventID)
	assert.Equal(t, 2, len(records), "2 passers")
	r := findPlayer(t, records, func(r model.PassingBoxScore) string { return r.PlayerID }, "00-0033077")
	assert.Equal(t, dalPhiEventID, r.EventID)
	assert.Equal(t, dalTeamID, r.TeamID)
	assert.Equal(t, "Dallas Cowboys", r.Team)
	assert.Equal(t, phiTeamID, r.OpponentID)
	assert.Equal(t, "Philadelphia Eagles", r.Opponent)
	assert.Equal(t, "32005052-4528-5723-d1b2-96e92ebc1241", r.PersonID)
	assert.Equal(t, "Dak Prescott", r.Player)
	assert.Equal(t, "D.Prescott", r.PlayerShortName)
	assert.Equal(t, "04", r.JerseyNumber)
	assert.Equal(t, int32(21), r.Completions)
	assert.Equal(t, int32(34), r.Attempts)
	assert.Equal(t, int32(188), r.Yards)
	assert.Equal(t, float32(61.8), r.CompletionPercent)
	assert.Equal(t, float32(5.53), r.YardsAverage)
	assert.Equal(t, int32(0), r.Touchdowns)
	assert.Equal(t, int32(0), r.Interceptions)
	assert.Equal(t, int32(32), r.Long)
	assert.Equal(t, int32(0), r.TimesSacked)
	assert.Equal(t, float32(76.6), r.Rating)
}

func TestPassingBoxScoreScraperPostSeason(t *testing.T) {
	// https://api.nfl.com/football/v2/stats/live/player-statistics/1ab6b0a3-f037-11f0-9442-5911216651e2
	// BUF @ DEN, AFC divisional round
	s := NewPassingBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, "2026-01-17", "1ab6b0a3-f037-11f0-9442-5911216651e2")
	assert.Equal(t, 2, len(records), "2 passers")
	r := findPlayer(t, records, func(r model.PassingBoxScore) string { return r.PlayerID }, "00-0034857")
	assert.Equal(t, "Buffalo Bills", r.Team)
	assert.Equal(t, "Denver Broncos", r.Opponent)
	assert.Equal(t, "Josh Allen", r.Player)
	assert.Equal(t, int32(25), r.Completions)
	assert.Equal(t, int32(39), r.Attempts)
	assert.Equal(t, int32(283), r.Yards)
	assert.Equal(t, int32(3), r.Touchdowns)
	assert.Equal(t, int32(2), r.Interceptions)
}

func TestPassingBoxScoreScraperProBowl(t *testing.T) {
	// https://api.nfl.com/football/v2/stats/live/player-statistics/10012015-0125-0020-8b38-7e57a77e8e95
	s := NewPassingBoxScoreScraper()
	s.Fetcher = testFetcher
	records := runBoxScore(t, s, proBowlDate, proBowlEventID)
	assert.Equal(t, 6, len(records), "6 passers")
	for _, r := range records {
		assert.Contains(t, []string{afcProBowlID, nfcProBowlID}, r.TeamID, "statlines carry the matchup's team IDs")
	}
	r := findPlayer(t, records, func(r model.PassingBoxScore) string { return r.PlayerID }, "00-0020531")
	assert.Equal(t, "Drew Brees", r.Player)
	assert.Equal(t, afcProBowlID, r.TeamID)
	assert.Equal(t, "AFC Pro Bowl Team", r.Team)
	assert.Equal(t, nfcProBowlID, r.OpponentID)
	assert.Equal(t, "NFC Pro Bowl Team", r.Opponent)
	assert.Equal(t, int32(13), r.Completions)
	assert.Equal(t, int32(21), r.Attempts)
	assert.Equal(t, int32(163), r.Yards)
}
