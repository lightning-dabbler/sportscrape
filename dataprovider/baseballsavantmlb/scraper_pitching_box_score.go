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

// PitchingBoxScoreScraperOption defines a configuration option for the scraper
type PitchingBoxScoreScraperOption func(*PitchingBoxScoreScraper)

// NewPitchingBoxScoreScraper creates a new PitchingBoxScoreScraper with the provided options
func NewPitchingBoxScoreScraper(options ...PitchingBoxScoreScraperOption) *PitchingBoxScoreScraper {
	s := &PitchingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.Init()

	return s
}

type PitchingBoxScoreScraper struct {
	EventDataScraper
}

func (s PitchingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballSavantMLBPitchingBoxScore
}

func (s PitchingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PitchingBoxScore] {
	context := s.ConstructContext(matchup)
	url := ConstructEventDataURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	gf, err := s.FetchGameFeed(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return sportscrape.EventDataOutput[model.PitchingBoxScore]{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	var data []model.PitchingBoxScore
	// home pitchers
	res, err := s.constructPitching("home", gf.HomePitchers, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput[model.PitchingBoxScore]{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	// away pitchers
	res, err = s.constructPitching("away", gf.AwayPitchers, gf, context)
	if err != nil {
		return sportscrape.EventDataOutput[model.PitchingBoxScore]{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	return sportscrape.EventDataOutput[model.PitchingBoxScore]{Error: err, Context: context, Output: data}
}

func (s PitchingBoxScoreScraper) constructPitching(team string, plays *jsonresponse.Plays, gf jsonresponse.GameFeed, context sportscrape.EventDataContext) ([]model.PitchingBoxScore, error) {
	if plays == nil {
		return nil, nil
	}
	var teamName, opponentName string
	var teamid, opponentid int64
	var boxscoreteam jsonresponse.BoxScoreTeam
	var data []model.PitchingBoxScore
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
			log.Printf("No pitching box score data available for player %s\n", playerid)
			continue
		}
		if player.Stats.Pitching.RunsScoredPer9 == "" || player.Stats.Pitching.StolenBasePercentage == "" {
			log.Printf("No pitching box score data available for player %s\n", playerid)
			continue
		}
		var ip float32

		ip, err = util.TextToFloat32(player.Stats.Pitching.InningsPitched)
		if err != nil {
			log.Printf("Issue parsing innings pitched %s\n", player.Stats.Pitching.InningsPitched)
			return nil, err
		}

		box_score := model.PitchingBoxScore{
			PullTimestamp:          context.PullTimestamp,
			PullTimestampParquet:   types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:              context.EventTime,
			EventTimeParquet:       types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			EventID:                eventid,
			TeamID:                 teamid,
			Team:                   teamName,
			Opponent:               opponentName,
			OpponentID:             opponentid,
			PlayerID:               player.Person.ID,
			Player:                 player.Person.Name,
			Position:               player.Position.Name,
			FlyOuts:                player.Stats.Pitching.FlyOuts,
			GroundOuts:             player.Stats.Pitching.GroundOuts,
			AirOuts:                player.Stats.Pitching.AirOuts,
			Runs:                   player.Stats.Pitching.Runs,
			Doubles:                player.Stats.Pitching.Doubles,
			Triples:                player.Stats.Pitching.Triples,
			HomeRuns:               player.Stats.Pitching.HomeRuns,
			Strikeouts:             player.Stats.Pitching.StrikeOuts,
			Walks:                  player.Stats.Pitching.BaseOnBalls,
			IntentionalWalks:       player.Stats.Pitching.IntentionalWalks,
			Hits:                   player.Stats.Pitching.Hits,
			HitByPitch:             player.Stats.Pitching.HitByPitch,
			AtBats:                 player.Stats.Pitching.AtBats,
			CaughtStealing:         player.Stats.Pitching.CaughtStealing,
			StolenBases:            player.Stats.Pitching.StolenBases,
			NumberOfPitches:        player.Stats.Pitching.NumberOfPitches,
			InningsPitched:         ip,
			Wins:                   player.Stats.Pitching.Wins,
			Losses:                 player.Stats.Pitching.Losses,
			Saves:                  player.Stats.Pitching.Saves,
			BlownSaves:             player.Stats.Pitching.BlownSaves,
			EarnedRuns:             player.Stats.Pitching.EarnedRuns,
			BattersFaced:           player.Stats.Pitching.BattersFaced,
			Outs:                   player.Stats.Pitching.Outs,
			Shutouts:               player.Stats.Pitching.Shutouts,
			Balls:                  player.Stats.Pitching.Balls,
			Strikes:                player.Stats.Pitching.Strikes,
			Balks:                  player.Stats.Pitching.Balks,
			WildPitches:            player.Stats.Pitching.WildPitches,
			Pickoffs:               player.Stats.Pitching.Pickoffs,
			RBI:                    player.Stats.Pitching.RBI,
			InheritedRunners:       player.Stats.Pitching.InheritedRunners,
			InheritedRunnersScored: player.Stats.Pitching.InheritedRunnersScored,
			CatchersInterference:   player.Stats.Pitching.CatchersInterference,
			SacBunts:               player.Stats.Pitching.SacBunts,
			SacFlies:               player.Stats.Pitching.SacFlies,
			PassedBall:             player.Stats.Pitching.PassedBall,
			PopOuts:                player.Stats.Pitching.PopOuts,
			LineOuts:               player.Stats.Pitching.LineOuts,
		}
		data = append(data, box_score)
	}
	return data, nil

}
