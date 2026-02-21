package foxsports

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

var pitchingHeaders []string = []string{"PITCHERS", "IP", "H", "R", "ER", "BB", "SO", "HR", "ERA"}

// MLBPitchingBoxScoreScraperOption defines a configuration option for the scraper
type MLBPitchingBoxScoreScraperOption func(*MLBPitchingBoxScoreScraper)

// MLBPitchingBoxScoreScraperParams sets the Params option
func MLBPitchingBoxScoreScraperParams(params map[string]string) MLBPitchingBoxScoreScraperOption {
	return func(s *MLBPitchingBoxScoreScraper) {
		s.Params = params
	}
}

// NewMLBPitchingBoxScoreScraper creates a new MLBPitchingBoxScoreScraper with the provided options
func NewMLBPitchingBoxScoreScraper(options ...MLBPitchingBoxScoreScraperOption) *MLBPitchingBoxScoreScraper {
	s := &MLBPitchingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.League = MLB
	s.Init()

	return s
}

type MLBPitchingBoxScoreScraper struct {
	EventDataScraper
}

func (s MLBPitchingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.FSMLBPitchingBoxScore
}

func (s *MLBPitchingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats] {
	start := time.Now().UTC()
	context := s.ConstructContext(matchup)

	var data []model.MLBPitchingBoxScoreStats
	// Construct event data URL
	log.Println("Constructing event data URL")
	url, err := s.ConstructEventDataURL(matchup.EventID)
	if err != nil {
		log.Println("Issue constructing event data URL")
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	// Fetch event data
	responseBody, err := s.FetchData(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	// Unmarshal JSON
	var responsePayload jsonresponse.MLBEventData
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}
	// Check for box score data
	if responsePayload.BoxScore == nil || responsePayload.BoxScore.BoxScoreSections == nil {
		log.Printf("No MLB pitching box score data available for event id: %d\n", matchup.EventID)
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Output: data, Context: context}
	}

	// Check that both Away and Home team box score stats are populated
	if responsePayload.BoxScore.BoxScoreSections.AwayStats == nil {
		log.Printf("No MLB pitching box score data available for away team (%s) for event id: %d\n", matchup.AwayTeamNameFull, matchup.EventID)
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Output: data, Context: context}
	}

	if responsePayload.BoxScore.BoxScoreSections.AwayStats == nil {
		log.Printf("No MLB pitching box score data available for home team (%s) for event id: %d\n", matchup.HomeTeamNameFull, matchup.EventID)
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Output: data, Context: context}
	}

	// validate MLBPitchingBoxScoreStats home and away positions
	uriSplit := strings.Split(responsePayload.BoxScore.BoxScoreSections.HomeStats.ContentURI, "/")
	actualHomeID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualHomeID != matchup.HomeTeamID {
		log.Printf("Home team ID, %d (%s), does not match expected, %d (%s)\n", actualHomeID, responsePayload.BoxScore.BoxScoreSections.HomeStats.Title, matchup.HomeTeamID, matchup.HomeTeamNameFull)
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}

	uriSplit = strings.Split(responsePayload.BoxScore.BoxScoreSections.AwayStats.ContentURI, "/")
	actualAwayID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualAwayID != matchup.AwayTeamID {
		log.Printf("Away team ID, %d (%s), does not match expected, %d (%s)\n", actualAwayID, responsePayload.BoxScore.BoxScoreSections.AwayStats.Title, matchup.AwayTeamID, matchup.AwayTeamNameFull)
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}

	// validate headers
	expectedHeadersSize := len(pitchingHeaders)

	// validate home pitching headers (index 2)
	actualHeaders := responsePayload.BoxScore.BoxScoreSections.HomeStats.BoxscoreItems[2].BoxscoreTable.Headers[0].Columns
	actualHeaderSize := len(actualHeaders)
	if actualHeaderSize != expectedHeadersSize {
		err = fmt.Errorf("home team pitching headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedHeadersSize)
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != pitchingHeaders[idx] {
			err = fmt.Errorf("home team pitching header '%s' unexpect at index %d. Expected %s", column.Text, idx, pitchingHeaders[idx])
			return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
		}
	}

	// validate away pitching headers (index 2)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayStats.BoxscoreItems[2].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedHeadersSize {
		err = fmt.Errorf("away team pitching headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedHeadersSize)
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != pitchingHeaders[idx] {
			err = fmt.Errorf("away team pitching header '%s' unexpect at index %d. Expected %s", column.Text, idx, pitchingHeaders[idx])
			return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
		}
	}
	stats, err := s.parsePitchingStats(responsePayload, context)
	if err != nil {
		return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Error: err, Context: context}
	}
	for _, obj := range stats {
		data = append(data, *obj)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchup.EventID, matchup.AwayTeamNameFull, matchup.HomeTeamNameFull, diff)
	return sportscrape.EventDataOutput[model.MLBPitchingBoxScoreStats]{Output: data, Context: context}
}

