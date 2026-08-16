package wnba_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for wnba.MatchupScraper
func ExampleMatchupScraper() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
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

// Example for wnba.MatchupScraper with a date range via WithMatchupEndDate
func ExampleMatchupScraper_dateRange() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-01"),
		wnba.WithMatchupEndDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
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

// Example for wnba.MatchupPeriodsScraper
func ExampleMatchupPeriodsScraper() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-14"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	periodsscraper := wnba.NewMatchupPeriodsScraper(
		wnba.WithMatchupPeriodsTimeout(2 * time.Minute),
	)
	periodsscraper.NetworkHeaders = wnba.NetworkHeaders
	periodsrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
			Scraper:     periodsscraper,
			Concurrency: 1,
		},
	)

	records, err := periodsrunner.Run(matchups)
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

// Example for wnba.PlayByPlayScraper
func ExamplePlayByPlayScraper() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-14"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	playbyplayscraper := wnba.NewPlayByPlayScraper(
		wnba.WithPlayByPlayTimeout(2 * time.Minute),
	)
	playbyplayscraper.NetworkHeaders = wnba.NetworkHeaders
	playbyplayrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper:     playbyplayscraper,
			Concurrency: 1,
		},
	)

	records, err := playbyplayrunner.Run(matchups)
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

// Example for wnba.BoxScoreTraditionalScraper full
func ExampleBoxScoreTraditionalScraper_full() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := wnba.NewBoxScoreTraditionalScraper(
		wnba.WithBoxScoreTraditionalTimeout(2*time.Minute),
		wnba.WithBoxScoreTraditionalPeriod(wnba.Full),
	)
	boxscorescraper.NetworkHeaders = wnba.NetworkHeaders
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreTraditional]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
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

// Example for wnba.BoxScoreAdvancedScraper full
func ExampleBoxScoreAdvancedScraper_full() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := wnba.NewBoxScoreAdvancedScraper(
		wnba.WithBoxScoreAdvancedTimeout(2*time.Minute),
		wnba.WithBoxScoreAdvancedPeriod(wnba.Full),
	)
	boxscorescraper.NetworkHeaders = wnba.NetworkHeaders
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreAdvanced]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
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

// Example for wnba.BoxScoreMiscScraper full
func ExampleBoxScoreMiscScraper_full() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := wnba.NewBoxScoreMiscScraper(
		wnba.WithBoxScoreMiscTimeout(2*time.Minute),
		wnba.WithBoxScoreMiscPeriod(wnba.Full),
	)
	boxscorescraper.NetworkHeaders = wnba.NetworkHeaders
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreMisc]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
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

// Example for wnba.BoxScoreScoringScraper h1
func ExampleBoxScoreScoringScraper_h1() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := wnba.NewBoxScoreScoringScraper(
		wnba.WithBoxScoreScoringTimeout(2*time.Minute),
		wnba.WithBoxScoreScoringPeriod(wnba.H1),
	)
	boxscorescraper.NetworkHeaders = wnba.NetworkHeaders
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreScoring]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
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

// Example for wnba.BoxScoreUsageScraper full
func ExampleBoxScoreUsageScraper_full() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := wnba.NewBoxScoreUsageScraper(
		wnba.WithBoxScoreUsageTimeout(2*time.Minute),
		wnba.WithBoxScoreUsagePeriod(wnba.Full),
	)
	boxscorescraper.NetworkHeaders = wnba.NetworkHeaders
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreUsage]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
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

// Example for wnba.BoxScoreFourFactorsScraper full
func ExampleBoxScoreFourFactorsScraper_full() {
	matchupScraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate("2026-08-03"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := wnba.NewBoxScoreFourFactorsScraper(
		wnba.WithBoxScoreFourFactorsTimeout(2*time.Minute),
		wnba.WithBoxScoreFourFactorsPeriod(wnba.Full),
	)
	boxscorescraper.NetworkHeaders = wnba.NetworkHeaders
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreFourFactors]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
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
