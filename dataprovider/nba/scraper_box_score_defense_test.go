package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestBoxScoreDefenseScraper(t *testing.T) {
	// https://www.nba.com/game/tor-vs-cle-0022500225/box-score?type=defense
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2025-11-13"),
		WithMatchupTimeout(1*time.Minute),
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
	boxscorescraper := NewBoxScoreDefenseScraper(
		WithBoxScoreDefenseTimeout(1 * time.Minute),
	)
	boxscorescraper.DocumentRetriever = matchupScraper.DocumentRetriever
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreDefense]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(tested)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 23, n_records, "23 stat lines")
	mitchellTested := false
	for _, s := range records {
		if s.PlayerName == "Donovan Mitchell" {
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
			assert.Equal(t, "G", s.Position)
			assert.Equal(t, true, s.Starter)
			assert.Equal(t, float32(13.67), s.MatchupMinutes)
			assert.Equal(t, float32(70.2), s.PartialPossessions)
			assert.Equal(t, int32(0), s.SwitchesOn)
			assert.Equal(t, int32(17), s.PlayerPoints)
			assert.Equal(t, int32(5), s.DefensiveRebounds)
			assert.Equal(t, int32(2), s.MatchupAssists)
			assert.Equal(t, int32(0), s.MatchupTurnovers)
			assert.Equal(t, int32(1), s.Steals)
			assert.Equal(t, int32(2), s.Blocks)
			assert.Equal(t, int32(7), s.MatchupFieldGoalsMade)
			assert.Equal(t, int32(9), s.MatchupFieldGoalsAttempted)
			assert.Equal(t, float32(0.778), s.MatchupFieldGoalPercentage)
			assert.Equal(t, int32(0), s.MatchupThreePointersMade)
			assert.Equal(t, int32(1), s.MatchupThreePointersAttempted)
			assert.Equal(t, float32(0), s.MatchupThreePointerPercentage)
			mitchellTested = true
		}
	}
	assert.True(t, mitchellTested, "Donovan Mitchell statline tested")
}
