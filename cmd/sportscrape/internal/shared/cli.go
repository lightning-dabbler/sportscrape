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

	// --concurrency
	concurrency, err := cmd.Flags().GetInt("concurrency")
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

	var date, year, feedstring string
	var timeoutDuration time.Duration

	switch provider {
	case "espn":
		switch league {
		case "ufc":
			// --year
			year, err = cmd.Flags().GetString("year")
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
	}

	switch provider {
	case "sportsreference", "espn", "nba":
		// --timeout
		timeout, err := cmd.Flags().GetInt("timeout")
		if err != nil {
			return err
		}
		timeoutDuration = time.Duration(timeout) * time.Second
	}

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
			Concurrency:    concurrency,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
		}
	case "sportsreference":
		e = &feed.SportsReferenceExtractor{
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
			Concurrency:    concurrency,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
		}
	case "espn":
		e = &feed.ESPNMMAExtractor{
			Feed:           feedstring,
			Year:           year,
			Timeout:        timeoutDuration,
			Concurrency:    concurrency,
			OutputPath:     destination,
			Format:         fileFormat,
			S3Config:       s3config,
			ParquetOptions: parquetOptions,
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
