package mlb_test

import (
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreference/mlb/model"
)

// Example for mlb.MatchupRunner
func ExampleMatchupRunner() {
	date := "2024-10-30"
	// Instantiate MatchupRunner
	runner := mlb.NewMatchupRunner(
		mlb.WithMatchupTimeout(2 * time.Minute),
	)
	// Retrieve MLB matchups associated with date
	matchups := runner.GetMatchups(date)
	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup.(model.MLBMatchup))
	}
}
