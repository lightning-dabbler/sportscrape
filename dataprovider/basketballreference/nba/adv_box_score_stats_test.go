//go:build integration

package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/stretchr/testify/assert"
)

func TestGetAdvBoxScoreStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	date := "2025-02-19"
	matchupRunner := NewMatchupRunner(
		WithMatchupTimeout(2 * time.Minute),
	)
	matchups := matchupRunner.GetMatchups(date)
	boxScoreRunner := NewAdvBoxScoreRunner(
		WithAdvBoxScoreTimeout(4*time.Minute),
		WithAdvBoxScoreConcurrency(1),
	)
	advBoxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)
	numAwayPlayers := 0
	numHomePlayers := 0
	numAwayDNP := 0
	numHomeDNP := 0
	expectedNumAwayPlayers := 12
	expectedNumHomePlayers := 15
	expectedAwayDNP := 2
	expectedHomeDNP := 5
	for _, s := range advBoxScoreStats {
		stats := s.(model.NBAAdvBoxScoreStats)
		if stats.Team == "LA Lakers" {
			// Actual home players
			numHomePlayers += 1
			if stats.MinutesPlayed == 0 {
				// Actual home players that did not play
				numHomeDNP += 1
			}
			// Sample player stats from home team (LeBron James)
			if stats.Player == "LeBron James" {
				assert.Equal(t, true, stats.Starter, "LeBron was a starter")
				assert.Equal(t, "jamesle01", stats.PlayerID, "LeBron's player ID")
				assert.Equal(t, "Charlotte", stats.Opponent, "LeBron's opposing team")
				assert.Equal(t, "https://www.basketball-reference.com/players/j/jamesle01.html", stats.PlayerLink, "LeBron's player link")
				assert.Equal(t, float32(37.9), stats.MinutesPlayed, "LeBron minutes played")
				assert.Equal(t, float32(0.568), stats.TrueShootingPercentage, "LeBron true shooting percentage")
				assert.Equal(t, float32(0.545), stats.EffectiveFieldGoalPercentage, "LeBron effective field goal percentage")
				assert.Equal(t, float32(0.5), stats.ThreePointAttemptRate, "LeBron three point attempt rate")
				assert.Equal(t, float32(0.091), stats.FreeThrowAttemptRate, "LeBron free throw attempt rate")
				assert.Equal(t, float32(2.5), stats.OffensiveReboundPercentage, "LeBron offensive rebound percentage")
				assert.Equal(t, float32(16.2), stats.DefensiveReboundPercentage, "LeBron defensive rebound percentage")
				assert.Equal(t, float32(9.1), stats.TotalReboundPercentage, "LeBron total rebound percentage")
				assert.Equal(t, float32(57.2), stats.AssistPercentage, "LeBron assist percentage")
				assert.Equal(t, float32(1.2), stats.StealPercentage, "LeBron steal percentage")
				assert.Equal(t, float32(6.7), stats.BlockPercentage, "LeBron block percentage")
				assert.Equal(t, float32(8.0), stats.TurnoverPercentage, "LeBron turnover percentage")
				assert.Equal(t, float32(27.8), stats.UsagePercentage, "LeBron usage percentage")
				assert.Equal(t, 122, stats.OffensiveRating, "LeBron offensive rating")
				assert.Equal(t, 97, stats.DefensiveRating, "LeBron defensive rating")
				assert.Equal(t, float32(14.3), stats.BoxPlusMinus, "LeBron box plus minus")
			}
		} else {
			// Actual away players
			numAwayPlayers += 1
			if stats.MinutesPlayed == 0 {
				// Actual away players that did not play
				numAwayDNP += 1
			}
			// Sample player stats from away team (Gabe Vincent)
			if stats.Player == "Gabe Vincent" {
				assert.Equal(t, false, stats.Starter, "Gabe was a reserve player")
				assert.Equal(t, "vincega01", stats.PlayerID, "Gabe's player ID")
				assert.Equal(t, "LA Lakers", stats.Opponent, "Gabe's opposing team")
				assert.Equal(t, "https://www.basketball-reference.com/players/v/vincega01.html", stats.PlayerLink, "Gabe's player link")
				assert.Equal(t, float32(20.2), stats.MinutesPlayed, "Gabe minutes played")
				assert.Equal(t, float32(0.429), stats.TrueShootingPercentage, "Gabe true shooting percentage")
				assert.Equal(t, float32(0.429), stats.EffectiveFieldGoalPercentage, "Gabe effective field goal percentage")
				assert.Equal(t, float32(0.857), stats.ThreePointAttemptRate, "Gabe three point attempt rate")
				assert.Equal(t, 0, stats.FreeThrowAttemptRate, "Gabe free throw attempt rate")
				assert.Equal(t, float32(9.5), stats.OffensiveReboundPercentage, "Gabe offensive rebound percentage")
				assert.Equal(t, float32(5.1), stats.DefensiveReboundPercentage, "Gabe defensive rebound percentage")
				assert.Equal(t, float32(7.3), stats.TotalReboundPercentage, "Gabe total rebound percentage")
				assert.Equal(t, float32(14.7), stats.AssistPercentage, "Gabe assist percentage")
				assert.Equal(t, 0, stats.StealPercentage, "Gabe steal percentage")
				assert.Equal(t, 0, stats.BlockPercentage, "Gabe block percentage")
				assert.Equal(t, float32(12.5), stats.TurnoverPercentage, "Gabe turnover percentage")
				assert.Equal(t, float32(16.8), stats.UsagePercentage, "Gabe usage percentage")
				assert.Equal(t, 95, stats.OffensiveRating, "Gabe offensive rating")
				assert.Equal(t, 107, stats.DefensiveRating, "Gabe defensive rating")
				assert.Equal(t, float32(-6.0), stats.BoxPlusMinus, "Gabe box plus minus")
			}
		}

	}
	assert.Equal(t, expectedNumAwayPlayers, numAwayPlayers, "Assert number of away players on Charlotte")
	assert.Equal(t, expectedNumHomePlayers, numHomePlayers, "Assert number of home players on LA Lakers")
	assert.Equal(t, expectedAwayDNP, numAwayDNP, "Assert number of away players on Charlotte that did not play")
	assert.Equal(t, expectedHomeDNP, numHomeDNP, "Assert number of home players on LA Lakers that did not play")
}
