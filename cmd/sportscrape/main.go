package main

import (
	"log/slog"
	"os"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/cli"
	"github.com/lightning-dabbler/sportscrape/version"

	"github.com/spf13/cobra"
)

func embedLoggerFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String("log-level", "info", "log level: debug, info, warn, error")
	cmd.PersistentFlags().String("log-format", "text", "log format: text, json")
}

func initLogger(cmd *cobra.Command) error {
	logLevel, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return err
	}
	logFormat, err := cmd.Flags().GetString("log-format")
	if err != nil {
		return err
	}

	var level slog.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		return err
	}

	opts := &slog.HandlerOptions{Level: level}
	var handler slog.Handler
	switch logFormat {
	case "json":
		handler = slog.NewJSONHandler(os.Stderr, opts)
	default:
		handler = slog.NewTextHandler(os.Stderr, opts)
	}

	slog.SetDefault(slog.New(handler))
	return nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "sportscrape",
		Short:   "Extract sports data",
		Long:    "Extract sports data from supported providers",
		Version: version.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initLogger(cmd)
		},
	}
	embedLoggerFlag(rootCmd)
	// Store subcommands (foxsports, sportsreference, baseballsavant, espn, nba)
	rootCmd.AddCommand(
		cli.CreateFSCmd(),
		cli.CreateSRCmd(),
		cli.CreateBaseballSavantCmd(),
		cli.CreateESPNCmd(),
		cli.CreateNBACmd(),
	)
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
