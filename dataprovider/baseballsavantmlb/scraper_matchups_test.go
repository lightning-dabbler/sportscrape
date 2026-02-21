//go:build integration

package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestMatchupScraper_NBA(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchupscraper := NewMatchupScraper(
		MatchupScraperDate("2024-10-18"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	n_matchups := len(matchups)
	assert.Equal(t, 2, n_matchups, "2 events")
	testMatchup := matchups[1]
	assert.Equal(t, int64(775311), testMatchup.EventID)
	assert.Equal(t, "Final", testMatchup.Status)
	assert.Equal(t, int64(5), testMatchup.VenueID)
	assert.Equal(t, "Progressive Field", testMatchup.VenueName)

	assert.Equal(t, int64(103), testMatchup.HomeTeamLeagueID)
	assert.Equal(t, "American League", testMatchup.HomeTeamLeagueName)
	assert.Equal(t, int64(202), testMatchup.HomeTeamDivisionID)
	assert.Equal(t, "American League Central", testMatchup.HomeTeamDivisionName)
	assert.Equal(t, int64(114), testMatchup.HomeTeamID)
	assert.Equal(t, "CLE", testMatchup.HomeTeamAbbreviation)
	assert.Equal(t, "Cleveland Guardians", testMatchup.HomeTeamName)
	assert.Equal(t, "Gavin Williams", *testMatchup.HomeStartingPitcher)
	assert.Equal(t, int64(668909), *testMatchup.HomeStartingPitcherID)
	assert.Equal(t, "R", *testMatchup.HomeStartingPitcherPitchHand)
	assert.Equal(t, int32(1), testMatchup.HomeWins)
	assert.Equal(t, int32(3), testMatchup.HomeLosses)
	assert.Equal(t, int32(6), *testMatchup.HomeScore)

	assert.Equal(t, int64(103), testMatchup.AwayTeamLeagueID)
	assert.Equal(t, "American League", testMatchup.AwayTeamLeagueName)
	assert.Equal(t, int64(201), testMatchup.AwayTeamDivisionID)
	assert.Equal(t, "American League East", testMatchup.AwayTeamDivisionName)
	assert.Equal(t, int64(147), testMatchup.AwayTeamID)
	assert.Equal(t, "NYY", testMatchup.AwayTeamAbbreviation)
	assert.Equal(t, "New York Yankees", testMatchup.AwayTeamName)
	assert.Equal(t, "Luis Gil", *testMatchup.AwayStartingPitcher)
	assert.Equal(t, int64(661563), *testMatchup.AwayStartingPitcherID)
	assert.Equal(t, "R", *testMatchup.AwayStartingPitcherPitchHand)
	assert.Equal(t, int32(3), testMatchup.AwayWins)
	assert.Equal(t, int32(1), testMatchup.AwayLosses)
	assert.Equal(t, int32(8), *testMatchup.AwayScore)

	assert.Equal(t, int64(147), *testMatchup.Loser)
	assert.Equal(t, "L", testMatchup.GameType)
	assert.Equal(t, "League Championship Series", testMatchup.SeriesDescription)
	assert.Equal(t, int32(7), testMatchup.GamesInSeries)
	assert.Equal(t, int32(4), testMatchup.SeriesGameNumber)
	assert.Equal(t, int32(2024), testMatchup.Season)

}
