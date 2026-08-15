package feed

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	WNBAConcurrencyOptions string = strings.Join([]string{
		"'advanced-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'traditional-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'scoring-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'usage-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'misc-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'four-factors-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'play-by-play'",
	}, ", ")
	WNBAOptions string = fmt.Sprintf("'matchup', 'matchup-periods', %s", WNBAConcurrencyOptions)
	ErrWNBA     error  = fmt.Errorf("invalid wnba.com feed, valid options: %s", WNBAOptions)
)

type WNBAExtractor struct {
	Feed           string
	Date           string
	EndDate        string
	Timeout        time.Duration
	Concurrency    int
	OutputPath     string
	Format         string
	S3Config       exporters.S3Config
	ParquetOptions []exporters.ParquetConfigOption
}

func (e *WNBAExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "matchup", "matchup-periods",
		"advanced-box-score", "advanced-box-score-q1", "advanced-box-score-q2", "advanced-box-score-q3", "advanced-box-score-q4", "advanced-box-score-h1", "advanced-box-score-h2", "advanced-box-score-ot",
		"traditional-box-score", "traditional-box-score-q1", "traditional-box-score-q2", "traditional-box-score-q3", "traditional-box-score-q4", "traditional-box-score-h1", "traditional-box-score-h2", "traditional-box-score-ot",
		"scoring-box-score", "scoring-box-score-q1", "scoring-box-score-q2", "scoring-box-score-q3", "scoring-box-score-q4", "scoring-box-score-h1", "scoring-box-score-h2", "scoring-box-score-ot",
		"usage-box-score", "usage-box-score-q1", "usage-box-score-q2", "usage-box-score-q3", "usage-box-score-q4", "usage-box-score-h1", "usage-box-score-h2", "usage-box-score-ot",
		"misc-box-score", "misc-box-score-q1", "misc-box-score-q2", "misc-box-score-q3", "misc-box-score-q4", "misc-box-score-h1", "misc-box-score-h2", "misc-box-score-ot",
		"four-factors-box-score", "four-factors-box-score-q1", "four-factors-box-score-q2", "four-factors-box-score-q3", "four-factors-box-score-q4", "four-factors-box-score-h1", "four-factors-box-score-h2", "four-factors-box-score-ot",
		"play-by-play":
		return nil
	default:
		return ErrWNBA
	}
}

func (e *WNBAExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "matchup":
		return e.scrapeMatchup(ctx)
	case "matchup-periods":
		return e.scrapeMatchupPeriods(ctx)
	case "advanced-box-score", "advanced-box-score-q1", "advanced-box-score-q2", "advanced-box-score-q3", "advanced-box-score-q4", "advanced-box-score-h1", "advanced-box-score-h2", "advanced-box-score-ot":
		return e.scrapeAdvancedBoxScore(ctx, e.period())
	case "traditional-box-score", "traditional-box-score-q1", "traditional-box-score-q2", "traditional-box-score-q3", "traditional-box-score-q4", "traditional-box-score-h1", "traditional-box-score-h2", "traditional-box-score-ot":
		return e.scrapeTraditionalBoxScore(ctx, e.period())
	case "scoring-box-score", "scoring-box-score-q1", "scoring-box-score-q2", "scoring-box-score-q3", "scoring-box-score-q4", "scoring-box-score-h1", "scoring-box-score-h2", "scoring-box-score-ot":
		return e.scrapeScoringBoxScore(ctx, e.period())
	case "usage-box-score", "usage-box-score-q1", "usage-box-score-q2", "usage-box-score-q3", "usage-box-score-q4", "usage-box-score-h1", "usage-box-score-h2", "usage-box-score-ot":
		return e.scrapeUsageBoxScore(ctx, e.period())
	case "misc-box-score", "misc-box-score-q1", "misc-box-score-q2", "misc-box-score-q3", "misc-box-score-q4", "misc-box-score-h1", "misc-box-score-h2", "misc-box-score-ot":
		return e.scrapeMiscBoxScore(ctx, e.period())
	case "four-factors-box-score", "four-factors-box-score-q1", "four-factors-box-score-q2", "four-factors-box-score-q3", "four-factors-box-score-q4", "four-factors-box-score-h1", "four-factors-box-score-h2", "four-factors-box-score-ot":
		return e.scrapeFourFactorsBoxScore(ctx, e.period())
	case "play-by-play":
		return e.scrapePlayByPlay(ctx)
	default:
		return fmt.Errorf("%w: %q", ErrUnsupportedFeed, e.Feed)
	}
}

