package mma_test

import (
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

// Example for mma.ESPNMMAMatchupScraper
func ExampleESPNMMAMatchupScraper() {
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: mma.ESPNMMAMatchupScraper{Year: "2024", League: "ufc", BaseScraper: scraper.BaseScraper{Timeout: time.Second * 10}},
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup)
	}
}

// Example for mma.ESPNMMAFightDetailsScraper
func ExampleESPNMMAFightDetailsScraper() {
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: mma.ESPNMMAMatchupScraper{Year: "2024", League: "ufc", BaseScraper: scraper.BaseScraper{Timeout: time.Second * 10}},
		},
	)

	result, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	eventRunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.FightDetails]{
			Scraper: mma.ESPNMMAFightDetailsScraper{League: "ufc", BaseScraper: scraper.BaseScraper{Timeout: 10 * time.Second}},
		},
	)

	fightDetails, err := eventRunner.Run(result[0:10])
	if err != nil {
		panic(err)
	}
	for _, fight := range fightDetails {
		fmt.Printf("%#v\n", fight)
	}
}
