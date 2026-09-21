//go:build integration

package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMLBBattingLineupScraper(t *testing.T) {
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

	lineupscraper := NewBattingLineupScraper()

	lineuprunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BattingLineup]{
			Scraper:     lineupscraper,
			Concurrency: 1,
		},
	)
	lineup, err := lineuprunner.Run(matchups)
	assert.NoError(t, err)
	assert.Equal(t, 32, len(lineup), "32 lineup entries")
	GavinLuxTested := false
	ShoheiOhtaniTested := false
	GleyberTorresTested := false
	for _, s := range lineup {
		switch s.Player {
		case "Gavin Lux":
			GavinLuxTested = true
			assert.Equal(t, int64(775296), s.EventID)
			assert.Equal(t, int64(119), s.TeamID)
			assert.Equal(t, "Los Angeles Dodgers", s.Team)
			assert.Equal(t, "New York Yankees", s.Opponent)
			assert.Equal(t, int64(147), s.OpponentID)
			assert.Equal(t, int64(666158), s.PlayerID)
			assert.Equal(t, "Second Base", s.Position)
			// Chris Taylor's substitution is listed ahead of Gavin Lux
			assert.Equal(t, int32(10), s.LineupOrder)
		case "Shohei Ohtani":
			ShoheiOhtaniTested = true
			assert.Equal(t, int64(775296), s.EventID)
			assert.Equal(t, int64(119), s.TeamID)
			assert.Equal(t, int64(660271), s.PlayerID)
			assert.Equal(t, "Designated Hitter", s.Position)
			assert.Equal(t, int32(1), s.LineupOrder)
		case "Gleyber Torres":
			GleyberTorresTested = true
			assert.Equal(t, int64(775296), s.EventID)
			assert.Equal(t, int64(147), s.TeamID)
			assert.Equal(t, "New York Yankees", s.Team)
			assert.Equal(t, "Los Angeles Dodgers", s.Opponent)
			assert.Equal(t, int64(119), s.OpponentID)
			assert.Equal(t, int64(650402), s.PlayerID)
			assert.Equal(t, "Second Base", s.Position)
			assert.Equal(t, int32(1), s.LineupOrder)
		}
	}
	assert.True(t, GavinLuxTested, "Gavin Lux lineup entry tested")
	assert.True(t, ShoheiOhtaniTested, "Shohei Ohtani lineup entry tested")
	assert.True(t, GleyberTorresTested, "Gleyber Torres lineup entry tested")
}
