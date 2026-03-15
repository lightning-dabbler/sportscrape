package basketballreferencenba_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Deprecated: basketball-reference.com provider is deprecated.
// Example for basketballreferencenba.MatchupRunner
func ExampleMatchupRunner() {
	date := "2025-02-20"
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchupscraper.NetworkHeaders = basketballreferencenba.NetworkHeaders
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.NBAMatchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		log.Println(err)
		return
	}
	for _, matchup := range matchups {
		jsonBytes, err := json.MarshalIndent(matchup, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Deprecated: basketball-reference.com provider is deprecated.
// Example for basketballreferencenba.BasicBoxScoreScraper Full basic box score stats
func ExampleBasicBoxScoreScraper_full() {
	date := "2025-02-19"
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchupscraper.NetworkHeaders = basketballreferencenba.NetworkHeaders
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.NBAMatchup]{
			Scraper:   matchupscraper,
			KeepAlive: true,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		log.Println(err)
		return
	}
	boxscorescraper := basketballreferencenba.NewBasicBoxScoreScraper(
		basketballreferencenba.WithBasicBoxScoreTimeout(4*time.Minute),
		basketballreferencenba.WithBasicBoxScorePeriod(basketballreferencenba.Full),
	)
	boxscorescraper.DocumentRetriever = matchupscraper.DocumentRetriever
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.NBAMatchup, model.NBABasicBoxScoreStats]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)
	basicBoxScoreStats, err := boxscorerunner.Run(matchups)
	if err != nil {
		log.Println(err)
		return
	}
	for _, stats := range basicBoxScoreStats {
		jsonBytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Deprecated: basketball-reference.com provider is deprecated.
// Example for basketballreferencenba.BasicBoxScoreScraper Q2 basic box score stats
func ExampleBasicBoxScoreScraper_q2() {
	date := "2025-02-19"
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchupscraper.NetworkHeaders = basketballreferencenba.NetworkHeaders
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.NBAMatchup]{
			Scraper:   matchupscraper,
			KeepAlive: true,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		log.Println(err)
		return
	}
	boxscorescraper := basketballreferencenba.NewBasicBoxScoreScraper(
		basketballreferencenba.WithBasicBoxScoreTimeout(4*time.Minute),
		basketballreferencenba.WithBasicBoxScorePeriod(basketballreferencenba.Q2),
	)
	boxscorescraper.DocumentRetriever = matchupscraper.DocumentRetriever
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.NBAMatchup, model.NBABasicBoxScoreStats]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)
	basicBoxScoreStats, err := boxscorerunner.Run(matchups)
	if err != nil {
		log.Println(err)
		return
	}
	for _, stats := range basicBoxScoreStats {
		jsonBytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Deprecated: basketball-reference.com provider is deprecated.
// Example for basketballreferencenba.BasicBoxScoreScraper H2 basic box score stats
func ExampleBasicBoxScoreScraper_h2() {
	date := "2025-02-19"
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchupscraper.NetworkHeaders = basketballreferencenba.NetworkHeaders
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.NBAMatchup]{
			Scraper:   matchupscraper,
			KeepAlive: true,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		log.Println(err)
		return
	}
	boxscorescraper := basketballreferencenba.NewBasicBoxScoreScraper(
		basketballreferencenba.WithBasicBoxScoreTimeout(4*time.Minute),
		basketballreferencenba.WithBasicBoxScorePeriod(basketballreferencenba.H2),
	)
	boxscorescraper.DocumentRetriever = matchupscraper.DocumentRetriever
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.NBAMatchup, model.NBABasicBoxScoreStats]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)
	basicBoxScoreStats, err := boxscorerunner.Run(matchups)
	if err != nil {
		log.Println(err)
		return
	}
	for _, stats := range basicBoxScoreStats {
		jsonBytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Deprecated: basketball-reference.com provider is deprecated.
// Example for basketballreferencenba.AdvBoxScoreScraper
func ExampleAdvBoxScoreScraper() {
	date := "2025-02-19"
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchupscraper.NetworkHeaders = basketballreferencenba.NetworkHeaders
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.NBAMatchup]{
			Scraper:   matchupscraper,
			KeepAlive: true,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		log.Println(err)
		return
	}
	boxscorescraper := basketballreferencenba.NewAdvBoxScoreScraper(
		basketballreferencenba.WithAdvBoxScoreTimeout(4 * time.Minute),
	)
	boxscorescraper.DocumentRetriever = matchupscraper.DocumentRetriever
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.NBAMatchup, model.NBAAdvBoxScoreStats]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)
	advBoxScoreStats, err := boxscorerunner.Run(matchups)
	if err != nil {
		log.Println(err)
		return
	}
	for _, stats := range advBoxScoreStats {
		jsonBytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
