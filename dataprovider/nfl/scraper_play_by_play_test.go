//go:build integration

package nfl

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runPlayByPlay(t *testing.T, matchup model.Matchup) []model.PlayByPlay {
	t.Helper()
	scraper := NewPlayByPlayScraper()
	scraper.Fetcher = testFetcher
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper: scraper,
		},
	)
	plays, err := eventrunner.Run([]model.Matchup{matchup})
	require.NoError(t, err)
	return plays
}

func TestPlayByPlayScraper(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/f5908b6d-311e-11f0-b670-ae1250fadad1?includeDriveChart=true
	// DAL @ PHI
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2025-09-04", "f5908b6d-311e-11f0-b670-ae1250fadad1")
	plays := runPlayByPlay(t, matchup)
	require.Equal(t, 179, len(plays), "179 plays")

	byID := make(map[int64]model.PlayByPlay, len(plays))
	drives := make(map[int32]bool)
	for _, play := range plays {
		byID[play.PlayID] = play
		if play.DriveSequence != nil {
			drives[*play.DriveSequence] = true
		}
		assert.Equal(t, "f5908b6d-311e-11f0-b670-ae1250fadad1", play.EventID)
		assert.Equal(t, "10403700-b939-3cbd-3d16-24d4d6742fa2", play.HomeTeamID)
		assert.Equal(t, "10401200-a308-98ca-ad5f-95df2fefea68", play.AwayTeamID)
	}
	assert.Equal(t, 16, len(drives), "16 drives")

	// game start: not part of a drive
	start, exists := byID[1]
	require.True(t, exists, "play 1 should exist")
	assert.Equal(t, "GAME_START", start.PlayType)
	assert.Nil(t, start.Clock)
	assert.Nil(t, start.DriveSequence)
	assert.Nil(t, start.DriveTeamID)

	// (14:54) J.Williams left tackle to PHI 46 for 7 yards (C.DeJean).
	rush, exists := byID[71]
	require.True(t, exists, "play 71 should exist")
	assert.Equal(t, float64(71), rush.PlaySequenceNumber)
	assert.Equal(t, int32(1), rush.Quarter)
	require.NotNil(t, rush.Clock)
	assert.Equal(t, "14:54", *rush.Clock)
	assert.Equal(t, "RUSH", rush.PlayType)
	assert.Equal(t, int32(1), rush.Down)
	assert.Equal(t, int32(10), rush.Distance)
	require.NotNil(t, rush.YardLine)
	assert.Equal(t, "DAL 47", *rush.YardLine)
	assert.Equal(t, int32(7), rush.YardsGained)
	require.NotNil(t, rush.Description)
	assert.Equal(t, "(14:54) J.Williams left tackle to PHI 46 for 7 yards (C.DeJean).", *rush.Description)
	require.NotNil(t, rush.PlayStartTime)
	assert.Equal(t, time.Date(2025, 9, 5, 0, 27, 9, 830000000, time.UTC), *rush.PlayStartTime)
	require.NotNil(t, rush.PlayEndTime)
	assert.Equal(t, time.Date(2025, 9, 5, 0, 27, 13, 513000000, time.UTC), *rush.PlayEndTime)
	assert.False(t, rush.Scored)
	assert.Nil(t, rush.ScoringTeamID)
	assert.False(t, rush.Deleted)
	require.NotNil(t, rush.DriveSequence)
	assert.Equal(t, int32(1), *rush.DriveSequence)
	require.NotNil(t, rush.DriveTeamID)
	assert.Equal(t, "10401200-a308-98ca-ad5f-95df2fefea68", *rush.DriveTeamID)
	require.NotNil(t, rush.DriveStart)
	assert.Equal(t, "Kickoff", *rush.DriveStart)
	require.NotNil(t, rush.DriveResult)
	assert.Equal(t, "Touchdown", *rush.DriveResult)

	// B.Aubrey extra point is GOOD
	pat, exists := byID[270]
	require.True(t, exists, "play 270 should exist")
	assert.Equal(t, "XP_KICK", pat.PlayType)
	assert.True(t, pat.Scored)
	assert.Equal(t, "PAT", pat.ScoringPlayType)
	require.NotNil(t, pat.ScoringTeamID)
	assert.Equal(t, "10401200-a308-98ca-ad5f-95df2fefea68", *pat.ScoringTeamID)

	// deleted plays have quarter 0
	deleted, exists := byID[94]
	require.True(t, exists, "play 94 should exist")
	assert.True(t, deleted.Deleted)
	assert.Equal(t, int32(0), deleted.Quarter)
	assert.Equal(t, "UNSPECIFIED", deleted.PlayType)
}

func TestPlayByPlayScraperOvertime(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/f6ced5de-311e-11f0-b670-ae1250fadad1?includeDriveChart=true
	// GB @ DAL, 40-40 (OT tie)
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2025-09-28", "f6ced5de-311e-11f0-b670-ae1250fadad1")
	plays := runPlayByPlay(t, matchup)
	require.Equal(t, 231, len(plays), "231 plays")
	var overtime int
	for _, play := range plays {
		if play.Quarter == 5 {
			overtime++
		}
	}
	assert.Equal(t, 32, overtime, "32 overtime plays")
}

func TestPlayByPlayScraperPostSeasonOvertime(t *testing.T) {
	// https://api.nfl.com/experience/v2/gamedetails/1ab6b0a3-f037-11f0-9442-5911216651e2?includeDriveChart=true
	// BUF @ DEN, 30-33 (OT), AFC divisional round
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	matchup := retrieveMatchup(t, "2026-01-17", "1ab6b0a3-f037-11f0-9442-5911216651e2")
	plays := runPlayByPlay(t, matchup)
	require.Equal(t, 220, len(plays), "220 plays")
	var overtime int
	drives := make(map[int32]bool)
	for _, play := range plays {
		if play.Quarter == 5 {
			overtime++
		}
		if play.DriveSequence != nil {
			drives[*play.DriveSequence] = true
		}
	}
	assert.Equal(t, 28, overtime, "28 overtime plays")
	assert.Equal(t, 24, len(drives), "24 drives")
}
