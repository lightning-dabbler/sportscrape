package nba_test

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreference/nba/model"
)

// Example for nba.GetMatchups
func ExampleGetMatchups() {
	date := "2025-02-20"
	matchups := nba.GetMatchups(date)
	for _, matchup := range matchups {
		fmt.Printf("%#v\n", matchup.(model.NBAMatchup))
	}
}

// Example for nba.GetBasicBoxScoreStats
func ExampleGetBasicBoxScoreStats() {
	date := "2025-02-19"
	matchups := nba.GetMatchups(date)
	basicBoxScoreStats := nba.GetBasicBoxScoreStats(1, matchups...)

	for _, stats := range basicBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBABasicBoxScoreStats))
	}
}

// Example for nba.GetAdvBoxScoreStats
func ExampleGetAdvBoxScoreStats() {
	date := "2025-02-19"
	matchups := nba.GetMatchups(date)
	advBoxScoreStats := nba.GetAdvBoxScoreStats(1, matchups...)

	for _, stats := range advBoxScoreStats {
		fmt.Printf("%#v\n", stats.(model.NBAAdvBoxScoreStats))
	}
}
