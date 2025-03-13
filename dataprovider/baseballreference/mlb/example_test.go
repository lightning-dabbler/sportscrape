package mlb_test

import (
	"fmt"
	"log"
	"time"

	"encoding/json"

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

// Example for mlb.BattingBoxScoreRunner
func ExampleBattingBoxScoreRunner() {
	date := "2024-10-13"
	// Instantiate MatchupRunner
	matchRunner := mlb.NewMatchupRunner(
		mlb.WithMatchupTimeout(2 * time.Minute),
	)
	// Retrieve MLB matchups associated with date
	matchups := matchRunner.GetMatchups(date)
	// Instantiate BattingBoxScoreRunner
	boxScoreRunner := mlb.NewBattingBoxScoreRunner(
		mlb.WithBattingBoxScoreTimeout(4*time.Minute),
		mlb.WithBattingBoxScoreConcurrency(1),
	)
	// Retrieve MLB batting box score stats associated with matchups
	boxScoreStats := boxScoreRunner.GetBoxScoresStats(matchups...)
	// Output each statline as pretty json
	for _, stats := range boxScoreStats {
		jsonBytes, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}
