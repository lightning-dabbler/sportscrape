package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func CreateNHLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nhl",
		Short: "Extract NHL data",
		Long:  "Extract NHL data from api-web.nhle.com",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "nhl", "")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.NHLConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.NHLOptions))
	shared.EmbedDateFlag(cmd)
	shared.EmbedTimeoutFlag(cmd)
	shared.EmbedFetchRetryFlags(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)
	return cmd
}
