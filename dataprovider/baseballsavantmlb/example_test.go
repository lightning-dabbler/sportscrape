package baseballsavantmlb_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb"
)

// Example for baseballsavantmlb.MatchupScraper
func ExampleMatchupScraper() {
	date := "2025-06-25"
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate(date),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
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

// Example for baseballsavantmlb.FieldingBoxScoreScraper
func ExampleFieldingBoxScoreScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := baseballsavantmlb.NewFieldingBoxScoreScraper()

	boxscorerunner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerScraper(boxscorescraper),
		sportscrape.EventDataRunnerConcurrency(1),
	)

	stats, err := boxscorerunner.Run(matchups...)
	if err != nil {
		panic(err)
	}

	// Output each statline as pretty json
	for _, statline := range stats {
		jsonBytes, err := json.MarshalIndent(statline, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}

}

// Example for baseballsavantmlb.BattingBoxScoreScraper
func ExampleBattingBoxScoreScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := baseballsavantmlb.NewBattingBoxScoreScraper()

	boxscorerunner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerScraper(boxscorescraper),
		sportscrape.EventDataRunnerConcurrency(1),
	)

	stats, err := boxscorerunner.Run(matchups...)
	if err != nil {
		panic(err)
	}

	// Output each statline as pretty json
	for _, statline := range stats {
		jsonBytes, err := json.MarshalIndent(statline, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}

}

// Example for baseballsavantmlb.PitchingBoxScoreScraper
func ExamplePitchingBoxScoreScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := baseballsavantmlb.NewPitchingBoxScoreScraper()

	boxscorerunner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerScraper(boxscorescraper),
		sportscrape.EventDataRunnerConcurrency(1),
	)

	stats, err := boxscorerunner.Run(matchups...)
	if err != nil {
		panic(err)
	}

	// Output each statline as pretty json
	for _, statline := range stats {
		jsonBytes, err := json.MarshalIndent(statline, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for baseballsavantmlb.PlayByPlayScraper
func ExamplePlayByPlayScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupscraper),
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	playbyplayscraper := baseballsavantmlb.NewPlayByPlayScraper()

	playbyplayrunner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerScraper(playbyplayscraper),
		sportscrape.EventDataRunnerConcurrency(1),
	)

	plays, err := playbyplayrunner.Run(matchups...)
	if err != nil {
		panic(err)
	}

	// Output each play as pretty json
	for _, play := range plays {
		jsonBytes, err := json.MarshalIndent(play, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}

}
