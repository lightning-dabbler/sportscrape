package baseballsavantmlb

import (
	"log"
	"strconv"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/xitongsys/parquet-go/types"
)

// BattingLineupScraperOption defines a configuration option for the scraper
type BattingLineupScraperOption func(*BattingLineupScraper)

// NewBattingLineupScraper creates a new BattingLineupScraper with the provided options
func NewBattingLineupScraper(options ...BattingLineupScraperOption) *BattingLineupScraper {
	s := &BattingLineupScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// BattingLineupScraper scrapes the batter lineup (`away_lineup` and `home_lineup`) from the baseballsavant.mlb.com game feed.
type BattingLineupScraper struct {
	EventDataScraper
}

func (s *BattingLineupScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballSavantMLBBattingLineup
}

func (s *BattingLineupScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.BattingLineup] {
	context := s.ConstructContext(matchup)
	url := ConstructEventDataURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	gf, err := s.FetchGameFeed(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return sportscrape.EventDataOutput[model.BattingLineup]{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	var data []model.BattingLineup
	// home batters
	res := s.constructBattingLineup("home", gf.HomeLineup, gf, context)
	if res != nil {
		data = append(data, res...)
	}
	// away batters
	res = s.constructBattingLineup("away", gf.AwayLineup, gf, context)
	if res != nil {
		data = append(data, res...)
	}
	return sportscrape.EventDataOutput[model.BattingLineup]{Error: err, Context: context, Output: data}
}

func (s *BattingLineupScraper) constructBattingLineup(team string, lineup jsonresponse.Lineup, gf jsonresponse.GameFeed, context sportscrape.EventDataContext) []model.BattingLineup {
	if len(lineup) == 0 {
		// e.g. the game has not started or the lineup has not been posted yet
		return nil
	}
	var teamName, opponentName string
	var teamid, opponentid int64
	var boxscoreteam jsonresponse.BoxScoreTeam
	var data []model.BattingLineup
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
	for i, playerid := range lineup {
		fmtID := s.FmtID(strconv.FormatInt(playerid, 10))
		player, exists := boxscoreteam.Players[fmtID]
		if !exists {
			log.Printf("No batter lineup data available for player %d\n", playerid)
			continue
		}
		data = append(data, model.BattingLineup{
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			EventID:              eventid,
			TeamID:               teamid,
			Team:                 teamName,
			Opponent:             opponentName,
			OpponentID:           opponentid,
			PlayerID:             playerid,
			Player:               player.Person.Name,
			Position:             player.Position.Name,
			LineupOrder:          int32(i + 1),
		})
	}
	return data
}
