package nba_test

import (
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for nba.MatchupScraper
func ExampleMatchupScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup)
	}
}

// Example for nba.MatchupPeriodsScraper
func ExampleMatchupPeriodsScraper() {
	matchupScraper := nba.NewMatchupPeriodsScraper(
		nba.WithMatchupPeriodsDate("2025-06-05"),
		nba.WithMatchupPeriodsTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.MatchupPeriods]{
			Scraper: matchupScraper,
		},
	)

	records, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, period := range records {
		fmt.Printf("%#v\n", period)
	}
}

// Example for nba.PlayByPlayScraper
func ExamplePlayByPlayScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	playbyplayscraper := nba.NewPlayByPlayScraper(
		nba.WithPlayByPlayTimeout(2 * time.Minute),
	)

	playbyplayrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper:     playbyplayscraper,
			Concurrency: 1,
		},
	)

	records, err := playbyplayrunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreUsageScraper full
func ExampleBoxScoreUsageScraper_full() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreUsageScraper(
		nba.WithBoxScoreUsageTimeout(2*time.Minute),
		nba.WithBoxScoreUsagePeriod(nba.Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreUsage]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreUsageScraper h2
func ExampleBoxScoreUsageScraper_h2() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreUsageScraper(
		nba.WithBoxScoreUsageTimeout(2*time.Minute),
		nba.WithBoxScoreUsagePeriod(nba.H2),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreUsage]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreTraditionalScraper q1
func ExampleBoxScoreTraditionalScraper_q1() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreTraditionalScraper(
		nba.WithBoxScoreTraditionalTimeout(2*time.Minute),
		nba.WithBoxScoreTraditionalPeriod(nba.Q1),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreTraditional]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreAdvancedScraper full
func ExampleBoxScoreAdvancedScraper_full() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-11"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreAdvancedScraper(
		nba.WithBoxScoreAdvancedTimeout(2*time.Minute),
		nba.WithBoxScoreAdvancedPeriod(nba.Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreAdvanced]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreScoringScraper h1
func ExampleBoxScoreScoringScraper_h1() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreScoringScraper(
		nba.WithBoxScoreScoringTimeout(2*time.Minute),
		nba.WithBoxScoreScoringPeriod(nba.H1),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreScoring]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreMiscScraper full
func ExampleBoxScoreMiscScraper_full() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreMiscScraper(
		nba.WithBoxScoreMiscTimeout(2*time.Minute),
		nba.WithBoxScoreMiscPeriod(nba.Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreMisc]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreFourFactorsScraper full
func ExampleBoxScoreFourFactorsScraper_full() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreFourFactorsScraper(
		nba.WithBoxScoreFourFactorsTimeout(2*time.Minute),
		nba.WithBoxScoreFourFactorsPeriod(nba.Full),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreFourFactors]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreLiveScraper
func ExampleBoxScoreLiveScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-12-10"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreLiveScraper(
		nba.WithBoxScoreLiveTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreLive]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreTrackingScraper
func ExampleBoxScoreTrackingScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreTrackingScraper(
		nba.WithBoxScoreTrackingTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreTracking]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreMatchupsScraper
func ExampleBoxScoreMatchupsScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreMatchupsScraper(
		nba.WithBoxScoreMatchupsTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreMatchups]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreDefenseScraper
func ExampleBoxScoreDefenseScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreDefenseScraper(
		nba.WithBoxScoreDefenseTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreDefense]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}

// Example for nba.BoxScoreHustleScraper
func ExampleBoxScoreHustleScraper() {
	matchupScraper := nba.NewMatchupScraper(
		nba.WithMatchupDate("2025-06-05"),
		nba.WithMatchupTimeout(2*time.Minute),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupScraper,
		},
	)

	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := nba.NewBoxScoreHustleScraper(
		nba.WithBoxScoreHustleTimeout(2 * time.Minute),
	)

	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreHustle]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	records, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}
	// Output each statline as pretty json
	for _, record := range records {
		fmt.Printf("%#v\n", record)
	}
}
