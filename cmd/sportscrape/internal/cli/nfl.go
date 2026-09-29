package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl"

	"github.com/spf13/cobra"
)

func CreateNFLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nfl",
		Short: "Extract NFL data",
		Long: fmt.Sprintf("Extract NFL data from nfl.com\n\n"+
			"The nfl.com web client credentials used to authenticate with api.nfl.com can be overridden by setting both the %s and %s environment variables.",
			nfl.ClientKeyEnv, nfl.ClientSecretEnv),
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "nfl", "")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.NFLConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.NFLOptions))
	shared.EmbedDateFlag(cmd)
	shared.EmbedTimeoutFlag(cmd)
	shared.EmbedFetchRetryFlags(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)
	return cmd
}
