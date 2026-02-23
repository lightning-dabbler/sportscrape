package feed

import (
	"context"
	"fmt"
	"strings"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	BaseballSavantConcurrencyOptions string = strings.Join([]string{
		"'pitching-box-score'",
		"'batting-box-score'",
		"'fielding-box-score'",
		"'play-by-play'",
	}, ", ")

	BaseballSavantOptions string = fmt.Sprintf("'matchup', %s", BaseballSavantConcurrencyOptions)
	ErrBaseballSavant     error  = fmt.Errorf("valid options: %s", BaseballSavantOptions)
)

type BaseballSavantExtractor struct {
	Feed           string
	Date           string
	Concurrency    int
	OutputPath     string
	Format         string
	S3Config       exporters.S3Config
	ParquetOptions []exporters.ParquetConfigOption
}

func (e *BaseballSavantExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "matchup", "pitching-box-score", "batting-box-score", "fielding-box-score", "play-by-play":
		return nil
	default:
		return fmt.Errorf("unsupported feed %q for baseball savant. %w", e.Feed, ErrBaseballSavant)
	}
}

func (e *BaseballSavantExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "matchup":
		return e.scrapeMatchup(ctx)
	case "pitching-box-score":
		return e.scrapePitchingBoxScore(ctx)
	case "batting-box-score":
		return e.scrapeBattingBoxScore(ctx)
	case "fielding-box-score":
		return e.scrapeFieldingBoxScore(ctx)
	case "play-by-play":
		return e.scrapePlayByPlay(ctx)
	default:
		return fmt.Errorf("unsupported feed %q for baseball savant. %w", e.Feed, ErrBaseballSavant)
	}
}

func (e *BaseballSavantExtractor) retrieveMatchup() ([]model.Matchup, error) {
	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: baseballsavantmlb.NewMatchupScraper(
				baseballsavantmlb.MatchupScraperDate(e.Date),
			),
		},
	)

	return matchuprunner.Run()
}

func (e *BaseballSavantExtractor) scrapeMatchup(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, m, e.ParquetOptions...)
}

func (e *BaseballSavantExtractor) scrapePitchingBoxScore(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := baseballsavantmlb.NewPitchingBoxScoreScraper()
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PitchingBoxScore]{
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

func (e *BaseballSavantExtractor) scrapeBattingBoxScore(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := baseballsavantmlb.NewBattingBoxScoreScraper()
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BattingBoxScore]{
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

func (e *BaseballSavantExtractor) scrapeFieldingBoxScore(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := baseballsavantmlb.NewFieldingBoxScoreScraper()
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.FieldingBoxScore]{
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

func (e *BaseballSavantExtractor) scrapePlayByPlay(ctx context.Context) error {
	m, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	eventdatascraper := baseballsavantmlb.NewPlayByPlayScraper()
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
