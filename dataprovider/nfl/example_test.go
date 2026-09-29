package nfl_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for nfl.MatchupScraper
func ExampleMatchupScraper() {
	matchupscraper := nfl.NewMatchupScraper(
		nfl.WithMatchupDate("2025-09-04"),
	)
	// Optional: retry fetches that fail (e.g. rate limited) up to 5 times
	matchupscraper.FetchAttempts = 5
	matchupscraper.FetchRetryBackoff = 5 * time.Second
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each matchup as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for nfl.MatchupPeriodsScraper
func ExampleMatchupPeriodsScraper() {
	matchupscraper := nfl.NewMatchupScraper(
		nfl.WithMatchupDate("2025-09-28"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	eventdatascraper := nfl.NewMatchupPeriodsScraper()
	eventdatarunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
			Scraper:     eventdatascraper,
			Concurrency: 4,
		},
	)
	records, err := eventdatarunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each period as pretty json
	for _, record := range records {
		jsonBytes, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for nfl.PassingBoxScoreScraper
// The other box score scrapers (Rushing, Receiving, Defense, Kicking, Kickoff, Punting, KickReturn, PuntReturn,
// Fumbles, Interceptions) are used the same way
func ExamplePassingBoxScoreScraper() {
	matchupscraper := nfl.NewMatchupScraper(
		nfl.WithMatchupDate("2025-09-04"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	eventdatascraper := nfl.NewPassingBoxScoreScraper()
	eventdatarunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PassingBoxScore]{
			Scraper:     eventdatascraper,
			Concurrency: 1,
		},
	)
	records, err := eventdatarunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each statline as pretty json
	for _, record := range records {
		jsonBytes, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for nfl.PlayByPlayScraper
func ExamplePlayByPlayScraper() {
	matchupscraper := nfl.NewMatchupScraper(
		nfl.WithMatchupDate("2025-09-04"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	eventdatascraper := nfl.NewPlayByPlayScraper()
	eventdatarunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper:     eventdatascraper,
			Concurrency: 1,
		},
	)
	records, err := eventdatarunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each play as pretty json
	for _, record := range records {
		jsonBytes, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for nfl.PlayByPlayStatsScraper
func ExamplePlayByPlayStatsScraper() {
	matchupscraper := nfl.NewMatchupScraper(
		nfl.WithMatchupDate("2025-09-04"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	eventdatascraper := nfl.NewPlayByPlayStatsScraper()
	eventdatarunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlayStat]{
			Scraper:     eventdatascraper,
			Concurrency: 1,
		},
	)
	records, err := eventdatarunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each play stat as pretty json
	for _, record := range records {
		jsonBytes, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
