package feed

import (
	"context"
	"fmt"
	"time"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	SRNBAConcurrencyOptions string = "'box-score-stats', 'q1-box-score-stats', 'q2-box-score-stats', 'q3-box-score-stats', 'q4-box-score-stats', 'h1-box-score-stats', 'h2-box-score-stats', 'adv-box-score-stats'"
	SRNBAOptions            string = fmt.Sprintf("'matchup', %s", SRNBAConcurrencyOptions)
	ErrSRNBA                error  = fmt.Errorf("invalid nba feed, valid options: %s", SRNBAOptions)
)

type SportsReferenceExtractor struct {
	Feed              string
	Date              string
	Timeout           time.Duration
	Concurrency       int
	OutputPath        string
	Format            string
	S3Config          exporters.S3Config
	ParquetOptions    []exporters.ParquetConfigOption
	nbaMatchupScraper *basketballreferencenba.MatchupScraper
}

func (e *SportsReferenceExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "nba-matchup",
		"nba-box-score-stats",
		"nba-q1-box-score-stats",
		"nba-q2-box-score-stats",
		"nba-q3-box-score-stats",
		"nba-q4-box-score-stats",
		"nba-h1-box-score-stats",
		"nba-h2-box-score-stats",
		"nba-adv-box-score-stats":
		return nil
	default:
		return ErrSRNBA
	}
}

func (e *SportsReferenceExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "nba-matchup":
		return e.scrapeMatchup(ctx)
	case "nba-box-score-stats":
		return e.scrapeBasicBoxScore(ctx, basketballreferencenba.Full)
	case "nba-q1-box-score-stats":
		return e.scrapeBasicBoxScore(ctx, basketballreferencenba.Q1)
	case "nba-q2-box-score-stats":
		return e.scrapeBasicBoxScore(ctx, basketballreferencenba.Q2)
	case "nba-q3-box-score-stats":
		return e.scrapeBasicBoxScore(ctx, basketballreferencenba.Q3)
	case "nba-q4-box-score-stats":
		return e.scrapeBasicBoxScore(ctx, basketballreferencenba.Q4)
	case "nba-h1-box-score-stats":
		return e.scrapeBasicBoxScore(ctx, basketballreferencenba.H1)
	case "nba-h2-box-score-stats":
		return e.scrapeBasicBoxScore(ctx, basketballreferencenba.H2)
	case "nba-adv-box-score-stats":
		return e.scrapeAdvBoxScore(ctx)
	default:
		return fmt.Errorf("%w: %q", ErrUnsupportedFeed, e.Feed)
	}
}

func (e *SportsReferenceExtractor) retrieveMatchup(close bool) ([]model.NBAMatchup, error) {
	scraper := basketballreferencenba.NewMatchupScraper(
		basketballreferencenba.WithMatchupDate(e.Date),
		basketballreferencenba.WithMatchupTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = basketballreferencenba.NetworkHeaders

	m, err := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.NBAMatchup]{
			Scraper: scraper,
			Close:   close,
		},
	).Run()
	if err != nil {
		scraper.Close()
		return m, err
	}
	e.nbaMatchupScraper = scraper
	return m, err
}

func (e *SportsReferenceExtractor) scrapeMatchup(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, matchups, e.ParquetOptions...)
}

func (e *SportsReferenceExtractor) scrapeBasicBoxScore(ctx context.Context, period basketballreferencenba.Period) error {
	matchups, err := e.retrieveMatchup(false)
	if err != nil {
		return err
	}
	scraper := basketballreferencenba.NewBasicBoxScoreScraper(
		basketballreferencenba.WithBasicBoxScorePeriod(period),
		basketballreferencenba.WithBasicBoxScoreTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.nbaMatchupScraper.DocumentRetriever

	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.NBAMatchup, model.NBABasicBoxScoreStats]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *SportsReferenceExtractor) scrapeAdvBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(false)
	if err != nil {
		return err
	}
	scraper := basketballreferencenba.NewAdvBoxScoreScraper(
		basketballreferencenba.WithAdvBoxScoreTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.nbaMatchupScraper.DocumentRetriever
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.NBAMatchup, model.NBAAdvBoxScoreStats]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}
