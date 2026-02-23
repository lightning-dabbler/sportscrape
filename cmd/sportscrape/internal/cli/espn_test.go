//go:build unit

package cli

import (
	"testing"
)

func TestCreateESPNCmd(t *testing.T) {
	cmd := CreateESPNCmd()

	t.Run("command metadata", func(t *testing.T) {
		if cmd.Use != "espn" {
			t.Errorf("Use = %q, want %q", cmd.Use, "espn")
		}
	})

	t.Run("subcommands", func(t *testing.T) {
		subcommands := map[string]bool{}
		for _, sub := range cmd.Commands() {
			subcommands[sub.Use] = true
		}
		if !subcommands["ufc"] {
			t.Error("expected subcommand 'ufc' to be registered")
		}
	})
}

func TestCreateESPNUFCCmd(t *testing.T) {
	cmd := createESPNUFCCmd()

	t.Run("command metadata", func(t *testing.T) {
		if cmd.Use != "ufc" {
			t.Errorf("Use = %q, want %q", cmd.Use, "ufc")
		}
	})

	t.Run("required flags exist", func(t *testing.T) {
		flags := []string{
			"feed",
			"year",
			"concurrency",
			"timeout",
			"destination",
			"file-format",
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
			{"concurrency", "1"},
			{"timeout", "120"},
			{"file-format", "jsonl"},
			{"parquet-compression", "SNAPPY"},
			{"parquet-row-group-size", "134217728"},
			{"parquet-page-size", "8192"},
			{"parquet-write-parallelism", "1"},
			{"aws-region", "us-east-1"},
			{"aws-endpoint", ""},
			{"destination", ""},
			{"feed", ""},
			{"year", ""},
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
