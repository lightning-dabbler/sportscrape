//go:build unit

package cli

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCreateFSCmd(t *testing.T) {
	cmd := CreateFSCmd()

	t.Run("command metadata", func(t *testing.T) {
		if cmd.Use != "foxsports" {
			t.Errorf("Use = %q, want %q", cmd.Use, "foxsports")
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
		for _, expected := range []string{"mlb", "nba", "wnba"} {
			if !subcommands[expected] {
				t.Errorf("expected subcommand %q to be registered", expected)
			}
		}
	})
}

func testFSSubCmd(t *testing.T, name string, createFn func() *cobra.Command) {
	t.Helper()
	cmd := createFn()

	t.Run("command metadata", func(t *testing.T) {
		if cmd.Use != name {
			t.Errorf("Use = %q, want %q", cmd.Use, name)
		}
	})

	t.Run("required flags exist", func(t *testing.T) {
		flags := []string{
			"feed",
			"date",
			"concurrency",
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

func TestCreateFSMLBCmd(t *testing.T) {
	testFSSubCmd(t, "mlb", createFSMLBCmd)
}

func TestCreateFSNBACmd(t *testing.T) {
	testFSSubCmd(t, "nba", createFSNBACmd)
}

func TestCreateFSWNBACmd(t *testing.T) {
	testFSSubCmd(t, "wnba", createFSWNBACmd)
}
