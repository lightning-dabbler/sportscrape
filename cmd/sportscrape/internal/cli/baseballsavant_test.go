//go:build unit

package cli

import (
	"testing"
)

func TestCreateBaseballSavantCmd(t *testing.T) {
	cmd := CreateBaseballSavantCmd()

	t.Run("command metadata", func(t *testing.T) {
		if cmd.Use != "baseballsavant" {
			t.Errorf("Use = %q, want %q", cmd.Use, "baseballsavant")
		}
		if len(cmd.Aliases) == 0 {
			t.Error("expected aliases to be set")
		}
	})

	t.Run("required flags exist", func(t *testing.T) {
		flags := []string{
			"feed",
			"date",
			"destination",
			"file-format",
			"concurrency",
			"parquet-compression",
			"parquet-row-group-size",
			"parquet-page-size",
			"parquet-write-parallelism",
			"aws-region",
			"aws-endpoint",
		}
		for _, flag := range flags {
			if cmd.Flags().Lookup(flag) == nil {
				t.Errorf("expected flag --%s to be registered", flag)
			}
		}
	})

	t.Run("flag defaults", func(t *testing.T) {
		cases := []struct {
			flag string
			want string
		}{
			{"file-format", "jsonl"},
			{"concurrency", "1"},
			{"parquet-compression", "SNAPPY"},
			{"parquet-row-group-size", "134217728"}, // 128*1024*1024
			{"parquet-page-size", "8192"},           // 8*1024
			{"parquet-write-parallelism", "1"},
			{"aws-region", "us-east-1"},
			{"aws-endpoint", ""},
			{"destination", ""},
			{"date", ""},
			{"feed", ""},
		}
		for _, tc := range cases {
			t.Run(tc.flag, func(t *testing.T) {
				got := cmd.Flags().Lookup(tc.flag).DefValue
				if got != tc.want {
					t.Errorf("flag --%s default = %q, want %q", tc.flag, got, tc.want)
				}
			})
		}
	})
}
