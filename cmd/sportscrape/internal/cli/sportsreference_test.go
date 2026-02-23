//go:build unit

package cli

import (
	"testing"
)

func TestCreateSRCmd(t *testing.T) {
	cmd := CreateSRCmd()

	t.Run("command metadata", func(t *testing.T) {
		if cmd.Use != "sportsreference" {
			t.Errorf("Use = %q, want %q", cmd.Use, "sportsreference")
		}
		if len(cmd.Aliases) == 0 {
			t.Error("expected aliases to be set")
		}
	})

	t.Run("subcommands", func(t *testing.T) {
		subcommands := map[string]bool{}
		for _, sub := range cmd.Commands() {
			subcommands[sub.Use] = true
		}
		if !subcommands["nba"] {
			t.Error("expected subcommand 'nba' to be registered")
		}
	})
}

func TestCreateSRNBACmd(t *testing.T) {
	cmd := createSRNBACmd()

	t.Run("command metadata", func(t *testing.T) {
		if cmd.Use != "nba" {
			t.Errorf("Use = %q, want %q", cmd.Use, "nba")
		}
	})

	t.Run("required flags exist", func(t *testing.T) {
		flags := []string{
			"feed",
			"date",
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
