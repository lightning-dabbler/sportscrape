//go:build integration

package mma

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	scraper2 "github.com/lightning-dabbler/sportscrape/scraper"

	"github.com/stretchr/testify/assert"
)

func TestESPNMMMAMatchupScraper(T *testing.T) {
	scraper := ESPNMMAMatchupScraper{Year: "2024", League: "ufc", BaseScraper: scraper2.BaseScraper{Timeout: 3 * time.Minute}}

	matchupRunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(
			scraper,
		),
	)

	r, err := matchupRunner.Run()
	assert.NoError(T, err)
	output := r
	assert.NotEmpty(T, output)
	for _, untyped := range output {
		matchup := untyped.(model.Matchup)
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
