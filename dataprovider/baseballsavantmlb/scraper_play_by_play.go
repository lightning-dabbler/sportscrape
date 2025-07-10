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

// PlayByPlayScraperOption defines a configuration option for the scraper
type PlayByPlayScraperOption func(*PlayByPlayScraper)

// NewPlayByPlayScraper creates a new PlayByPlayScraper with the provided options
func NewPlayByPlayScraper(options ...PlayByPlayScraperOption) *PlayByPlayScraper {
	s := &PlayByPlayScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}
	s.Init()

	return s
}

type PlayByPlayScraper struct {
	EventDataScraper
}

func (s PlayByPlayScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballSavantMLBPlayByPlay
}

func (s PlayByPlayScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
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
	res, err := s.constructPlayByPlay(gf.HomePitchers, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	// away pitchers
	res, err = s.constructPlayByPlay(gf.AwayPitchers, context)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	if res != nil {
		data = append(data, res...)
	}
	return sportscrape.EventDataOutput{Error: err, Context: context, Output: data}
}

func (s PlayByPlayScraper) constructPlayByPlay(plays *jsonresponse.Plays, context sportscrape.EventDataContext) ([]interface{}, error) {
	if plays == nil {
		return nil, nil
	}
	var data []interface{}
	eventid := context.EventID.(int64)
	for _, events := range *plays {
		for _, event := range events {
			playbyplay := model.PlayByPlay{
				PullTimestamp:                  context.PullTimestamp,
				PullTimestampParquet:           types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
				EventTime:                      context.EventTime,
				EventTimeParquet:               types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
				EventID:                        eventid,
				PlayID:                         event.PlayID,
				Inning:                         event.Inning,
				AtBatNum:                       event.AtBatNum,
				Strikes:                        event.Strikes,
				Balls:                          event.Balls,
				Outs:                           event.Outs,
				PreStrikes:                     event.PreStrikes,
				PreBalls:                       event.PreBalls,
				BatterID:                       event.BatterID,
				BatterStand:                    event.BatterStand,
				BatterName:                     event.BatterName,
				PitcherID:                      event.PitcherID,
				PitcherThrow:                   event.PitcherThrow,
				PitcherName:                    event.PitcherName,
				PitchType:                      event.PitchType,
				PitchName:                      event.PitchName,
				CallName:                       event.CallName,
				CallDescription:                event.CallDescription,
				IsStrikeSwinging:               event.IsStrikeSwinging,
				PitchStartSpeed:                event.PitchStartSpeed,
				PitchEndSpeed:                  event.PitchEndSpeed,
				PitchNumber:                    event.PitchNumber,
				PitcherTotalPitches:            event.PitcherTotalPitches,
				PitcherTotalPitchesByPitchType: event.PitcherTotalPitchesByPitchType,
				GameTotalPitches:               event.GameTotalPitches,
			}
			if event.Result != nil {
				playbyplay.Result = event.Result
			}
			if event.ResultDescription != nil {
				playbyplay.ResultDescription = event.ResultDescription
			}
			if event.HitDistance != nil && *event.HitDistance != "" {
				hitdistance, err := util.TextToInt32(*event.HitDistance)
				if err != nil {
					log.Printf("Issue parsing hit distance %s\n", *event.HitDistance)
					return nil, err
				}
				playbyplay.HitDistance = &hitdistance
			}
			if event.HitSpeed != nil && *event.HitSpeed != "" {
				hitspeed, err := util.TextToFloat32(*event.HitSpeed)
				if err != nil {
					log.Printf("Issue parsing hit speed %s\n", *event.HitSpeed)
					return nil, err
				}
				playbyplay.HitSpeed = &hitspeed
			}

			data = append(data, playbyplay)
		}
	}
	return data, nil
}
