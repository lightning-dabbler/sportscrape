//go:build integration

package nfl

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInjuriesScraper(t *testing.T) {
	// https://api.nfl.com/football/v2/injuries?limit=500&season=2025&seasonType=REG
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	scraper := NewInjuriesScraper(WithInjuriesSeason("2025"))
	scraper.Fetcher = testFetcher
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Injury]{
			Scraper: scraper,
		},
	)
	injuries, err := matchuprunner.Run()
	require.NoError(t, err)
	// 5783 REG + 285 POST as of 2026-10-06
	require.GreaterOrEqual(t, len(injuries), 6000)

	type key struct {
		seasonType string
		week       int32
		personID   string
	}
	seen := make(map[key]bool, len(injuries))
	weeks := map[string]map[int32]bool{}
	for _, injury := range injuries {
		k := key{injury.SeasonType, injury.Week, injury.PersonID}
		assert.False(t, seen[k], "duplicate %v", k)
		seen[k] = true
		if weeks[injury.SeasonType] == nil {
			weeks[injury.SeasonType] = map[int32]bool{}
		}
		weeks[injury.SeasonType][injury.Week] = true
		assert.Equal(t, int32(2025), injury.Season)
		assert.NotEmpty(t, injury.TeamAbbreviation, "team abbreviation of %s", injury.TeamID)
		assert.NotEmpty(t, injury.PlayerID)
		assert.NotEmpty(t, injury.PersonID)
	}
	assert.Len(t, weeks["REG"], 18, "REG weeks")
	assert.Len(t, weeks["POST"], 4, "POST weeks")

	// Tyler Bass, BUF, 2025 REG week 1
	var bass *model.Injury
	for i := range injuries {
		if injuries[i].SeasonType == "REG" && injuries[i].Week == 1 && injuries[i].PlayerID == "00-0036162" {
			bass = &injuries[i]
		}
	}
	require.NotNil(t, bass)
	assert.Equal(t, "Tyler Bass", bass.Player)
	assert.Equal(t, "32004241-5355-5802-e27b-153ecc1a567e", bass.PersonID)
	assert.Equal(t, "10400610-c40e-a673-1743-2ce2a5d5d731", bass.TeamID)
	assert.Equal(t, "Buffalo Bills", bass.Team)
	assert.Equal(t, "BUF", bass.TeamAbbreviation)
	assert.Equal(t, "K", bass.Position)
	require.NotNil(t, bass.InjuryStatus)
	assert.Equal(t, "OUT", *bass.InjuryStatus)
	require.NotNil(t, bass.Injuries)
	assert.Equal(t, "Hip, Groin", *bass.Injuries)
	require.NotNil(t, bass.Practices)
	assert.Equal(t, "Hip, Groin", *bass.Practices)
	require.NotNil(t, bass.PracticeStatus)
	assert.Equal(t, "DIDNOT", *bass.PracticeStatus)
	require.NotNil(t, bass.PracticeDays)
	assert.Equal(t, "2025-09-03:LIMITED,2025-09-04:DIDNOT,2025-09-05:DIDNOT", *bass.PracticeDays)
}
