package foxsports

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

const (
	probablePitcherTitle = "PROBABLE STARTING PITCHERS"
	regexExpr            = `^(\d+\.\d+)\sERA`
)

var re = regexp.MustCompile(regexExpr)

// MLBProbableStartingPitcherScraperOption defines a configuration option for the scraper
type MLBProbableStartingPitcherScraperOption func(*MLBProbableStartingPitcherScraper)

// MLBProbableStartingPitcherScraperParams sets the Params option
func MLBProbableStartingPitcherScraperParams(params map[string]string) MLBProbableStartingPitcherScraperOption {
	return func(s *MLBProbableStartingPitcherScraper) {
		s.Params = params
	}
}

// NewMLBProbableStartingPitcherScraper creates a new MLBProbableStartingPitcherScraper with the provided options
func NewMLBProbableStartingPitcherScraper(options ...MLBProbableStartingPitcherScraperOption) *MLBProbableStartingPitcherScraper {
	s := &MLBProbableStartingPitcherScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.League = MLB
	s.Init()

	return s
}

type MLBProbableStartingPitcherScraper struct {
	EventDataScraper
}

func (s MLBProbableStartingPitcherScraper) Feed() sportscrape.Feed {
	return sportscrape.FSMLBProbableStartingPitcher
}

func (s *MLBProbableStartingPitcherScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	var context sportscrape.EventDataContext
	context.AwayTeam = matchupModel.AwayTeamNameFull
	context.AwayID = matchupModel.AwayTeamID
	context.HomeTeam = matchupModel.HomeTeamNameFull
	context.HomeID = matchupModel.HomeTeamID
	context.EventID = matchupModel.EventID
	context.EventTime = matchupModel.EventTime

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
	if responsePayload.FeaturedPairing == nil {
		log.Printf("No probable starting pitcher data available for event %d\n", matchupModel.EventID)
		return sportscrape.EventDataOutput{Context: context}
	}
	if responsePayload.FeaturedPairing.Title != probablePitcherTitle {
		err = fmt.Errorf("unknown title '%s', expected '%s'", responsePayload.FeaturedPairing.Title, probablePitcherTitle)
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}

	pitcher, err := s.pitcher("home", responsePayload, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if pitcher != nil {
		data = append(data, *pitcher)
	}

	pitcher, err = s.pitcher("away", responsePayload, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if pitcher != nil {
		data = append(data, *pitcher)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchupModel.EventID, matchupModel.AwayTeamNameFull, matchupModel.HomeTeamNameFull, diff)
	return sportscrape.EventDataOutput{Output: data, Context: context}
}

func (s *MLBProbableStartingPitcherScraper) era(rawStatline string) (float32, error) {
	matches := re.FindStringSubmatch(rawStatline)
	if len(matches) != 2 {
		return 0, fmt.Errorf("%s does not match ERA pattern %s in featuredPairing", rawStatline, regexExpr)
	}
	era, err := util.TextToFloat32(matches[1])
	if err != nil {
		return 0, err
	}
	return era, nil
}

func (s *MLBProbableStartingPitcherScraper) pitcher(team string, responsePayload jsonresponse.MLBMatchupComparison, context sportscrape.EventDataContext) (*model.MLBProbableStartingPitcher, error) {
	var name, era, playerid, teamName string

	switch team {
	case "home":
		name = responsePayload.FeaturedPairing.HomePitcher.Name
		teamName = context.HomeTeam
	default:
		name = responsePayload.FeaturedPairing.AwayPitcher.Name
		teamName = context.AwayTeam
	}

	if name == "TBD" {
		log.Printf("Skipping: %s starting pitcher not announced at %s", teamName, context.URL)
	} else {
		probablePitchers := &model.MLBProbableStartingPitcher{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventID:              context.EventID,
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			TeamNameFull:         teamName,
		}

		switch team {
		case "home":
			probablePitchers.TeamID = context.HomeID
			probablePitchers.StartingPitcherRecord = responsePayload.FeaturedPairing.HomePitcher.StatLine1
			probablePitchers.StartingPitcher = responsePayload.FeaturedPairing.HomePitcher.Player
			playerid = responsePayload.FeaturedPairing.HomePitcher.EntityLink.Layout.Tokens.ID
			era = responsePayload.FeaturedPairing.HomePitcher.StatLine2
		default:
			probablePitchers.TeamID = context.AwayID
			probablePitchers.StartingPitcherRecord = responsePayload.FeaturedPairing.AwayPitcher.StatLine1
			probablePitchers.StartingPitcher = responsePayload.FeaturedPairing.AwayPitcher.Player
			playerid = responsePayload.FeaturedPairing.AwayPitcher.EntityLink.Layout.Tokens.ID
			era = responsePayload.FeaturedPairing.AwayPitcher.StatLine2
		}
		// StartingPitcherID
		playerID, err := util.TextToInt64(playerid)
		if err != nil {
			return nil, err
		}
		probablePitchers.StartingPitcherID = playerID

		// StartingPitcherERA
		era, err := s.era(era)
		if err != nil {
			return nil, err
		}
		probablePitchers.StartingPitcherERA = era
		return probablePitchers, nil
	}

	return nil, nil
}