func (s *MLBPitchingBoxScoreScraper) parsePitchingStats(responsePayload jsonresponse.MLBEventData, context sportscrape.EventDataContext) ([]*model.MLBPitchingBoxScoreStats, error) {
	var stats []*model.MLBPitchingBoxScoreStats
	homeid := context.HomeID.(int64)
	awayid := context.AwayID.(int64)
	eventid := context.EventID.(int64)

	// Home
	for idx, record := range responsePayload.BoxScore.BoxScoreSections.HomeStats.BoxscoreItems[2].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return stats, err
		}

		statline := &model.MLBPitchingBoxScoreStats{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              eventid,
			TeamID:               homeid,
			Team:                 context.HomeTeam,
			OpponentID:           awayid,
			Opponent:             context.AwayTeam,
			PitchingOrder:        int32(idx + 1),
		}
		err = s.parseStatline(statline, record)
		if err != nil {
			return stats, err
		}
		stats = append(stats, statline)
	}

	// Away
	for idx, record := range responsePayload.BoxScore.BoxScoreSections.AwayStats.BoxscoreItems[2].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return stats, err
		}
		statline := &model.MLBPitchingBoxScoreStats{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              eventid,
			OpponentID:           homeid,
			Opponent:             context.HomeTeam,
			TeamID:               awayid,
			Team:                 context.AwayTeam,
			PitchingOrder:        int32(idx + 1),
		}
		err = s.parseStatline(statline, record)
		if err != nil {
			return stats, err
		}
		stats = append(stats, statline)
	}
	return stats, nil

}

func (s *MLBPitchingBoxScoreScraper) parseStatline(stats *model.MLBPitchingBoxScoreStats, statline jsonresponse.BoxScoreStatline) error {
	var err error
	stats.Player = statline.EntityLink.Player
	stats.Record = statline.Columns[0].Superscript

	// InningsPitched
	if statline.Columns[1].Text == "-" {
		stats.InningsPitched = 0
	} else {
		stats.InningsPitched, err = util.TextToFloat32(statline.Columns[1].Text)
		if err != nil {
			return err
		}
	}
	// HitsAllowed
	if statline.Columns[2].Text == "-" {
		stats.HitsAllowed = 0
	} else {
		stats.HitsAllowed, err = util.TextToInt32(statline.Columns[2].Text)
		if err != nil {
			return err
		}
	}
	// RunsAllowed
	if statline.Columns[3].Text == "-" {
		stats.RunsAllowed = 0
	} else {
		stats.RunsAllowed, err = util.TextToInt32(statline.Columns[3].Text)
		if err != nil {
			return err
		}
	}
	// EarnedRunsAllowed
	if statline.Columns[4].Text == "-" {
		stats.EarnedRunsAllowed = 0
	} else {
		stats.EarnedRunsAllowed, err = util.TextToInt32(statline.Columns[4].Text)
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
	// HomeRunsAllowed
	if statline.Columns[7].Text == "-" {
		stats.HomeRunsAllowed = 0
	} else {
		stats.HomeRunsAllowed, err = util.TextToInt32(statline.Columns[7].Text)
		if err != nil {
			return err
		}
	}
	// EarnedRunAverage
	if statline.Columns[8].Text == "-" {
		stats.EarnedRunAverage = 0
	} else {
		stats.EarnedRunAverage, err = util.TextToFloat32(statline.Columns[8].Text)
		if err != nil {
			return err
		}
	}
	return nil
}
