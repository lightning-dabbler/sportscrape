//go:build integration

package mlb

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb/model"
	"github.com/stretchr/testify/assert"
)

func TestGetPitchingBoxScoreStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	date := "2024-10-30"

	matchupRunner := NewMatchupRunner(
		WithMatchupTimeout(5 * time.Minute),
	)
	matchups := matchupRunner.GetMatchups(date)
	boxScoreRunner := NewPitchingBoxScoreRunner(
		WithPitchingBoxScoreTimeout(6*time.Minute),
		WithPitchingBoxScoreConcurrency(1),
	)
	boxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)

	playerToTest := map[string]bool{
		"Luke Weaver":   false,
		"Gerrit Cole":   false,
		"Jack Flaherty": false,
	}
	assert.Equal(t, 13, len(boxScoreStats), "13 statlines")
	for _, statline := range boxScoreStats {
		stats := statline.(model.MLBPitchingBoxScoreStats)
		if stats.Player == "Luke Weaver" {
			playerToTest["Luke Weaver"] = true
			assert.Equal(t, "New York Yankees", stats.Team)
			assert.Equal(t, "Los Angeles Dodgers", stats.Opponent)
			assert.Equal(t, "weavelu01", stats.PlayerID)
			assert.Equal(t, "https://www.baseball-reference.com/players/w/weavelu01.shtml", stats.PlayerLink)
			assert.Equal(t, int32(4), stats.PitchingOrder)
			assert.Equal(t, float32(1.1), stats.InningsPitched)
			assert.Equal(t, int32(1), stats.HitsAllowed)
			assert.Equal(t, int32(0), stats.RunsAllowed)
			assert.Equal(t, int32(0), stats.EarnedRunsAllowed)
			assert.Equal(t, int32(1), stats.Walks)
			assert.Equal(t, int32(1), stats.Strikeouts)
			assert.Equal(t, int32(0), stats.HomeRunsAllowed)
			assert.Equal(t, float32(1.76), stats.EarnedRunAverage)
			assert.Equal(t, int32(7), stats.BattersFaced)
			assert.Equal(t, int32(22), stats.PitchesPerPlateAppearance)
			assert.Equal(t, int32(13), stats.Strikes)
			assert.Equal(t, int32(7), stats.StrikesByContact)
			assert.Equal(t, int32(3), stats.StrikesSwinging)
			assert.Equal(t, int32(2), stats.StrikesLooking)
			assert.Equal(t, int32(1), stats.GroundBalls)
			assert.Equal(t, int32(3), stats.FlyBalls)
			assert.Equal(t, int32(0), stats.LineDrives)
			assert.Equal(t, int32(0), stats.UnknownBattedBallType)
			assert.Nil(t, stats.GameScore)
			assert.Equal(t, int32(3), *stats.InheritedRunners)
			assert.Equal(t, int32(2), *stats.InheritedScore)
			assert.Equal(t, float32(-0.096), stats.WinProbabilityAdded)
			assert.Equal(t, float32(2.33), stats.AverageLeverageIndex)
			assert.Equal(t, float32(-2.25), stats.ChampionshipWinProbabilityAdded)
			assert.Equal(t, float32(85.68), stats.AverageChampionshipLeverageIndex)
			assert.Equal(t, float32(-0.1), stats.BaseOutRunsSaved)
		} else if stats.Player == "Gerrit Cole" {
			playerToTest["Gerrit Cole"] = true
			assert.Equal(t, "New York Yankees", stats.Team)
			assert.Equal(t, "Los Angeles Dodgers", stats.Opponent)
			assert.Equal(t, "colege01", stats.PlayerID)
			assert.Equal(t, "https://www.baseball-reference.com/players/c/colege01.shtml", stats.PlayerLink)
			assert.Equal(t, int32(1), stats.PitchingOrder)
			assert.Equal(t, float32(6.2), stats.InningsPitched)
			assert.Equal(t, int32(4), stats.HitsAllowed)
			assert.Equal(t, int32(5), stats.RunsAllowed)
			assert.Equal(t, int32(0), stats.EarnedRunsAllowed)
			assert.Equal(t, int32(4), stats.Walks)
			assert.Equal(t, int32(6), stats.Strikeouts)
			assert.Equal(t, int32(0), stats.HomeRunsAllowed)
			assert.Equal(t, float32(2.17), stats.EarnedRunAverage)
			assert.Equal(t, int32(30), stats.BattersFaced)
			assert.Equal(t, int32(108), stats.PitchesPerPlateAppearance)
			assert.Equal(t, int32(76), stats.Strikes)
			assert.Equal(t, int32(46), stats.StrikesByContact)
			assert.Equal(t, int32(9), stats.StrikesSwinging)
			assert.Equal(t, int32(21), stats.StrikesLooking)
			assert.Equal(t, int32(7), stats.GroundBalls)
			assert.Equal(t, int32(13), stats.FlyBalls)
			assert.Equal(t, int32(5), stats.LineDrives)
			assert.Equal(t, int32(0), stats.UnknownBattedBallType)
			assert.Equal(t, int32(58), *stats.GameScore)
			assert.Nil(t, stats.InheritedRunners)
			assert.Nil(t, stats.InheritedScore)
			assert.Equal(t, float32(-0.101), stats.WinProbabilityAdded)
			assert.Equal(t, float32(1.04), stats.AverageLeverageIndex)
			assert.Equal(t, float32(-2.10), stats.ChampionshipWinProbabilityAdded)
			assert.Equal(t, float32(38.13), stats.AverageChampionshipLeverageIndex)
			assert.Equal(t, float32(-1.8), stats.BaseOutRunsSaved)
		} else if stats.Player == "Jack Flaherty" {
			playerToTest["Jack Flaherty"] = true
			assert.Equal(t, "Los Angeles Dodgers", stats.Team)
			assert.Equal(t, "New York Yankees", stats.Opponent)
			assert.Equal(t, "flaheja01", stats.PlayerID)
			assert.Equal(t, "https://www.baseball-reference.com/players/f/flaheja01.shtml", stats.PlayerLink)
			assert.Equal(t, int32(1), stats.PitchingOrder)
			assert.Equal(t, float32(1.1), stats.InningsPitched)
			assert.Equal(t, int32(4), stats.HitsAllowed)
			assert.Equal(t, int32(4), stats.RunsAllowed)
			assert.Equal(t, int32(4), stats.EarnedRunsAllowed)
			assert.Equal(t, int32(1), stats.Walks)
			assert.Equal(t, int32(1), stats.Strikeouts)
			assert.Equal(t, int32(2), stats.HomeRunsAllowed)
			assert.Equal(t, float32(7.36), stats.EarnedRunAverage)
			assert.Equal(t, int32(9), stats.BattersFaced)
			assert.Equal(t, int32(35), stats.PitchesPerPlateAppearance)
			assert.Equal(t, int32(18), stats.Strikes)
			assert.Equal(t, int32(7), stats.StrikesByContact)
			assert.Equal(t, int32(3), stats.StrikesSwinging)
			assert.Equal(t, int32(8), stats.StrikesLooking)
			assert.Equal(t, int32(2), stats.GroundBalls)
			assert.Equal(t, int32(5), stats.FlyBalls)
			assert.Equal(t, int32(3), stats.LineDrives)
			assert.Equal(t, int32(0), stats.UnknownBattedBallType)
			assert.Equal(t, int32(30), *stats.GameScore)
			assert.Nil(t, stats.InheritedRunners)
			assert.Nil(t, stats.InheritedScore)
			assert.Equal(t, float32(-0.293), stats.WinProbabilityAdded)
			assert.Equal(t, float32(0.61), stats.AverageLeverageIndex)
			assert.Equal(t, float32(-6.06), stats.ChampionshipWinProbabilityAdded)
			assert.Equal(t, float32(22.46), stats.AverageChampionshipLeverageIndex)
			assert.Equal(t, float32(-3.5), stats.BaseOutRunsSaved)
		}
		assert.Equal(t, "NYA202410300", stats.EventID, "The event ID should be the same across all records as there's only one matchup for this event date")
	}
	assert.True(t, playerToTest["Luke Weaver"], "Luke Weaver's pitching stats were collected")
	assert.True(t, playerToTest["Gerrit Cole"], "Gerrit Cole's pitching stats were collected")
	assert.True(t, playerToTest["Jack Flaherty"], "Jack Flaherty's pitching stats were collected")
}
