package nhl_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for nhl.MatchupScraper
func ExampleMatchupScraper() {
	matchupscraper := nhl.NewMatchupScraper(
		nhl.WithMatchupDate("2024-11-12"),
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

// Example for nhl.MatchupPeriodsScraper
func ExampleMatchupPeriodsScraper() {
	matchupscraper := nhl.NewMatchupScraper(
		nhl.WithMatchupDate("2026-09-24"),
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

	eventdatascraper := nhl.NewMatchupPeriodsScraper()
	eventdatarunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
			Scraper:     eventdatascraper,
			Concurrency: 1,
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

// Example for nhl.SkaterBoxScoreScraper
func ExampleSkaterBoxScoreScraper() {
	matchupscraper := nhl.NewMatchupScraper(
		nhl.WithMatchupDate("2024-11-12"),
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

	eventdatascraper := nhl.NewSkaterBoxScoreScraper()
	eventdatarunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.SkaterBoxScore]{
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

// Example for nhl.GoalieBoxScoreScraper
func ExampleGoalieBoxScoreScraper() {
	matchupscraper := nhl.NewMatchupScraper(
		nhl.WithMatchupDate("2024-11-12"),
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

	eventdatascraper := nhl.NewGoalieBoxScoreScraper()
	eventdatarunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.GoalieBoxScore]{
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

// Example for nhl.PlayByPlayScraper
func ExamplePlayByPlayScraper() {
	matchupscraper := nhl.NewMatchupScraper(
		nhl.WithMatchupDate("2024-11-12"),
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

	eventdatascraper := nhl.NewPlayByPlayScraper()
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
