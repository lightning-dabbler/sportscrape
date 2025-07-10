package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/stretchr/testify/assert"
)

func TestMLBPitchingBoxScoreScraper(t *testing.T) {
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

	boxscorescraper := NewPitchingBoxScoreScraper()
	boxscorerunner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerScraper(boxscorescraper),
		sportscrape.EventDataRunnerConcurrency(1),
	)
	boxScoreStats, err := boxscorerunner.Run(matchups...)
	assert.NoError(t, err)
	assert.Equal(t, 13, len(boxScoreStats), "13 statlines")
	GerritColeTested := false
	for _, statline := range boxScoreStats {
		s := statline.(model.PitchingBoxScore)
		if s.Player == "Gerrit Cole" {
			GerritColeTested = true
			assert.Equal(t, int64(775296), s.EventID)
			assert.Equal(t, int64(147), s.TeamID)
			assert.Equal(t, "New York Yankees", s.Team)
			assert.Equal(t, "Los Angeles Dodgers", s.Opponent)
			assert.Equal(t, int64(119), s.OpponentID)
			assert.Equal(t, int64(543037), s.PlayerID)
			assert.Equal(t, "Pitcher", s.Position)
			assert.Equal(t, int32(6), s.FlyOuts)
			assert.Equal(t, int32(6), s.GroundOuts)
			assert.Equal(t, int32(10), s.AirOuts)
			assert.Equal(t, int32(5), s.Runs)
			assert.Equal(t, int32(1), s.Doubles)
			assert.Equal(t, int32(0), s.Triples)
			assert.Equal(t, int32(0), s.HomeRuns)
			assert.Equal(t, int32(6), s.Strikeouts)
			assert.Equal(t, int32(4), s.Walks)
			assert.Equal(t, int32(0), s.IntentionalWalks)
			assert.Equal(t, int32(4), s.Hits)
			assert.Equal(t, int32(0), s.HitByPitch)
			assert.Equal(t, int32(26), s.AtBats)
			assert.Equal(t, int32(0), s.CaughtStealing)
			assert.Equal(t, int32(0), s.StolenBases)
			assert.Equal(t, int32(108), s.NumberOfPitches)
			assert.Equal(t, float32(6.2), s.InningsPitched)
			assert.Equal(t, int32(0), s.Wins)
			assert.Equal(t, int32(0), s.Losses)
			assert.Equal(t, int32(0), s.Saves)
			assert.Equal(t, int32(0), s.BlownSaves)
			assert.Equal(t, int32(0), s.EarnedRuns)
			assert.Equal(t, int32(30), s.BattersFaced)
			assert.Equal(t, int32(20), s.Outs)
			assert.Equal(t, int32(0), s.Shutouts)
			assert.Equal(t, int32(32), s.Balls)
			assert.Equal(t, int32(76), s.Strikes)
			assert.Equal(t, int32(0), s.Balks)
			assert.Equal(t, int32(0), s.WildPitches)
			assert.Equal(t, int32(0), s.Pickoffs)
			assert.Equal(t, int32(5), s.RBI)
			assert.Equal(t, int32(0), s.InheritedRunners)
			assert.Equal(t, int32(0), s.InheritedRunnersScored)
			assert.Equal(t, int32(0), s.CatchersInterference)
			assert.Equal(t, int32(0), s.SacBunts)
			assert.Equal(t, int32(0), s.SacFlies)
			assert.Equal(t, int32(0), s.PassedBall)
			assert.Equal(t, int32(1), s.PopOuts)
			assert.Equal(t, int32(3), s.LineOuts)
		}
	}
	assert.True(t, GerritColeTested, "Gerrit Cole statline tested")
}
