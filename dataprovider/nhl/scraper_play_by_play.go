package nhl

import (
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// PlayByPlayScraperOption defines a configuration option for PlayByPlayScraper
type PlayByPlayScraperOption func(*PlayByPlayScraper)

// NewPlayByPlayScraper creates a new PlayByPlayScraper with the provided options
func NewPlayByPlayScraper(options ...PlayByPlayScraperOption) *PlayByPlayScraper {
	s := &PlayByPlayScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

type PlayByPlayScraper struct {
	EventDataScraper
}

func (s *PlayByPlayScraper) Feed() sportscrape.Feed {
	return sportscrape.NHLPlayByPlay
}

func (s *PlayByPlayScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PlayByPlay] {
	context := s.ConstructContext(matchup)
	url := ConstructPlayByPlayURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	pbp, err := fetchJSON[jsonresponse.PlayByPlay](url, s.Fetcher)
	if err != nil {
		return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
	}
	var data []model.PlayByPlay
	for _, play := range pbp.Plays {
		record := model.PlayByPlay{
			PullTimestamp:         pullTimestamp,
			PullTimestampParquet:  pullTimestampParquet,
			EventID:               matchup.EventID,
			EventTime:             matchup.EventTime,
			EventTimeParquet:      matchup.EventTimeParquet,
			HomeTeamID:            matchup.HomeTeamID,
			AwayTeamID:            matchup.AwayTeamID,
			PlayEventID:           play.EventID,
			SortOrder:             play.SortOrder,
			Period:                play.PeriodDescriptor.Number,
			PeriodType:            play.PeriodDescriptor.PeriodType,
			SituationCode:         play.SituationCode,
			HomeTeamDefendingSide: play.HomeTeamDefendingSide,
			TypeCode:              play.TypeCode,
			TypeDescKey:           play.TypeDescKey,
		}
		record.TimeInPeriod, err = util.TransformMinutesPlayed(play.TimeInPeriod)
		if err != nil {
			return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
		}
		record.TimeRemaining, err = util.TransformMinutesPlayed(play.TimeRemaining)
		if err != nil {
			return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
		}
		if d := play.Details; d != nil {
			record.EventOwnerTeamID = d.EventOwnerTeamID
			record.XCoord = d.XCoord
			record.YCoord = d.YCoord
			record.ZoneCode = d.ZoneCode
			record.ShotType = d.ShotType
			record.Reason = d.Reason
			record.SecondaryReason = d.SecondaryReason
			record.PenaltyTypeCode = d.TypeCode
			record.PenaltyDescKey = d.DescKey
			record.PenaltyDuration = d.Duration
			record.AwayScore = d.AwayScore
			record.HomeScore = d.HomeScore
			record.AwaySOG = d.AwaySOG
			record.HomeSOG = d.HomeSOG
			record.GoalInGame = d.GoalInGame
			record.ScoringPlayerID = d.ScoringPlayerID
			record.ScoringPlayerTotal = d.ScoringPlayerTotal
			record.Assist1PlayerID = d.Assist1PlayerID
			record.Assist1PlayerTotal = d.Assist1PlayerTotal
			record.Assist2PlayerID = d.Assist2PlayerID
			record.Assist2PlayerTotal = d.Assist2PlayerTotal
			record.GoalieInNetID = d.GoalieInNetID
			record.ShootingPlayerID = d.ShootingPlayerID
			record.BlockingPlayerID = d.BlockingPlayerID
			record.HittingPlayerID = d.HittingPlayerID
			record.HitteePlayerID = d.HitteePlayerID
			record.WinningPlayerID = d.WinningPlayerID
			record.LosingPlayerID = d.LosingPlayerID
			record.CommittedByPlayerID = d.CommittedByPlayerID
			record.DrawnByPlayerID = d.DrawnByPlayerID
			record.ServedByPlayerID = d.ServedByPlayerID
			record.PlayerID = d.PlayerID
		}
		data = append(data, record)
	}
	return sportscrape.EventDataOutput[model.PlayByPlay]{Context: context, Output: data}
}
