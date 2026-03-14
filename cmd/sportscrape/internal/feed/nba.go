package feed

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/dataprovider/nba"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/runner"
)

var (
	NBAConcurrencyOptions string = strings.Join([]string{
		"'live-box-score'",
		"'advanced-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'traditional-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'scoring-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'usage-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'misc-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'four-factors-box-score[-q1|-q2|-q3|-q4|-h1|-h2|-ot]'",
		"'hustle-box-score'",
		"'matchups-box-score'",
		"'defense-box-score'",
		"'tracking-box-score'",
		"'play-by-play'",
	}, ", ")
	NBAOptions string = fmt.Sprintf("'matchup', 'matchup-periods', %s", NBAConcurrencyOptions)
	ErrNBA     error  = fmt.Errorf("invalid nba.com feed, valid options: %s", NBAOptions)
)

type NBAExtractor struct {
	Feed           string
	Date           string
	Timeout        time.Duration
	Concurrency    int
	OutputPath     string
	Format         string
	S3Config       exporters.S3Config
	ParquetOptions []exporters.ParquetConfigOption
	matchupScraper *nba.MatchupScraper
}

func (e *NBAExtractor) ValidateFeed() error {
	if err := exporters.ValidateFormat(e.Format); err != nil {
		return err
	}
	switch e.Feed {
	case "matchup", "matchup-periods",
		"live-box-score",
		"advanced-box-score", "advanced-box-score-q1", "advanced-box-score-q2", "advanced-box-score-q3", "advanced-box-score-q4", "advanced-box-score-h1", "advanced-box-score-h2", "advanced-box-score-ot",
		"traditional-box-score", "traditional-box-score-q1", "traditional-box-score-q2", "traditional-box-score-q3", "traditional-box-score-q4", "traditional-box-score-h1", "traditional-box-score-h2", "traditional-box-score-ot",
		"scoring-box-score", "scoring-box-score-q1", "scoring-box-score-q2", "scoring-box-score-q3", "scoring-box-score-q4", "scoring-box-score-h1", "scoring-box-score-h2", "scoring-box-score-ot",
		"usage-box-score", "usage-box-score-q1", "usage-box-score-q2", "usage-box-score-q3", "usage-box-score-q4", "usage-box-score-h1", "usage-box-score-h2", "usage-box-score-ot",
		"misc-box-score", "misc-box-score-q1", "misc-box-score-q2", "misc-box-score-q3", "misc-box-score-q4", "misc-box-score-h1", "misc-box-score-h2", "misc-box-score-ot",
		"four-factors-box-score", "four-factors-box-score-q1", "four-factors-box-score-q2", "four-factors-box-score-q3", "four-factors-box-score-q4", "four-factors-box-score-h1", "four-factors-box-score-h2", "four-factors-box-score-ot",
		"hustle-box-score", "matchups-box-score", "defense-box-score", "tracking-box-score",
		"play-by-play":
		return nil
	default:
		return ErrNBA
	}
}

func (e *NBAExtractor) Scrape(ctx context.Context) error {
	switch e.Feed {
	case "matchup":
		return e.scrapeMatchup(ctx)
	case "matchup-periods":
		return e.scrapeMatchupPeriods(ctx)
	case "live-box-score":
		return e.scrapeLiveBoxScore(ctx)
	case "hustle-box-score":
		return e.scrapeHustleBoxScore(ctx)
	case "matchups-box-score":
		return e.scrapeMatchupsBoxScore(ctx)
	case "defense-box-score":
		return e.scrapeDefenseBoxScore(ctx)
	case "tracking-box-score":
		return e.scrapeTrackingBoxScore(ctx)
	case "play-by-play":
		return e.scrapePlayByPlay(ctx)
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
	default:
		return fmt.Errorf("%w: %q", ErrUnsupportedFeed, e.Feed)
	}
}

// period derives the nba.Period from the feed suffix.
func (e *NBAExtractor) period() nba.Period {
	switch {
	case strings.HasSuffix(e.Feed, "-q1"):
		return nba.Q1
	case strings.HasSuffix(e.Feed, "-q2"):
		return nba.Q2
	case strings.HasSuffix(e.Feed, "-q3"):
		return nba.Q3
	case strings.HasSuffix(e.Feed, "-q4"):
		return nba.Q4
	case strings.HasSuffix(e.Feed, "-h1"):
		return nba.H1
	case strings.HasSuffix(e.Feed, "-h2"):
		return nba.H2
	case strings.HasSuffix(e.Feed, "-ot"):
		return nba.AllOT
	default:
		return nba.Full
	}
}

