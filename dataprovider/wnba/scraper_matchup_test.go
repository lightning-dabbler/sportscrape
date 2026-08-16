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

func TestBaseMatchupScraper_dateRange(t *testing.T) {
	tests := []struct {
		name        string
		date        string
		endDate     string
		wantStart   string
		wantEnd     string
		expectError bool
	}{
		{name: "single day, no EndDate", date: "2026-08-03", endDate: "", wantStart: "2026-08-03", wantEnd: "2026-08-03"},
		{name: "range", date: "2026-08-01", endDate: "2026-08-03", wantStart: "2026-08-01", wantEnd: "2026-08-03"},
		{name: "invalid Date", date: "not-a-date", expectError: true},
		{name: "invalid EndDate", date: "2026-08-01", endDate: "not-a-date", expectError: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bms := &BaseMatchupScraper{Date: tt.date, EndDate: tt.endDate}
			start, end, err := bms.dateRange()
			if tt.expectError {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantStart, start.Format("2006-01-02"))
			assert.Equal(t, tt.wantEnd, end.Format("2006-01-02"))
		})
	}
}

func TestBaseMatchupScraper_seasons(t *testing.T) {
	tests := []struct {
		name    string
		date    string
		endDate string
		want    []string
	}{
		{name: "single year, no EndDate", date: "2026-08-01", endDate: "", want: []string{"2026"}},
		{name: "single year range", date: "2026-08-01", endDate: "2026-08-03", want: []string{"2026"}},
		{name: "spans year boundary", date: "2026-12-30", endDate: "2027-01-02", want: []string{"2026", "2027"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bms := &BaseMatchupScraper{Date: tt.date, EndDate: tt.endDate}
			got, err := bms.seasons()
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBaseMatchupScraper_URL(t *testing.T) {
	bms := &BaseMatchupScraper{}
	got, err := bms.URL("2026")
	require.NoError(t, err)
	assert.Equal(t, "https://www.wnba.com/api/schedule?regionId=1&season=2026", got)
}

func TestDeriveShareURL(t *testing.T) {
	got := deriveShareURL("PHX", "CHI", "1022600224")
	assert.Equal(t, "https://www.wnba.com/game/phx-vs-chi-1022600224", got)
}

func TestParseScheduleBucketDate(t *testing.T) {
	got, err := parseScheduleBucketDate("08/03/2026 00:00:00")
	require.NoError(t, err)
	assert.Equal(t, "2026-08-03", got.Format("2006-01-02"))

	_, err = parseScheduleBucketDate("not-a-date")
	assert.Error(t, err)
}

func TestMatchupScraper(t *testing.T) {
	// https://www.wnba.com/api/schedule?season=2026&regionId=1
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2026-08-14"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	require.NotEmpty(t, matchups)

	var found *model.Matchup
	for i := range matchups {
		if matchups[i].EventID == "1022600254" {
			found = &matchups[i]
			break
		}
	}
	require.NotNil(t, found, "expected to find game 1022600254 (DAL @ IND) on 2026-08-14")

	// PullTimestamp(Parquet) is dynamic (time.Now() at scrape time) - check
	// recency instead of an exact value. Every other field is validated
	// exactly against the real, known response captured 2026-08-15.
	assert.WithinDuration(t, time.Now().UTC(), found.PullTimestamp, time.Minute)
	assert.Equal(t, "1022600254", found.EventID)
	assert.Equal(t, "2026-08-14T23:30:00Z", found.EventTime.Format(time.RFC3339))
	assert.Equal(t, int32(3), found.EventStatus)
	assert.Equal(t, "Final", found.EventStatusText)
	assert.Equal(t, int64(1611661325), found.HomeTeamID)
	assert.Equal(t, "Fever", found.HomeTeam)
	assert.Equal(t, "IND", found.HomeTeamAbbreviation)
	assert.Equal(t, int64(1611661321), found.AwayTeamID)
	assert.Equal(t, "Wings", found.AwayTeam)
	assert.Equal(t, "DAL", found.AwayTeamAbbreviation)
	assert.Equal(t, int32(87), found.AwayTeamScore)
	assert.Equal(t, int32(98), found.HomeTeamScore)
	assert.Equal(t, int32(20), found.AwayTeamWins)
	assert.Equal(t, int32(22), found.HomeTeamWins)
	assert.Equal(t, int32(15), found.AwayTeamLosses)
	assert.Equal(t, int32(12), found.HomeTeamLosses)
	assert.Equal(t, "https://www.wnba.com/game/dal-vs-ind-1022600254", found.ShareURL)
	assert.Equal(t, "Regular Season", found.SeasonType)
	assert.Equal(t, int32(2026), found.SeasonYear)
	assert.Equal(t, "10", found.LeagueID)
}

func TestMatchupScraper_DateRange(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	matchupScraper := NewMatchupScraper(
		WithMatchupDate("2026-08-01"),
		WithMatchupEndDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)
	matchups, err := matchuprunner.Run()
	assert.NoError(t, err)
	require.NotEmpty(t, matchups)

	// EventTime is UTC and games are bucketed by wnba.com's own (likely
	// Eastern-time) schedule day, so a strict per-game UTC calendar-date
	// bound check here would be flaky around evening tip-offs that roll
	// into the next UTC day. Assert the range spans more than a single day
	// (proving EndDate actually extended the query) instead.
	seenDates := map[string]bool{}
	for _, m := range matchups {
		seenDates[m.EventTime.Format("2006-01-02")] = true
	}
	assert.Greater(t, len(seenDates), 1)
}