// period derives the wnba.Period from the feed suffix.
func (e *WNBAExtractor) period() wnba.Period {
	switch {
	case strings.HasSuffix(e.Feed, "-q1"):
		return wnba.Q1
	case strings.HasSuffix(e.Feed, "-q2"):
		return wnba.Q2
	case strings.HasSuffix(e.Feed, "-q3"):
		return wnba.Q3
	case strings.HasSuffix(e.Feed, "-q4"):
		return wnba.Q4
	case strings.HasSuffix(e.Feed, "-h1"):
		return wnba.H1
	case strings.HasSuffix(e.Feed, "-h2"):
		return wnba.H2
	case strings.HasSuffix(e.Feed, "-ot"):
		return wnba.AllOT
	default:
		return wnba.Full
	}
}

// retrieveMatchup fetches matchups via the JSON schedule API. Unlike NBA's
// retrieveMatchup, there is no browser session to keep alive or share with
// the (chromedp-based) box score scrapers - the matchup scraper is plain
// HTTP JSON, so each box score scrape independently launches and tears down
// its own chromedp session via its normal Init()/Close() lifecycle.
func (e *WNBAExtractor) retrieveMatchup() ([]model.Matchup, error) {
	scraper := wnba.NewMatchupScraper(
		wnba.WithMatchupDate(e.Date),
		wnba.WithMatchupEndDate(e.EndDate),
	)
	return runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper: scraper,
		},
	).Run()
}

func (e *WNBAExtractor) scrapeMatchup(ctx context.Context) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, matchups, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapeMatchupPeriods(ctx context.Context) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewMatchupPeriodsScraper(wnba.WithMatchupPeriodsTimeout(e.Timeout))
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.MatchupPeriods]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapeAdvancedBoxScore(ctx context.Context, period wnba.Period) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewBoxScoreAdvancedScraper(
		wnba.WithBoxScoreAdvancedPeriod(period),
		wnba.WithBoxScoreAdvancedTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreAdvanced]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapeTraditionalBoxScore(ctx context.Context, period wnba.Period) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewBoxScoreTraditionalScraper(
		wnba.WithBoxScoreTraditionalPeriod(period),
		wnba.WithBoxScoreTraditionalTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreTraditional]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapeScoringBoxScore(ctx context.Context, period wnba.Period) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewBoxScoreScoringScraper(
		wnba.WithBoxScoreScoringPeriod(period),
		wnba.WithBoxScoreScoringTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreScoring]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapeUsageBoxScore(ctx context.Context, period wnba.Period) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewBoxScoreUsageScraper(
		wnba.WithBoxScoreUsagePeriod(period),
		wnba.WithBoxScoreUsageTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreUsage]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapeMiscBoxScore(ctx context.Context, period wnba.Period) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewBoxScoreMiscScraper(
		wnba.WithBoxScoreMiscPeriod(period),
		wnba.WithBoxScoreMiscTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreMisc]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapeFourFactorsBoxScore(ctx context.Context, period wnba.Period) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewBoxScoreFourFactorsScraper(
		wnba.WithBoxScoreFourFactorsPeriod(period),
		wnba.WithBoxScoreFourFactorsTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreFourFactors]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *WNBAExtractor) scrapePlayByPlay(ctx context.Context) error {
	matchups, err := e.retrieveMatchup()
	if err != nil {
		return err
	}
	scraper := wnba.NewPlayByPlayScraper(wnba.WithPlayByPlayTimeout(e.Timeout))
	scraper.NetworkHeaders = wnba.NetworkHeaders
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.PlayByPlay]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}
