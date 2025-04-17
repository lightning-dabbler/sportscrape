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
	"github.com/xitongsys/parquet-go/types"
)

var battingHeaders []string = []string{"BATTERS", "AB", "R", "H", "RBI", "BB", "SO", "LOB", "AVG"}

type MLBBattingBoxScoreScraper struct {
	EventDataScraper
}

func (s *MLBBattingBoxScoreScraper) Scrape(matchup interface{}) OutputWrapper {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	var context Context
	context.AwayTeam = matchupModel.AwayTeamNameFull
	context.AwayID = matchupModel.AwayTeamID
	context.HomeTeam = matchupModel.HomeTeamNameFull
	context.HomeID = matchupModel.HomeTeamID
	context.EventID = matchupModel.EventID
	context.EventTime = matchupModel.EventTime

	var data []interface{}
	// Construct event data URL
	log.Println("Constructing event data URL")
	url, err := s.ConstructEventDataURL(matchupModel.EventID)
	if err != nil {
		log.Println("Issue constructing event data URL")
		return OutputWrapper{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	// Fetch event data
	responseBody, err := s.FetchEventData(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return OutputWrapper{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	// Unmarshal JSON
	var responsePayload jsonresponse.MLBEventData
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	// Check for box score data
	if responsePayload.BoxScore == nil || responsePayload.BoxScore.BoxScoreSections == nil {
		log.Printf("No MLB batting box score data available for event id: %d\n", matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// Check that both Away and Home team box score stats are populated
	if responsePayload.BoxScore.BoxScoreSections.AwayStats == nil {
		log.Printf("No MLB batting box score data available for away team (%s) for event id: %d\n", matchupModel.AwayTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	if responsePayload.BoxScore.BoxScoreSections.AwayStats == nil {
		log.Printf("No MLB batting box score data available for home team (%s) for event id: %d\n", matchupModel.HomeTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// validate MLBBattingBoxScoreStats home and away positions
	uriSplit := strings.Split(responsePayload.BoxScore.BoxScoreSections.HomeStats.ContentURI, "/")
	actualHomeID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualHomeID != matchupModel.HomeTeamID {
		log.Printf("Home team ID, %d (%s), does not match expected, %d (%s)\n", actualHomeID, responsePayload.BoxScore.BoxScoreSections.HomeStats.Title, matchupModel.HomeTeamID, matchupModel.HomeTeamNameFull)
		return OutputWrapper{Error: err, Context: context}
	}

	uriSplit = strings.Split(responsePayload.BoxScore.BoxScoreSections.AwayStats.ContentURI, "/")
	actualAwayID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualAwayID != matchupModel.AwayTeamID {
		log.Printf("Away team ID, %d (%s), does not match expected, %d (%s)\n", actualAwayID, responsePayload.BoxScore.BoxScoreSections.AwayStats.Title, matchupModel.AwayTeamID, matchupModel.AwayTeamNameFull)
		return OutputWrapper{Error: err, Context: context}
	}

	// validate headers
	expectedHeadersSize := len(battingHeaders)

	// validate home batting headers (index 0)
	actualHeaders := responsePayload.BoxScore.BoxScoreSections.HomeStats.BoxscoreItems[0].BoxscoreTable.Headers[0].Columns
	actualHeaderSize := len(actualHeaders)
	if actualHeaderSize != expectedHeadersSize {
		err = fmt.Errorf("Home team starter headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedHeadersSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != battingHeaders[idx] {
			err = fmt.Errorf("Home team starter header '%s' unexpect at index %d. Expected %s", column.Text, idx, battingHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate away batting headers (index 0)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayStats.BoxscoreItems[0].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedHeadersSize {
		err = fmt.Errorf("Away team starter headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedHeadersSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != battingHeaders[idx] {
			err = fmt.Errorf("Away team starter header '%s' unexpect at index %d. Expected %s", column.Text, idx, battingHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}
	stats, err := s.parseBattingStats(responsePayload, context)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	for _, obj := range stats {
		data = append(data, *obj)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchupModel.EventID, matchupModel.AwayTeamNameFull, matchupModel.HomeTeamNameFull, diff)
	return OutputWrapper{Output: data, Context: context}
}

func (s *MLBBattingBoxScoreScraper) parseBattingStats(responsePayload jsonresponse.MLBEventData, context Context) ([]*model.MLBBattingBoxScoreStats, error) {
	var stats []*model.MLBBattingBoxScoreStats

	// Home
	for _, record := range responsePayload.BoxScore.BoxScoreSections.HomeStats.BoxscoreItems[0].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return stats, err
		}

		statline := &model.MLBBattingBoxScoreStats{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              context.EventID,
			TeamID:               context.HomeID,
			Team:                 context.HomeTeam,
			OpponentID:           context.AwayID,
			Opponent:             context.AwayTeam,
		}
		err = s.parseStatline(statline, record)
		if err != nil {
			return stats, err
		}
		stats = append(stats, statline)
	}

	// Away
	for _, record := range responsePayload.BoxScore.BoxScoreSections.AwayStats.BoxscoreItems[0].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return stats, err
		}
		statline := &model.MLBBattingBoxScoreStats{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              context.EventID,
			OpponentID:           context.HomeID,
			Opponent:             context.HomeTeam,
			TeamID:               context.AwayID,
			Team:                 context.AwayTeam,
		}
		err = s.parseStatline(statline, record)
		if err != nil {
			return stats, err
		}
		stats = append(stats, statline)
	}
	return stats, nil

}

func (s *MLBBattingBoxScoreScraper) parseStatline(stats *model.MLBBattingBoxScoreStats, statline jsonresponse.BoxScoreStatline) error {
	var err error
	stats.Player = statline.EntityLink.Player
	stats.Position = *statline.Columns[0].Superscript
	// AtBat
	if statline.Columns[1].Text == "-" {
		stats.AtBat = 0
	} else {
		stats.AtBat, err = util.TextToInt32(statline.Columns[1].Text)
		if err != nil {
			return err
		}
	}
	// Runs
	if statline.Columns[2].Text == "-" {
		stats.Runs = 0
	} else {
		stats.Runs, err = util.TextToInt32(statline.Columns[2].Text)
		if err != nil {
			return err
		}
	}
	// Hits
	if statline.Columns[3].Text == "-" {
		stats.Hits = 0
	} else {
		stats.Hits, err = util.TextToInt32(statline.Columns[3].Text)
		if err != nil {
			return err
		}
	}
	// RunsBattedIn
	if statline.Columns[4].Text == "-" {
		stats.RunsBattedIn = 0
	} else {
		stats.RunsBattedIn, err = util.TextToInt32(statline.Columns[4].Text)
		if err != nil {
			return err
		}
	}
	// Walks
	if statline.Columns[5].Text == "-" {
		stats.Walks = 0
	} else {
		stats.Walks, err = util.TextToInt32(statline.Columns[5].Text)
		if err != nil {
			return err
		}
	}
	// Strikeouts
	if statline.Columns[6].Text == "-" {
		stats.Strikeouts = 0
	} else {
		stats.Strikeouts, err = util.TextToInt32(statline.Columns[6].Text)
		if err != nil {
			return err
		}
	}
	// LeftOnBase
	if statline.Columns[7].Text == "-" {
		stats.LeftOnBase = 0
	} else {
		stats.LeftOnBase, err = util.TextToInt32(statline.Columns[7].Text)
		if err != nil {
			return err
		}
	}
	// BattingAverage
	if statline.Columns[8].Text == "-" {
		stats.BattingAverage = 0
	} else {
		stats.BattingAverage, err = util.TextToFloat32(statline.Columns[8].Text)
		if err != nil {
			return err
		}
	}
	return nil
}
