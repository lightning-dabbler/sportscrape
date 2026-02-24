package exporters

import "github.com/xitongsys/parquet-go/parquet"

type ParquetConfig struct {
	CompressionType *parquet.CompressionCodec
	RowGroupSize    *int64
	PageSize        *int64
	Parallelism     int64
}

type ParquetConfigOption func(*ParquetConfig)

func WithCompressionType(c parquet.CompressionCodec) ParquetConfigOption {
	return func(cfg *ParquetConfig) { cfg.CompressionType = &c }
}

func WithRowGroupSize(n int64) ParquetConfigOption {
	return func(cfg *ParquetConfig) { cfg.RowGroupSize = &n }
}

func WithPageSize(n int64) ParquetConfigOption {
	return func(cfg *ParquetConfig) { cfg.PageSize = &n }
}

func WithParallelism(n int64) ParquetConfigOption {
	return func(cfg *ParquetConfig) { cfg.Parallelism = n }
}

func newParquetConfig(opts ...ParquetConfigOption) ParquetConfig {
	cfg := ParquetConfig{}
	for _, o := range opts {
		o(&cfg)
	}
	return cfg
}
