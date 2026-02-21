package foxsports

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

const (
	oddsMoneyLineTitle = "TEAM TO WIN"
)

// MLBOddsMoneyLineScraperOption defines a configuration option for the scraper
type MLBOddsMoneyLineScraperOption func(*MLBOddsMoneyLineScraper)

// MLBOddsMoneyLineScraperParams sets the Params option
func MLBOddsMoneyLineScraperParams(params map[string]string) MLBOddsMoneyLineScraperOption {
	return func(s *MLBOddsMoneyLineScraper) {
		s.Params = params
	}
}

// NewMLBOddsMoneyLineScraper creates a new MLBOddsMoneyLineScraper with the provided options
func NewMLBOddsMoneyLineScraper(options ...MLBOddsMoneyLineScraperOption) *MLBOddsMoneyLineScraper {
	s := &MLBOddsMoneyLineScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.League = MLB
	s.Init()

	return s
}

type MLBOddsMoneyLineScraper struct {
	EventDataScraper
}

func (s MLBOddsMoneyLineScraper) Feed() sportscrape.Feed {
	return sportscrape.FSMLBOddsMoneyLine
}

func (s *MLBOddsMoneyLineScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.MLBOddsMoneyLine] {
	start := time.Now().UTC()
	context := s.ConstructContext(matchup)

	var data []model.MLBOddsMoneyLine
	// Construct event data URL
	log.Println("Constructing event data URL")
	url, err := s.ConstructMatchupComparisonURL(matchup.EventID)
	if err != nil {
		log.Println("Issue constructing matchup comparison URL")
		return sportscrape.EventDataOutput[model.MLBOddsMoneyLine]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	// Fetch event data
	responseBody, err := s.FetchData(url)
	if err != nil {
		log.Println("Issue fetching matchup comparison")
		return sportscrape.EventDataOutput[model.MLBOddsMoneyLine]{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	// Unmarshal JSON
	var responsePayload jsonresponse.MLBMatchupComparison
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.MLBOddsMoneyLine]{Error: err, Context: context}
	}
	if responsePayload.BetSection == nil {
		log.Printf("No betting odds data available for event %d\n", matchup.EventID)
		return sportscrape.EventDataOutput[model.MLBOddsMoneyLine]{Context: context}
	}
	if responsePayload.BetSection.Name != betSectionTitle {
		err = fmt.Errorf("unknown title '%s'. expected '%s'", responsePayload.BetSection.Name, betSectionTitle)
		return sportscrape.EventDataOutput[model.MLBOddsMoneyLine]{Error: err, Context: context}
	}

	odds, err := s.record(matchup, responsePayload, context)
	if err != nil {
		return sportscrape.EventDataOutput[model.MLBOddsMoneyLine]{Error: err, Context: context}
	}

	if odds != nil {
		data = append(data, *odds)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchup.EventID, matchup.AwayTeamNameFull, matchup.HomeTeamNameFull, diff)
	return sportscrape.EventDataOutput[model.MLBOddsMoneyLine]{Output: data, Context: context}
}

func (s *MLBOddsMoneyLineScraper) record(matchup model.Matchup, responsePayload jsonresponse.MLBMatchupComparison, context sportscrape.EventDataContext) (*model.MLBOddsMoneyLine, error) {
	var oddsText string
	for _, bet := range responsePayload.BetSection.Bets {
		if bet.Model.Subtitle != oddsMoneyLineTitle {
			continue
		}
		record := &model.MLBOddsMoneyLine{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventID:              context.EventID.(int64),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			HomeTeamID:           context.HomeID.(int64),
			HomeTeamNameFull:     context.HomeTeam,
			AwayTeamID:           context.AwayID.(int64),
			AwayTeamNameFull:     context.AwayTeam,
		}
		n := len(bet.Model.Odds)
		if n != 2 {
			return nil, fmt.Errorf("%d mlb odds money line items identified. expected 2", n)
		}
		for _, oddsItem := range bet.Model.Odds {
			oddsText = *oddsItem.Text
			odds, err := util.TextToInt32(oddsText)
			if err != nil {
				return nil, err
			}
			if matchup.AwayTeamAbbreviation == oddsItem.SubText {
				record.AwayTeamOdds = odds
			} else if matchup.HomeTeamAbbreviation == oddsItem.SubText {
				record.HomeTeamOdds = odds
			} else {
				return nil, fmt.Errorf("unexpected team abbreviation identifed '%s'. expected %s or %s", oddsItem.SubText, matchup.HomeTeamAbbreviation, matchup.AwayTeamAbbreviation)
			}
		}
		return record, nil
	}
	return nil, nil
}
