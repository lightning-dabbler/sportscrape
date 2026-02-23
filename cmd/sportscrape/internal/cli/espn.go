package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func createESPNUFCCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ufc",
		Short: "Scrape and download UFC data",
		Long:  "Scrape and download UFC data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "espn", "ufc")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.ESPNMMAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to scrape and export. Options: %s", feed.ESPNMMAOptions))
	cmd.Flags().String("year", "", "YYYY year to scrape.")
	shared.EmbedTimeoutFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)

	return cmd
}

func CreateESPNCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "espn",
		Short: "Scrape ESPN sports data",
		Long:  "Scrape ESPN sports data",
	}
	// Store subcommands (ufc)
	cmd.AddCommand(createESPNUFCCmd())
	return cmd
}
