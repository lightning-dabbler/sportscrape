//go:build integration

package eventdata

import (
	"fmt"
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
	"github.com/stretchr/testify/assert"
)

func TestNBABoxScoreScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get matchups
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.NBA),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2025-04-07"}),
	)

	matchups := matchupRunner.GetMatchups()

	// Get boxscore data
	scraper := NBABoxScoreScraper{}
	scraper.League = foxsports.NBA
	runner := NewRunner(
		RunnerName("NBA Box score stats"),
		RunnerConcurrency(2),
		RunnerScraper(
			&scraper,
		),
	)
	boxScoreStats := runner.RunEventsDataScraper(matchups...)
	n_stats := len(boxScoreStats)
	assert.Equal(t, 41, n_stats, "41 statlines")
	// 43197 (Sacramento Kings vs Detroit Pistons) scraped for 20
	// 43198 (Philadelphia 76ers vs Miami Heat) scraped for 21
	expectedStatlineCountGroupedByEventID := map[int64]int{
		43197: 20,
		43198: 21,
	}
	actualStatlineCountGroupedByEventID := make(map[int64]int)
	KeonEllisTested := false

	for _, statline := range boxScoreStats {
		s := statline.(model.NBABoxScoreStats)
		actualStatlineCountGroupedByEventID[s.EventID] += 1
		if s.Player == "Keon Ellis" {
			KeonEllisTested = true
			assert.Equal(t, int64(43197), s.EventID, "EventID")
			assert.Equal(t, int64(26), s.TeamID, "TeamID")
			assert.Equal(t, "Sacramento Kings", s.Team, "Team")
			assert.Equal(t, int64(12), s.OpponentID, "OpponentID")
			assert.Equal(t, "Detroit Pistons", s.Opponent, "Opponent")
			assert.Equal(t, int64(3708), s.PlayerID, "PlayerID")
			assert.Equal(t, "Keon Ellis", s.Player, "Player")
			assert.Equal(t, "SG", *s.Position, "Position")
			assert.Equal(t, true, s.Starter, "Starter")
			assert.Equal(t, int32(37), s.MinutesPlayed, "MinutesPlayed")
			assert.Equal(t, int32(1), s.FieldGoalsMade, "FieldGoalsMade")
			assert.Equal(t, int32(2), s.FieldGoalAttempts, "FieldGoalAttempts")
			assert.Equal(t, int32(0), s.ThreePointsMade, "ThreePointsMade")
			assert.Equal(t, int32(1), s.ThreePointAttempts, "ThreePointAttempts")
			assert.Equal(t, int32(0), s.FreeThrowsMade, "FreeThrowsMade")
			assert.Equal(t, int32(0), s.FreeThrowAttempts, "FreeThrowAttempts")
			assert.Equal(t, int32(2), s.OffensiveRebounds, "OffensiveRebounds")
			assert.Equal(t, int32(4), s.DefensiveRebounds, "DefensiveRebounds")
			assert.Equal(t, int32(6), s.TotalRebounds, "TotalRebounds")
			assert.Equal(t, int32(0), s.Assists, "Assists")
			assert.Equal(t, int32(3), s.Steals, "Steals")
			assert.Equal(t, int32(0), s.Blocks, "Blocks")
			assert.Equal(t, int32(0), s.Turnovers, "Turnovers")
			assert.Equal(t, int32(3), s.PersonalFouls, "PersonalFouls")
			assert.Equal(t, int32(2), s.Points, "Points")
		}
	}
	assert.True(t, KeonEllisTested, "Keon Ellis statline tested")
	for eventID, count := range actualStatlineCountGroupedByEventID {
		val, exists := expectedStatlineCountGroupedByEventID[eventID]
		assert.True(t, exists, fmt.Sprintf("Event ID %d is in expected list", eventID))
		assert.Equal(t, val, count)
	}
	// 2019-10-06
	// Issue: https://github.com/lightning-dabbler/sportscrape/issues/64

	matchupRunner = matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.NBA),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2019-10-06"}),
	)

	matchups = matchupRunner.GetMatchups()

	// Get boxscore data
	scraper = NBABoxScoreScraper{}
	scraper.League = foxsports.NBA
	runner = NewRunner(
		RunnerName("NBA Box score stats"),
		RunnerConcurrency(2),
		RunnerScraper(
			&scraper,
		),
	)
	boxScoreStats = runner.RunEventsDataScraper(matchups...)
	n_stats = len(boxScoreStats)
	assert.Equal(t, 69, n_stats, "69 statlines")
}
