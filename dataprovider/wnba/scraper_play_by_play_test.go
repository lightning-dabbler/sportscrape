//go:build integration

package wnba

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlayByPlayScraper(t *testing.T) {
	// https://www.wnba.com/game/dal-vs-ind-1022600254/play-by-play?period=All
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(WithMatchupDate("2026-08-14"))
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{Scraper: matchupScraper},
	)
	matchups, err := matchuprunner.Run()
	require.NoError(t, err)

	var matchup model.Matchup
	var found bool
	for _, m := range matchups {
		if m.EventID == "1022600254" {
			matchup = m
			found = true
			break
		}
	}
	require.True(t, found, "expected to find game 1022600254 (DAL @ IND) on 2026-08-14")

	pbpscraper := NewPlayByPlayScraper(WithPlayByPlayTimeout(3 * time.Minute))
	pbpscraper.NetworkHeaders = NetworkHeaders
	pbprunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper:     pbpscraper,
			Concurrency: 1,
		},
	)

	records, err := pbprunner.Run([]model.Matchup{matchup})
	assert.NoError(t, err)
	require.NotEmpty(t, records, "expected at least one play-by-play action")

	// Last action: validate every field against the real, known response
	// captured 2026-08-15 (end of 4th period).
	last := records[len(records)-1]
	assert.WithinDuration(t, time.Now().UTC(), last.PullTimestamp, time.Minute)
	assert.Equal(t, "1022600254", last.EventID)
	assert.Equal(t, "2026-08-14T23:30:00Z", last.EventTime.Format(time.RFC3339))
	assert.Equal(t, int32(3), last.EventStatus)
	assert.Equal(t, "Final", last.EventStatusText)
	assert.Equal(t, int32(560), last.ActionNumber)
	assert.Equal(t, float32(0), last.Clock)
	assert.Equal(t, int32(4), last.Period)
	assert.Equal(t, int64(0), last.TeamID)
	assert.Equal(t, "", last.TeamAbbreviation)
	assert.Equal(t, int64(0), last.PersonID)
	assert.Equal(t, "", last.PlayerName)
	assert.Equal(t, "", last.PlayerNameInitial)
	assert.Equal(t, float32(0), last.ShotDistance)
	assert.Equal(t, "", last.ShotResult)
	assert.Equal(t, int32(0), last.IsFieldGoal)
	assert.Equal(t, "98", last.ScoreHome)
	assert.Equal(t, "87", last.ScoreAway)
	assert.Equal(t, int32(185), last.PointsTotal)
	assert.Equal(t, "", last.Location)
	assert.Equal(t, "End of 4th Period (9:40 PM EST)", last.Description)
	assert.Equal(t, "period", last.ActionType)
	assert.Equal(t, "end", last.SubType)
	assert.Equal(t, int32(0), last.ShotValue)
	assert.Equal(t, int32(405), last.ActionID)

	// action_id 9 (a made shot): validate every field against the real,
	// known response captured 2026-08-15.
	var madeShot *model.PlayByPlay
	for i := range records {
		if records[i].ActionID == 9 {
			madeShot = &records[i]
			break
		}
	}
	require.NotNil(t, madeShot, "expected to find action_id 9 (Boston putback layup)")
	assert.Equal(t, "1022600254", madeShot.EventID)
	assert.Equal(t, int32(13), madeShot.ActionNumber)
	assert.Equal(t, float32(8.95), madeShot.Clock)
	assert.Equal(t, int32(1), madeShot.Period)
	assert.Equal(t, int64(1611661325), madeShot.TeamID)
	assert.Equal(t, "IND", madeShot.TeamAbbreviation)
	assert.Equal(t, int64(1641648), madeShot.PersonID)
	assert.Equal(t, "Boston", madeShot.PlayerName)
	assert.Equal(t, "A. Boston", madeShot.PlayerNameInitial)
	assert.Equal(t, float32(2), madeShot.ShotDistance)
	assert.Equal(t, "Made", madeShot.ShotResult)
	assert.Equal(t, int32(1), madeShot.IsFieldGoal)
	assert.Equal(t, "2", madeShot.ScoreHome)
	assert.Equal(t, "0", madeShot.ScoreAway)
	assert.Equal(t, int32(2), madeShot.PointsTotal)
	assert.Equal(t, "h", madeShot.Location)
	assert.Equal(t, "Boston 2' Putback Layup (2 PTS)", madeShot.Description)
	assert.Equal(t, "Made Shot", madeShot.ActionType)
	assert.Equal(t, "Putback Layup Shot", madeShot.SubType)
	assert.Equal(t, int32(2), madeShot.ShotValue)
}
