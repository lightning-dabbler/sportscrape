package baseballsavantmlb_test

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

// Example for baseballsavantmlb.MatchupScraper
func ExampleMatchupScraper() {
	date := "2025-06-25"
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate(date),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
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

// Example for baseballsavantmlb.FieldingBoxScoreScraper
func ExampleFieldingBoxScoreScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := baseballsavantmlb.NewFieldingBoxScoreScraper()
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.FieldingBoxScore]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)
	stats, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each statline as pretty json
	for _, statline := range stats {
		fmt.Printf("%#v\n", statline)
	}
}

// Example for baseballsavantmlb.BattingBoxScoreScraper
func ExampleBattingBoxScoreScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := baseballsavantmlb.NewBattingBoxScoreScraper()
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BattingBoxScore]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	stats, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each statline as pretty json
	for _, statline := range stats {
		fmt.Printf("%#v\n", statline)
	}

}

// Example for baseballsavantmlb.PitchingBoxScoreScraper
func ExamplePitchingBoxScoreScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	boxscorescraper := baseballsavantmlb.NewPitchingBoxScoreScraper()
	boxscorerunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PitchingBoxScore]{
			Scraper:     boxscorescraper,
			Concurrency: 1,
		},
	)

	stats, err := boxscorerunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each statline as pretty json
	for _, statline := range stats {
		fmt.Printf("%#v\n", statline)
	}
}

// Example for baseballsavantmlb.PlayByPlayScraper
func ExamplePlayByPlayScraper() {
	matchupscraper := baseballsavantmlb.NewMatchupScraper(
		baseballsavantmlb.MatchupScraperDate("2024-10-30"),
	)
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)
	matchups, err := matchuprunner.Run()
	if err != nil {
		panic(err)
	}

	playbyplayscraper := baseballsavantmlb.NewPlayByPlayScraper()
	playbyplayrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Scraper:     playbyplayscraper,
			Concurrency: 1,
		},
	)

	plays, err := playbyplayrunner.Run(matchups)
	if err != nil {
		panic(err)
	}

	// Output each play as pretty json
	for _, play := range plays {
		fmt.Printf("%#v\n", play)
	}

}
