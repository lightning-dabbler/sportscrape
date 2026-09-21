//go:build integration

package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMLBPitchingLineupScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchupscraper := NewMatchupScraper(
		MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)

	lineupscraper := NewPitchingLineupScraper()
	lineuprunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PitchingLineup]{
			Scraper:     lineupscraper,
			Concurrency: 1,
		},
	)

	lineup, err := lineuprunner.Run(matchups)
	assert.NoError(t, err)
	assert.Equal(t, 13, len(lineup), "13 lineup entries")
	GerritColeTested := false
	JackFlahertyTested := false
	WalkerBuehlerTested := false
	for _, s := range lineup {
		switch s.Player {
		case "Gerrit Cole":
			GerritColeTested = true
			assert.Equal(t, int64(775296), s.EventID)
			assert.Equal(t, int64(147), s.TeamID)
			assert.Equal(t, "New York Yankees", s.Team)
			assert.Equal(t, "Los Angeles Dodgers", s.Opponent)
			assert.Equal(t, int64(119), s.OpponentID)
			assert.Equal(t, int64(543037), s.PlayerID)
			assert.Equal(t, "Pitcher", s.Position)
			assert.Equal(t, int32(1), s.LineupOrder)
		case "Jack Flaherty":
			JackFlahertyTested = true
			assert.Equal(t, int64(775296), s.EventID)
			assert.Equal(t, int64(119), s.TeamID)
			assert.Equal(t, "Los Angeles Dodgers", s.Team)
			assert.Equal(t, "New York Yankees", s.Opponent)
			assert.Equal(t, int64(147), s.OpponentID)
			assert.Equal(t, int64(656427), s.PlayerID)
			assert.Equal(t, "Pitcher", s.Position)
			assert.Equal(t, int32(1), s.LineupOrder)
		case "Walker Buehler":
			WalkerBuehlerTested = true
			assert.Equal(t, int64(621111), s.PlayerID)
			assert.Equal(t, "Pitcher", s.Position)
			assert.Equal(t, int32(8), s.LineupOrder)
		}
	}
	assert.True(t, GerritColeTested, "Gerrit Cole lineup entry tested")
	assert.True(t, JackFlahertyTested, "Jack Flaherty lineup entry tested")
	assert.True(t, WalkerBuehlerTested, "Walker Buehler lineup entry tested")
}
