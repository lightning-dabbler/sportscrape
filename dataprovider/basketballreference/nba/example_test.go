package nba_test

import (
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
)

// Example for nba.MatchupRunner
func ExampleMatchupRunner() {
	date := "2025-02-20"
	// Instantiate MatchupRunner
	runner := nba.NewMatchupRunner(
		nba.WithMatchupTimeout(2 * time.Minute),
	)
	// Retrieve NBA matchups associated with date
	matchups := runner.GetMatchups(date)
	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup.(model.NBAMatchup))
	}
}

// Example for nba.BasicBoxScoreRunner
func ExampleBasicBoxScoreRunner() {
	date := "2025-02-19"
	// Instantiate MatchupRunner
	matchupRunner := nba.NewMatchupRunner(
		nba.WithMatchupTimeout(2 * time.Minute),
	)
	// Retrieve NBA matchups associated with date
	matchups := matchupRunner.GetMatchups(date)
	// Instantiate BasicBoxScoreRunner
	boxScoreRunner := nba.NewBasicBoxScoreRunner(
		nba.WithBasicBoxScoreTimeout(4*time.Minute),
		nba.WithBasicBoxScoreConcurrency(1),
	)
	// Retrieve NBA basic box score stats associated with matchups
	basicBoxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)

	for _, stats := range basicBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBABasicBoxScoreStats))
	}
}

// Example for nba.AdvBoxScoreRunner
func ExampleAdvBoxScoreRunner() {
	date := "2025-02-19"
	// Instantiate MatchupRunner
	matchupRunner := nba.NewMatchupRunner(
		nba.WithMatchupTimeout(2 * time.Minute),
	)
	// Retrieve NBA matchups associated with date
	matchups := matchupRunner.GetMatchups(date)
	// Instantiate AdvBoxScoreRunner
	boxScoreRunner := nba.NewAdvBoxScoreRunner(
		nba.WithAdvBoxScoreTimeout(4*time.Minute),
		nba.WithAdvBoxScoreConcurrency(1),
	)
	// Retrieve NBA advanced box score stats associated with matchups
	advBoxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)

	for _, stats := range advBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBAAdvBoxScoreStats))
	}
}
