package eventdata

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"time"

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

type MLBProbableStartingPitcherScraper struct {
	EventDataScraper
}

func (s *MLBProbableStartingPitcherScraper) Scrape(matchup interface{}) OutputWrapper {
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
	url, err := s.ConstructMatchupComparisonURL(matchupModel.EventID)
	if err != nil {
		log.Println("Issue constructing matchup comparison URL")
		return OutputWrapper{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	// Fetch event data
	responseBody, err := s.FetchData(url)
	if err != nil {
		log.Println("Issue fetching matchup comparison")
		return OutputWrapper{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	// Unmarshal JSON
	var responsePayload jsonresponse.MLBMatchupComparison
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	if responsePayload.FeaturedPairing == nil {
		log.Printf("No probable starting pitcher data available for event %d\n", matchupModel.EventID)
		return OutputWrapper{Context: context}
	}
	if responsePayload.FeaturedPairing.Title != probablePitcherTitle {
		err = fmt.Errorf("unknown title '%s', expected '%s'", responsePayload.FeaturedPairing.Title, probablePitcherTitle)
		return OutputWrapper{Error: err, Context: context}
	}

	probablePitchers := model.MLBProbableStartingPitcher{
		PullTimestamp:             context.PullTimestamp,
		PullTimestampParquet:      types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
		EventID:                   context.EventID,
		EventTime:                 context.EventTime,
		EventTimeParquet:          types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
		HomeTeamID:                context.HomeID,
		HomeTeamNameFull:          context.HomeTeam,
		HomeStartingPitcher:       responsePayload.FeaturedPairing.HomePitcher.Player,
		HomeStartingPitcherRecord: responsePayload.FeaturedPairing.HomePitcher.StatLine1,
		AwayTeamID:                context.AwayID,
		AwayTeamNameFull:          context.AwayTeam,
		AwayStartingPitcher:       responsePayload.FeaturedPairing.AwayPitcher.Player,
		AwayStartingPitcherRecord: responsePayload.FeaturedPairing.AwayPitcher.StatLine1,
	}
	// HomeStartingPitcherID
	playerID, err := util.TextToInt64(responsePayload.FeaturedPairing.HomePitcher.EntityLink.Layout.Tokens.ID)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	probablePitchers.HomeStartingPitcherID = playerID
	// AwayStartingPitcherID
	playerID, err = util.TextToInt64(responsePayload.FeaturedPairing.AwayPitcher.EntityLink.Layout.Tokens.ID)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	probablePitchers.AwayStartingPitcherID = playerID
	// HomeStartingPitcherERA
	era, err := s.era(responsePayload.FeaturedPairing.HomePitcher.StatLine2)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	probablePitchers.HomeStartingPitcherERA = era
	// AwayStartingPitcherERA
	era, err = s.era(responsePayload.FeaturedPairing.AwayPitcher.StatLine2)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	probablePitchers.AwayStartingPitcherERA = era

	data = append(data, probablePitchers)
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchupModel.EventID, matchupModel.AwayTeamNameFull, matchupModel.HomeTeamNameFull, diff)
	return OutputWrapper{Output: data, Context: context}
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
