//go:build integration

package baseballsavantmlb

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/stretchr/testify/assert"
)

func TestMatchupScraper_NBA(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchupscraper := NewMatchupScraper(
		MatchupScraperDate("2024-10-18"),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	n_matchups := len(matchups)
	assert.Equal(t, 2, n_matchups, "2 events")
	testMatchup := matchups[1].(model.Matchup)
	assert.Equal(t, int64(775311), testMatchup.EventID)
	assert.Equal(t, "Final", testMatchup.Status)

	assert.Equal(t, int64(114), testMatchup.HomeTeamID)
	assert.Equal(t, "CLE", testMatchup.HomeTeamAbbreviation)
	assert.Equal(t, "Cleveland Guardians", testMatchup.HomeTeamName)
	assert.Equal(t, "Gavin Williams", *testMatchup.HomeStartingPitcher)
	assert.Equal(t, int64(668909), *testMatchup.HomeStartingPitcherID)
	assert.Equal(t, int32(1), testMatchup.HomeWins)
	assert.Equal(t, int32(3), testMatchup.HomeLosses)
	assert.Equal(t, int32(6), *testMatchup.HomeScore)

	assert.Equal(t, int64(147), testMatchup.AwayTeamID)
	assert.Equal(t, "NYY", testMatchup.AwayTeamAbbreviation)
	assert.Equal(t, "New York Yankees", testMatchup.AwayTeamName)
	assert.Equal(t, "Luis Gil", *testMatchup.AwayStartingPitcher)
	assert.Equal(t, int64(661563), *testMatchup.AwayStartingPitcherID)
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
