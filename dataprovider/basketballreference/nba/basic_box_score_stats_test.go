//go:build integration

package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
	"github.com/stretchr/testify/assert"
)

func TestGetFullBasicBoxScoreStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	date := "2025-02-19"
	matchupRunner := NewMatchupRunner(
		WithMatchupTimeout(5 * time.Minute),
	)
	matchups := matchupRunner.GetMatchups(date)
	boxScoreRunner := NewBasicBoxScoreRunner(
		WithBasicBoxScoreTimeout(5*time.Minute),
		WithBasicBoxScoreConcurrency(1),
	)
	basicBoxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)
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
				assert.Equal(t, true, stats.Starter, "LaMelo is a starter")
				assert.Equal(t, "ballla01", stats.PlayerID, "LaMelo's player ID")
				assert.Equal(t, "LA Lakers", stats.Opponent, "LaMelo's opposing team")
				assert.Equal(t, "https://www.basketball-reference.com/players/b/ballla01.html", stats.PlayerLink, "LaMelo's player link")
				assert.Equal(t, float32(33.25), stats.MinutesPlayed, "LaMelo minutes played")
				assert.Equal(t, int32(9), stats.FieldGoalsMade, "LaMelo field goals made")
				assert.Equal(t, int32(19), stats.FieldGoalAttempts, "LaMelo field goal attempts")
				assert.Equal(t, float32(0.474), stats.FieldGoalPercentage, "LaMelo field goal percentage")
				assert.Equal(t, int32(5), stats.ThreePointsMade, "LaMelo three points made")
				assert.Equal(t, int32(13), stats.ThreePointAttempts, "LaMelo three point attempts")
				assert.Equal(t, float32(0.385), stats.ThreePointPercentage, "LaMelo three point percentage")
				assert.Equal(t, int32(4), stats.FreeThrowsMade, "LaMelo free thows made")
				assert.Equal(t, int32(4), stats.FreeThrowAttempts, "LaMelo free throw attempts")
				assert.Equal(t, float32(1), stats.FreeThrowPercentage, "LaMelo free throw percentage")
				assert.Equal(t, int32(0), stats.OffensiveRebounds, "LaMelo offensive rebounds")
				assert.Equal(t, int32(5), stats.DefensiveRebounds, "LaMelo defensive rebounds")
				assert.Equal(t, int32(5), stats.TotalRebounds, "LaMelo total rebounds")
				assert.Equal(t, int32(6), stats.Assists, "LaMelo assists")
				assert.Equal(t, int32(1), stats.Steals, "LaMelo steals")
				assert.Equal(t, int32(0), stats.Blocks, "LaMelo blocks")
				assert.Equal(t, int32(3), stats.Turnovers, "LaMelo turnovers")
				assert.Equal(t, int32(1), stats.PersonalFouls, "LaMelo personal fouls")
				assert.Equal(t, int32(27), stats.Points, "LaMelo points")
				assert.Equal(t, float32(20.6), stats.GameScore, "LaMelo game score")
				assert.Equal(t, int32(8), stats.PlusMinus, "LaMelo plus minus")
			}
		}

	}
	assert.Equal(t, expectedNumAwayPlayers, numAwayPlayers, "Assert number of away players on Charlotte")
	assert.Equal(t, expectedNumHomePlayers, numHomePlayers, "Assert number of home players on LA Lakers")
	assert.Equal(t, expectedAwayDNP, numAwayDNP, "Assert number of away players on Charlotte that did not play")
	assert.Equal(t, expectedHomeDNP, numHomeDNP, "Assert number of home players on LA Lakers that did not play")
}

