package foxsports_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for foxsports.MatchupScraper NBA
func ExampleMatchupScraper_nba() {
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.NBA),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
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

// Example for foxsports.MatchupScraper MLB
func ExampleMatchupScraper_mlb() {
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.MLB),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-08-02"}),
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

// Example for foxsports.MatchupScraper NCAAB
func ExampleMatchupScraper_ncaab() {
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.NCAAB),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2025-01-10"}),
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

// Example for foxsports.MatchupScraper NFL
func ExampleMatchupScraper_nfl() {
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.NFL),
		foxsports.MatchupScraperSegmenter(&foxsports.NFLSegmenter{Year: 2024, Week: 4, Season: foxsports.POSTSEASON}),
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

// Example for foxsports.NBABoxScoreScraper
func ExampleNBABoxScoreScraper() {
	// Get matchups
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.NBA),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
	)

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	// Get boxscore data
	eventdatascraper := foxsports.NewNBABoxScoreScraper()
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(4),
		runner.EventDataRunnerScraper(
			eventdatascraper,
		),
	)
	boxScoreStats, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	for _, statline := range boxScoreStats {
		fmt.Printf("%#v\n", statline.(model.NBABoxScoreStats))
	}
}

// Example for foxsports.MLBBattingBoxScoreScraper
func ExampleMLBBattingBoxScoreScraper() {
	// Get matchups
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.MLB),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	// Get boxscore data
	eventdatascraper := foxsports.NewMLBBattingBoxScoreScraper()
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(4),
		runner.EventDataRunnerScraper(
			eventdatascraper,
		),
	)

	boxScoreStats, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	for _, statline := range boxScoreStats {
		fmt.Printf("%#v\n", statline.(model.MLBBattingBoxScoreStats))
	}
}

// Example for foxsports.MLBPitchingBoxScoreScraper
func ExampleMLBPitchingBoxScoreScraper() {
	// Get matchups
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.MLB),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	// Get boxscore data
	eventdatascraper := foxsports.NewMLBPitchingBoxScoreScraper()
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(4),
		runner.EventDataRunnerScraper(
			eventdatascraper,
		),
	)

	boxScoreStats, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, statline := range boxScoreStats {
		jsonBytes, err := json.MarshalIndent(statline, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for foxsports.MLBProbableStartingPitcherScraper
func ExampleMLBProbableStartingPitcherScraper() {
	// Get matchups
	matchupScraper := foxsports.NewMatchupScraper(
		foxsports.MatchupScraperLeague(foxsports.MLB),
		foxsports.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-17"}),
	)

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	// Get starting pitcher data
	eventdatascraper := foxsports.NewMLBProbableStartingPitcherScraper()
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(4),
		runner.EventDataRunnerScraper(
			eventdatascraper,
		),
	)

	probablePitchers, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	for _, event := range probablePitchers {
		fmt.Printf("%#v\n", event.(model.MLBProbableStartingPitcher))
	}
}
