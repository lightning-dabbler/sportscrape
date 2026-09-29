package nfl

import (
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/xitongsys/parquet-go/types"
)

// PlayByPlayStatsScraperOption defines a configuration option for PlayByPlayStatsScraper
type PlayByPlayStatsScraperOption func(*PlayByPlayStatsScraper)

// NewPlayByPlayStatsScraper creates a new PlayByPlayStatsScraper with the provided options
func NewPlayByPlayStatsScraper(options ...PlayByPlayStatsScraperOption) *PlayByPlayStatsScraper {
	s := &PlayByPlayStatsScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// PlayByPlayStatsScraper scrapes the stats credited to players (or only teams) on each play of the matchup's drive chart,
// one record per stat (see PlayByPlayScraper for the plays and stat_type.go for the stat type codes)
type PlayByPlayStatsScraper struct {
	EventDataScraper
}

func (s *PlayByPlayStatsScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLPlayByPlayStats
}

func (s *PlayByPlayStatsScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PlayByPlayStat] {
	context := s.ConstructContext(matchup)
	url := ConstructGameDetailsURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	details, err := fetchJSON[jsonresponse.GameDetails](url, s.Fetcher)
	if err != nil {
		return sportscrape.EventDataOutput[model.PlayByPlayStat]{Error: err, Context: context}
	}
	// Drive chart is not available yet (the game is further out)
	if details.DriveChart == nil {
		return sportscrape.EventDataOutput[model.PlayByPlayStat]{Context: context}
	}

	// stats are always credited to the matchup's home or away team
	teams := map[string]string{
		matchup.HomeTeamID: matchup.HomeTeam,
		matchup.AwayTeamID: matchup.AwayTeam,
	}
	var data []model.PlayByPlayStat
	for _, play := range details.DriveChart.Plays {
		for _, stat := range play.Stats {
			record := model.PlayByPlayStat{
				PullTimestamp:        pullTimestamp,
				PullTimestampParquet: pullTimestampParquet,
				EventID:              matchup.EventID,
				EventTime:            matchup.EventTime,
				EventTimeParquet:     matchup.EventTimeParquet,
				PlayID:               play.PlayID,
				PlayStatID:           stat.PlayStatID,
				Quarter:              play.Quarter,
				StatType:             stat.StatType,
				StatTypeDescription:  StatTypeDescription(stat.StatType),
				Yards:                stat.Yards,
				TeamID:               stat.TeamID,
				Team:                 teams[stat.TeamID],
				PlayerID:             stat.GSISPlayerID,
				PersonID:             stat.PersonID,
				PlayerShortName:      stat.GSISPlayerName,
				JerseyNumber:         stat.GSISPlayerJerseyNumber,
			}
			// team stats (e.g. first down rushing) have no player
			if stat.GSISPlayerName != nil {
				var personID string
				if stat.PersonID != nil {
					personID = *stat.PersonID
				}
				name := s.playerName(personID, *stat.GSISPlayerName)
				record.Player = &name
			}
			data = append(data, record)
		}
	}
	return sportscrape.EventDataOutput[model.PlayByPlayStat]{Context: context, Output: data}
}