func (e *NBAExtractor) retrieveMatchup(keepAlive bool) ([]model.Matchup, error) {
	scraper := nba.NewMatchupScraper(
		nba.WithMatchupDate(e.Date),
		nba.WithMatchupTimeout(e.Timeout),
	)
	scraper.NetworkHeaders = nba.NetworkHeaders
	m, err := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Matchup]{
			Scraper:   scraper,
			KeepAlive: keepAlive,
		},
	).Run()
	if err != nil {
		scraper.Close()
		return m, err
	}
	e.matchupScraper = scraper
	return m, err
}

func (e *NBAExtractor) scrapeMatchup(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(false)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, matchups, e.ParquetOptions...)
}

func (e *NBAExtractor) scrapeMatchupPeriods(ctx context.Context) error {
	matchups, err := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.MatchupPeriods]{
			Scraper: nba.NewMatchupPeriodsScraper(
				nba.WithMatchupPeriodsDate(e.Date),
				nba.WithMatchupPeriodsTimeout(e.Timeout),
			),
		},
	).Run()
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, matchups, e.ParquetOptions...)
}

func (e *NBAExtractor) scrapeLiveBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreLiveScraper(nba.WithBoxScoreLiveTimeout(e.Timeout))
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreLive]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *NBAExtractor) scrapeAdvancedBoxScore(ctx context.Context, period nba.Period) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreAdvancedScraper(
		nba.WithBoxScoreAdvancedPeriod(period),
		nba.WithBoxScoreAdvancedTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
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

func (e *NBAExtractor) scrapeTraditionalBoxScore(ctx context.Context, period nba.Period) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreTraditionalScraper(
		nba.WithBoxScoreTraditionalPeriod(period),
		nba.WithBoxScoreTraditionalTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
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

func (e *NBAExtractor) scrapeScoringBoxScore(ctx context.Context, period nba.Period) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreScoringScraper(
		nba.WithBoxScoreScoringPeriod(period),
		nba.WithBoxScoreScoringTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
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

func (e *NBAExtractor) scrapeUsageBoxScore(ctx context.Context, period nba.Period) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreUsageScraper(
		nba.WithBoxScoreUsagePeriod(period),
		nba.WithBoxScoreUsageTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
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

func (e *NBAExtractor) scrapeMiscBoxScore(ctx context.Context, period nba.Period) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreMiscScraper(
		nba.WithBoxScoreMiscPeriod(period),
		nba.WithBoxScoreMiscTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
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

func (e *NBAExtractor) scrapeFourFactorsBoxScore(ctx context.Context, period nba.Period) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreFourFactorsScraper(
		nba.WithBoxScoreFourFactorsPeriod(period),
		nba.WithBoxScoreFourFactorsTimeout(e.Timeout),
	)
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
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

func (e *NBAExtractor) scrapeHustleBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreHustleScraper(nba.WithBoxScoreHustleTimeout(e.Timeout))
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreHustle]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *NBAExtractor) scrapeMatchupsBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreMatchupsScraper(nba.WithBoxScoreMatchupsTimeout(e.Timeout))
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreMatchups]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *NBAExtractor) scrapeDefenseBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreDefenseScraper(nba.WithBoxScoreDefenseTimeout(e.Timeout))
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreDefense]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *NBAExtractor) scrapeTrackingBoxScore(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewBoxScoreTrackingScraper(nba.WithBoxScoreTrackingTimeout(e.Timeout))
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
	records, err := runner.NewEventDataRunner(
		runner.EventDataRunnerConfig[model.Matchup, model.BoxScoreTracking]{
			Concurrency: e.Concurrency,
			Scraper:     scraper,
		},
	).Run(matchups)
	if err != nil {
		return err
	}
	return exporters.BuildAndWrite(ctx, e.OutputPath, e.Format, e.S3Config, records, e.ParquetOptions...)
}

func (e *NBAExtractor) scrapePlayByPlay(ctx context.Context) error {
	matchups, err := e.retrieveMatchup(true)
	if err != nil {
		return err
	}
	scraper := nba.NewPlayByPlayScraper(nba.WithPlayByPlayTimeout(e.Timeout))
	scraper.DocumentRetriever = e.matchupScraper.DocumentRetriever
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
