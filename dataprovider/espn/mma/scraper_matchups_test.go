//go:build integration

package mma

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/lightning-dabbler/sportscrape/scraper"

	"github.com/stretchr/testify/assert"
)

func TestESPNMMMAMatchupScraper(T *testing.T) {
	if testing.Short() {
		T.Skip("Skipping integration test")
	}
	matchupscraper := ESPNMMAMatchupScraper{Year: "2024", League: "ufc", BaseScraper: scraper.BaseScraper{Timeout: 3 * time.Minute}}
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	r, err := matchuprunner.Run()
	assert.NoError(T, err)
	output := r
	assert.NotEmpty(T, output)
	for _, matchup := range output {
		if matchup.EventID == "600039853" {
			matchup.PullTimestamp = time.Time{}
			matchup.PullTimestampParquet = 0
			assert.Equal(T,
				model.Matchup{
					PullTimestamp:          time.Time{},
					PullTimestampParquet:   0,
					EventID:                "600039853",
					EventTime:              time.Date(2024, time.September, 14, 23, 30, 0, 0, time.UTC),
					EventTimeParquet:       1726356600000,
					LeagueID:               "3321",
					LeagueName:             "Ultimate Fighting Championship",
					Date:                   "",
					Completed:              true,
					Link:                   "/mma/fightcenter/_/id/600039853/league/ufc",
					Name:                   "UFC 306 – Riyadh Season Noche UFC: O’Malley vs. Dvalishvili",
					IsPostponedOrCancelled: false,
					StatusID:               "3",
					StatusState:            "post",
					StatusDetail:           "Final"},
				matchup,
			)
			return
		}

	}
	assert.Fail(T, "Matchup not found")
}
