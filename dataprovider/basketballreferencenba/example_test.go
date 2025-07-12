package basketballreferencenba_test

import (
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for nba.MatchupRunner
func ExampleMatchupRunner() {
	date := "2025-02-20"
	// Instantiate MatchupRunner
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	runner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupscraper),
	)
	// Retrieve NBA matchups associated with date
	matchups, err := runner.Run()
	if err != nil {
		panic(err)
	}

	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup.(model.NBAMatchup))
	}
}

// Example for basketballreferencenba.BasicBoxScoreScraper Full basic box score stats
func ExampleBasicBoxScoreScraper_full() {
	date := "2025-02-19"
	// Instantiate MatchupRunner
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupscraper),
	)
	// Retrieve NBA matchups associated with date
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Instantiate BasicBoxScoreScraper
	boxscorescraper := basketballreferencenba.NewBasicBoxScoreScraper(
		basketballreferencenba.WithBasicBoxScoreTimeout(4*time.Minute),
		basketballreferencenba.WithBasicBoxScorePeriod(basketballreferencenba.Full),
	)
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(1),
		runner.EventDataRunnerScraper(boxscorescraper),
	)
	// Retrieve NBA basic box score stats associated with matchups
	basicBoxScoreStats, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	for _, stats := range basicBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBABasicBoxScoreStats))
	}
}

// Example for basketballreferencenba.BasicBoxScoreScraper Q2 basic box score stats
func ExampleBasicBoxScoreScraper_q2() {
	date := "2025-02-19"
	// Instantiate MatchupRunner
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupscraper),
	)
	// Retrieve NBA matchups associated with date
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Instantiate BasicBoxScoreScraper
	boxscorescraper := basketballreferencenba.NewBasicBoxScoreScraper(
		basketballreferencenba.WithBasicBoxScoreTimeout(4*time.Minute),
		basketballreferencenba.WithBasicBoxScorePeriod(basketballreferencenba.Q2),
	)
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(1),
		runner.EventDataRunnerScraper(boxscorescraper),
	)
	// Retrieve NBA basic box score stats associated with matchups
	basicBoxScoreStats, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	for _, stats := range basicBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBABasicBoxScoreStats))
	}
}

// Example for basketballreferencenba.BasicBoxScoreScraper H2 basic box score stats
func ExampleBasicBoxScoreScraper_h2() {
	date := "2025-02-19"
	// Instantiate MatchupRunner
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupscraper),
	)
	// Retrieve NBA matchups associated with date
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Instantiate BasicBoxScoreScraper
	boxscorescraper := basketballreferencenba.NewBasicBoxScoreScraper(
		basketballreferencenba.WithBasicBoxScoreTimeout(4*time.Minute),
		basketballreferencenba.WithBasicBoxScorePeriod(basketballreferencenba.H2),
	)
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(1),
		runner.EventDataRunnerScraper(boxscorescraper),
	)
	// Retrieve NBA basic box score stats associated with matchups
	basicBoxScoreStats, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	for _, stats := range basicBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBABasicBoxScoreStats))
	}
}

// Example for basketballreferencenba.AdvBoxScoreScraper
func ExampleAdvBoxScoreScraper() {
	date := "2025-02-19"
	// Instantiate MatchupRunner
	matchupscraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(date),
		basketballreferencenba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerScraper(matchupscraper),
	)
	// Retrieve NBA matchups associated with date
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Instantiate AdvBoxScoreScraper
	boxscorescraper := basketballreferencenba.NewAdvBoxScoreScraper(
		basketballreferencenba.WithAdvBoxScoreTimeout(4 * time.Minute),
	)
	runner := runner.NewEventDataRunner(
		runner.EventDataRunnerConcurrency(1),
		runner.EventDataRunnerScraper(boxscorescraper),
	)
	// Retrieve NBA basic box score stats associated with matchups
	advBoxScoreStats, err := runner.Run(matchups...)
	if err != nil {
		panic(err)
	}
	for _, stats := range advBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBAAdvBoxScoreStats))
	}
}