func TestGetH1BasicBoxScoreStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	date := "2025-02-19"
	matchupRunner := NewMatchupRunner(
		WithMatchupTimeout(4 * time.Minute),
	)
	matchups := matchupRunner.GetMatchups(date)
	boxScoreRunner := NewBasicBoxScoreRunner(
		WithBasicBoxScoreTimeout(4*time.Minute),
		WithBasicBoxScoreConcurrency(1),
		WithBasicBoxScorePeriod(H1),
	)
	basicBoxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)

	playerToTest := map[string]bool{
		"LaMelo Ball":  false,
		"Luka Dončić":  false,
		"Gabe Vincent": false,
	}

	for _, s := range basicBoxScoreStats {
		stats := s.(model.NBABasicBoxScoreStats)
		if stats.Player == "LaMelo Ball" {
			playerToTest["LaMelo Ball"] = true
			assert.Equal(t, true, stats.Starter, "LaMelo is a starter")
			assert.Equal(t, "ballla01", stats.PlayerID, "LaMelo's player ID")
			assert.Equal(t, "LA Lakers", stats.Opponent, "LaMelo's opposing team")
			assert.Equal(t, "https://www.basketball-reference.com/players/b/ballla01.html", stats.PlayerLink, "LaMelo's player link")
			assert.Equal(t, float32(14.78), stats.MinutesPlayed, "LaMelo minutes played")
			assert.Equal(t, int32(4), stats.FieldGoalsMade, "LaMelo field goals made")
			assert.Equal(t, int32(9), stats.FieldGoalAttempts, "LaMelo field goal attempts")
			assert.Equal(t, float32(0.444), stats.FieldGoalPercentage, "LaMelo field goal percentage")
			assert.Equal(t, int32(1), stats.ThreePointsMade, "LaMelo three points made")
			assert.Equal(t, int32(4), stats.ThreePointAttempts, "LaMelo three point attempts")
			assert.Equal(t, float32(0.250), stats.ThreePointPercentage, "LaMelo three point percentage")
			assert.Equal(t, int32(0), stats.FreeThrowsMade, "LaMelo free thows made")
			assert.Equal(t, int32(0), stats.FreeThrowAttempts, "LaMelo free throw attempts")
			assert.Equal(t, float32(0), stats.FreeThrowPercentage, "LaMelo free throw percentage")
			assert.Equal(t, int32(0), stats.OffensiveRebounds, "LaMelo offensive rebounds")
			assert.Equal(t, int32(0), stats.DefensiveRebounds, "LaMelo defensive rebounds")
			assert.Equal(t, int32(0), stats.TotalRebounds, "LaMelo total rebounds")
			assert.Equal(t, int32(3), stats.Assists, "LaMelo assists")
			assert.Equal(t, int32(0), stats.Steals, "LaMelo steals")
			assert.Equal(t, int32(0), stats.Blocks, "LaMelo blocks")
			assert.Equal(t, int32(2), stats.Turnovers, "LaMelo turnovers")
			assert.Equal(t, int32(1), stats.PersonalFouls, "LaMelo personal fouls")
			assert.Equal(t, int32(9), stats.Points, "LaMelo points")
			assert.Equal(t, float32(4.0), stats.GameScore, "LaMelo game score")
			assert.Equal(t, int32(-3), stats.PlusMinus, "LaMelo plus minus")
		} else if stats.Player == "Luka Dončić" {
			playerToTest["Luka Dončić"] = true
			assert.Equal(t, true, stats.Starter, "Luka is a starter")
			assert.Equal(t, "doncilu01", stats.PlayerID, "Luka's player ID")
			assert.Equal(t, "Charlotte", stats.Opponent, "Luka's opposing team")
			assert.Equal(t, "https://www.basketball-reference.com/players/d/doncilu01.html", stats.PlayerLink, "Luka's player link")
			assert.Equal(t, float32(15.55), stats.MinutesPlayed, "Luka minutes played")
			assert.Equal(t, int32(2), stats.FieldGoalsMade, "Luka field goals made")
			assert.Equal(t, int32(9), stats.FieldGoalAttempts, "Luka field goal attempts")
			assert.Equal(t, float32(0.222), stats.FieldGoalPercentage, "Luka field goal percentage")
			assert.Equal(t, int32(0), stats.ThreePointsMade, "Luka three points made")
			assert.Equal(t, int32(5), stats.ThreePointAttempts, "Luka three point attempts")
			assert.Equal(t, float32(0), stats.ThreePointPercentage, "Luka three point percentage")
			assert.Equal(t, int32(0), stats.FreeThrowsMade, "Luka free thows made")
			assert.Equal(t, int32(0), stats.FreeThrowAttempts, "Luka free throw attempts")
			assert.Equal(t, float32(0), stats.FreeThrowPercentage, "Luka free throw percentage")
			assert.Equal(t, int32(0), stats.OffensiveRebounds, "Luka offensive rebounds")
			assert.Equal(t, int32(7), stats.DefensiveRebounds, "Luka defensive rebounds")
			assert.Equal(t, int32(7), stats.TotalRebounds, "Luka total rebounds")
			assert.Equal(t, int32(5), stats.Assists, "Luka assists")
			assert.Equal(t, int32(0), stats.Steals, "Luka steals")
			assert.Equal(t, int32(0), stats.Blocks, "Luka blocks")
			assert.Equal(t, int32(5), stats.Turnovers, "Luka turnovers")
			assert.Equal(t, int32(1), stats.PersonalFouls, "Luka personal fouls")
			assert.Equal(t, int32(4), stats.Points, "Luka points")
			assert.Equal(t, float32(-1.3), stats.GameScore, "Luka game score")
			assert.Equal(t, int32(10), stats.PlusMinus, "Luka plus minus")

		} else if stats.Player == "Gabe Vincent" {
			playerToTest["Gabe Vincent"] = true
			assert.Equal(t, false, stats.Starter, "Gabe is a reserves player")
			assert.Equal(t, "vincega01", stats.PlayerID, "Gabe's player ID")
			assert.Equal(t, "Charlotte", stats.Opponent, "Gabe's opposing team")
			assert.Equal(t, "https://www.basketball-reference.com/players/v/vincega01.html", stats.PlayerLink, "Gabe's player link")
			assert.Equal(t, float32(10.65), stats.MinutesPlayed, "Gabe minutes played")
			assert.Equal(t, int32(2), stats.FieldGoalsMade, "Gabe field goals made")
			assert.Equal(t, int32(5), stats.FieldGoalAttempts, "Gabe field goal attempts")
			assert.Equal(t, float32(0.4), stats.FieldGoalPercentage, "Gabe field goal percentage")
			assert.Equal(t, int32(2), stats.ThreePointsMade, "Gabe three points made")
			assert.Equal(t, int32(5), stats.ThreePointAttempts, "Gabe three point attempts")
			assert.Equal(t, float32(0.4), stats.ThreePointPercentage, "Gabe three point percentage")
			assert.Equal(t, int32(0), stats.FreeThrowsMade, "Gabe free thows made")
			assert.Equal(t, int32(0), stats.FreeThrowAttempts, "Gabe free throw attempts")
			assert.Equal(t, float32(0), stats.FreeThrowPercentage, "Gabe free throw percentage")
			assert.Equal(t, int32(1), stats.OffensiveRebounds, "Gabe offensive rebounds")
			assert.Equal(t, int32(1), stats.DefensiveRebounds, "Gabe defensive rebounds")
			assert.Equal(t, int32(2), stats.TotalRebounds, "Gabe total rebounds")
			assert.Equal(t, int32(1), stats.Assists, "Gabe assists")
			assert.Equal(t, int32(0), stats.Steals, "Gabe steals")
			assert.Equal(t, int32(0), stats.Blocks, "Gabe blocks")
			assert.Equal(t, int32(1), stats.Turnovers, "Gabe turnovers")
			assert.Equal(t, int32(2), stats.PersonalFouls, "Gabe personal fouls")
			assert.Equal(t, int32(6), stats.Points, "Gabe points")
			assert.Equal(t, float32(3.2), stats.GameScore, "Gabe game score")
			assert.Equal(t, int32(10), stats.PlusMinus, "Gabe plus minus")

		}
	}
	assert.True(t, playerToTest["LaMelo Ball"], "LaMelo Ball's stats were collected")
	assert.True(t, playerToTest["Luka Dončić"], "Luka Dončić's stats were collected")
	assert.True(t, playerToTest["Gabe Vincent"], "Gabe Vincent's stats were collected")
}

