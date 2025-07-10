//go:build integration

package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/stretchr/testify/assert"
)

func TestMLBBattingBoxScoreScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchupscraper := NewMatchupScraper(
		MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	boxscorescraper := NewBattingBoxScoreScraper()
	boxscorerunner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerScraper(boxscorescraper),
		sportscrape.EventDataRunnerConcurrency(1),
	)
	boxScoreStats, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	assert.Equal(t, 18, len(boxScoreStats), "18 statlines")
	GavinLuxTested := false
	for _, statline := range boxScoreStats {
		s := statline.(model.BattingBoxScore)
		if s.Player == "Gavin Lux" {
			GavinLuxTested = true
			assert.Equal(t, int64(775296), s.EventID)
			assert.Equal(t, int64(119), s.TeamID)
			assert.Equal(t, "Los Angeles Dodgers", s.Team)
			assert.Equal(t, "New York Yankees", s.Opponent)
			assert.Equal(t, int64(147), s.OpponentID)
			assert.Equal(t, int64(666158), s.PlayerID)
			assert.Equal(t, "Second Base", s.Position)
			assert.Equal(t, int32(2), s.FlyOuts)
			assert.Equal(t, int32(0), s.GroundOuts)
			assert.Equal(t, int32(2), s.AirOuts)
			assert.Equal(t, int32(0), s.Runs)
			assert.Equal(t, int32(0), s.Doubles)
			assert.Equal(t, int32(0), s.Triples)
			assert.Equal(t, int32(0), s.HomeRuns)
			assert.Equal(t, int32(1), s.Strikeouts)
			assert.Equal(t, int32(1), s.Walks)
			assert.Equal(t, int32(0), s.IntentionalWalks)
			assert.Equal(t, int32(0), s.Hits)
			assert.Equal(t, int32(0), s.HitByPitch)
			assert.Equal(t, int32(2), s.AtBats)
			assert.Equal(t, int32(0), s.CaughtStealing)
			assert.Equal(t, int32(0), s.StolenBases)
			assert.Equal(t, int32(0), s.GroundIntoDoublePlay)
			assert.Equal(t, int32(0), s.GroundIntoTriplePlay)
			assert.Equal(t, int32(4), s.PlateAppearances)
			assert.Equal(t, int32(0), s.TotalBases)
			assert.Equal(t, int32(1), s.RBI)
			assert.Equal(t, int32(3), s.LeftOnBase)
			assert.Equal(t, int32(0), s.SacBunts)
			assert.Equal(t, int32(1), s.SacFlies)
			assert.Equal(t, int32(0), s.CatchersInterference)
			assert.Equal(t, int32(0), s.Pickoffs)
			assert.Equal(t, float32(0), s.AtBatsPerHomeRun)
			assert.Equal(t, int32(0), s.PopOuts)
			assert.Equal(t, int32(0), s.LineOuts)
		}
	}
	assert.True(t, GavinLuxTested, "Gavin Lux statline tested")

}
