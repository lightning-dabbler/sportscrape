package matchup_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
	matchuputil "github.com/lightning-dabbler/sportscrape/util/runner/matchup"
)

// Example for matchup.Runner NBA
func ExampleRunner_nba() {
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.NBA),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("NBA Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for matchup.Runner MLB
func ExampleRunner_mlb() {
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.MLB),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-08-02"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("MLB Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for matchup.Runner NCAAB
func ExampleRunner_ncaab() {
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.NCAAB),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2025-01-10"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("NCAAB Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for matchup.Runner NFL
func ExampleRunner_nfl() {
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.NFL),
		matchup.ScraperSegmenter(&foxsports.NFLSegmenter{Year: 2024, Week: 4, Season: foxsports.POSTSEASON}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("NCAAB Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
