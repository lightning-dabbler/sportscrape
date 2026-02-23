package shared

import "github.com/spf13/cobra"

func EmbedFileFormatFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("file-format", "f", "jsonl", "The file format to export data. Options: parquet, jsonl")
}

func EmbedParquetFlags(cmd *cobra.Command) {
	cmd.Flags().String("parquet-compression", "SNAPPY", "Parquet compression options: 'UNCOMPRESSED','SNAPPY','GZIP','LZO','BROTLI','LZ4','ZSTD', and 'LZ4_RAW'")
	cmd.Flags().Int64("parquet-row-group-size", 128*1024*1024, "Row groups are horizontal partitions of data that allow for larger sequential I/O operations. 128MB default.")
	cmd.Flags().Int64("parquet-page-size", 8*1024, "Data pages are the smallest indivisible units in Parquet. Smaller pages allow more fine-grained reading (useful for single row lookups), while larger pages have less overhead from headers but potentially more parsing overhead. 8KB default.")
	cmd.Flags().Int64("parquet-write-parallelism", 1, "Max number of concurrent goroutines to write parquet file.")
}

func EmbedS3Flags(cmd *cobra.Command) {
	cmd.Flags().String("aws-region", "us-east-1", "Region of bucket")
	cmd.Flags().String("aws-endpoint", "", "Custom endpoint URL for S3-compatible storage. Leave empty to use AWS S3.")
}

func EmbedDestinationFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("destination", "d", "", "The location to write to")
}

func EmbedTimeoutFlag(cmd *cobra.Command) {
	cmd.Flags().IntP("timeout", "t", 120, "The chromium request timeout (in seconds).")
}

func EmbedDateFlag(cmd *cobra.Command) {
	cmd.Flags().String("date", "", "YYYY-MM-DD date to scrape.")
}
