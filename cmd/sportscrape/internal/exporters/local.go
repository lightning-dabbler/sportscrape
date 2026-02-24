package exporters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

type LocalExporter[E any] struct {
	Destination   string
	ParquetConfig ParquetConfig
}

func NewLocalExporter[E any](destination string, opts ...ParquetConfigOption) *LocalExporter[E] {
	return &LocalExporter[E]{
		Destination:   destination,
		ParquetConfig: newParquetConfig(opts...),
	}
}

func (e *LocalExporter[E]) WriteParquet(ctx context.Context, records []E) error {
	if len(records) == 0 {
		return fmt.Errorf("no records supplied")
	}
	if err := ensureDir(e.Destination); err != nil {
		return err
	}
	fw, err := local.NewLocalFileWriter(e.Destination)
	if err != nil {
		return fmt.Errorf("open local parquet writer %s: %w", e.Destination, err)
	}
	defer fw.Close()

	cfg := e.ParquetConfig
	if cfg.Parallelism == 0 {
		cfg.Parallelism = 1
	}

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

func (e *LocalExporter[E]) WriteJSONL(ctx context.Context, records []E) error {
	if len(records) == 0 {
		return fmt.Errorf("no records supplied")
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	for i, record := range records {
		if err := enc.Encode(record); err != nil {
			return fmt.Errorf("encode record %d: %w", i, err)
		}
	}

	if err := ensureDir(e.Destination); err != nil {
		return err
	}
	if err := os.WriteFile(e.Destination, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", e.Destination, err)
	}
	slog.Info("File written", "destination", e.Destination)
	return nil
}

func ensureDir(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", filepath.Dir(path), err)
	}
	return nil
}
