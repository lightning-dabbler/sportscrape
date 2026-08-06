package mlb_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/mlb"
	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/mlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for mlb.WeatherScraper
func ExampleWeatherScraper() {
	date := "2026-07-30"
	weatherscraper := mlb.NewWeatherScraper(
		mlb.WeatherScraperDate(date),
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
