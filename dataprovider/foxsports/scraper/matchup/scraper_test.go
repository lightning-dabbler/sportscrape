//go:build integration

package matchup

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/runner/matchup"
	"github.com/stretchr/testify/assert"
)

func TestGetNBAMatchup(t *testing.T) {
	// https://api.foxsports.com/bifrost/v1/nba/scoreboard/segment/20250406?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	scraper := NewScraper(
		ScraperLeague(foxsports.NBA),
		ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2025-04-06"}),
	)

	runner := matchup.NewRunner(
		matchup.RunnerName("NBA Matchups"),
		matchup.RunnerScraper(scraper),
	)

	matchups := runner.RunMatchupsScraper()
	n_matchups := len(matchups)
	assert.Equal(t, 11, n_matchups, "11 events")
	testMatchup := matchups[4].(model.Matchup)
	assert.Equal(t, int64(43190), testMatchup.EventID)
	assert.Equal(t, time.Date(2025, time.April, 6, 22, 0, 0, 0, time.UTC), testMatchup.EventTime) // 2025-04-06T22:00:00Z
	assert.Equal(t, int32(3), testMatchup.EventStatus)
	assert.Equal(t, "FINAL", testMatchup.StatusLine)
	assert.Equal(t, int64(8), testMatchup.HomeTeamID)
	assert.Equal(t, "ATL", testMatchup.HomeTeamAbbreviation)
	assert.Equal(t, "Hawks", testMatchup.HomeTeamNameLong)
	assert.Equal(t, "Atlanta Hawks", testMatchup.HomeTeamNameFull)
	assert.Equal(t, "37-41", testMatchup.HomeRecord)
	assert.Equal(t, int32(147), testMatchup.HomeScore)
	assert.Nil(t, testMatchup.HomeRank)
	assert.Equal(t, int64(20), testMatchup.AwayTeamID)
	assert.Equal(t, "UTA", testMatchup.AwayTeamAbbreviation)
	assert.Equal(t, "Jazz", testMatchup.AwayTeamNameLong)
	assert.Equal(t, "Utah Jazz", testMatchup.AwayTeamNameFull)
	assert.Equal(t, "16-63", testMatchup.AwayRecord)
	assert.Equal(t, int32(134), testMatchup.AwayScore)
	assert.Nil(t, testMatchup.AwayRank)
	assert.Equal(t, int64(20), *testMatchup.Loser) // Utah lost
	assert.Equal(t, false, testMatchup.IsPlayoff)

}

func TestGetMLBMatchup(t *testing.T) {
	// https://api.foxsports.com/bifrost/v1/mlb/scoreboard/segment/20241018?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	scraper := NewScraper(
		ScraperLeague(foxsports.MLB),
		ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-18"}),
	)

	runner := matchup.NewRunner(
		matchup.RunnerName("MLB Matchups"),
		matchup.RunnerScraper(scraper),
	)

	matchups := runner.RunMatchupsScraper()

	n_matchups := len(matchups)
	assert.Equal(t, 2, n_matchups, "2 events")
	testMatchup := matchups[0].(model.Matchup)
	assert.Equal(t, int64(91212), testMatchup.EventID)
	assert.Equal(t, time.Date(2024, time.October, 18, 21, 8, 0, 0, time.UTC), testMatchup.EventTime) // 2024-10-18T21:08:00Z
	assert.Equal(t, int32(3), testMatchup.EventStatus)
	assert.Equal(t, "FINAL", testMatchup.StatusLine)
	assert.Equal(t, int64(17), testMatchup.HomeTeamID)
	assert.Equal(t, "NYM", testMatchup.HomeTeamAbbreviation)
	assert.Equal(t, "Mets", testMatchup.HomeTeamNameLong)
	assert.Equal(t, "New York Mets", testMatchup.HomeTeamNameFull)
	assert.Equal(t, "89-73", testMatchup.HomeRecord)
	assert.Equal(t, int32(12), testMatchup.HomeScore)
	assert.Nil(t, testMatchup.HomeRank)
	assert.Equal(t, int64(24), testMatchup.AwayTeamID)
	assert.Equal(t, "LAD", testMatchup.AwayTeamAbbreviation)
	assert.Equal(t, "Dodgers", testMatchup.AwayTeamNameLong)
	assert.Equal(t, "Los Angeles Dodgers", testMatchup.AwayTeamNameFull)
	assert.Equal(t, "98-64", testMatchup.AwayRecord)
	assert.Equal(t, int32(6), testMatchup.AwayScore)
	assert.Nil(t, testMatchup.AwayRank)
	assert.Equal(t, int64(24), *testMatchup.Loser) // Dodgers lost
	assert.Equal(t, true, testMatchup.IsPlayoff)

}

