package shared

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/exporters"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"

	"github.com/spf13/cobra"
	"github.com/xitongsys/parquet-go/parquet"
)

func Run(cmd *cobra.Command, provider, league string) error {
	start := time.Now().UTC()
	// --destination
	destination, err := cmd.Flags().GetString("destination")
	if err != nil {
		return err
	}
	if destination == "" {
		return fmt.Errorf("--destination is required and cannot be empty")
	}
	parsedDestination, err := url.Parse(destination)
	if err != nil {
		return err
	}
	err = exporters.SupportedDestination(parsedDestination)
	if err != nil {
		return err
	}

	// --file-format
	fileFormat, err := cmd.Flags().GetString("file-format")
	if err != nil {
		return err
	}

	// --parquet-compression
	parquetCompression, err := cmd.Flags().GetString("parquet-compression")
	if err != nil {
		return err
	}
	compression, err := parquet.CompressionCodecFromString(parquetCompression)
	if err != nil {
		return err
	}

	// --parquet-row-group-size
	parquetRowGroupSize, err := cmd.Flags().GetInt64("parquet-row-group-size")
	if err != nil {
		return err
	}

	// --parquet-page-size
	parquetPageSize, err := cmd.Flags().GetInt64("parquet-page-size")
	if err != nil {
		return err
	}

	// --parquet-write-parallelism
	parquetWriteParallelism, err := cmd.Flags().GetInt64("parquet-write-parallelism")
	if err != nil {
		return err
	}

	// --feed
	rawFeed, err := cmd.Flags().GetString("feed")
	if err != nil {
		return err
	}

	// s3 arguments

	// aws-region
	awsRegion, err := cmd.Flags().GetString("aws-region")
	if err != nil {
		return err
	}
	// aws-endpoint
	awsEndpoint, err := cmd.Flags().GetString("aws-endpoint")
	if err != nil {
		return err
	}

	var date, year, feedstring, endDate string
	var timeoutDuration time.Duration
	var fetchAttempts int
	var fetchRetryBackoff time.Duration
	var concurrency int

	switch provider {
	case "foxsports", "baseballsavant", "espn", "nba", "wnba", "nhl":
		// --concurrency
		concurrency, err = cmd.Flags().GetInt("concurrency")
		if err != nil {
			return err
		}
	}

	switch provider {
	case "espn":
		switch league {
		case "ufc":
			// --year
			year, err = cmd.Flags().GetString("year")
			if err != nil {
				return err
			}
			// --fetch-attempts, --fetch-retry-backoff
			fetchAttempts, fetchRetryBackoff, err = fetchRetryFlags(cmd)
			if err != nil {
				return err
			}
		}
	default:
		// --date
		date, err = cmd.Flags().GetString("date")
		if err != nil {
			return err
		}
		if provider == "nhl" {
			// --fetch-attempts, --fetch-retry-backoff
			fetchAttempts, fetchRetryBackoff, err = fetchRetryFlags(cmd)
			if err != nil {
				return err
			}
		}
		if provider == "wnba" {
			// --end-date
			endDate, err = cmd.Flags().GetString("end-date")
			if err != nil {
				return err
			}
		}
	}

	// --timeout
	timeout, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		return err
	}
	if timeout <= 0 {
		return fmt.Errorf("--timeout must be greater than 0, got %d", timeout)
	}
	timeoutDuration = time.Duration(timeout) * time.Second

	switch league {
	case "mlb":
		feedstring = "mlb-" + rawFeed
	case "nba":
		feedstring = "nba-" + rawFeed
	case "wnba":
		feedstring = "wnba-" + rawFeed
	case "ufc":
		feedstring = "ufc-" + rawFeed
	default:
		feedstring = rawFeed
	}

	// parquet options

	parquetOptions := []exporters.ParquetConfigOption{
		exporters.WithCompressionType(compression),
		exporters.WithRowGroupSize(parquetRowGroupSize),
		exporters.WithPageSize(parquetPageSize),
		exporters.WithParallelism(parquetWriteParallelism),
	}

	if fileFormat == "parquet" {
		slog.Debug("Parquet config", "compression_type", parquetCompression, "row_group_size", parquetRowGroupSize, "page_size", parquetPageSize, "write_parallelism", parquetWriteParallelism)
	}

	s3config := exporters.S3Config{
		Endpoint: awsEndpoint,
		Region:   awsRegion,
	}
	var e feed.ProviderExtractor
	switch provider {
	case "foxsports":
		e = &feed.FoxSportsExtractor{
			Feed:           feedstring,
			Date:           date,
			Timeout:        timeoutDuration,
			Concurrency:    concurrency,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
		}
	case "baseballsavant":
		e = &feed.BaseballSavantExtractor{
			Feed:           feedstring,
			Date:           date,
			Timeout:        timeoutDuration,
			Concurrency:    concurrency,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
		}
	case "nhl":
		e = &feed.NHLExtractor{
			Feed:              feedstring,
			Date:              date,
			FetchAttempts:     fetchAttempts,
			FetchRetryBackoff: fetchRetryBackoff,
			Timeout:           timeoutDuration,
			Concurrency:       concurrency,
			OutputPath:        destination,
			Format:            fileFormat,
			S3Config:          s3config,
			ParquetOptions:    parquetOptions,
		}
	case "propfinder":
		e = &feed.PropFinderExtractor{
			Feed:           feedstring,
			Date:           date,
			Timeout:        timeoutDuration,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
		}
	case "espn":
		e = &feed.ESPNMMAExtractor{
			Feed:              feedstring,
			Year:              year,
			Timeout:           timeoutDuration,
			FetchAttempts:     fetchAttempts,
			FetchRetryBackoff: fetchRetryBackoff,
			Concurrency:       concurrency,
			OutputPath:        destination,
			Format:            fileFormat,
			S3Config:          s3config,
			ParquetOptions:    parquetOptions,
		}
	case "nba":
		e = &feed.NBAExtractor{
			Feed:           feedstring,
			Date:           date,
			Timeout:        timeoutDuration,
			Concurrency:    concurrency,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
		}
	case "wnba":
		e = &feed.WNBAExtractor{
			Feed:           feedstring,
			Date:           date,
			EndDate:        endDate,
			Timeout:        timeoutDuration,
			Concurrency:    concurrency,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
		}
	default:
		return fmt.Errorf("unsupported provider %s", provider)
	}
	err = e.ValidateFeed()
	if err != nil {
		return err
	}

	err = e.Scrape(context.TODO())
	if err != nil {
		return err
	}

	diff := time.Now().UTC().Sub(start)
	slog.Info("Data extraction complete", "duration", diff)
	return nil
}

// fetchRetryFlags reads and validates --fetch-attempts (must be greater than 0) and
// --fetch-retry-backoff (seconds, must be 0 or greater; 0 retries without a delay).
func fetchRetryFlags(cmd *cobra.Command) (int, time.Duration, error) {
	attempts, err := cmd.Flags().GetInt("fetch-attempts")
	if err != nil {
		return 0, 0, err
	}
	if attempts <= 0 {
		return 0, 0, fmt.Errorf("--fetch-attempts must be greater than 0, got %d", attempts)
	}
	backoffSeconds, err := cmd.Flags().GetInt("fetch-retry-backoff")
	if err != nil {
		return 0, 0, err
	}
	if backoffSeconds < 0 {
		return 0, 0, fmt.Errorf("--fetch-retry-backoff must be 0 or greater, got %d", backoffSeconds)
	}
	return attempts, time.Duration(backoffSeconds) * time.Second, nil
}
