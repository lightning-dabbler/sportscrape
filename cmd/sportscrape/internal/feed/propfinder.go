package feed

import (
	"context"
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/mlb"
	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/mlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	PropFinderMLBOptions string = "'weather'"
	ErrPropFinderMLB     error  = fmt.Errorf("valid options: %s", PropFinderMLBOptions)
)

type PropFinderExtractor struct {
	Feed           string
	Date           string
	OutputPath     string
	Format         string
	S3Config       exporters.S3Config
	ParquetOptions []exporters.ParquetConfigOption
}

func (e *PropFinderExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "mlb-weather":
		return nil
	default:
		return fmt.Errorf("unsupported feed %q for prop finder. %w", e.Feed, ErrPropFinderMLB)
	}
}

func (e *PropFinderExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "mlb-weather":
		return e.scrapeMLBWeather(ctx)
	default:
		return fmt.Errorf("unsupported feed %q for prop finder. %w", e.Feed, ErrPropFinderMLB)
	}
}

func (e *PropFinderExtractor) scrapeMLBWeather(ctx context.Context) error {
	weatherrunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Weather]{
			Scraper: mlb.NewWeatherScraper(
				mlb.WeatherScraperDate(e.Date),
			),
		},
	)
	m, err := weatherrunner.Run()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, m, e.ParquetOptions...)
}
