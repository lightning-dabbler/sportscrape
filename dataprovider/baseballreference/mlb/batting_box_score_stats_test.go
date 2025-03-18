//go:build integration

package mlb

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb/model"
	"github.com/stretchr/testify/assert"
)

func TestGetBattingBoxScoreStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	date := "2024-10-13"

	matchupRunner := NewMatchupRunner(
		WithMatchupTimeout(2 * time.Minute),
	)
	matchups := matchupRunner.GetMatchups(date)
	boxScoreRunner := NewBattingBoxScoreRunner(
		WithBattingBoxScoreTimeout(4*time.Minute),
		WithBattingBoxScoreConcurrency(1),
	)
	boxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)

	playerToTest := map[string]bool{
		"Harrison Bader":    false,
		"Shohei Ohtani":     false,
		"Enrique Hernández": false,
	}
	assert.Equal(t, 24, len(boxScoreStats), "24 statlines")
	expectedNumPlayersNoAtBat := 2
	actualNumPlayersNoAtBat := 0
	for _, statline := range boxScoreStats {
		stats := statline.(model.MLBBattingBoxScoreStats)
		if stats.Player == "Harrison Bader" {
			playerToTest["Harrison Bader"] = true
			assert.Equal(t, "New York Mets", stats.Team)
			assert.Equal(t, "Los Angeles Dodgers", stats.Opponent)
			assert.Equal(t, "baderha01", stats.PlayerID)
			assert.Equal(t, "https://www.baseball-reference.com/players/b/baderha01.shtml", stats.PlayerLink)
			assert.Equal(t, "CF", stats.Position)
			assert.Equal(t, 0, stats.AtBat)
			assert.Equal(t, 0, stats.Runs)
			assert.Equal(t, 0, stats.Hits)
			assert.Equal(t, 0, stats.RunsBattedIn)
			assert.Equal(t, 0, stats.Walks)
			assert.Equal(t, 0, stats.Strikeouts)
			assert.Equal(t, 0, stats.PlateAppearances)
			assert.Equal(t, float32(0.167), *stats.BattingAverage)
			assert.Equal(t, float32(0.167), *stats.OnBasePercentage)
			assert.Equal(t, float32(0.167), *stats.SluggingPercentage)
			assert.Equal(t, float32(0.333), *stats.OnBasePlusSlugging)
			assert.Nil(t, stats.PitchesPerPlateAppearance)
			assert.Nil(t, stats.Strikes)
			assert.Nil(t, stats.WinProbabilityAdded)
			assert.Nil(t, stats.AverageLeverageIndex)
			assert.Nil(t, stats.SumPositiveWinProbabilityAdded)
			assert.Nil(t, stats.SumNegativeWinProbabilityAdded)
			assert.Nil(t, stats.ChampionshipWinProbabilityAdded)
			assert.Nil(t, stats.AverageChampionshipLeverageIndex)
			assert.Nil(t, stats.BaseOutRunsAdded)
			assert.Equal(t, 1, stats.Putout)
			assert.Equal(t, 0, stats.Assist)

		} else if stats.Player == "Shohei Ohtani" {
			playerToTest["Shohei Ohtani"] = true
			assert.Equal(t, "Los Angeles Dodgers", stats.Team)
			assert.Equal(t, "New York Mets", stats.Opponent)
			assert.Equal(t, "ohtansh01", stats.PlayerID)
			assert.Equal(t, "https://www.baseball-reference.com/players/o/ohtansh01.shtml", stats.PlayerLink)
			assert.Equal(t, "DH", stats.Position)
			assert.Equal(t, 4, stats.AtBat)
			assert.Equal(t, 2, stats.Runs)
			assert.Equal(t, 2, stats.Hits)
			assert.Equal(t, 1, stats.RunsBattedIn)
			assert.Equal(t, 1, stats.Walks)
			assert.Equal(t, 0, stats.Strikeouts)
			assert.Equal(t, 5, stats.PlateAppearances)
			assert.Equal(t, float32(0.250), *stats.BattingAverage)
			assert.Equal(t, float32(0.333), *stats.OnBasePercentage)
			assert.Equal(t, float32(0.375), *stats.SluggingPercentage)
			assert.Equal(t, float32(0.708), *stats.OnBasePlusSlugging)
			assert.Equal(t, 14, *stats.PitchesPerPlateAppearance)
			assert.Equal(t, 6, *stats.Strikes)
			assert.Equal(t, float32(0.065), *stats.WinProbabilityAdded)
			assert.Equal(t, float32(0.42), *stats.AverageLeverageIndex)
			assert.Equal(t, float32(0.099), *stats.SumPositiveWinProbabilityAdded)
			assert.Equal(t, float32(-0.035), *stats.SumNegativeWinProbabilityAdded)
			assert.Equal(t, float32(0.95), *stats.ChampionshipWinProbabilityAdded)
			assert.Equal(t, float32(11.01), *stats.AverageChampionshipLeverageIndex)
			assert.Equal(t, float32(2.1), *stats.BaseOutRunsAdded)
			assert.Equal(t, 0, stats.Putout)
			assert.Equal(t, 0, stats.Assist)
		} else if stats.Player == "Enrique Hernández" {
			playerToTest["Enrique Hernández"] = true
			assert.Equal(t, "CF-2B-3B", stats.Position)
		}

		if stats.AtBat == 0 {
			actualNumPlayersNoAtBat += 1
		}

		assert.Equal(t, "LAN202410130", stats.EventID, "The event ID should be the same across all records as there's only one matchup for this event date")

	}
	assert.Equal(t, expectedNumPlayersNoAtBat, actualNumPlayersNoAtBat, "Only two players across teams have zero at bat stats")
	assert.True(t, playerToTest["Harrison Bader"], "Harrison Bader's batting stats were collected")
	assert.True(t, playerToTest["Shohei Ohtani"], "Shohei Ohtani's batting stats were collected")
	assert.True(t, playerToTest["Enrique Hernández"], "Enrique Hernández's batting stats were collected")
}
