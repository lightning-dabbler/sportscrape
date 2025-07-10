package baseballsavantmlb

import (
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/xitongsys/parquet-go/types"
)

// FieldingBoxScoreScraperOption defines a configuration option for the scraper
type FieldingBoxScoreScraperOption func(*FieldingBoxScoreScraper)

// NewFieldingBoxScoreScraper creates a new FieldingBoxScoreScraper with the provided options
func NewFieldingBoxScoreScraper(options ...FieldingBoxScoreScraperOption) *FieldingBoxScoreScraper {
	s := &FieldingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.Init()

	return s
}

type FieldingBoxScoreScraper struct {
	EventDataScraper
}

func (s FieldingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballSavantMLBFieldingBoxScore
}

func (s FieldingBoxScoreScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	matchupmodel := matchup.(model.Matchup)
	context := s.ConstructContext(matchupmodel)
	url := ConstructEventDataURL(matchupmodel.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	gf, err := s.FetchGameFeed(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	var data []interface{}
	// home pitchers
	res, err := s.constructFielding("home", gf.HomePitchers, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	// home batters
	res, err = s.constructFielding("home", gf.HomeBatters, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	// away pitchers
	res, err = s.constructFielding("away", gf.AwayPitchers, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	// away batters
	res, err = s.constructFielding("away", gf.AwayBatters, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	return sportscrape.EventDataOutput{Error: err, Context: context, Output: data}
}

func (s FieldingBoxScoreScraper) constructFielding(team string, plays *jsonresponse.Plays, gf jsonresponse.GameFeed, context sportscrape.EventDataContext) ([]interface{}, error) {
	if plays == nil {
		return nil, nil
	}
	var teamName, opponentName string
	var teamid, opponentid int64
	var boxscoreteam jsonresponse.BoxScoreTeam
	var data []interface{}
	eventid := context.EventID.(int64)
	switch team {
	case "home":
		teamName = context.HomeTeam
		teamid = context.HomeID.(int64)
		boxscoreteam = gf.BoxScore.Teams.Home
		opponentName = context.AwayTeam
		opponentid = context.AwayID.(int64)
	default:
		teamName = context.AwayTeam
		teamid = context.AwayID.(int64)
		boxscoreteam = gf.BoxScore.Teams.Away
		opponentName = context.HomeTeam
		opponentid = context.HomeID.(int64)
	}
	for playerid := range *plays {
		fmtID := s.FmtID(playerid)
		player, exists := boxscoreteam.Players[fmtID]
		if !exists {
			log.Printf("No fielding box score data available for player %s\n", playerid)
			continue
		}
		if player.Stats.Fielding.Fielding == "" || player.Stats.Fielding.StolenBasePercentage == "" {
			log.Printf("No fielding box score data available for player %s\n", playerid)
			continue
		}

		box_score := model.FieldingBoxScore{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			EventID:              eventid,
			TeamID:               teamid,
			Team:                 teamName,
			Opponent:             opponentName,
			OpponentID:           opponentid,
			PlayerID:             player.Person.ID,
			Player:               player.Person.Name,
			Position:             player.Position.Name,
			CaughtStealing:       player.Stats.Fielding.CaughtStealing,
			StolenBases:          player.Stats.Fielding.StolenBases,
			Assists:              player.Stats.Fielding.Assists,
			Putouts:              player.Stats.Fielding.PutOuts,
			Errors:               player.Stats.Fielding.Errors,
			Chances:              player.Stats.Fielding.Chances,
			PassedBall:           player.Stats.Fielding.PassedBall,
			Pickoffs:             player.Stats.Fielding.Pickoffs,
		}
		data = append(data, box_score)
	}
	return data, nil

}
