package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreMatchupsScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2025-06-05"),
		WithMatchupTimeout(3*time.Minute),
	)
	matchupScraper.NetworkHeaders = NetworkHeaders
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper:   matchupScraper,
			KeepAlive: true,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		matchupScraper.Close()
		t.Fatal(err)
	}
	boxscorescraper := NewBoxScoreMatchupsScraper(
		WithBoxScoreMatchupsTimeout(3 * time.Minute),
	)
	boxscorescraper.DocumentRetriever = matchupScraper.DocumentRetriever
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreMatchups]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 189, n_records, "189 stat lines")
	shaiTested := false
	for _, s := range records {
		if s.PlayerName == "Shai Gilgeous-Alexander" && s.OpponentPlayerName == "Tyrese Haliburton" {
			assert.Equal(t, "0042400401", s.EventID)
			assert.Equal(t, int32(3), s.EventStatus)
			assert.Equal(t, "Final", s.EventStatusText)
			assert.Equal(t, int64(1610612760), s.TeamID)
			assert.Equal(t, "Thunder", s.TeamName)
			assert.Equal(t, "Oklahoma City Thunder", s.TeamNameFull)
			assert.Equal(t, int64(1610612754), s.OpponentID)
			assert.Equal(t, "Pacers", s.OpponentName)
			assert.Equal(t, "Indiana Pacers", s.OpponentNameFull)
			assert.Equal(t, int64(1628983), s.PlayerID)
			assert.Equal(t, "Shai Gilgeous-Alexander", s.PlayerName)
			assert.Equal(t, "G", s.Position)
			assert.Equal(t, true, s.Starter)

			assert.Equal(t, int64(1630169), s.OpponentPlayerID)
			assert.Equal(t, "Tyrese Haliburton", s.OpponentPlayerName)
			assert.Equal(t, float32(0.65), s.MatchupMinutes)
			assert.Equal(t, float32(39), s.MatchupMinutesSort)
			assert.Equal(t, float32(4.3), s.PartialPossessions)
			assert.Equal(t, float32(0.0491803), s.PercentageDefenderTotalTime)
			assert.Equal(t, float32(0.0463734), s.PercentageOffensiveTotalTime)
			assert.Equal(t, float32(0.048), s.PercentageTotalTimeBothOn)
			assert.Equal(t, int32(0), s.SwitchesOn)
			assert.Equal(t, int32(9), s.PlayerPoints)
			assert.Equal(t, int32(9), s.TeamPoints)
			assert.Equal(t, int32(0), s.MatchupAssists)
			assert.Equal(t, int32(0), s.MatchupPotentialAssists)
			assert.Equal(t, int32(0), s.MatchupTurnovers)
			assert.Equal(t, int32(0), s.MatchupBlocks)
			assert.Equal(t, int32(2), s.MatchupFieldGoalsMade)
			assert.Equal(t, int32(2), s.MatchupFieldGoalsAttempted)
			assert.Equal(t, float32(1), s.MatchupFieldGoalsPercentage)
			assert.Equal(t, int32(1), s.MatchupThreePointersMade)
			assert.Equal(t, int32(1), s.MatchupThreePointersAttempted)
			assert.Equal(t, float32(1), s.MatchupThreePointersPercentage)
			assert.Equal(t, int32(0), s.HelpBlocks)
			assert.Equal(t, int32(0), s.HelpFieldGoalsMade)
			assert.Equal(t, int32(0), s.HelpFieldGoalsAttempted)
			assert.Equal(t, float32(0), s.HelpFieldGoalsPercentage)
			assert.Equal(t, int32(4), s.MatchupFreeThrowsMade)
			assert.Equal(t, int32(4), s.MatchupFreeThrowsAttempted)
			assert.Equal(t, int32(1), s.ShootingFouls)
			shaiTested = true
		}
	}
	assert.True(t, shaiTested, "Shai Gilgeous-Alexander statline tested")
}
