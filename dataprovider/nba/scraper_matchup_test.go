//go:build integration

package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMatchupScraper(t *testing.T) {
	// https://www.nba.com/games?date=2025-06-05
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(3*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	n_matchups := len(matchups)
	matchup := matchups[0]
	assert.Equal(t, 1, n_matchups, "1 event")
	assert.Equal(t, "0042400401", matchup.EventID)
	assert.Equal(t, int32(3), matchup.EventStatus)
	assert.Equal(t, "Final", matchup.EventStatusText)
	assert.Equal(t, int64(1610612760), matchup.HomeTeamID)
	assert.Equal(t, "Thunder", matchup.HomeTeam)
	assert.Equal(t, "OKC", matchup.HomeTeamAbbreviation)
	assert.Equal(t, int64(1610612754), matchup.AwayTeamID)
	assert.Equal(t, "Pacers", matchup.AwayTeam)
	assert.Equal(t, "IND", matchup.AwayTeamAbbreviation)
	assert.Equal(t, int32(111), matchup.AwayTeamScore)
	assert.Equal(t, int32(110), matchup.HomeTeamScore)
	assert.Equal(t, int32(1), matchup.AwayTeamWins)
	assert.Equal(t, int32(0), matchup.HomeTeamWins)
	assert.Equal(t, int32(0), matchup.AwayTeamLosses)
	assert.Equal(t, int32(1), matchup.HomeTeamLosses)
	assert.Equal(t, "https://www.nba.com/game/ind-vs-okc-0042400401", matchup.ShareURL)
	assert.Equal(t, "Playoffs", matchup.SeasonType)
	assert.Equal(t, "2024-25", matchup.SeasonYear)
	assert.Equal(t, "00", matchup.LeagueID)
}
