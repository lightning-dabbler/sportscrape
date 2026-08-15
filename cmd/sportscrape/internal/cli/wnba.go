package cli

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/feed"
	"github.com/lightning-dabbler/sportscrape/cmd/sportscrape/internal/shared"

	"github.com/spf13/cobra"
)

func CreateWNBACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wnba",
		Short: "Extract WNBA data from wnba.com",
		Long:  "Extract WNBA data from wnba.com",
		RunE: func(cmd *cobra.Command, args []string) error {
			return shared.Run(cmd, "wnba", "")
		},
	}
	cmd.Flags().IntP("concurrency", "c", 1, fmt.Sprintf("Max number of concurrent goroutines. Dependent on data feed (%s)", feed.WNBAConcurrencyOptions))
	cmd.Flags().String("feed", "", fmt.Sprintf("The data feed to extract. Options: %s", feed.WNBAOptions))
	shared.EmbedDateFlag(cmd)
	shared.EmbedEndDateFlag(cmd)
	shared.EmbedTimeoutFlag(cmd)
	shared.EmbedDestinationFlag(cmd)
	shared.EmbedFileFormatFlag(cmd)
	shared.EmbedParquetFlags(cmd)
	shared.EmbedS3Flags(cmd)

	return cmd
}
