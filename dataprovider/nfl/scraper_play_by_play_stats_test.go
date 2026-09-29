//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runPlayByPlayStats(t *testing.T, matchup model.Matchup) []model.PlayByPlayStat {
	t.Helper()
	scraper := NewPlayByPlayStatsScraper()
	scraper.Fetcher = testFetcher
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlayStat]{
			Scraper: scraper,
		},
	)
	stats, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	return stats
}

// findStat returns the stat of statType on playID
func findStat(t *testing.T, stats []model.PlayByPlayStat, playID int64, statType int32) model.PlayByPlayStat {
	t.Helper()
	for _, stat := range stats {
		if stat.PlayID == playID && stat.StatType == statType {
			return stat
		}
	}
	t.Fatalf("stat type %d not found on play %d", statType, playID)
	return model.PlayByPlayStat{}
}

func TestPlayByPlayStatsScraper(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/f5908b6d-311e-11f0-b670-ae1250fadad1?includeDriveChart=true
	// DAL @ PHI
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, dalPhiDate, dalPhiEventID)
	stats := runPlayByPlayStats(t, matchup)
	require.Equal(t, 612, len(stats), "612 stats")
	playStatIDs := make(map[string]bool, len(stats))
	for _, stat := range stats {
		assert.Equal(t, dalPhiEventID, stat.EventID)
		assert.NotNil(t, stat.StatTypeDescription, "stat type %d should be decoded", stat.StatType)
		assert.Contains(t, []string{"Dallas Cowboys", "Philadelphia Eagles"}, stat.Team)
		playStatIDs[stat.PlayStatID] = true
	}
	assert.Equal(t, len(stats), len(playStatIDs), "play stat IDs are unique")

	// (14:54) J.Williams left tackle to PHI 46 for 7 yards (C.DeJean).
	rush := findStat(t, stats, 71, 10)
	assert.Equal(t, "10330059-8430-0710-0100-001c3dd1f15d", rush.PlayStatID)
	assert.Equal(t, int32(1), rush.Quarter)
	require.NotNil(t, rush.StatTypeDescription)
	assert.Equal(t, "rushing yards", *rush.StatTypeDescription)
	require.NotNil(t, rush.Yards)
	assert.Equal(t, int32(7), *rush.Yards)
	assert.Equal(t, dalTeamID, rush.TeamID)
	assert.Equal(t, "Dallas Cowboys", rush.Team)
	require.NotNil(t, rush.PlayerID)
	assert.Equal(t, "00-0036997", *rush.PlayerID)
	require.NotNil(t, rush.Player)
	assert.Equal(t, "Javonte Williams", *rush.Player)
	require.NotNil(t, rush.PlayerShortName)
	assert.Equal(t, "J.Williams", *rush.PlayerShortName)
	require.NotNil(t, rush.JerseyNumber)
	assert.Equal(t, "33", *rush.JerseyNumber)

	tackle := findStat(t, stats, 71, 79)
	require.NotNil(t, tackle.StatTypeDescription)
	assert.Equal(t, "solo tackle", *tackle.StatTypeDescription)
	assert.Nil(t, tackle.Yards)
	assert.Equal(t, phiTeamID, tackle.TeamID)
	assert.Equal(t, "Philadelphia Eagles", tackle.Team)
	require.NotNil(t, tackle.Player)
	assert.Equal(t, "Cooper DeJean", *tackle.Player)

	// D.Prescott pass to C.Lamb for 32 yards on third down: team stats have no player
	firstDown := findStat(t, stats, 188, 4)
	require.NotNil(t, firstDown.StatTypeDescription)
	assert.Equal(t, "first down passing", *firstDown.StatTypeDescription)
	assert.Equal(t, dalTeamID, firstDown.TeamID)
	assert.Equal(t, "Dallas Cowboys", firstDown.Team)
	assert.Nil(t, firstDown.PlayerID)
	assert.Nil(t, firstDown.PersonID)
	assert.Nil(t, firstDown.Player)
	assert.Nil(t, firstDown.PlayerShortName)
	thirdDown := findStat(t, stats, 188, 6)
	require.NotNil(t, thirdDown.StatTypeDescription)
	assert.Equal(t, "third down converted", *thirdDown.StatTypeDescription)
	airYards := findStat(t, stats, 188, 111)
	require.NotNil(t, airYards.Yards)
	assert.Equal(t, int32(26), *airYards.Yards)
	require.NotNil(t, airYards.Player)
	assert.Equal(t, "Dak Prescott", *airYards.Player)
}

func TestPlayByPlayStatsScraperPostSeasonOvertime(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/1ab6b0a3-f037-11f0-9442-5911216651e2?includeDriveChart=true
	// BUF @ DEN, 30-33 (OT), AFC divisional round
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2026-01-17", "1ab6b0a3-f037-11f0-9442-5911216651e2")
	stats := runPlayByPlayStats(t, matchup)
	var overtime int
	var allenRushingYards int32
	for _, stat := range stats {
		if stat.Quarter == 5 {
			overtime++
		}
		// rushing yards, rushing touchdown and secondary runner rushing yards add up to the box score's rushing yards
		if stat.PlayerID != nil && *stat.PlayerID == "00-0034857" && stat.Yards != nil && (stat.StatType == 10 || stat.StatType == 11 || stat.StatType == 12) {
			allenRushingYards += *stat.Yards
		}
	}
	assert.Equal(t, 107, overtime, "107 overtime stats")
	assert.Equal(t, int32(66), allenRushingYards, "Josh Allen rushing yards")
}