func TestGetQ3BasicBoxScoreStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	date := "2025-02-19"
	matchupRunner := NewMatchupRunner(
		WithMatchupTimeout(4 * time.Minute),
	)
	matchups := matchupRunner.GetMatchups(date)
	boxScoreRunner := NewBasicBoxScoreRunner(
		WithBasicBoxScoreTimeout(4*time.Minute),
		WithBasicBoxScoreConcurrency(1),
		WithBasicBoxScorePeriod(Q3),
	)
	basicBoxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)
	playerToTest := map[string]bool{
		"Tidjane Salaün": false,
		"Rui Hachimura":  false,
	}

	for _, s := range basicBoxScoreStats {
		stats := s.(model.NBABasicBoxScoreStats)
		if stats.Player == "Tidjane Salaün" {
			playerToTest["Tidjane Salaün"] = true
			assert.Equal(t, false, stats.Starter, "Tidjane is a reserves player")
			assert.Equal(t, "salauti01", stats.PlayerID, "Tidjane's player ID")
			assert.Equal(t, "LA Lakers", stats.Opponent, "Tidjane's opposing team")
			assert.Equal(t, "https://www.basketball-reference.com/players/s/salauti01.html", stats.PlayerLink, "Tidjane's player link")
			assert.Equal(t, float32(0), stats.MinutesPlayed, "Tidjane minutes played")
			assert.Equal(t, int32(0), stats.FieldGoalsMade, "Tidjane field goals made")
			assert.Equal(t, int32(0), stats.FieldGoalAttempts, "Tidjane field goal attempts")
			assert.Equal(t, float32(0), stats.FieldGoalPercentage, "Tidjane field goal percentage")
			assert.Equal(t, int32(0), stats.ThreePointsMade, "Tidjane three points made")
			assert.Equal(t, int32(0), stats.ThreePointAttempts, "Tidjane three point attempts")
			assert.Equal(t, float32(0), stats.ThreePointPercentage, "Tidjane three point percentage")
			assert.Equal(t, int32(0), stats.FreeThrowsMade, "Tidjane free thows made")
			assert.Equal(t, int32(0), stats.FreeThrowAttempts, "Tidjane free throw attempts")
			assert.Equal(t, float32(0), stats.FreeThrowPercentage, "Tidjane free throw percentage")
			assert.Equal(t, int32(0), stats.OffensiveRebounds, "Tidjane offensive rebounds")
			assert.Equal(t, int32(0), stats.DefensiveRebounds, "Tidjane defensive rebounds")
			assert.Equal(t, int32(0), stats.TotalRebounds, "Tidjane total rebounds")
			assert.Equal(t, int32(0), stats.Assists, "Tidjane assists")
			assert.Equal(t, int32(0), stats.Steals, "Tidjane steals")
			assert.Equal(t, int32(0), stats.Blocks, "Tidjane blocks")
			assert.Equal(t, int32(0), stats.Turnovers, "Tidjane turnovers")
			assert.Equal(t, int32(0), stats.PersonalFouls, "Tidjane personal fouls")
			assert.Equal(t, int32(0), stats.Points, "Tidjane points")
			assert.Equal(t, float32(0), stats.GameScore, "Tidjane game score")
			assert.Equal(t, int32(0), stats.PlusMinus, "Tidjane plus minus")
		} else if stats.Player == "Rui Hachimura" {
			playerToTest["Rui Hachimura"] = true
			assert.Equal(t, true, stats.Starter, "Rui is a starter")
			assert.Equal(t, "hachiru01", stats.PlayerID, "Rui's player ID")
			assert.Equal(t, "Charlotte", stats.Opponent, "Rui's opposing team")
			assert.Equal(t, "https://www.basketball-reference.com/players/h/hachiru01.html", stats.PlayerLink, "Rui's player link")
			assert.Equal(t, float32(10.27), stats.MinutesPlayed, "Rui minutes played")
			assert.Equal(t, int32(2), stats.FieldGoalsMade, "Rui field goals made")
			assert.Equal(t, int32(5), stats.FieldGoalAttempts, "Rui field goal attempts")
			assert.Equal(t, float32(.4), stats.FieldGoalPercentage, "Rui field goal percentage")
			assert.Equal(t, int32(0), stats.ThreePointsMade, "Rui three points made")
			assert.Equal(t, int32(2), stats.ThreePointAttempts, "Rui three point attempts")
			assert.Equal(t, float32(0), stats.ThreePointPercentage, "Rui three point percentage")
			assert.Equal(t, int32(0), stats.FreeThrowsMade, "Rui free thows made")
			assert.Equal(t, int32(0), stats.FreeThrowAttempts, "Rui free throw attempts")
			assert.Equal(t, float32(0), stats.FreeThrowPercentage, "Rui free throw percentage")
			assert.Equal(t, int32(1), stats.OffensiveRebounds, "Rui offensive rebounds")
			assert.Equal(t, int32(0), stats.DefensiveRebounds, "Rui defensive rebounds")
			assert.Equal(t, int32(1), stats.TotalRebounds, "Rui total rebounds")
			assert.Equal(t, int32(0), stats.Assists, "Rui assists")
			assert.Equal(t, int32(0), stats.Steals, "Rui steals")
			assert.Equal(t, int32(0), stats.Blocks, "Rui blocks")
			assert.Equal(t, int32(0), stats.Turnovers, "Rui turnovers")
			assert.Equal(t, int32(0), stats.PersonalFouls, "Rui personal fouls")
			assert.Equal(t, int32(4), stats.Points, "Rui points")
			assert.Equal(t, float32(2), stats.GameScore, "Rui game score")
			assert.Equal(t, int32(-5), stats.PlusMinus, "Rui plus minus")

		}

	}
	assert.True(t, playerToTest["Tidjane Salaün"], "Tidjane Salaün's stats were collected")
	assert.True(t, playerToTest["Rui Hachimura"], "Rui Hachimura's stats were collected")
}
