package eventdata

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
)

var NBABoxscoreStatsHeaders map[string][]string = map[string][]string{
	"starters": {"STARTERS", "MIN", "OFF", "DEF", "REB", "AST", "STL", "BLK", "TO", "PF", "PTS"},
	"bench":    {"BENCH", "MIN", "OFF", "DEF", "REB", "AST", "STL", "BLK", "TO", "PF", "PTS"},
	"shooting": {"SHOOTING", "FG", "3FG", "FT", "PTS"},
}

type NBABoxScoreScraper struct {
	EventDataScraper
}

func (s *NBABoxScoreScraper) Scrape(matchup interface{}) OutputWrapper {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	var context Context
	context.AwayTeam = matchupModel.AwayTeamNameFull
	context.HomeTeam = matchupModel.HomeTeamNameFull
	context.EventID = matchupModel.EventID
	var data []interface{}
	// Construct event data URL
	log.Println("Constructing event data URL")
	url, err := s.ConstructEventDataURL(matchupModel.EventID)
	if err != nil {
		log.Println("Issue constructing event data URL")
		return OutputWrapper{Error: err, Context: context}
	}
	context.URL = url
	// Fetch event data
	responseBody, err := s.FetchEventData(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return OutputWrapper{Error: err, Context: context}
	}
	// Unmarshal JSON
	var responsePayload jsonresponse.NBAEventData
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	// Check for box score data
	if responsePayload.BoxScore == nil || responsePayload.BoxScore.BoxScoreSections == nil {
		log.Printf("No NBA box score data available for event id: %d\n", matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// Check that both Away and Home team box score stats are populated
	if responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats == nil {
		log.Printf("No NBA box score data available for away team (%s) for event id: %d\n", matchupModel.AwayTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	if responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats == nil {
		log.Printf("No NBA box score data available for home team (%s) for event id: %d\n", matchupModel.HomeTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// validate NBABoxScoreStats home and away positions
	uriSplit := strings.Split(responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.ContentURI, "/")
	actualHomeID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualHomeID != matchupModel.HomeTeamID {
		log.Printf("Home team ID, %d (%s), does not match expected, %d (%s)\n", actualHomeID, responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.Title, matchupModel.HomeTeamID, matchupModel.HomeTeamNameFull)
		return OutputWrapper{Error: err, Context: context}
	}

	uriSplit = strings.Split(responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.ContentURI, "/")
	actualAwayID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualAwayID != matchupModel.AwayTeamID {
		log.Printf("Away team ID, %d (%s), does not match expected, %d (%s)\n", actualAwayID, responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.Title, matchupModel.AwayTeamID, matchupModel.AwayTeamNameFull)
		return OutputWrapper{Error: err, Context: context}
	}

	// validate headers
	expectedStartersHeaders := NBABoxscoreStatsHeaders["starters"]
	expectedStarterHeaderSize := len(expectedStartersHeaders)
	expectedBenchHeaders := NBABoxscoreStatsHeaders["bench"]
	expectedBenchHeaderSize := len(expectedBenchHeaders)
	expectedShootingHeaders := NBABoxscoreStatsHeaders["shooting"]
	expectedShootingHeaderSize := len(expectedShootingHeaders)

	// validate home starter headers (index 0)
	actualHeaders := responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[0].BoxscoreTable.Headers[0].Columns
	actualHeaderSize := len(actualHeaders)
	if actualHeaderSize != expectedStarterHeaderSize {
		err = fmt.Errorf("Home team starter headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedStarterHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedStartersHeaders[idx] {
			err = fmt.Errorf("Home team starter header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedStartersHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate home bench headers (index 1)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[1].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedBenchHeaderSize {
		err = fmt.Errorf("Home team bench headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedBenchHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedBenchHeaders[idx] {
			err = fmt.Errorf("Home team bench header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedBenchHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate home shooting headers (index 3)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[3].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedShootingHeaderSize {
		err = fmt.Errorf("Home team shooting headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedShootingHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedShootingHeaders[idx] {
			err = fmt.Errorf("Home team shooting header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedShootingHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate away starter headers (index 0)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[0].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedStarterHeaderSize {
		err = fmt.Errorf("Away team starter headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedStarterHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedStartersHeaders[idx] {
			err = fmt.Errorf("Away team starter header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedStartersHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate away bench headers (index 1)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[1].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedBenchHeaderSize {
		err = fmt.Errorf("Away team bench headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedBenchHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedBenchHeaders[idx] {
			err = fmt.Errorf("Away team bench header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedBenchHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate away shooting headers (index 3)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[3].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedShootingHeaderSize {
		err = fmt.Errorf("Away team shooting headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedShootingHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedShootingHeaders[idx] {
			err = fmt.Errorf("Away team shooting header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedShootingHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// parse and develop playerMap of relevant statlines
	playerMap, err := s.parseBoxScoreStats(responsePayload)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}

	// Allocate each statline to data output
	for _, obj := range playerMap {
		data = append(data, obj)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchupModel.EventID, matchupModel.AwayTeamNameFull, matchupModel.HomeTeamNameFull, diff)
	return OutputWrapper{Output: data, Context: context}
}

func (s *NBABoxScoreScraper) parseBoxScoreStats(responsePayload jsonresponse.NBAEventData) (map[int64]model.NBABoxScoreStats, error) {
	var playerMap map[int64]model.NBABoxScoreStats
	return playerMap, nil

}
