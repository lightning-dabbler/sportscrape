//go:build integration

package nba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestPlayByPlayScraper(t *testing.T) {
	// https://www.nba.com/games?date=2025-06-11
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2025-06-11"),
		WithMatchupTimeout(3*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	playbyplayscraper := NewPlayByPlayScraper(
		WithPlayByPlayTimeout(2 * time.Minute),
	)

	playbyplayrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(playbyplayscraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := playbyplayrunner.Run(matchups...)
	assert.NoError(t, err)
	n_records := len(records)
	assert.Equal(t, 521, n_records, "521 plays")
	testRecord := records[517].(model.PlayByPlay)

	assert.Equal(t, "0042400403", testRecord.EventID)
	assert.Equal(t, int32(3), testRecord.EventStatus)
	assert.Equal(t, "Final", testRecord.EventStatusText)
	assert.Equal(t, int32(745), testRecord.ActionNumber)
	assert.Equal(t, "PT00M07.70S", testRecord.Clock)
	assert.Equal(t, int32(4), testRecord.Period)
	assert.Equal(t, int64(1610612754), testRecord.TeamID)
	assert.Equal(t, "IND", testRecord.TeamAbbreviation)
	assert.Equal(t, int64(1627783), testRecord.PersonID)
	assert.Equal(t, "Siakam", testRecord.PlayerName)
	assert.Equal(t, "P. Siakam", testRecord.PlayerNameInitial)
	assert.Equal(t, int32(1), testRecord.ShotDistance)
	assert.Equal(t, "Made", testRecord.ShotResult)
	assert.Equal(t, int32(1), testRecord.IsFieldGoal)
	assert.Equal(t, "116", testRecord.ScoreHome)
	assert.Equal(t, "107", testRecord.ScoreAway)
	assert.Equal(t, int32(223), testRecord.PointsTotal)
	assert.Equal(t, "h", testRecord.Location)
	assert.Equal(t, "Siakam 1' Driving Layup (21 PTS)", testRecord.Description)
	assert.Equal(t, "Made Shot", testRecord.ActionType)
	assert.Equal(t, "Driving Layup Shot", testRecord.SubType)
	assert.Equal(t, int32(2), testRecord.ShotValue)
	assert.Equal(t, int32(518), testRecord.ActionID)
}
