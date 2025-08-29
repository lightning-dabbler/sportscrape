package foxsports

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

const (
	betSectionTitle    = "ODDS"
	oddsTotalTitle     = "TOTAL"
	oddsTotalLineRegex = `(?:UNDER|OVER)\s+(\d+(?:\.\d)?)` // e.g "OVER 9"| "UNDER 8.5"
)

var oddsTotalLineRe = regexp.MustCompile(oddsTotalLineRegex)

// MLBOddsTotalScraperOption defines a configuration option for the scraper
type MLBOddsTotalScraperOption func(*MLBOddsTotalScraper)

// MLBOddsTotalScraperParams sets the Params option
func MLBOddsTotalScraperParams(params map[string]string) MLBOddsTotalScraperOption {
	return func(s *MLBOddsTotalScraper) {
		s.Params = params
	}
}

// NewMLBOddsTotalScraper creates a new MLBOddsTotalScraper with the provided options
func NewMLBOddsTotalScraper(options ...MLBOddsTotalScraperOption) *MLBOddsTotalScraper {
	s := &MLBOddsTotalScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.League = MLB
	s.Init()

	return s
}

type MLBOddsTotalScraper struct {
	EventDataScraper
}

func (s MLBOddsTotalScraper) Feed() sportscrape.Feed {
	return sportscrape.FSMLBOddsTotal
}

func (s *MLBOddsTotalScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	context := s.ConstructContext(matchupModel)

	var data []interface{}
	// Construct event data URL
	log.Println("Constructing event data URL")
	url, err := s.ConstructMatchupComparisonURL(matchupModel.EventID)
	if err != nil {
		log.Println("Issue constructing matchup comparison URL")
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	// Fetch event data
	responseBody, err := s.FetchData(url)
	if err != nil {
		log.Println("Issue fetching matchup comparison")
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	// Unmarshal JSON
	var responsePayload jsonresponse.MLBMatchupComparison
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if responsePayload.BetSection == nil {
		log.Printf("No betting odds data available for event %d\n", matchupModel.EventID)
		return sportscrape.EventDataOutput{Context: context}
	}
	if responsePayload.BetSection.Name != betSectionTitle {
		err = fmt.Errorf("unknown title '%s'. expected '%s'", responsePayload.BetSection.Name, betSectionTitle)
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}

	odds, err := s.record(matchupModel, responsePayload, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}

	if odds != nil {
		data = append(data, *odds)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchupModel.EventID, matchupModel.AwayTeamNameFull, matchupModel.HomeTeamNameFull, diff)
	return sportscrape.EventDataOutput{Output: data, Context: context}
}

func (s *MLBOddsTotalScraper) parseLine(lineText string) (float32, error) {
	matches := oddsTotalLineRe.FindStringSubmatch(lineText)
	if len(matches) != 2 {
		return 0, fmt.Errorf("'%s' does not match odds total line pattern %s", lineText, oddsTotalLineRegex)
	}
	line, err := util.TextToFloat32(matches[1])
	if err != nil {
		return 0, err
	}
	return line, nil
}

func (s *MLBOddsTotalScraper) record(matchupModel model.Matchup, responsePayload jsonresponse.MLBMatchupComparison, context sportscrape.EventDataContext) (*model.MLBOddsTotal, error) {
	var lineText, oddsText string
	var line float32
	var err error
	for _, bet := range responsePayload.BetSection.Bets {
		if bet.Model.Subtitle != oddsTotalTitle {
			continue
		}
		record := &model.MLBOddsTotal{
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
		var n int
		n = len(bet.Model.Odds)
		if n != 2 {
			return nil, fmt.Errorf("%d mlb odds total items identified. expected 2", n)
		}
		parsedLines := []float32{}
		for _, oddsItem := range bet.Model.Odds {
			lineText = oddsItem.SubText
			oddsText = oddsItem.Text
			if strings.HasPrefix(lineText, "UNDER") {
				line, err = s.parseLine(lineText)
				if err != nil {
					return nil, err
				}
				parsedLines = append(parsedLines, line)
				odds, err := util.TextToInt32(oddsText)
				if err != nil {
					return nil, err
				}
				record.UnderOdds = odds

			} else if strings.HasPrefix(lineText, "OVER") {
				line, err = s.parseLine(lineText)
				if err != nil {
					return nil, err
				}
				parsedLines = append(parsedLines, line)
				odds, err := util.TextToInt32(oddsText)
				if err != nil {
					return nil, err
				}
				record.OverOdds = odds

			} else {
				return nil, fmt.Errorf("unrecognized prefix in odds total line text %s. expected OVER or UNDER", lineText)
			}
		}
		n = len(parsedLines)
		if n != 2 {
			return nil, fmt.Errorf("expected total lines to be parsed twice was instead parsed %d times", n)
		}
		if parsedLines[0] != parsedLines[1] {
			return nil, fmt.Errorf("expected equal line results (%f != %f)", parsedLines[0], parsedLines[1])
		}
		record.TotalLine = parsedLines[0]
		return record, nil
	}
	return nil, nil
}
