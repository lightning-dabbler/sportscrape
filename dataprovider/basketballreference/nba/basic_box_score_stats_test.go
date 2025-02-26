//go:build integration

package nba

import (
	"testing"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/stretchr/testify/assert"
)

func TestGetBasicBoxScoreStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	date := "2025-02-19"
	matchups := GetMatchups(date)
	basicBoxScoreStats := GetBasicBoxScoreStats(1, matchups...)
	numAwayPlayers := 0
	numHomePlayers := 0
	numAwayDNP := 0
	numHomeDNP := 0
	expectedNumAwayPlayers := 12
	expectedNumHomePlayers := 15
	expectedAwayDNP := 2
	expectedHomeDNP := 5
	for _, s := range basicBoxScoreStats {
		stats := s.(model.NBABasicBoxScoreStats)
		if stats.Team == "LA Lakers" {
			// Actual home players
			numHomePlayers += 1
			if stats.MinutesPlayed == 0 {
				// Actual home players that did not play
				numHomeDNP += 1
			}
		} else {
			// Actual away players
			numAwayPlayers += 1
			if stats.MinutesPlayed == 0 {
				// Actual away players that did not play
				numAwayDNP += 1
			}
			// Sample player stats from away team (LaMelo Ball)
			if stats.Player == "LaMelo Ball" {
				assert.Equal(t, true, stats.Starter, "LaMelo was a starter")
				assert.Equal(t, "ballla01", stats.PlayerID, "LaMelo's player ID")
				assert.Equal(t, "LA Lakers", stats.Opponent, "LaMelo's opposing team")
				assert.Equal(t, "https://www.basketball-reference.com/players/b/ballla01.html", stats.PlayerLink, "LaMelo's player link")
				assert.Equal(t, float32(33.25), stats.MinutesPlayed, "LaMelo minutes played")
				assert.Equal(t, 9, stats.FieldGoalsMade, "LaMelo field goals made")
				assert.Equal(t, 19, stats.FieldGoalAttempts, "LaMelo field goal attempts")
				assert.Equal(t, float32(0.474), stats.FieldGoalPercentage, "LaMelo field goal percentage")
				assert.Equal(t, 5, stats.ThreePointsMade, "LaMelo three points made")
				assert.Equal(t, 13, stats.ThreePointAttempts, "LaMelo three point attempts")
				assert.Equal(t, float32(0.385), stats.ThreePointPercentage, "LaMelo three point percentage")
				assert.Equal(t, 4, stats.FreeThrowsMade, "LaMelo free thows made")
				assert.Equal(t, 4, stats.FreeThrowAttempts, "LaMelo free throw attempts")
				assert.Equal(t, float32(1), stats.FreeThrowPercentage, "LaMelo free throw percentage")
				assert.Equal(t, 0, stats.OffensiveRebounds, "LaMelo offensive rebounds")
				assert.Equal(t, 5, stats.DefensiveRebounds, "LaMelo defensive rebounds")
				assert.Equal(t, 5, stats.TotalRebounds, "LaMelo total rebounds")
				assert.Equal(t, 6, stats.Assists, "LaMelo assists")
				assert.Equal(t, 1, stats.Steals, "LaMelo steals")
				assert.Equal(t, 0, stats.Blocks, "LaMelo blocks")
				assert.Equal(t, 3, stats.Turnovers, "LaMelo turnovers")
				assert.Equal(t, 1, stats.PersonalFouls, "LaMelo personal fouls")
				assert.Equal(t, 27, stats.Points, "LaMelo points")
				assert.Equal(t, float32(20.6), stats.GameScore, "LaMelo game score")
				assert.Equal(t, 8, stats.PlusMinus, "LaMelo plus minus")
			}
		}

	}
	assert.Equal(t, expectedNumAwayPlayers, numAwayPlayers, "Assert number of away players on Charlotte")
	assert.Equal(t, expectedNumHomePlayers, numHomePlayers, "Assert number of home players on LA Lakers")
	assert.Equal(t, expectedAwayDNP, numAwayDNP, "Assert number of away players on Charlotte that did not play")
	assert.Equal(t, expectedHomeDNP, numHomeDNP, "Assert number of home players on LA Lakers that did not play")
}
