package propfinder_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder"
	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for propfinder.WeatherScraper
func ExampleWeatherScraper() {
	date := "2026-07-30"
	weatherscraper := propfinder.NewWeatherScraper(
		propfinder.WeatherScraperDate(date),
	)
	weatherrunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Weather]{
			Scraper: weatherscraper,
		},
	)
	weather, err := weatherrunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each hourly weather reading as pretty json
	for _, w := range weather {
		jsonBytes, err := json.MarshalIndent(w, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
