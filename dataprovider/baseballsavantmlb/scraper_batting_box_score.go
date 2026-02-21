package baseballsavantmlb

import (
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// BattingBoxScoreScraperOption defines a configuration option for the scraper
type BattingBoxScoreScraperOption func(*BattingBoxScoreScraper)

// NewBattingBoxScoreScraper creates a new BattingBoxScoreScraper with the provided options
func NewBattingBoxScoreScraper(options ...BattingBoxScoreScraperOption) *BattingBoxScoreScraper {
	s := &BattingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.Init()

	return s
}

type BattingBoxScoreScraper struct {
	EventDataScraper
}

func (s BattingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballSavantMLBBattingBoxScore
}

func (s BattingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.BattingBoxScore] {
	context := s.ConstructContext(matchup)
	url := ConstructEventDataURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	gf, err := s.FetchGameFeed(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return sportscrape.EventDataOutput[model.BattingBoxScore]{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	var data []model.BattingBoxScore
	// home batters
	res, err := s.constructBatting("home", gf.HomeBatters, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput[model.BattingBoxScore]{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	// away batters
	res, err = s.constructBatting("away", gf.AwayBatters, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput[model.BattingBoxScore]{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	return sportscrape.EventDataOutput[model.BattingBoxScore]{Error: err, Context: context, Output: data}
}

func (s BattingBoxScoreScraper) constructBatting(team string, plays *jsonresponse.Plays, gf jsonresponse.GameFeed, context sportscrape.EventDataContext) ([]model.BattingBoxScore, error) {
	if plays == nil {
		return nil, nil
	}
	var teamName, opponentName string
	var teamid, opponentid int64
	var boxscoreteam jsonresponse.BoxScoreTeam
	var data []model.BattingBoxScore
	var err error
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
			log.Printf("No batting box score data available for player %s\n", playerid)
			continue
		}
		if player.Stats.Batting.AtBatsPerHomeRun == "" || player.Stats.Batting.StolenBasePercentage == "" {
			log.Printf("No batting box score data available for player %s\n", playerid)
			continue
		}
		var ab_hr float32

		ab_hr_str := player.Stats.Batting.AtBatsPerHomeRun

		if ab_hr_str == "-.--" {
			ab_hr = float32(0)
		} else {
			ab_hr, err = util.TextToFloat32(player.Stats.Batting.AtBatsPerHomeRun)
			if err != nil {
				log.Printf("Issue parsing AB/HR %s\n", player.Stats.Batting.AtBatsPerHomeRun)
				return nil, err
			}
		}

		box_score := model.BattingBoxScore{
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
			FlyOuts:              player.Stats.Batting.FlyOuts,
			GroundOuts:           player.Stats.Batting.GroundOuts,
			AirOuts:              player.Stats.Batting.AirOuts,
			Runs:                 player.Stats.Batting.Runs,
			Doubles:              player.Stats.Batting.Doubles,
			Triples:              player.Stats.Batting.Triples,
			HomeRuns:             player.Stats.Batting.HomeRuns,
			Strikeouts:           player.Stats.Batting.StrikeOuts,
			Walks:                player.Stats.Batting.BaseOnBalls,
			IntentionalWalks:     player.Stats.Batting.IntentionalWalks,
			Hits:                 player.Stats.Batting.Hits,
			HitByPitch:           player.Stats.Batting.HitByPitch,
			AtBats:               player.Stats.Batting.AtBats,
			CaughtStealing:       player.Stats.Batting.CaughtStealing,
			StolenBases:          player.Stats.Batting.StolenBases,
			GroundIntoDoublePlay: player.Stats.Batting.GroundIntoDoublePlay,
			GroundIntoTriplePlay: player.Stats.Batting.GroundIntoTriplePlay,
			PlateAppearances:     player.Stats.Batting.PlateAppearances,
			TotalBases:           player.Stats.Batting.TotalBases,
			RBI:                  player.Stats.Batting.RBI,
			LeftOnBase:           player.Stats.Batting.LeftOnBase,
			SacBunts:             player.Stats.Batting.SacBunts,
			SacFlies:             player.Stats.Batting.SacFlies,
			CatchersInterference: player.Stats.Batting.CatchersInterference,
			Pickoffs:             player.Stats.Batting.Pickoffs,
			AtBatsPerHomeRun:     ab_hr,
			PopOuts:              player.Stats.Batting.PopOuts,
			LineOuts:             player.Stats.Batting.LineOuts,
		}
		data = append(data, box_score)
	}
	return data, nil

}
