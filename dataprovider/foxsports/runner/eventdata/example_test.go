package eventdata_test

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/eventdata"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
)

// Example for eventdata.NBABoxScoreScraper
func ExampleNBABoxScoreScraper_nba() {
	// Get matchups
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.NBA),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
	)

	matchups := matchupRunner.GetMatchups()

	// Get boxscore data
	scraper := eventdata.NBABoxScoreScraper{}
	scraper.League = foxsports.NBA
	runner := eventdata.NewRunner(
		eventdata.RunnerName("NBA Box score stats"),
		eventdata.RunnerConcurrency(4),
		eventdata.RunnerScraper(
			&scraper,
		),
	)
	boxScoreStats := runner.RunEventsDataScraper(matchups...)
	for _, statline := range boxScoreStats {
		fmt.Printf("%#v\n", statline.(model.NBABoxScoreStats))
	}
}
