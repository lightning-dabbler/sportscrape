package exporters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xitongsys/parquet-go-source/s3v2"
	"github.com/xitongsys/parquet-go/writer"
)

type AWSConfigOption func(*aws.Config)

func WithRegion(region string) AWSConfigOption {
	return func(c *aws.Config) { c.Region = region }
}

func WithEndpoint(endpoint string) AWSConfigOption {
	return func(c *aws.Config) {
		if endpoint != "" {
			c.BaseEndpoint = aws.String(endpoint)
		}
	}
}

func WithStaticCredentials(accessKey, secretKey string) AWSConfigOption {
	return func(c *aws.Config) {
		c.Credentials = credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	}
}

func newAWSConfig(opts ...AWSConfigOption) aws.Config {
	cfg := aws.Config{Region: "us-east-1"}
	for _, o := range opts {
		o(&cfg)
	}
	return cfg
}

type s3Path struct {
	bucket string
	key    string
}

func parseS3Path(raw string) (s3Path, error) {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme != "s3" {
		return s3Path{}, fmt.Errorf("invalid s3 path %q: must be s3://bucket/key", raw)
	}
	return s3Path{
		bucket: u.Host,
		key:    strings.TrimLeft(u.Path, "/"),
	}, nil
}

type S3Exporter[E any] struct {
	Destination   string
	ParquetConfig ParquetConfig
	AWSConfig     aws.Config
}

func NewS3Exporter[E any](destination string, parquetOpts []ParquetConfigOption, awsOpts ...AWSConfigOption) *S3Exporter[E] {
	return &S3Exporter[E]{
		Destination:   destination,
		ParquetConfig: newParquetConfig(parquetOpts...),
		AWSConfig:     newAWSConfig(awsOpts...),
	}
}

func (e *S3Exporter[E]) WriteParquet(ctx context.Context, records []E) error {
	if len(records) == 0 {
		return fmt.Errorf("no records supplied")
	}
	path, err := parseS3Path(e.Destination)
	if err != nil {
		return err
	}

	cfg := e.ParquetConfig
	if cfg.Parallelism == 0 {
		cfg.Parallelism = 1
	}

	client := s3.NewFromConfig(e.AWSConfig)
	fw, err := s3v2.NewS3FileWriterWithClient(ctx, client, path.bucket, path.key, nil)
	if err != nil {
		return fmt.Errorf("open s3 parquet writer %s: %w", e.Destination, err)
	}
	defer fw.Close()

	var zero E
	pw, err := writer.NewParquetWriter(fw, &zero, cfg.Parallelism)
	if err != nil {
		return fmt.Errorf("create parquet writer: %w", err)
	}
	if cfg.RowGroupSize != nil {
		pw.RowGroupSize = *cfg.RowGroupSize
	}
	if cfg.PageSize != nil {
		pw.PageSize = *cfg.PageSize
	}
	if cfg.CompressionType != nil {
		pw.CompressionType = *cfg.CompressionType
	}

	for i, record := range records {
		if err := pw.Write(record); err != nil {
			return fmt.Errorf("write record %d: %w", i, err)
		}
	}
	if err := pw.WriteStop(); err != nil {
		return fmt.Errorf("finalise parquet: %w", err)
	}
	slog.Info("File written", "destination", e.Destination)
	return nil
}

func (e *S3Exporter[E]) WriteJSONL(ctx context.Context, records []E) error {
	if len(records) == 0 {
		return fmt.Errorf("no records supplied")
	}
	path, err := parseS3Path(e.Destination)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	for i, record := range records {
		if err := enc.Encode(record); err != nil {
			return fmt.Errorf("encode record %d: %w", i, err)
		}
	}

	client := s3.NewFromConfig(e.AWSConfig)
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(path.bucket),
		Key:           aws.String(path.key),
		Body:          bytes.NewReader(buf.Bytes()),
		ContentLength: aws.Int64(int64(buf.Len())),
		ContentType:   aws.String("application/jsonl"),
	})
	if err != nil {
		return fmt.Errorf("put object %s: %w", e.Destination, err)
	}
	slog.Info("File written", "destination", e.Destination)
	return nil
}
