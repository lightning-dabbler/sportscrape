package runner_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner"
)

// Example for runner.GeneralMatchupRunner NBA
func ExampleGeneralMatchupRunner_nba() {
	matchupRunner := runner.NewGeneralMatchupRunner(
		runner.GeneralMatchupLeague(foxsports.NBA),
		runner.GeneralMatchupSegmenter(&foxsports.GeneralSementer{Date: "2023-01-10"}),
	)

	matchups := matchupRunner.GetMatchups()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for runner.GeneralMatchupRunner MLB
func ExampleGeneralMatchupRunner_mlb() {
	matchupRunner := runner.NewGeneralMatchupRunner(
		runner.GeneralMatchupLeague(foxsports.MLB),
		runner.GeneralMatchupSegmenter(&foxsports.GeneralSementer{Date: "2023-08-02"}),
	)

	matchups := matchupRunner.GetMatchups()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for runner.GeneralMatchupRunner NCAAB
func ExampleGeneralMatchupRunner_ncaab() {
	matchupRunner := runner.NewGeneralMatchupRunner(
		runner.GeneralMatchupLeague(foxsports.NCAAB),
		runner.GeneralMatchupSegmenter(&foxsports.GeneralSementer{Date: "2025-01-10"}),
	)

	matchups := matchupRunner.GetMatchups()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for runner.GeneralMatchupRunner NFL
func ExampleGeneralMatchupRunner_nfl() {
	matchupRunner := runner.NewGeneralMatchupRunner(
		runner.GeneralMatchupLeague(foxsports.NFL),
		runner.GeneralMatchupSegmenter(&foxsports.NFLSementer{Year: 2024, Week: 4, Season: foxsports.POSTSEASON}),
	)

	matchups := matchupRunner.GetMatchups()
	// Output each statline as pretty json
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
