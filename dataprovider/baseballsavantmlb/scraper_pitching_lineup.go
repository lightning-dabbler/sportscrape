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

// PitchingLineupScraperOption defines a configuration option for the scraper
type PitchingLineupScraperOption func(*PitchingLineupScraper)

// NewPitchingLineupScraper creates a new PitchingLineupScraper with the provided options
func NewPitchingLineupScraper(options ...PitchingLineupScraperOption) *PitchingLineupScraper {
	s := &PitchingLineupScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// PitchingLineupScraper scrapes the pitcher lineup (`away_pitcher_lineup` and `home_pitcher_lineup`) from the baseballsavant.mlb.com game feed.
type PitchingLineupScraper struct {
	EventDataScraper
}

func (s *PitchingLineupScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballSavantMLBPitchingLineup
}

func (s *PitchingLineupScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PitchingLineup] {
	context := s.ConstructContext(matchup)
	url := ConstructEventDataURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	gf, err := s.FetchGameFeed(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return sportscrape.EventDataOutput[model.PitchingLineup]{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	var data []model.PitchingLineup
	// home pitchers
	res := s.constructPitchingLineup("home", gf.HomePitcherLineup, gf, context)
	if res != nil {
		data = append(data, res...)
	}
	// away pitchers
	res = s.constructPitchingLineup("away", gf.AwayPitcherLineup, gf, context)
	if res != nil {
		data = append(data, res...)
	}
	return sportscrape.EventDataOutput[model.PitchingLineup]{Error: err, Context: context, Output: data}
}

func (s *PitchingLineupScraper) constructPitchingLineup(team string, lineup jsonresponse.Lineup, gf jsonresponse.GameFeed, context sportscrape.EventDataContext) []model.PitchingLineup {
	if len(lineup) == 0 {
		// e.g. the game has not started or the lineup has not been posted yet
		return nil
	}
	var teamName, opponentName string
	var teamid, opponentid int64
	var boxscoreteam jsonresponse.BoxScoreTeam
	var data []model.PitchingLineup
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
			log.Printf("No pitcher lineup data available for player %d\n", playerid)
			continue
		}
		data = append(data, model.PitchingLineup{
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
