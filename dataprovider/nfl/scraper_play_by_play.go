package nfl

import (
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
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

// PlayByPlayScraper scrapes every play of the matchup's drive chart (all drives), with each play's drive attached
type PlayByPlayScraper struct {
	EventDataScraper
}

func (s *PlayByPlayScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLPlayByPlay
}

func (s *PlayByPlayScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PlayByPlay] {
	context := s.ConstructContext(matchup)
	url := ConstructGameDetailsURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	details, err := fetchJSON[jsonresponse.GameDetails](url, s.Fetcher)
	if err != nil {
		return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
	}
	// Drive chart is not available yet (the game is further out)
	if details.DriveChart == nil {
		return sportscrape.EventDataOutput[model.PlayByPlay]{Context: context}
	}
	drives := make(map[int32]jsonresponse.Drive, len(details.DriveChart.Drives))
	for _, drive := range details.DriveChart.Drives {
		drives[drive.Sequence] = drive
	}

	var data []model.PlayByPlay
	for _, play := range details.DriveChart.Plays {
		record := model.PlayByPlay{
			PullTimestamp:                pullTimestamp,
			PullTimestampParquet:         pullTimestampParquet,
			EventID:                      matchup.EventID,
			EventTime:                    matchup.EventTime,
			EventTimeParquet:             matchup.EventTimeParquet,
			HomeTeamID:                   matchup.HomeTeamID,
			AwayTeamID:                   matchup.AwayTeamID,
			PlayID:                       play.PlayID,
			PlaySequenceNumber:           play.PlaySequenceNumber,
			Quarter:                      play.Quarter,
			Clock:                        play.ClockTime,
			PlayType:                     play.PlayType,
			Down:                         play.Down,
			Distance:                     play.YardsRemaining,
			YardLine:                     play.YardLine,
			YardsGained:                  play.YardsGained,
			PrePlayByPlay:                play.PrePlayByPlay,
			Description:                  play.PlayDescription,
			DescriptionWithJerseyNumbers: play.PlayDescriptionWithJerseyNos,
			IsGoalToGo:                   play.PlayIsGoalToGo,
			IsEndOfQuarter:               play.PlayIsEndOfQuarter,
			Scored:                       play.PlayScored,
			ScoringPlayType:              play.ScoringPlayType,
			ScoringTeamID:                play.ScoringTeamID,
			SpecialTeamsPlayType:         play.SpecialTeamsPlayType,
			NextPlayType:                 play.NextPlayType,
			NextPlayIsGoalToGo:           play.NextPlayIsGoalToGo,
			Deleted:                      play.PlayDeleted,
		}
		if play.PlayStartTime != nil {
			start, err := util.RFC3339ToTime(*play.PlayStartTime)
			if err != nil {
				return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
			}
			startParquet := types.TimeToTIMESTAMP_MILLIS(start, true)
			record.PlayStartTime = &start
			record.PlayStartTimeParquet = &startParquet
		}
		if play.PlayEndTime != nil {
			end, err := util.RFC3339ToTime(*play.PlayEndTime)
			if err != nil {
				return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
			}
			endParquet := types.TimeToTIMESTAMP_MILLIS(end, true)
			record.PlayEndTime = &end
			record.PlayEndTimeParquet = &endParquet
		}
		// driveSequence is 0 when the play isn't part of a drive
		if drive, exists := drives[play.DriveSequence]; exists && play.DriveSequence > 0 {
			record.DriveSequence = &drive.Sequence
			record.DriveTeamID = &drive.TeamID
			record.DriveStart = &drive.StartedDescription
			record.DriveResult = &drive.EndedDescription
		}
		data = append(data, record)
	}
	return sportscrape.EventDataOutput[model.PlayByPlay]{Context: context, Output: data}
}
