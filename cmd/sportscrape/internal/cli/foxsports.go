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
		Short: "Extract MLB data",
		Long:  "Extract MLB data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "foxsports", "mlb")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.FSMLBConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.FSMLBOptions))

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
		Short: "Extract NBA data",
		Long:  "Extract NBA data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "foxsports", "nba")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.FSNBAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.FSNBAOptions))

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
		Short: "Extract WNBA data",
		Long:  "Extract WNBA data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "foxsports", "wnba")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.FSNBAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.FSNBAOptions))

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
		Short:   "Extract Fox sports data",
		Long:    "Extract Fox sports data",
	}
	// Store subcommands (mlb, nba, wnba)
	cmd.AddCommand(createFSMLBCmd(), createFSNBACmd(), createFSWNBACmd())
	return cmd
}
