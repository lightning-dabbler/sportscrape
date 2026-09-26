package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreMatchupsScraper(t *testing.T) {
	// https://www.nba.com/game/tor-vs-cle-0022500225/box-score?type=matchups
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2025-11-13"),
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
	// 2025-11-13 has 3 games; only scrape TOR @ CLE
	var tested []model.Matchup
	for _, m := range matchups {
		if m.EventID == "0022500225" {
			tested = append(tested, m)
		}
	}
	if len(tested) != 1 {
		matchupScraper.Close()
		t.Fatalf("expected event 0022500225 on 2025-11-13, got %d matching matchups", len(tested))
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

	records, err := boxscorerunner.Run(tested)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 211, n_records, "211 stat lines")
	mitchellTested := false
	for _, s := range records {
		if s.PlayerName == "Donovan Mitchell" && s.OpponentPlayerName == "RJ Barrett" {
			assert.Equal(t, "0022500225", s.EventID)
			assert.Equal(t, int32(3), s.EventStatus)
			assert.Equal(t, "Final", s.EventStatusText)
			assert.Equal(t, int64(1610612739), s.TeamID)
			assert.Equal(t, "Cavaliers", s.TeamName)
			assert.Equal(t, "Cleveland Cavaliers", s.TeamNameFull)
			assert.Equal(t, int64(1610612761), s.OpponentID)
			assert.Equal(t, "Raptors", s.OpponentName)
			assert.Equal(t, "Toronto Raptors", s.OpponentNameFull)
			assert.Equal(t, int64(1628378), s.PlayerID)
			assert.Equal(t, "Donovan Mitchell", s.PlayerName)
			assert.Equal(t, "G", s.Position)
			assert.Equal(t, true, s.Starter)

			assert.Equal(t, int64(1629628), s.OpponentPlayerID)
			assert.Equal(t, "RJ Barrett", s.OpponentPlayerName)
			assert.Equal(t, float32(2.8), s.MatchupMinutes)
			assert.Equal(t, float32(168), s.MatchupMinutesSort)
			assert.Equal(t, float32(12.5), s.PartialPossessions)
			assert.Equal(t, float32(0.2645669), s.PercentageDefenderTotalTime)
			assert.Equal(t, float32(0.2148338), s.PercentageOffensiveTotalTime)
			assert.Equal(t, float32(0.327), s.PercentageTotalTimeBothOn)
			assert.Equal(t, int32(0), s.SwitchesOn)
			assert.Equal(t, int32(6), s.PlayerPoints)
			assert.Equal(t, int32(10), s.TeamPoints)
			assert.Equal(t, int32(3), s.MatchupAssists)
			assert.Equal(t, int32(0), s.MatchupPotentialAssists)
			assert.Equal(t, int32(1), s.MatchupTurnovers)
			assert.Equal(t, int32(0), s.MatchupBlocks)
			assert.Equal(t, int32(1), s.MatchupFieldGoalsMade)
			assert.Equal(t, int32(3), s.MatchupFieldGoalsAttempted)
			assert.Equal(t, float32(0.333), s.MatchupFieldGoalsPercentage)
			assert.Equal(t, int32(0), s.MatchupThreePointersMade)
			assert.Equal(t, int32(0), s.MatchupThreePointersAttempted)
			assert.Equal(t, float32(0), s.MatchupThreePointersPercentage)
			assert.Equal(t, int32(0), s.HelpBlocks)
			assert.Equal(t, int32(0), s.HelpFieldGoalsMade)
			assert.Equal(t, int32(0), s.HelpFieldGoalsAttempted)
			assert.Equal(t, float32(0), s.HelpFieldGoalsPercentage)
			assert.Equal(t, int32(4), s.MatchupFreeThrowsMade)
			assert.Equal(t, int32(5), s.MatchupFreeThrowsAttempted)
			assert.Equal(t, int32(2), s.ShootingFouls)
			mitchellTested = true
		}
	}
	assert.True(t, mitchellTested, "Donovan Mitchell vs RJ Barrett statline tested")
}
