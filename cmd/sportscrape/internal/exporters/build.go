package exporters

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
)

type S3Config struct {
	Endpoint string
	Region   string
}

var (
	ErrUnsupportedDestination = fmt.Errorf("unsupported destination")
	ErrUnsupportedFormat      = fmt.Errorf("unsupported format")
)

func ValidateFormat(format string) error {
	switch format {
	case "jsonl", "parquet":
		return nil
	default:
		return fmt.Errorf("%w: %q", ErrUnsupportedFormat, format)
	}
}

// Build returns the appropriate Exporter based on the outputPath scheme.
// Local paths (no scheme or file://) return a LocalExporter.
// s3:// paths return an S3Exporter.
func Build[E any](outputPath string, s3cfg S3Config, parquetOpts ...ParquetConfigOption) (Exporter[E], error) {
	destination, err := url.Parse(outputPath)
	if err != nil {
		return nil, fmt.Errorf("invalid output path %q: %w", outputPath, err)
	}

	switch destination.Scheme {
	case "file", "":
		return NewLocalExporter[E](outputPath, parquetOpts...), nil
	case "s3":
		return NewS3Exporter[E](outputPath, parquetOpts,
			WithEndpoint(s3cfg.Endpoint),
			WithRegion(s3cfg.Region),
		), nil
	}
	return nil, fmt.Errorf("%w: %q", ErrUnsupportedDestination, destination.Scheme)
}

// BuildAndWrite infers the destination from outputPath (no scheme or file:// → local, s3:// → S3),
// writes records in the given format ("jsonl" or "parquet"), and returns
// ErrUnsupportedDestination or ErrUnsupportedFormat if either is unrecognised.
func BuildAndWrite[E any](ctx context.Context, outputPath, format string, s3cfg S3Config, records []E, parquetOpts ...ParquetConfigOption) error {
	if len(records) == 0 {
		slog.Info("no data to write", "output_path", outputPath)
		return nil
	}
	exp, err := Build[E](outputPath, s3cfg, parquetOpts...)
	if err != nil {
		return err
	}
	switch format {
	case "jsonl":
		return exp.WriteJSONL(ctx, records)
	case "parquet":
		return exp.WriteParquet(ctx, records)
	default:
		return fmt.Errorf("%w: %q", ErrUnsupportedFormat, format)
	}
}
