package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func CreateNBACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nba",
		Short: "Extract NBA data from nba.com",
		Long:  "Extract NBA data from nba.com",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "nba", "")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.NBAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.NBAOptions))
	shared.EmbedDateFlag(cmd)
	shared.EmbedTimeoutFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)

	return cmd
}
