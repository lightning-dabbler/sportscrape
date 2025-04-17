package eventdata_test

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/eventdata"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
)

// Example for eventdata.NBABoxScoreScraper
func ExampleNBABoxScoreScraper() {
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

// Example for eventdata.MLBBattingBoxScoreScraper
func ExampleMLBBattingBoxScoreScraper() {
	// Get matchups
	matchupRunner := matchup.NewGeneralMatchupRunner(
		matchup.GeneralMatchupLeague(foxsports.MLB),
		matchup.GeneralMatchupSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchups := matchupRunner.GetMatchups()

	// Get boxscore data
	scraper := eventdata.MLBBattingBoxScoreScraper{}
	scraper.League = foxsports.MLB
	runner := eventdata.NewRunner(
		eventdata.RunnerName("MLB Batting box score stats"),
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
