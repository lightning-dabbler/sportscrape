package scraper_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/scraper"
)

// Example for scraper.MatchupScraper NBA
func ExampleMatchupScraper_nba() {
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.NBA),
		scraper.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
	)
	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
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

// Example for scraper.MatchupScraper MLB
func ExampleMatchupScraper_mlb() {
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.MLB),
		scraper.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-08-02"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
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

// Example for scraper.MatchupScraper NCAAB
func ExampleMatchupScraper_ncaab() {
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.NCAAB),
		scraper.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2025-01-10"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
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

// Example for scraper.MatchupScraper NFL
func ExampleMatchupScraper_nfl() {
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.NFL),
		scraper.MatchupScraperSegmenter(&foxsports.NFLSegmenter{Year: 2024, Week: 4, Season: foxsports.POSTSEASON}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
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

// Example for scraper.NBABoxScoreScraper
func ExampleNBABoxScoreScraper() {
	// Get matchups
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.NBA),
		scraper.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
	if err != nil {
		panic(err)
	}

	// Get boxscore data
	eventdatascraper := scraper.NBABoxScoreScraper{}
	eventdatascraper.League = foxsports.NBA
	runner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerConcurrency(4),
		sportscrape.EventDataRunnerScraper(
			&eventdatascraper,
		),
	)
	boxScoreStats, err := runner.RunEventsDataScraper(matchups...)
	if err != nil {
		panic(err)
	}
	for _, statline := range boxScoreStats {
		fmt.Printf("%#v\n", statline.(model.NBABoxScoreStats))
	}
}

// Example for scraper.MLBBattingBoxScoreScraper
func ExampleMLBBattingBoxScoreScraper() {
	// Get matchups
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.MLB),
		scraper.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
	if err != nil {
		panic(err)
	}

	// Get boxscore data
	eventdatascraper := scraper.MLBBattingBoxScoreScraper{}
	eventdatascraper.League = foxsports.MLB
	runner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerConcurrency(4),
		sportscrape.EventDataRunnerScraper(
			&eventdatascraper,
		),
	)

	boxScoreStats, err := runner.RunEventsDataScraper(matchups...)
	if err != nil {
		panic(err)
	}
	for _, statline := range boxScoreStats {
		fmt.Printf("%#v\n", statline.(model.MLBBattingBoxScoreStats))
	}
}

// Example for scraper.MLBPitchingBoxScoreScraper
func ExampleMLBPitchingBoxScoreScraper() {
	// Get matchups
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.MLB),
		scraper.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
	if err != nil {
		panic(err)
	}

	// Get boxscore data
	eventdatascraper := scraper.MLBPitchingBoxScoreScraper{}
	eventdatascraper.League = foxsports.MLB
	runner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerConcurrency(4),
		sportscrape.EventDataRunnerScraper(
			&eventdatascraper,
		),
	)

	boxScoreStats, err := runner.RunEventsDataScraper(matchups...)
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

// Example for scraper.MLBProbableStartingPitcherScraper
func ExampleMLBProbableStartingPitcherScraper() {
	// Get matchups
	matchupScraper := scraper.NewMatchupScraper(
		scraper.MatchupScraperLeague(foxsports.MLB),
		scraper.MatchupScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-17"}),
	)

	matchuprunner := sportscrape.NewMatchupRunner(
		sportscrape.MatchupRunnerScraper(matchupScraper),
	)

	matchups, err := matchuprunner.RunMatchupsScraper()
	if err != nil {
		panic(err)
	}

	// Get starting pitcher data
	eventdatascraper := scraper.MLBProbableStartingPitcherScraper{}
	eventdatascraper.League = foxsports.MLB
	runner := sportscrape.NewEventDataRunner(
		sportscrape.EventDataRunnerConcurrency(4),
		sportscrape.EventDataRunnerScraper(
			&eventdatascraper,
		),
	)

	probablePitchers, err := runner.RunEventsDataScraper(matchups...)
	if err != nil {
		panic(err)
	}
	for _, event := range probablePitchers {
		fmt.Printf("%#v\n", event.(model.MLBProbableStartingPitcher))
	}
}
