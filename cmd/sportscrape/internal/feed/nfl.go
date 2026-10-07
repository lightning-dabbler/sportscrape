package feed

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"
	"github.com/lightning-dabbler/sportscrape/scraper"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	NFLConcurrencyOptions string = strings.Join([]string{
		"'matchup-periods'",
		"'passing-box-score'",
		"'rushing-box-score'",
		"'receiving-box-score'",
		"'defense-box-score'",
		"'kicking-box-score'",
		"'kickoff-box-score'",
		"'punting-box-score'",
		"'kick-return-box-score'",
		"'punt-return-box-score'",
		"'fumbles-box-score'",
		"'interceptions-box-score'",
		"'play-by-play'",
		"'play-by-play-stats'",
	}, ", ")

	NFLOptions string = fmt.Sprintf("'matchup', 'injuries', %s", NFLConcurrencyOptions)
	ErrNFL     error  = fmt.Errorf("valid options: %s", NFLOptions)
)

type NFLExtractor struct {
	Feed              string
	Date              string
	Year              string
	FetchAttempts     int
	FetchRetryBackoff time.Duration
	Timeout           time.Duration
	Concurrency       int
	OutputPath        string
	Format            string
	S3Config          exporters.S3Config
	ParquetOptions    []exporters.ParquetConfigOption
}

func (e *NFLExtractor) fetcher() nfl.Fetcher {
	return nfl.Fetcher{
		FetchAttempts:     e.FetchAttempts,
		FetchRetryBackoff: e.FetchRetryBackoff,
		Timeout:           e.Timeout,
	}
}

func (e *NFLExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "matchup", "injuries", "matchup-periods", "passing-box-score", "rushing-box-score", "receiving-box-score",
		"defense-box-score", "kicking-box-score", "kickoff-box-score", "punting-box-score", "kick-return-box-score",
		"punt-return-box-score", "fumbles-box-score", "interceptions-box-score", "play-by-play", "play-by-play-stats":
		return nil
	default:
		return fmt.Errorf("unsupported feed %q for nfl. %w", e.Feed, ErrNFL)
	}
}

func (e *NFLExtractor) Scrape(ctx context.Context) error {
	fetcher := e.fetcher()
	switch e.Feed {
	case "matchup":
		return e.scrapeMatchup(ctx)
	case "injuries":
		return e.scrapeInjuries(ctx)
	case "matchup-periods":
		s := nfl.NewMatchupPeriodsScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "passing-box-score":
		s := nfl.NewPassingBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "rushing-box-score":
		s := nfl.NewRushingBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "receiving-box-score":
		s := nfl.NewReceivingBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "defense-box-score":
		s := nfl.NewDefenseBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "kicking-box-score":
		s := nfl.NewKickingBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "kickoff-box-score":
		s := nfl.NewKickoffBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "punting-box-score":
		s := nfl.NewPuntingBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "kick-return-box-score":
		s := nfl.NewKickReturnBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "punt-return-box-score":
		s := nfl.NewPuntReturnBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "fumbles-box-score":
		s := nfl.NewFumblesBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "interceptions-box-score":
		s := nfl.NewInterceptionsBoxScoreScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "play-by-play":
		s := nfl.NewPlayByPlayScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	case "play-by-play-stats":
		s := nfl.NewPlayByPlayStatsScraper()
		s.Fetcher = fetcher
		return scrapeNFLEventData(ctx, e, s)
	default:
		return fmt.Errorf("unsupported feed %q for nfl. %w", e.Feed, ErrNFL)
	}
}

func (e *NFLExtractor) retrieveMatchup() ([]model.Matchup, error) {
	matchupscraper := nfl.NewMatchupScraper(
		nfl.WithMatchupDate(e.Date),
	)
	matchupscraper.Fetcher = e.fetcher()
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)

	return matchuprunner.Run()
}

func (e *NFLExtractor) scrapeMatchup(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, m, e.ParquetOptions...)
}

func (e *NFLExtractor) scrapeInjuries(ctx context.Context) error {
	injuriesscraper := nfl.NewInjuriesScraper(
		nfl.WithInjuriesSeason(e.Year),
	)
	injuriesscraper.Fetcher = e.fetcher()
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Injury]{
			Scraper: injuriesscraper,
		},
	)
	injuries, err := matchuprunner.Run()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, injuries, e.ParquetOptions...)
}

// scrapeNFLEventData runs eventdatascraper against the date's matchups and writes the records
func scrapeNFLEventData[E any](ctx context.Context, e *NFLExtractor, eventdatascraper scraper.EventDataScraper[model.Matchup, E]) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, E]{
			Concurrency: e.Concurrency,
			Scraper:     eventdatascraper,
		},
	)
	records, err := eventrunner.Run(m)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}
