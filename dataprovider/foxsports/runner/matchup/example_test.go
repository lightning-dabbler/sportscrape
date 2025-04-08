package matchup_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
)

// Example for matchup.GeneralMatchupRunner NBA
func ExampleGeneralMatchupRunner_nba() {
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.NBA),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
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

// Example for matchup.GeneralMatchupRunner MLB
func ExampleGeneralMatchupRunner_mlb() {
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.MLB),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2023-08-02"}),
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

// Example for matchup.GeneralMatchupRunner NCAAB
func ExampleGeneralMatchupRunner_ncaab() {
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.NCAAB),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2025-01-10"}),
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

// Example for matchup.GeneralMatchupRunner NFL
func ExampleGeneralMatchupRunner_nfl() {
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.NFL),
		matchup.GeneralMatchupSegmenter(&foxsports.NFLSegmenter{Year: 2024, Week: 4, Season: foxsports.POSTSEASON}),
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
