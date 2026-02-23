package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func createFSMLBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mlb",
		Short: "Scrape and download MLB data",
		Long:  "Scrape and download MLB data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "foxsports", "mlb")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.FSMLBConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to scrape and export. Options: %s", feed.FSMLBOptions))

	shared.EmbedDateFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)

	return cmd
}

func createFSNBACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nba",
		Short: "Scrape and download NBA data",
		Long:  "Scrape and download NBA data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "foxsports", "nba")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.FSNBAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to scrape and export. Options: %s", feed.FSNBAOptions))

	shared.EmbedDateFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)

	return cmd
}

func createFSWNBACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wnba",
		Short: "Scrape and download WNBA data",
		Long:  "Scrape and download WNBA data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "foxsports", "wnba")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.FSNBAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to scrape and export. Options: %s", feed.FSNBAOptions))

	shared.EmbedDateFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)

	return cmd
}

func CreateFSCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "foxsports",
		Aliases: []string{"fs"},
		Short:   "Scrape fox sports data",
		Long:    "Scrape fox sports data",
	}
	// Store subcommands (mlb, nba)
	cmd.AddCommand(createFSMLBCmd(), createFSNBACmd(), createFSWNBACmd())
	return cmd
}
