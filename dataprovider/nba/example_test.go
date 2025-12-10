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
	for _, record := range records {
		jsonBytes, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for nba.BoxScoreUsageScraper full
func ExampleBoxScoreUsageScraper_full() {
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

	boxscorescraper := nba.NewBoxScoreUsageScraper(
		nba.WithBoxScoreUsageTimeout(2*time.Minute),
		nba.WithBoxScoreUsagePeriod(nba.Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
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

// Example for nba.BoxScoreUsageScraper h2
func ExampleBoxScoreUsageScraper_h2() {
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

	boxscorescraper := nba.NewBoxScoreUsageScraper(
		nba.WithBoxScoreUsageTimeout(2*time.Minute),
		nba.WithBoxScoreUsagePeriod(nba.H2),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
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

// Example for nba.BoxScoreTraditionalScraper q1
func ExampleBoxScoreTraditionalScraper_q1() {
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

	boxscorescraper := nba.NewBoxScoreTraditionalScraper(
		nba.WithBoxScoreTraditionalTimeout(2*time.Minute),
		nba.WithBoxScoreTraditionalPeriod(nba.Q1),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
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

// Example for nba.BoxScoreAdvancedScraper full
func ExampleBoxScoreAdvancedScraper_full() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-11"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreAdvancedScraper(
		nba.WithBoxScoreAdvancedTimeout(2*time.Minute),
		nba.WithBoxScoreAdvancedPeriod(nba.Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
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

// Example for nba.BoxScoreScoringScraper h1
func ExampleBoxScoreScoringScraper_h1() {
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

	boxscorescraper := nba.NewBoxScoreScoringScraper(
		nba.WithBoxScoreScoringTimeout(2*time.Minute),
		nba.WithBoxScoreScoringPeriod(nba.H1),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
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

// Example for nba.BoxScoreMiscScraper full
func ExampleBoxScoreMiscScraper_full() {
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

	boxscorescraper := nba.NewBoxScoreMiscScraper(
		nba.WithBoxScoreMiscTimeout(2*time.Minute),
		nba.WithBoxScoreMiscPeriod(nba.Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(boxscorescraper),
		runner.EventDataRunnerConcurrency(1),
	)

	records, err := boxscorerunner.Run(matchups...)
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
