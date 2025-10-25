package mma_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma"
	"github.com/lightning-dabbler/sportscrape/runner"
)
// Example for mma.ESPNMMAMatchupScraper
func ExampleESPNMMAMatchupScraper() {
	matchupRunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(
			mma.ESPNMMAMatchupScraper{Year: "2024"},
		),
	)

	matchups, err := matchupRunner.Run()
	if err != nil {
		panic(err)
	}
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
// Example for mma.ESPNMMAFightDetailsScraper
func ExampleESPNMMAFightDetailsScraper() {

	matchupRunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(
			mma.ESPNMMAMatchupScraper{Year: "2024"},
		),
	)

	result, err := matchupRunner.Run()
	if err != nil {
		panic(err)
	}

	eventRunner := runner.NewEventDataRunner(
		runner.EventDataRunnerScraper(mma.ESPNMMAFightDetailsScraper{}),
	)

	fightDetails, err := eventRunner.Run(result[0:10]...)
	if err != nil {
		panic(err)
	}
	for _, matchup := range fightDetails {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
