package nba_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for nba.MatchupScraper
func ExampleMatchupScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for nba.MatchupPeriodsScraper
func ExampleMatchupPeriodsScraper() {
	matchupScraper := nba.NewMatchupPeriodsScraper(
		nba.WithMatchupPeriodsDate("2025-06-05"),
		nba.WithMatchupPeriodsTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	records, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, period := range records {
		jsonBytes, err := json.MarshalIndent(period, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for nba.PlayByPlayScraper
func ExamplePlayByPlayScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	playbyplayscraper := nba.NewPlayByPlayScraper(
		nba.WithPlayByPlayTimeout(2 * time.Minute),
	)

	playbyplayrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(playbyplayscraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := playbyplayrunner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, period := range records {
		jsonBytes, err := json.MarshalIndent(period, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
