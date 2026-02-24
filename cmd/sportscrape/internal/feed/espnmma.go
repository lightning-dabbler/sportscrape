package feed

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma"
	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	ESPNMMAConcurrencyOptions string = "'fight-details'"
	ESPNMMAOptions            string = fmt.Sprintf("'matchups', %s", ESPNMMAConcurrencyOptions)
	ErrESPNUFC                error  = fmt.Errorf("invalid ufc feed, valid options: %s", ESPNMMAOptions)
)

type ESPNMMAExtractor struct {
	Feed           string
	Year           string
	Timeout        time.Duration
	Concurrency    int
	OutputPath     string
	Format         string
	S3Config       exporters.S3Config
	ParquetOptions []exporters.ParquetConfigOption
}

func (e *ESPNMMAExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "ufc-matchups", "ufc-fight-details":
		return nil
	default:
		switch {
		case strings.HasPrefix(e.Feed, "ufc-"):
			return ErrESPNUFC
		}
		return fmt.Errorf("%w: %q", ErrUnsupportedFeed, e.Feed)
	}
}

func (e *ESPNMMAExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "ufc-matchups":
		return e.scrapeMatchup(ctx, "ufc")
	case "ufc-fight-details":
		return e.scrapeFightDetails(ctx, "ufc")
	default:
		return fmt.Errorf("%w: %q", ErrUnsupportedFeed, e.Feed)
	}
}

func (e *ESPNMMAExtractor) retrieveMatchup(league string) ([]model.Matchup, error) {
	matchupscraper := mma.ESPNMMAMatchupScraper{}
	matchupscraper.Timeout = e.Timeout
	matchupscraper.League = league
	matchupscraper.Year = e.Year

	matchuprunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: matchupscraper,
		},
	)

	return matchuprunner.Run()
}

func (e *ESPNMMAExtractor) scrapeMatchup(ctx context.Context, league string) error {
	m, err := e.retrieveMatchup(league)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, m, e.ParquetOptions...)
}

func (e *ESPNMMAExtractor) scrapeFightDetails(ctx context.Context, league string) error {
	m, err := e.retrieveMatchup(league)
	if err != nil {
		return err
	}
	fightdetailsscraper := mma.ESPNMMAFightDetailsScraper{}
	fightdetailsscraper.Timeout = e.Timeout
	fightdetailsscraper.League = league
	eventrunner := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.FightDetails]{
			Concurrency: e.Concurrency,
			Scraper:     fightdetailsscraper,
		},
	)
	records, err := eventrunner.Run(m)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}
