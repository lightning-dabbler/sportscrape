package baseballreferencemlb_test

import (
	"fmt"
	"log"
	"time"

	"encoding/json"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for baseballreferencemlb.MatchupScraper
func ExampleMatchupRunner() {
	date := "2024-10-30"

	matchupscraper := baseballreferencemlb.NewMatchupScraper(
		baseballreferencemlb.WithMatchupDate(date),
		baseballreferencemlb.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.MLBMatchup]{
			Scraper: matchupscraper,
		},
	)
	// Retrieve MLB matchups associated with date
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup)
	}
}

// Example for baseballreferencemlb.BattingBoxScoreScraper
func ExampleBattingBoxScoreScraper() {
	date := "2024-10-13"
	// Instantiate MatchupRunner
	matchupscraper := baseballreferencemlb.NewMatchupScraper(
		baseballreferencemlb.WithMatchupDate(date),
		baseballreferencemlb.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.MLBMatchup]{
			Scraper: matchupscraper,
		},
	)
	// Retrieve MLB matchups associated with date
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Instantiate BattingBoxScoreScraper
	boxscorescraper := baseballreferencemlb.NewBattingBoxScoreScraper(
		baseballreferencemlb.WithBattingBoxScoreTimeout(4 * time.Minute),
	)
	boxScoreRunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.MLBMatchup, model.MLBBattingBoxScoreStats]{
			Concurrency: 1,
			Scraper:     boxscorescraper,
		},
	)
	// Retrieve MLB batting box score stats associated with matchups
	boxScoreStats, err := boxScoreRunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, stats := range boxScoreStats {
		jsonBytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for baseballreferencemlb.PitchingBoxScoreScraper
func ExamplePitchingBoxScoreScraper() {
	date := "2024-10-30"
	// Instantiate MatchupRunner
	matchupscraper := baseballreferencemlb.NewMatchupScraper(
		baseballreferencemlb.WithMatchupDate(date),
		baseballreferencemlb.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.MLBMatchup]{
			Scraper: matchupscraper,
		},
	)
	// Retrieve MLB matchups associated with date
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	// Instantiate PitchingBoxScoreScraper
	boxscorescraper := baseballreferencemlb.NewPitchingBoxScoreScraper(
		baseballreferencemlb.WithPitchingBoxScoreTimeout(4 * time.Minute),
	)
	boxScoreRunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.MLBMatchup, model.MLBPitchingBoxScoreStats]{
			Concurrency: 1,
			Scraper:     boxscorescraper,
		},
	)
	// Retrieve MLB pitching box score stats associated with matchups
	boxScoreStats, err := boxScoreRunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, stats := range boxScoreStats {
		jsonBytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}

}
