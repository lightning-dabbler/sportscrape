package eventdata_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/eventdata"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/runner/matchup"
	eventdatautil "github.com/lightning-dabbler/sportscrape/util/runner/eventdata"
	matchuputil "github.com/lightning-dabbler/sportscrape/util/runner/matchup"
)

// Example for eventdata.NBABoxScoreScraper
func ExampleNBABoxScoreScraper() {
	// Get matchups
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.NBA),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2023-01-10"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("NBA Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()

	// Get boxscore data
	scraper := eventdata.NBABoxScoreScraper{}
	scraper.League = foxsports.NBA
	runner := eventdatautil.NewRunner(
		eventdatautil.RunnerName("NBA Box score stats"),
		eventdatautil.RunnerConcurrency(4),
		eventdatautil.RunnerScraper(
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
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.MLB),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("MLB Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()

	// Get boxscore data
	scraper := eventdata.MLBBattingBoxScoreScraper{}
	scraper.League = foxsports.MLB
	runner := eventdatautil.NewRunner(
		eventdatautil.RunnerName("MLB Batting box score stats"),
		eventdatautil.RunnerConcurrency(4),
		eventdatautil.RunnerScraper(
			&scraper,
		),
	)
	boxScoreStats := runner.RunEventsDataScraper(matchups...)
	for _, statline := range boxScoreStats {
		fmt.Printf("%#v\n", statline.(model.MLBBattingBoxScoreStats))
	}
}

// Example for eventdata.MLBPitchingBoxScoreScraper
func ExampleMLBPitchingBoxScoreScraper() {
	// Get matchups
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.MLB),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-30"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("MLB Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()

	// Get boxscore data
	scraper := eventdata.MLBPitchingBoxScoreScraper{}
	scraper.League = foxsports.MLB
	runner := eventdatautil.NewRunner(
		eventdatautil.RunnerName("MLB Pitching box score stats"),
		eventdatautil.RunnerConcurrency(4),
		eventdatautil.RunnerScraper(
			&scraper,
		),
	)
	boxScoreStats := runner.RunEventsDataScraper(matchups...)
	// Output each statline as pretty json
	for _, statline := range boxScoreStats {
		jsonBytes, err := json.MarshalIndent(statline, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v\n", err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// Example for eventdata.MLBProbableStartingPitcherScraper
func ExampleMLBProbableStartingPitcherScraper() {
	// Get matchups
	matchupScraper := matchup.NewScraper(
		matchup.ScraperLeague(foxsports.MLB),
		matchup.ScraperSegmenter(&foxsports.GeneralSegmenter{Date: "2024-10-17"}),
	)

	matchuprunner := matchuputil.NewRunner(
		matchuputil.RunnerName("MLB Matchups"),
		matchuputil.RunnerScraper(matchupScraper),
	)

	matchups := matchuprunner.RunMatchupsScraper()

	// Get starting pitcher data
	scraper := eventdata.MLBProbableStartingPitcherScraper{}
	scraper.League = foxsports.MLB
	runner := eventdatautil.NewRunner(
		eventdatautil.RunnerName("MLB probable starting pitchers"),
		eventdatautil.RunnerConcurrency(4),
		eventdatautil.RunnerScraper(
			&scraper,
		),
	)
	probablePitchers := runner.RunEventsDataScraper(matchups...)
	for _, event := range probablePitchers {
		fmt.Printf("%#v\n", event.(model.MLBProbableStartingPitcher))
	}
}
