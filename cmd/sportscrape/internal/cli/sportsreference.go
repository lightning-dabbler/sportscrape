package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func createSRNBACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nba",
		Short: "Extract NBA data",
		Long:  "Extract NBA data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "sportsreference", "nba")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.SRNBAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.SRNBAOptions))
	shared.EmbedDateFlag(cmd)
	shared.EmbedTimeoutFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)

	return cmd
}

func CreateSRCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sportsreference",
		Aliases: []string{"sr"},
		Short:   "Extract sports reference data",
		Long:    "Extract sports reference data",
	}
	// Store subcommands (nba)
	cmd.AddCommand(createSRNBACmd())
	return cmd
}
