package feed

import (
	"context"
	"fmt"
	"strings"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/lightning-dabbler/sportscrape/util"
)

var (
	FSMLBConcurrencyOptions string = "'pitching-box-score', 'batting-box-score', 'probable-starting-pitcher', 'odds-total', 'odds-money-line'"
	FSMLBOptions            string = fmt.Sprintf("'matchup', %s", FSMLBConcurrencyOptions)
	FSNBAConcurrencyOptions string = "'box-score-stats'"
	FSNBAOptions            string = fmt.Sprintf("'matchup', %s", FSNBAConcurrencyOptions)
	ErrFSMLB                error  = fmt.Errorf("invalid mlb feed, valid options: %s", FSMLBOptions)
	ErrFSNBA                error  = fmt.Errorf("invalid nba feed, valid options: %s", FSNBAOptions)
	ErrFSWNBA               error  = fmt.Errorf("invalid wnba feed, valid options: %s", FSNBAOptions)
	ErrFSNFLDateFmt         error  = fmt.Errorf("date for NFL feed should be of the form '{year}-{week}-{seasontype}' e.g. 2024-4-2")
)

type FoxSportsExtractor struct {
	Feed           string
	Date           string
	Concurrency    int
	OutputPath     string
	Format         string
	S3Config       exporters.S3Config
	ParquetOptions []exporters.ParquetConfigOption
}

func (e *FoxSportsExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "mlb-matchup", "mlb-pitching-box-score", "mlb-batting-box-score",
		"mlb-probable-starting-pitcher", "mlb-odds-total", "mlb-odds-money-line",
		"nba-matchup", "nba-box-score-stats",
		"wnba-matchup", "wnba-box-score-stats",
		"ncaab-matchup", "nfl-matchup":
		return nil
	default:
		switch {
		case strings.HasPrefix(e.Feed, "mlb-"):
			return ErrFSMLB
		case strings.HasPrefix(e.Feed, "wnba-"):
			return ErrFSWNBA
		default:
			return ErrFSNBA
		}
	}
}

func (e *FoxSportsExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "mlb-matchup":
		return e.scrapeMatchup(ctx, foxsports.MLB)
	case "nba-matchup":
		return e.scrapeMatchup(ctx, foxsports.NBA)
	case "wnba-matchup":
		return e.scrapeMatchup(ctx, foxsports.WNBA)
	case "ncaab-matchup":
		return e.scrapeMatchup(ctx, foxsports.NCAAB)
	case "nfl-matchup":
		return e.scrapeNFLMatchup(ctx)
	case "nba-box-score-stats":
		return e.scrapeNBABoxScore(ctx, foxsports.NBA)
	case "wnba-box-score-stats":
		return e.scrapeNBABoxScore(ctx, foxsports.WNBA)
	case "mlb-batting-box-score":
		return e.scrapeMLBBattingBoxScore(ctx)
	case "mlb-pitching-box-score":
		return e.scrapeMLBPitchingBoxScore(ctx)
	case "mlb-probable-starting-pitcher":
		return e.scrapeMLBProbableStartingPitcher(ctx)
	case "mlb-odds-total":
		return e.scrapeMLBOddsTotal(ctx)
	case "mlb-odds-money-line":
		return e.scrapeMLBOddsMoneyLine(ctx)
	default:
		return fmt.Errorf("unsupported feed %q. %w", e.Feed, ErrUnsupportedFeed)
	}
}

func (e *FoxSportsExtractor) retrieveMatchup(league foxsports.League) ([]model.Matchup, error) {
	return runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: &foxsports.MatchupScraper{
				League:    league,
				Segmenter: &foxsports.GeneralSegmenter{Date: e.Date},
			},
		},
	).Run()
}

func (e *FoxSportsExtractor) scrapeMatchup(ctx context.Context, league foxsports.League) error {
	matchups, err := e.retrieveMatchup(league)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, matchups, e.ParquetOptions...)
}

func (e *FoxSportsExtractor) scrapeNFLMatchup(ctx context.Context) error {
	parts := strings.Split(e.Date, "-")
	if len(parts) != 3 {
		return ErrFSNFLDateFmt
	}
	year, err := util.TextToInt32(parts[0])
	if err != nil {
		return err
	}
	week, err := util.TextToInt32(parts[1])
	if err != nil {
		return err
	}
	seasontype, err := util.TextToInt(parts[2])
	if err != nil {
		return err
	}
	matchups, err := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: &foxsports.MatchupScraper{
				League: foxsports.NFL,
				Segmenter: &foxsports.NFLSegmenter{
					Season: foxsports.SeasonType(seasontype),
					Week:   week,
					Year:   year,
				},
			},
		},
	).Run()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, matchups, e.ParquetOptions...)
}

func (e *FoxSportsExtractor) scrapeNBABoxScore(ctx context.Context, league foxsports.League) error {
	matchups, err := e.retrieveMatchup(league)
	if err != nil {
		return err
	}
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.NBABoxScoreStats]{
			Concurrency: e.Concurrency,
			Scraper:     foxsports.NewNBABoxScoreScraper(foxsports.NBABoxScoreScraperLeague(league)),
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *FoxSportsExtractor) scrapeMLBBattingBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(foxsports.MLB)
	if err != nil {
		return err
	}
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MLBBattingBoxScoreStats]{
			Concurrency: e.Concurrency,
			Scraper:     foxsports.NewMLBBattingBoxScoreScraper(),
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *FoxSportsExtractor) scrapeMLBPitchingBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(foxsports.MLB)
	if err != nil {
		return err
	}
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MLBPitchingBoxScoreStats]{
			Concurrency: e.Concurrency,
			Scraper:     foxsports.NewMLBPitchingBoxScoreScraper(),
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *FoxSportsExtractor) scrapeMLBProbableStartingPitcher(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(foxsports.MLB)
	if err != nil {
		return err
	}
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MLBProbableStartingPitcher]{
			Concurrency: e.Concurrency,
			Scraper:     foxsports.NewMLBProbableStartingPitcherScraper(),
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *FoxSportsExtractor) scrapeMLBOddsTotal(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(foxsports.MLB)
	if err != nil {
		return err
	}
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MLBOddsTotal]{
			Concurrency: e.Concurrency,
			Scraper:     foxsports.NewMLBOddsTotalScraper(),
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *FoxSportsExtractor) scrapeMLBOddsMoneyLine(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(foxsports.MLB)
	if err != nil {
		return err
	}
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MLBOddsMoneyLine]{
			Concurrency: e.Concurrency,
			Scraper:     foxsports.NewMLBOddsMoneyLineScraper(),
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}