func TestGetNFLMatchup(t *testing.T) {
	// https://api.foxsports.com/bifrost/v1/nfl/scoreboard/segment/2024-4-2?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	scraper := NewScraper(
		ScraperLeague(foxsports.NFL),
		ScraperSegmenter(&foxsports.NFLSegmenter{Year: 2024, Week: 3, Season: foxsports.POSTSEASON}),
	)

	runner := matchup.NewRunner(
		matchup.RunnerName("NFL Matchups"),
		matchup.RunnerScraper(scraper),
	)

	matchups := runner.RunMatchupsScraper()
	n_matchups := len(matchups)
	assert.Equal(t, 2, n_matchups, "2 events")
	testMatchup := matchups[1].(model.Matchup)
	assert.Equal(t, int64(10701), testMatchup.EventID)
	assert.Equal(t, time.Date(2025, time.January, 26, 23, 30, 0, 0, time.UTC), testMatchup.EventTime) // 2025-01-26T23:30:00Z
	assert.Equal(t, int32(3), testMatchup.EventStatus)
	assert.Equal(t, "FINAL", testMatchup.StatusLine)
	assert.Equal(t, int64(11), testMatchup.HomeTeamID)
	assert.Equal(t, "KC", testMatchup.HomeTeamAbbreviation)
	assert.Equal(t, "Chiefs", testMatchup.HomeTeamNameLong)
	assert.Equal(t, "Kansas City Chiefs", testMatchup.HomeTeamNameFull)
	assert.Equal(t, "15-2", testMatchup.HomeRecord)
	assert.Equal(t, int32(32), testMatchup.HomeScore)
	assert.Nil(t, testMatchup.HomeRank)
	assert.Equal(t, int64(1), testMatchup.AwayTeamID)
	assert.Equal(t, "BUF", testMatchup.AwayTeamAbbreviation)
	assert.Equal(t, "Bills", testMatchup.AwayTeamNameLong)
	assert.Equal(t, "Buffalo Bills", testMatchup.AwayTeamNameFull)
	assert.Equal(t, "13-4", testMatchup.AwayRecord)
	assert.Equal(t, int32(29), testMatchup.AwayScore)
	assert.Nil(t, testMatchup.AwayRank)
	assert.Equal(t, int64(1), *testMatchup.Loser) // Bills lost
	assert.Equal(t, true, testMatchup.IsPlayoff)
}

func TestGetNCAABMatchup(t *testing.T) {
	// https://api.foxsports.com/bifrost/v1/cbk/scoreboard/segment/20250405?groupId=2&apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	scraper := NewScraper(
		ScraperLeague(foxsports.NCAAB),
		ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2025-04-05"}),
	)

	runner := matchup.NewRunner(
		matchup.RunnerName("NCAAB Matchup"),
		matchup.RunnerScraper(scraper),
	)

	matchups := runner.RunMatchupsScraper()

	n_matchups := len(matchups)
	assert.Equal(t, 4, n_matchups, "4 events")
	testMatchup := matchups[2].(model.Matchup)
	assert.Equal(t, int64(258066), testMatchup.EventID)
	assert.Equal(t, time.Date(2025, time.April, 5, 22, 9, 0, 0, time.UTC), testMatchup.EventTime) // 2025-04-05T22:09:00Z
	assert.Equal(t, int32(3), testMatchup.EventStatus)
	assert.Equal(t, "FINAL", testMatchup.StatusLine)
	assert.Equal(t, int64(237), testMatchup.HomeTeamID)
	assert.Equal(t, "FLA", testMatchup.HomeTeamAbbreviation)
	assert.Equal(t, "Florida", testMatchup.HomeTeamNameLong)
	assert.Equal(t, "Florida Gators", testMatchup.HomeTeamNameFull)
	assert.Equal(t, "35-4", testMatchup.HomeRecord)
	assert.Equal(t, int32(79), testMatchup.HomeScore)
	assert.Equal(t, int32(1), *testMatchup.HomeRank)
	assert.Equal(t, int64(245), testMatchup.AwayTeamID)
	assert.Equal(t, "AUB", testMatchup.AwayTeamAbbreviation)
	assert.Equal(t, "Auburn", testMatchup.AwayTeamNameLong)
	assert.Equal(t, "Auburn Tigers", testMatchup.AwayTeamNameFull)
	assert.Equal(t, "32-6", testMatchup.AwayRecord)
	assert.Equal(t, int32(73), testMatchup.AwayScore)
	assert.Equal(t, int32(1), *testMatchup.AwayRank)
	assert.Equal(t, int64(245), *testMatchup.Loser) // Auburn lost
	assert.Equal(t, false, testMatchup.IsPlayoff)
}
