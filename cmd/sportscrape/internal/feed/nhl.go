package feed

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	NHLConcurrencyOptions string = strings.Join([]string{
		"'matchup-periods'",
		"'skater-box-score'",
		"'goalie-box-score'",
		"'play-by-play'",
	}, ", ")

	NHLOptions string = fmt.Sprintf("'matchup', %s", NHLConcurrencyOptions)
	ErrNHL     error  = fmt.Errorf("valid options: %s", NHLOptions)
)

type NHLExtractor struct {
	Feed              string
	Date              string
	FetchAttempts     int
	FetchRetryBackoff time.Duration
	Timeout           time.Duration
	Concurrency       int
	OutputPath        string
	Format            string
	S3Config          exporters.S3Config
	ParquetOptions    []exporters.ParquetConfigOption
}

func (e *NHLExtractor) fetcher() nhl.Fetcher {
	return nhl.Fetcher{
		FetchAttempts:     e.FetchAttempts,
		FetchRetryBackoff: e.FetchRetryBackoff,
		Timeout:           e.Timeout,
	}
}

func (e *NHLExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "matchup", "matchup-periods", "skater-box-score", "goalie-box-score", "play-by-play":
		return nil
	default:
		return fmt.Errorf("unsupported feed %q for nhl. %w", e.Feed, ErrNHL)
	}
}

func (e *NHLExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "matchup":
		return e.scrapeMatchup(ctx)
	case "matchup-periods":
		return e.scrapeMatchupPeriods(ctx)
	case "skater-box-score":
		return e.scrapeSkaterBoxScore(ctx)
	case "goalie-box-score":
		return e.scrapeGoalieBoxScore(ctx)
	case "play-by-play":
		return e.scrapePlayByPlay(ctx)
	default:
		return fmt.Errorf("unsupported feed %q for nhl. %w", e.Feed, ErrNHL)
	}
}

func (e *NHLExtractor) retrieveMatchup() ([]model.Matchup, error) {
	matchupscraper := nhl.NewMatchupScraper(
		nhl.WithMatchupDate(e.Date),
	)
	matchupscraper.Fetcher = e.fetcher()
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)

	return matchuprunner.Run()
}

func (e *NHLExtractor) scrapeMatchup(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, m, e.ParquetOptions...)
}

func (e *NHLExtractor) scrapeMatchupPeriods(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := nhl.NewMatchupPeriodsScraper()
	eventdatascraper.Fetcher = e.fetcher()
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
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

func (e *NHLExtractor) scrapeSkaterBoxScore(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := nhl.NewSkaterBoxScoreScraper()
	eventdatascraper.Fetcher = e.fetcher()
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.SkaterBoxScore]{
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

func (e *NHLExtractor) scrapeGoalieBoxScore(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := nhl.NewGoalieBoxScoreScraper()
	eventdatascraper.Fetcher = e.fetcher()
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.GoalieBoxScore]{
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

func (e *NHLExtractor) scrapePlayByPlay(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := nhl.NewPlayByPlayScraper()
	eventdatascraper.Fetcher = e.fetcher()
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
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
